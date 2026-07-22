# Índice de decisiones arquitectónicas

Las siguientes decisiones están **aceptadas** y forman la línea base.

| ADR | Decisión | Estado |
| --- | --- | --- |
| [ADR-001](adr/ADR-001-persistence-under-srv.md) | Persistencia bajo `/srv` | Aceptada |
| [ADR-002](adr/ADR-002-docker-first.md) | Docker First | Aceptada |
| [ADR-003](adr/ADR-003-ubuntu-server.md) | Ubuntu Server LTS | Aceptada |
| [ADR-004](adr/ADR-004-single-repository.md) | Un único repositorio | Aceptada |
| [ADR-005](adr/ADR-005-restic-backblaze.md) | Restic + Backblaze B2 | Aceptada |
| [ADR-006](adr/ADR-006-tailscale.md) | Tailscale para acceso privado | Aceptada |
| [ADR-007](adr/ADR-007-caddy.md) | Caddy como proxy inverso | Aceptada |
| [ADR-008](adr/ADR-008-home-arpa.md) | `home.arpa` como dominio interno | Aceptada |

PostgreSQL y Redis son tecnologías de datos compartidas aprobadas. Su adopción concreta por servicio se documentará más adelante; un cambio arquitectónico requiere un ADR adicional.

## Política de ADR

Los ADR conservan el registro histórico. Una decisión aceptada no se edita para cambiar su significado: se crea un ADR que la sustituya o complemente y se enlazan ambos documentos.
