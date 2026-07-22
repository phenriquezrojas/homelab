# Homelab

Repositorio único para documentar y gestionar un homelab self-hosted reproducible y recuperable.
Hardware is disposable.
The repository is the source of truth.
Data is the only irreplaceable asset.

## Estado

El proyecto está en **Sprint -1: fundaciones documentales y organizacionales**. No hay servicios desplegados ni automatizaciones funcionales. Consulte [ROADMAP.md](ROADMAP.md) y [CURRENT_STATE.md](CURRENT_STATE.md).

## Decisiones vigentes

- Ubuntu Server LTS como sistema operativo objetivo.
- Docker como modelo de ejecución de servicios.
- Persistencia bajo `/srv` y dominio interno `home.arpa`.
- Tailscale para acceso remoto privado y Caddy como proxy inverso.
- PostgreSQL y Redis como tecnologías de datos compartidas.
- Restic con Backblaze B2 para respaldos externos.
- Un único repositorio como fuente de verdad.

Los detalles están en [adr/](adr/) y el índice ejecutivo en [DECISIONS.md](DECISIONS.md).

## Estructura prevista

```text
adr/        decisiones de arquitectura
assets/     recursos estáticos de documentación
backup/     definiciones y documentación de respaldos
bootstrap/  material declarativo de preparación del host
compose/    definiciones Docker Compose por dominio
diagrams/   diagramas de arquitectura y operación
docs/       documentación transversal y convenciones
restore/    procedimientos de restauración
runbooks/   guías operativas
scripts/    automatizaciones mantenidas
services/   configuración por servicio
tests/      validaciones
```

Las carpetas son un contrato organizacional; su contenido se añadirá en los sprints correspondientes. Consulte [docs/CONVENTIONS.md](docs/CONVENTIONS.md).

## Colaboración

1. Revise los ADR y el charter antes de proponer cambios de arquitectura.
2. Use Conventional Commits y mantenga cambios pequeños y trazables.
3. Registre decisiones duraderas mediante ADR nuevos; no reescriba decisiones aceptadas.
4. Nunca registre secretos, archivos de entorno ni material de respaldo.

## Documentación clave

- [Visión](VISION.md)
- [Project Charter](PROJECT_CHARTER.md)
- [Hoja de ruta](ROADMAP.md)
- [Estado actual](CURRENT_STATE.md)
- [Índice de decisiones](DECISIONS.md)
- [Convenciones](docs/CONVENTIONS.md)
- [Historial de cambios](CHANGELOG.md)

## Arquitectura

Phone
↓
Immich
↓
Restic
↓
Backblaze

## Licencia

El proyecto se distribuye bajo MIT. El texto completo de la licencia debe incorporarse antes de una distribución pública.
