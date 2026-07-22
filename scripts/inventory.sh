#!/usr/bin/env bash
#
# Host inventory collector for Homelab.
#
# This script only reads host state. It does not install packages, modify
# configuration, restart services, or contact remote systems.

REPORT_DIR="reports"
REPORT_DATE="$(date '+%Y-%m-%d_%H-%M-%S' 2>/dev/null || printf 'unknown-date')"
HOSTNAME_VALUE="$(hostname 2>/dev/null || printf 'unknown-host')"
REPORT_HOSTNAME="$(printf '%s' "$HOSTNAME_VALUE" | tr -cs '[:alnum:].-' '-' | sed 's/^-//; s/-$//')"
[ -n "$REPORT_HOSTNAME" ] || REPORT_HOSTNAME='unknown-host'
REPORT_FILE="${REPORT_DIR}/HOST_INVENTORY_${REPORT_HOSTNAME}_${REPORT_DATE}.md"
TEMP_OUTPUT="${REPORT_FILE}.tmp.$$"

cleanup() {
    rm -f "$TEMP_OUTPUT"
}

trap cleanup EXIT HUP INT TERM

mkdir -p "$REPORT_DIR" || {
    printf '%s\n' "Unable to create report directory: $REPORT_DIR" >&2
    exit 1
}

section() {
    printf '\n## %s\n\n' "$1"
}

not_available() {
    printf '%s\n' 'Not Available'
}

command_not_installed() {
    printf '%s\n' 'Command Not Installed'
}

permission_denied() {
    printf '%s\n' 'Permission Denied'
}

command_failed() {
    printf '%s\n' 'Command Failed'
}

# Classify a failed installed command from its captured diagnostic. This keeps
# command_section simple and avoids relying on implementation-specific codes.
command_failure_status() {
    if grep -qi 'permission denied' "$TEMP_OUTPUT" 2>/dev/null; then
        permission_denied
    else
        command_failed
    fi
}

code_start() {
    printf '%s\n' '```text'
}

code_end() {
    printf '%s\n' '```'
}

# Run a fixed command and its arguments. This function intentionally does not
# evaluate shell expressions or pipelines supplied by callers.
command_section() {
    title=$1
    command_name=$2
    shift 2

    section "$title"
    code_start
    if command -v "$command_name" >/dev/null 2>&1; then
        if "$command_name" "$@" >"$TEMP_OUTPUT" 2>&1; then
            if [ -s "$TEMP_OUTPUT" ]; then
                cat "$TEMP_OUTPUT"
            else
                not_available
            fi
        else
            command_failure_status
        fi
    else
        command_not_installed
    fi
    code_end
}

file_section() {
    title=$1
    file=$2

    section "$title"
    code_start
    if [ -r "$file" ]; then
        cat "$file"
    elif [ -e "$file" ]; then
        permission_denied
    else
        not_available
    fi
    code_end
}

ubuntu_section() {
    section 'Ubuntu'
    code_start
    if [ -r /etc/os-release ]; then
        awk -F= '/^(PRETTY_NAME|NAME|VERSION)=/ { print }' /etc/os-release
    else
        not_available
    fi
    code_end
}

summary_ubuntu() {
    if [ -r /etc/os-release ]; then
        awk -F= '$1 == "PRETTY_NAME" { value=$2; gsub(/^"|"$/, "", value); print value; exit }' /etc/os-release
    else
        not_available
    fi
}

summary_cpu() {
    if ! command -v lscpu >/dev/null 2>&1; then
        command_not_installed
    elif lscpu >"$TEMP_OUTPUT" 2>&1; then
        awk -F: '$1 == "Model name" { sub(/^[[:space:]]+/, "", $2); print $2; found=1; exit } $1 == "Architecture" { fallback=$2 } END { if (!found && fallback != "") { sub(/^[[:space:]]+/, "", fallback); print fallback } }' "$TEMP_OUTPUT"
    else
        command_failure_status
    fi
}

