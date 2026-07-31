#!/usr/bin/env bash
set -Eeuo pipefail

BOOTSTRAP_VERSION="1.1.0"
EOS_VERSION="1.0"
SCRIPT_NAME=$(basename "$0")

# Configuración de log
LOG_FILE="/var/log/homelab-bootstrap.log"
exec > >(tee -a "$LOG_FILE") 2>&1

# Flag para asegurar que apt update corra máximo una vez por ejecución, y solo si se requiere
APT_UPDATED=0

function update_apt_once() {
    if [ "$APT_UPDATED" -eq 0 ]; then
        log_info "Actualizando índices de apt..."
        apt-get update -q
        APT_UPDATED=1
    fi
}

function log_info() { echo -e "[INFO] $(date '+%Y-%m-%d %H:%M:%S') - $1"; }
function log_ok() { echo -e "[OK]   $(date '+%Y-%m-%d %H:%M:%S') - $1"; }
function log_warn() { echo -e "[WARN] $(date '+%Y-%m-%d %H:%M:%S') - $1"; }
function log_error() { echo -e "[ERROR] $(date '+%Y-%m-%d %H:%M:%S') - $1"; exit 1; }

GPG_TMP=""
REPO_TMP=""

function cleanup() {
    local rc=$1
    if [ -n "$GPG_TMP" ] && [ -f "$GPG_TMP" ]; then rm -f "$GPG_TMP"; fi
    if [ -n "$REPO_TMP" ] && [ -f "$REPO_TMP" ]; then rm -f "$REPO_TMP"; fi
    if [ "$rc" -ne 0 ]; then
        echo -e "[ERROR] $(date '+%Y-%m-%d %H:%M:%S') - El script terminó inesperadamente con código $rc. Revisar logs."
    fi
    exit "$rc"
}
trap 'cleanup $?' EXIT

# Nivel 1 de Idempotencia: Garantizar paquetes
function ensure_package_installed() {
    local pkg=$1
    if dpkg-query -W -f='${Status}' "$pkg" 2>/dev/null | grep -q "ok installed"; then
        log_info "Paquete $pkg ya está instalado."
    else
        log_info "Instalando paquete $pkg..."
        update_apt_once
        DEBIAN_FRONTEND=noninteractive apt-get install -y -q "$pkg"
    fi
}

# Nivel 1 de Idempotencia: Garantizar directorio (Convergencia Real)
function ensure_directory() {
    local dir=$1
    local expected_perms=${2:-755}
    local expected_owner="root:root"
    
    if [ ! -d "$dir" ]; then
        log_info "Creando directorio $dir..."
        mkdir -p "$dir"
    fi
    
    # Validar permisos y propietario
    local current_perms
    current_perms=$(stat -c "%a" "$dir")
    local current_owner
    current_owner=$(stat -c "%U:%G" "$dir")
    
    if [ "$current_perms" != "$expected_perms" ]; then
        log_info "Corrigiendo permisos de $dir (actual: $current_perms, esperado: $expected_perms)"
        chmod "$expected_perms" "$dir"
    else
        log_info "Permisos de $dir correctos ($current_perms)."
    fi
    
    if [ "$current_owner" != "$expected_owner" ]; then
        log_info "Corrigiendo propietario de $dir (actual: $current_owner, esperado: $expected_owner)"
        chown root:root "$dir"
    else
        log_info "Propietario de $dir correcto ($current_owner)."
    fi
}

# Nivel 2 de Idempotencia: Garantizar servicios
function ensure_service_running() {
    local svc=$1
    if systemctl is-active --quiet "$svc" && systemctl is-enabled --quiet "$svc"; then
        log_info "Servicio $svc ya está en ejecución y habilitado."
    else
        log_info "Habilitando e iniciando servicio $svc..."
        systemctl enable --now "$svc"
    fi
}

# Nivel 3 de Idempotencia: Validar componentes
function validate_component() {
    local component=$1
    log_info "Validando componente: $component"
    case "$component" in
        srv_structure)
            [ -d "/srv/homelab/app_data" ] || log_error "Falta /srv/homelab/app_data"
            [ -d "/srv/homelab/secrets" ] || log_error "Falta /srv/homelab/secrets"
            [ -d "/srv/homelab/backups" ] || log_error "Falta /srv/homelab/backups"
            local perms
            perms=$(stat -c "%a" /srv/homelab/secrets)
            [ "$perms" == "700" ] || log_error "Permisos incorrectos en /srv/homelab/secrets (esperado 700, actual $perms)"
            ;;
        base_dependencies)
            command -v curl >/dev/null || log_error "curl no instalado"
            command -v git >/dev/null || log_error "git no instalado"
            ;;
        docker)
            docker --version >/dev/null || log_error "Docker Engine no está funcional"
            docker compose version >/dev/null || log_error "Docker Compose plugin no está funcional"
            ;;
        user_docker_group)
            if id -nG "$BOOTSTRAP_USER" | grep -qw "docker"; then
                :
            else
                log_error "Usuario $BOOTSTRAP_USER no pertenece al grupo docker"
            fi
            ;;
        docker_functional)
            # Validar e instalar hello-world si no existe para no descargarlo en cada prueba
            docker image inspect hello-world >/dev/null 2>&1 || sudo -u "$BOOTSTRAP_USER" docker pull hello-world >/dev/null
            sudo -u "$BOOTSTRAP_USER" docker run --rm hello-world >/dev/null || log_error "Prueba funcional de docker falló para el usuario $BOOTSTRAP_USER"
            # No se elimina la imagen para mantener la idempotencia pura (evitar descargar siempre)
            ;;
        *)
            log_error "Componente desconocido para validación: $component"
            ;;
    esac
    log_ok "Componente $component validado correctamente."
}

