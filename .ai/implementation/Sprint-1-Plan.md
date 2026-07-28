---
Estado: Approved
Versión: 1.1
Autor: AI Agent
Relacionado: Sprint 1.A, Sprint 1.B, ADR-001, ADR-002, ADR-003, ADR-009
---

# Implementation Plan — Bootstrap del Host (Sprint 1)

## Objetivo

Definir el contrato técnico completo que Sprint 1.B deberá ejecutar para convertir un Ubuntu Server LTS recién instalado en una plataforma lista para despliegue de servicios containerizados.

## Contexto y ADR

Este plan se fundamenta en los siguientes ADR aprobados:

- [ADR-001](../../adr/ADR-001-persistence-under-srv.md) — Persistencia bajo `/srv`. El bootstrap es responsable de **garantizar** la existencia de la estructura `/srv/homelab` con sus permisos correctos. Si ya existe (caso del host actual, Sprint 0.B), la verifica y continúa. Si no existe (host nuevo), la crea. Este comportamiento hace al bootstrap completamente reproducible y auto-suficiente.
- [ADR-002](../../adr/ADR-002-docker-first.md) — Docker First. La instalación de Docker Engine es el entregable técnico central del Bootstrap.
- [ADR-003](../../adr/ADR-003-ubuntu-server.md) — Ubuntu Server LTS. El script se diseña y prueba exclusivamente sobre Ubuntu Server LTS. La versión exacta del host se determinará en la verificación de prerrequisitos del propio script.

## Prerrequisitos del Host

Antes de ejecutar el script de bootstrap, el host debe cumplir:

1. Ubuntu Server LTS instalado y con acceso a Internet.
2. Usuario con privilegios de `sudo` disponible.
3. Repositorio `phenriquezrojas/homelab` clonado en `~/homelab`.
4. Acceso SSH confirmado desde la máquina de gestión.

> [!NOTE]
> La estructura `/srv/homelab` ya **no es un prerrequisito externo**. El bootstrap es responsable de garantizarla. Si no existe, la crea con los permisos correctos. Si ya existe, la valida y continúa.

## Alcance

### Incluido en el Bootstrap

1. **Verificación de prerrequisitos:** Validar versión Ubuntu, usuario y conectividad antes de proceder.
2. **Garantía de estructura `/srv/homelab`:** Verificar si los directorios `app_data`, `secrets`, `backups` existen bajo `/srv/homelab`. Si no existen, crearlos con los permisos correctos (`700` para `secrets`, `755` para el resto, propietario `root`).
3. **Actualización del sistema:** `apt-get update` y `apt-get upgrade` para asegurar base de seguridad.
4. **Instalación de dependencias base:** Herramientas mínimas necesarias (`curl`, `git`, `ca-certificates`, `gnupg`).
5. **Instalación de Docker Engine:** Siguiendo el procedimiento oficial de Docker para Ubuntu (repositorio apt de Docker, no el paquete de Ubuntu).
6. **Instalación de Docker Compose Plugin:** Como plugin oficial de Docker (v2), no la versión standalone.
7. **Configuración del usuario:** Añadir el usuario operativo al grupo `docker` para evitar `sudo` en operaciones cotidianas.
8. **Verificación final del estado:** Confirmar que Docker y Docker Compose están operativos con comandos de prueba.
9. **Logging de la ejecución:** Registro de cada paso con timestamp en un archivo local.

### Excluido del Bootstrap

- Instalación de Tailscale.
- Instalación o configuración de Caddy.
- Creación de redes Docker nombradas para servicios.
- Configuración DNS (`home.arpa`).
- Despliegue de ningún servicio (ni Immich, ni PostgreSQL, ni Redis).
- Configuración de Restic o backups.
- Cualquier cambio de configuración en el enrutador físico.

## Diseño del Script

### Localización

```
bootstrap/bootstrap.sh
```

### Estructura de Flujo

El script seguirá este flujo secuencial y explícito:

```
[INICIO]
  │
  ├─ FASE 0: Verificación de prerrequisitos del SO
  │    ├─ Verificar que el SO es Ubuntu LTS
  │    ├─ Verificar conectividad a Internet
  │    └─ Si alguna verificación falla → ABORT con mensaje claro
  │
  ├─ FASE 1: Garantía de estructura /srv/homelab
  │    ├─ ensure_directory /srv/homelab            (crea si no existe)
  │    ├─ ensure_directory /srv/homelab/app_data   (crea si no existe, permisos 755)
  │    ├─ ensure_directory /srv/homelab/backups    (crea si no existe, permisos 755)
  │    ├─ ensure_directory /srv/homelab/secrets    (crea si no existe, permisos 700)
  │    └─ validate_component srv_structure
  │
  ├─ FASE 2: Actualización del sistema base
  │    ├─ apt-get update
  │    └─ apt-get upgrade -y
  │
  ├─ FASE 3: Instalación de dependencias base
  │    ├─ ensure_package_installed curl
  │    ├─ ensure_package_installed git
  │    ├─ ensure_package_installed ca-certificates
  │    ├─ ensure_package_installed gnupg
  │    └─ validate_component base_dependencies
  │
  ├─ FASE 4: Instalación de Docker Engine
  │    ├─ ensure_package_installed docker-ce (vía repositorio oficial Docker)
  │    ├─ ensure_package_installed docker-ce-cli
  │    ├─ ensure_package_installed containerd.io
  │    ├─ ensure_package_installed docker-buildx-plugin
  │    ├─ ensure_package_installed docker-compose-plugin
  │    ├─ ensure_service_running docker
  │    └─ validate_component docker
  │
  ├─ FASE 5: Configuración del usuario operativo
  │    ├─ ensure_user_in_group docker $BOOTSTRAP_USER
  │    └─ validate_component user_docker_group
  │
  ├─ FASE 6: Verificación funcional end-to-end
  │    ├─ docker run hello-world
  │    └─ validate_component docker_functional
  │
  └─ [FIN: Log de éxito con timestamp]
```