summary_ram() {
    if ! command -v free >/dev/null 2>&1; then
        command_not_installed
    elif free -h >"$TEMP_OUTPUT" 2>&1; then
        awk '/^Mem:/ { print $2; found=1 } END { if (!found) print "Not Available" }' "$TEMP_OUTPUT"
    else
        command_failure_status
    fi
}

summary_main_disk() {
    if ! command -v lsblk >/dev/null 2>&1; then
        command_not_installed
    elif lsblk -dn -o NAME,MODEL,SIZE >"$TEMP_OUTPUT" 2>&1; then
        sed -n '1p' "$TEMP_OUTPUT"
    else
        command_failure_status
    fi
}

summary_available_space() {
    if ! command -v df >/dev/null 2>&1; then
        command_not_installed
    elif df -hP / >"$TEMP_OUTPUT" 2>&1; then
        awk 'NR == 2 { print $4; found=1 } END { if (!found) print "Not Available" }' "$TEMP_OUTPUT"
    else
        command_failure_status
    fi
}

executive_summary() {
    summary_ubuntu_value="$(summary_ubuntu)"
    summary_kernel_value="$(uname -r 2>/dev/null || command_failed)"
    summary_cpu_value="$(summary_cpu)"
    summary_ram_value="$(summary_ram)"
    summary_disk_value="$(summary_main_disk)"
    summary_space_value="$(summary_available_space)"

    printf '\n## Resumen ejecutivo\n\n'
    printf '%s\n' '| Campo | Valor |'
    printf '%s\n' '| --- | --- |'
    printf '| Hostname | %s |\n' "$HOSTNAME_VALUE"
    printf '| Fecha | %s |\n' "$(date -Iseconds 2>/dev/null || date 2>/dev/null || printf 'Not Available')"
    printf '| Ubuntu | %s |\n' "$summary_ubuntu_value"
    printf '| Kernel | %s |\n' "$summary_kernel_value"
    printf '| CPU | %s |\n' "$summary_cpu_value"
    printf '| RAM | %s |\n' "$summary_ram_value"
    printf '| Disco principal | %s |\n' "$summary_disk_value"
    printf '| Espacio disponible | %s |\n' "$summary_space_value"
    printf '%s\n' '| Estado general | Best Effort |'
}

dns_section() {
    section 'DNS'
    code_start
    if command -v resolvectl >/dev/null 2>&1; then
        if resolvectl dns >"$TEMP_OUTPUT" 2>&1; then
            cat "$TEMP_OUTPUT"
        else
            command_failure_status
        fi
    elif [ -r /etc/resolv.conf ]; then
        awk '/^[[:space:]]*nameserver[[:space:]]/ { print }' /etc/resolv.conf
    else
        not_available
    fi
    code_end
}

compose_section() {
    section 'Docker Compose'
    code_start
    if command -v docker >/dev/null 2>&1; then
        if docker compose version >"$TEMP_OUTPUT" 2>&1; then
            cat "$TEMP_OUTPUT"
        else
            command_failure_status
        fi
    elif command -v docker-compose >/dev/null 2>&1; then
        if docker-compose version >"$TEMP_OUTPUT" 2>&1; then
            cat "$TEMP_OUTPUT"
        else
            command_failure_status
        fi
    else
        command_not_installed
    fi
    code_end
}

gpu_section() {
    section 'GPU'
    code_start
    if command -v lspci >/dev/null 2>&1; then
        if lspci >"$TEMP_OUTPUT" 2>&1; then
            gpu_lines="$(awk '{ line = tolower($0); if (line ~ /vga compatible controller|3d controller|display controller/) print }' "$TEMP_OUTPUT")"
            if [ -n "$gpu_lines" ]; then
                printf '%s\n' "$gpu_lines"
            else
                not_available
            fi
        else
            command_failure_status
        fi
    else
        command_not_installed
    fi
    code_end
}

