# Estado actual

**Fase:** Sprint -1 — Fundaciones documentales y organizacionales  
**Fecha de actualización:** 2026-07-22  
**Estado general:** base documental completada; implementación técnica no iniciada.

## Completado

- Visión, charter, roadmap y estado operativo definidos.
- Decisiones arquitectónicas registradas en ADR-001 a ADR-008.
- Convenciones de contribución, versión, nombres y Markdown establecidas.
- Estructura de carpetas reservada para sprints futuros.
- Reglas iniciales para evitar secretos y artefactos locales.

## No iniciado

- Host Ubuntu Server LTS y Docker.
- Docker Compose, redes, servicios o bases de datos.
- Tailscale, Caddy y `home.arpa`.
- Esquema físico de `/srv` y gestión de secretos.
- Restic/Backblaze B2 y pruebas de restauración.
- Monitorización, hardening, CI/CD y pruebas automatizadas.

## Restricciones vigentes

No se añadirán secretos, servicios funcionales, scripts funcionales, Docker Compose ni GitHub Actions durante Sprint -1. `bootstrap.sh` y `restore.sh` siguen siendo marcadores de posición.

## Próximo hito

Planificar Sprint 0: inventario del host, responsables, modelo de secretos, topología de acceso privada, estructura detallada de `/srv` y criterios de aceptación del bootstrap.