### Principio de Idempotencia del Bootstrap (Aprobado)

El Bootstrap delega la gestión del estado de paquetes al gestor nativo del sistema operativo (`apt`). Después de cada operación, ejecuta validaciones explícitas sobre el estado esperado del sistema. **El Bootstrap nunca considera una operación exitosa únicamente porque el comando terminó sin error; la operación solo se considera completada cuando las verificaciones posteriores demuestran que el estado objetivo fue alcanzado.**

Este principio se implementa mediante tres niveles de funciones:

| Función | Responsabilidad | Ejemplo |
|---|---|---|
| `ensure_package_installed <pkg>` | Delega en `apt`; instala si falta, no falla si ya existe | `ensure_package_installed docker-ce` |
| `ensure_service_running <svc>` | Verifica que el servicio está activo y habilitado; lo activa si no lo está | `ensure_service_running docker` |
| `validate_component <name>` | Verifica el estado funcional esperado del componente; aborta con error claro si falla | `validate_component docker` |

Aplicación por caso:
- `apt-get install` es idempotente por diseño: `ensure_package_installed` lo aprovecha directamente.
- `usermod -aG` NO es idempotente: se verifica con `id -nG $USER` antes de ejecutar.
- La clave GPG y el repositorio apt de Docker se verifican antes de añadirse.
- La estructura `/srv/homelab` se garantiza con `ensure_directory` que crea solo si no existe y valida permisos.

### Estrategia de Logging

- Todo output del script se duplica a `~/homelab-bootstrap.log` mediante `tee`.
- Cada paso imprime un prefijo `[INFO]`, `[OK]`, `[WARN]` o `[ERROR]`.
- El archivo de log incluye timestamp al inicio y al final de la ejecución.
- En caso de fallo, el log preserva el último error antes del abort.

### Estrategia de Rollback

El Bootstrap no tiene rollback automático porque las acciones son aditivas (instalaciones), no destructivas. La estrategia de reversión es:

| Acción | Reversión |
|---|---|
| `apt-get upgrade` | No revertible; riesgo bajo ya que es actualización de seguridad |
| Dependencias instaladas | `apt-get remove --purge <paquete>` |
| Docker Engine instalado | `apt-get remove --purge docker-ce docker-ce-cli containerd.io` + borrar repositorio y clave GPG |
| Usuario en grupo docker | `gpasswd -d $USER docker` |
| Estructura `/srv` | Creada por el propio script en FASE 1; si ya existía, no se modifica |

## Validaciones

| Validación | Comando de evidencia |
|---|---|
| Ubuntu LTS detectado | `lsb_release -d` |
| Dependencias instaladas | `which curl git` |
| Docker instalado | `docker --version` |
| Docker Compose Plugin | `docker compose version` |
| Usuario en grupo docker | `groups $USER` |
| Docker funcional | `docker run hello-world` |
| Log de ejecución generado | `cat ~/homelab-bootstrap.log` |

## Criterios de Éxito (Success Criteria)

Este plan se considera ejecutado exitosamente cuando Sprint 1.B pueda demostrar:

- [ ] El script `bootstrap/bootstrap.sh` existe en el repositorio y es ejecutable.
- [ ] El script ejecutado en el host `homelab` termina sin errores.
- [ ] `docker --version` retorna versión Docker CE ≥ 24.0.
- [ ] `docker compose version` retorna versión Compose Plugin ≥ 2.0.
- [ ] El usuario `peter` pertenece al grupo `docker`.
- [ ] `docker run hello-world` ejecuta correctamente.
- [ ] El archivo `~/homelab-bootstrap.log` existe y contiene la traza completa.
- [ ] El script puede ejecutarse por segunda vez sin errores (idempotencia).

## Riesgos y Mitigaciones

| Riesgo | Probabilidad | Impacto | Mitigación |
|---|---|---|---|
| La versión LTS del host no es la esperada | Media | Medio | El script verifica el SO en FASE 0 y aborta con mensaje claro |
| Docker ya estaba instalado manualmente con versión distinta | Baja | Alto | El script detecta instalación previa y alerta; no sobreescribe sin confirmación |
| Sin conectividad durante la instalación | Baja | Alto | FASE 0 verifica conectividad antes de proceder |
| El usuario ya está en el grupo docker | Media | Bajo | `id -nG` verifica antes del `usermod` |
| `/srv/homelab` no existe (host nuevo) | Media | Medio | FASE 1 la crea automáticamente con `ensure_directory`; no aborta |
| Cambio de URL del repositorio apt de Docker | Baja | Alto | El script debe obtener la URL de la documentación oficial; se documenta la URL en el script con comentario de versión |

## Trazabilidad

- **Fase:** Phase 1 — Bootstrap
- **Sprint:** Sprint 1
- **Tareas:** HL-0009, HL-0010, HL-0011, HL-0012
- **ADR relacionados:** ADR-001, ADR-002, ADR-003, ADR-009