smart_section() {
    section 'SMART'
    code_start
    if ! command -v smartctl >/dev/null 2>&1; then
        command_not_installed
    elif ! smartctl --scan-open >"$TEMP_OUTPUT" 2>&1; then
        command_failure_status
    else
        devices="$(awk '/^\/dev\// { print $1 }' "$TEMP_OUTPUT")"
        if [ -z "$devices" ]; then
            not_available
        else
            printf '%s\n' '# Detected devices'
            cat "$TEMP_OUTPUT"
            for device in $devices; do
                printf '\n# Health: %s\n' "$device"
                if ! smartctl -H "$device" >"$TEMP_OUTPUT" 2>&1; then
                    command_failure_status
                else
                    cat "$TEMP_OUTPUT"
                fi
            done
        fi
    fi
    code_end
}

firewall_section() {
    section 'Firewall'
    code_start
    found=0
    if command -v ufw >/dev/null 2>&1; then
        found=1
        ufw status verbose 2>&1 || not_available
    fi
    if command -v firewall-cmd >/dev/null 2>&1; then
        found=1
        firewall-cmd --state 2>&1 || not_available
    fi
    if command -v nft >/dev/null 2>&1; then
        found=1
        nft list ruleset 2>&1 || not_available
    fi
    if [ "$found" -eq 0 ]; then
        command_not_installed
    fi
    code_end
}

ssh_section() {
    section 'SSH'
    code_start
    if command -v systemctl >/dev/null 2>&1; then
        if systemctl is-active --quiet ssh 2>/dev/null; then
            printf '%s\n' 'ssh: active'
        elif systemctl is-active --quiet sshd 2>/dev/null; then
            printf '%s\n' 'sshd: active'
        else
            printf '%s\n' 'SSH service: inactive or not found'
        fi
    else
        command_not_installed
    fi
    code_end
}

local_users_section() {
    section 'Usuarios locales'
    code_start
    if [ -r /etc/passwd ]; then
        awk -F: '$3 >= 1000 && $3 < 65534 { printf "%s (uid=%s, shell=%s)\n", $1, $3, $7 }' /etc/passwd
    else
        not_available
    fi
    code_end
}

sudo_users_section() {
    section 'Usuarios sudo'
    code_start
    if command -v getent >/dev/null 2>&1; then
        result="$(getent group sudo 2>/dev/null || getent group wheel 2>/dev/null)"
        if [ -n "$result" ]; then
            printf '%s\n' "$result"
        else
            not_available
        fi
    else
        command_not_installed
    fi
    code_end
}

cron_section() {
    section 'Cron'
    code_start
    found=0
    if [ -r /etc/crontab ]; then
        found=1
        printf '%s\n' '# /etc/crontab'
        cat /etc/crontab
    fi
    if [ -d /etc/cron.d ]; then
        found=1
        printf '%s\n' '# /etc/cron.d'
        ls -1 /etc/cron.d 2>/dev/null || true
    fi
    if command -v crontab >/dev/null 2>&1; then
        found=1
        printf '%s\n' '# Current user crontab'
        crontab -l 2>/dev/null || printf '%s\n' 'Not Available'
    fi
    [ "$found" -eq 1 ] || not_available
    code_end
}

environment_section() {
    section 'Variables de entorno relevantes'
    printf '%s\n' 'Los valores se omiten deliberadamente; solo se informa si la variable está definida.'
    code_start
    if command -v env >/dev/null 2>&1; then
        env | awk -F= '
            $1 == "PATH" || $1 == "SHELL" || $1 == "LANG" || $1 == "TERM" ||
            $1 == "USER" || $1 == "LOGNAME" || $1 == "HOME" || $1 ~ /^LC_/ ||
            $1 ~ /^XDG_/ { print $1 "=<set>"; found=1 }
            END { if (!found) print "Not Available" }
        '
    else
        command_not_installed
    fi
    code_end
}

