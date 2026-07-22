# Hoja de ruta

La hoja de ruta se organiza por resultados y dependencias, no por fechas. Un sprint termina cuando sus criterios de salida se cumplen y su documentación queda actualizada.

| Sprint | Objetivo | Resultado esperado | Dependencias |
| --- | --- | --- | --- |
| -1 — Fundaciones | Definir base documental y organizacional. | Charter, visión, ADR, convenciones, estado y roadmap completos. | Ninguna |
| 0 — Preparación | Planificar host y repositorio. | Inventario del host, modelo de secretos, plan de red y estructura de `/srv` aprobados. | Sprint -1 |
| 1 — Bootstrap | Reproducir base del host. | Proceso validado para Ubuntu Server LTS, Docker y directorios persistentes. | Sprint 0 |
| 2 — Plataforma central | Establecer componentes compartidos. | Tailscale, Caddy, `home.arpa` y dependencias de datos operables. | Sprint 1 |
| 3 — Respaldo | Proteger y recuperar datos. | Restic a Backblaze B2 y restauración probada. | Sprint 2 |
| 4 — Monitorización | Obtener visibilidad. | Métricas, logs, alertas y runbooks. | Sprints 2 y 3 |
| 5 — Immich | Incorporar el primer servicio priorizado. | Servicio documentado, respaldado, monitorizado y recuperable. | Sprints 2, 3 y 4 |
| 6 — Hardening | Reducir riesgos. | Revisión de acceso, secretos, actualizaciones y configuración segura. | Sprints 2 a 5 |
| 7 — CI/CD | Automatizar validaciones y entregas seguras. | Validaciones y proceso de despliegue documentado. | Sprint 6 |
| 8 — Servicios | Ampliar de forma controlada. | Nuevos servicios evaluados con los mismos criterios. | Sprints 3, 4 y 6 |

## Criterios transversales

Antes de cerrar un sprint, los cambios deben estar documentados, no contener secretos y actualizar estado, changelog, runbooks y ADR cuando corresponda. Sprint -1 no autoriza implementación técnica.