# Determinar el usuario no root que está ejecutando el script a través de sudo
if [ -n "${SUDO_USER:-}" ]; then
    BOOTSTRAP_USER="$SUDO_USER"
else
    if [ "$USER" == "root" ]; then
        log_error "El script debe ser ejecutado vía sudo por el usuario operativo (ej. sudo ./bootstrap.sh), no directamente como root."
    fi
    BOOTSTRAP_USER="$USER"
fi

log_info "=========================================================="
log_info "INICIO DEL BOOTSTRAP DEL HOST HOMELAB ($SCRIPT_NAME v$BOOTSTRAP_VERSION - EOS v$EOS_VERSION)"
log_info "Usuario operativo detectado: $BOOTSTRAP_USER"
log_info "=========================================================="

# FASE 0: Verificación de prerrequisitos del SO
log_info "FASE 0: Verificando prerrequisitos..."
if [ "$EUID" -ne 0 ]; then
  log_error "Este script debe ser ejecutado con privilegios de root (ej. usando sudo)."
fi

OS_NAME=$(lsb_release -is 2>/dev/null || echo "Unknown")
if [ "$OS_NAME" != "Ubuntu" ]; then
    log_error "Sistema Operativo no soportado. Se requiere Ubuntu. (Detectado: $OS_NAME)"
fi

log_info "Verificando conectividad a Internet..."
if ! curl -fsSL --head https://download.docker.com > /dev/null 2>&1; then
    log_error "No hay conectividad a Internet o download.docker.com es inaccesible."
fi
log_ok "Prerrequisitos verificados."

# FASE 1: Garantía de estructura /srv/homelab
log_info "FASE 1: Garantía de estructura /srv/homelab..."
ensure_directory "/srv/homelab" 755
ensure_directory "/srv/homelab/app_data" 755
ensure_directory "/srv/homelab/backups" 755
ensure_directory "/srv/homelab/secrets" 700
validate_component srv_structure

# FASE 2: Actualización del sistema base (solo si es necesario en FASE 3/4)
log_info "FASE 2: Preparación del sistema base (apt index delay evaluation)..."
# La actualización se realiza de manera diferida mediante update_apt_once si se requiere instalar algo.

# FASE 3: Instalación de dependencias base
log_info "FASE 3: Instalación de dependencias base..."
ensure_package_installed curl
ensure_package_installed git
ensure_package_installed ca-certificates
ensure_package_installed gnupg
validate_component base_dependencies

# FASE 4: Instalación de Docker Engine
log_info "FASE 4: Instalación de Docker Engine..."
install -m 0755 -d /etc/apt/keyrings

GPG_TMP=$(mktemp)
curl -fsSL https://download.docker.com/linux/ubuntu/gpg -o "$GPG_TMP"
if [ -s "$GPG_TMP" ]; then
    if [ -f /etc/apt/keyrings/docker.asc ] && cmp -s "$GPG_TMP" /etc/apt/keyrings/docker.asc; then
        log_info "Clave GPG de Docker ya existe y es correcta."
    else
        log_info "Instalando/Actualizando clave GPG de Docker..."
        cp "$GPG_TMP" /etc/apt/keyrings/docker.asc
        chmod a+r /etc/apt/keyrings/docker.asc
    fi
else
    log_error "Descarga de clave GPG de Docker falló."
fi

REPO_TMP=$(mktemp)
echo "deb [arch=$(dpkg --print-architecture) signed-by=/etc/apt/keyrings/docker.asc] https://download.docker.com/linux/ubuntu $(lsb_release -cs) stable" > "$REPO_TMP"
if [ -f /etc/apt/sources.list.d/docker.list ] && cmp -s "$REPO_TMP" /etc/apt/sources.list.d/docker.list; then
    log_info "Repositorio apt de Docker ya existe y es correcto."
else
    log_info "Instalando/Actualizando repositorio apt de Docker..."
    cp "$REPO_TMP" /etc/apt/sources.list.d/docker.list
    update_apt_once
fi

ensure_package_installed docker-ce
ensure_package_installed docker-ce-cli
ensure_package_installed containerd.io
ensure_package_installed docker-buildx-plugin
ensure_package_installed docker-compose-plugin
ensure_service_running docker
validate_component docker

# FASE 5: Configuración del usuario operativo
log_info "FASE 5: Configuración del usuario operativo..."
if id -nG "$BOOTSTRAP_USER" | grep -qw "docker"; then
    log_info "Usuario $BOOTSTRAP_USER ya está en el grupo docker."
else
    log_info "Añadiendo usuario $BOOTSTRAP_USER al grupo docker..."
    usermod -aG docker "$BOOTSTRAP_USER"
    log_warn "NOTA: Para que los cambios de grupo tengan efecto, $BOOTSTRAP_USER debe reiniciar sesión o ejecutar 'newgrp docker'."
fi
validate_component user_docker_group

# FASE 6: Verificación funcional end-to-end
log_info "FASE 6: Verificación funcional end-to-end..."
validate_component docker_functional

log_info "=========================================================="
log_ok "FIN DEL BOOTSTRAP: El host está preparado."
log_info "=========================================================="
