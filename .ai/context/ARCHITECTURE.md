# Architecture Context

La arquitectura aprobada establece Ubuntu Server LTS como host, Docker como ejecución de servicios, persistencia bajo `/srv`, acceso privado mediante Tailscale, Caddy como proxy HTTP(S), `home.arpa` como dominio interno, PostgreSQL y Redis como tecnologías de datos compartidas, y Restic con Backblaze B2 para respaldo externo.

No describe una implementación actual: Sprint -0.7 sigue sin infraestructura desplegada. Las decisiones detalladas viven en [ADR](../../adr/) y su índice en [DECISIONS.md](../../DECISIONS.md).