risks_section() {
    section 'Riesgos detectados'
    printf '%s\n' 'Hallazgos heurísticos; requieren validación del operador.'
    code_start
    found=0

    if ! command -v smartctl >/dev/null 2>&1; then
        printf '%s\n' '- SMART no está disponible; no se puede evaluar la salud de discos.'
        found=1
    fi
    if command -v swapon >/dev/null 2>&1 && ! swapon --show 2>/dev/null | awk 'NR > 1 { found=1 } END { exit !found }'; then
        printf '%s\n' '- No se detectó swap activo.'
        found=1
    fi
    if command -v df >/dev/null 2>&1; then
        high_usage="$(df -P 2>/dev/null | awk 'NR > 1 { gsub(/%/, "", $5); if ($5 >= 85) print $6 "=" $5 "%" }')"
        if [ -n "$high_usage" ]; then
            printf '%s\n' '- Filesystems con uso igual o superior a 85%:'
            printf '%s\n' "$high_usage"
            found=1
        fi
    fi
    if command -v systemctl >/dev/null 2>&1 && systemctl is-active --quiet ssh 2>/dev/null; then
        printf '%s\n' '- SSH está activo; valide acceso por claves, usuarios autorizados y política de firewall.'
        found=1
    fi
    [ "$found" -eq 1 ] || printf '%s\n' 'No heuristic risks detected.'
    code_end
}

recommendations_section() {
    section 'Recomendaciones'
    printf '%s\n' '- Revise los hallazgos antes de ejecutar el Sprint 0; este reporte no altera el host.'
    printf '%s\n' '- Documente la capacidad, estructura y permisos de `/srv` antes del bootstrap.'
    printf '%s\n' '- Si SMART no está disponible, evalúe instalarlo en un sprint posterior según la política aprobada.'
    printf '%s\n' '- Verifique respaldos y restauraciones antes de confiar datos persistentes al host.'
}

{
    printf '%s\n' '# Host Inventory'
    if [ "$(id -u 2>/dev/null)" != '0' ]; then
        printf '%s\n' '' '> **WARNING**' '>' '> Some inventory sections may be incomplete because the script is not running with elevated privileges.'
    fi
    executive_summary
    printf '\nGenerated: %s\n' "$(date -Iseconds 2>/dev/null || date 2>/dev/null || printf 'Not Available')"
    printf 'Report file: `%s`\n' "$REPORT_FILE"

    command_section 'Hostname' hostname
    ubuntu_section
    command_section 'Kernel' uname -r
    command_section 'Arquitectura' uname -m
    command_section 'BIOS' dmidecode -t bios
    command_section 'Motherboard' dmidecode -t baseboard
    command_section 'CPU' lscpu
    command_section 'RAM' free -h
    gpu_section
    command_section 'Discos físicos' lsblk -d -o NAME,MODEL,SERIAL,SIZE,ROTA,TYPE
    smart_section
    command_section 'Particiones' lsblk -f
    command_section 'UUID' blkid
    command_section 'Filesystem' df -Th
    file_section '/etc/fstab' /etc/fstab
    command_section 'Uso de disco' df -hT
    command_section 'Uso de memoria' free -h
    command_section 'Swap' swapon --show
    command_section 'Load Average' uptime
    command_section 'Interfaces de red' ip -brief link
    command_section 'IP' ip -brief address
    command_section 'Gateway' ip route show default
    dns_section
    local_users_section
    sudo_users_section
    ssh_section
    firewall_section
    command_section 'Docker' docker version
    compose_section
    command_section 'Git' git --version
    cron_section
    command_section 'Systemd Timers' systemctl list-timers --all
    command_section 'Servicios activos' systemctl list-units --type=service --state=running
    command_section 'Puertos abiertos' ss -tulpn
    environment_section
    risks_section
    recommendations_section
} >"$REPORT_FILE"

printf '%s\n' "Host inventory written to: $REPORT_FILE"
