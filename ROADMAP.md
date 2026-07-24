# Hoja de ruta

La hoja de ruta se organiza por resultados y dependencias, no por fechas. Un sprint termina cuando sus criterios de salida se cumplen y su documentación queda actualizada.

| Sprint | Objetivo | Resultado esperado | Dependencias | Estado |
| --- | --- | --- | --- | --- |
| -1 — Fundaciones | Definir base documental y de gobierno. | Charter, visión, ADR, convenciones, estado y roadmap completos. | Ninguna | Completed |
| -0.5 — Host Inventory | Recoger evidencia del host. | Script de inventario `scripts/inventory.sh` y reporte local en Markdown. | Sprint -1 | Completed |
| -0.5.1 — PR Corrections | Atender observaciones del PR de inventario. | Corrección de privilegios y clasificación de comandos en script. | Sprint -0.5 | Completed |
| -0.6 — AI Framework | Definir marco de desarrollo para agentes de IA. | Criterios y directrices de desarrollo asistido por agentes (`.ai/context/AI_FRAMEWORK.md`). | Sprint -1 | Completed |
| -0.7 — EOS | Crear el Engineering Operating System (EOS). | Constitución, plan maestro, tareas y plantillas transversales en `.ai/`. | Sprints -1, -0.5.1, -0.6 | Completed |
| 0 — Preparación | Planificar host y repositorio. | Modelo de secretos, plan de red y estructura de `/srv` aprobados. | Sprints -1, -0.7 | Planned |
| 1 — Bootstrap | Reproducir base del host. | Proceso validado para Ubuntu Server LTS, Docker y directorios persistentes. | Sprint 0 | Planned |
| 2 — Plataforma central | Establecer componentes compartidos. | Tailscale, Caddy, `home.arpa` y dependencias de datos operables. | Sprint 1 | Planned |
| 3 — Respaldo | Proteger y recuperar datos. | Restic a Backblaze B2 y restauración probada. | Sprint 2 | Planned |
| 4 — Monitorización | Obtener visibilidad. | Métricas, logs, alertas y runbooks. | Sprints 2 y 3 | Planned |
| 5 — Immich | Incorporar el primer servicio priorizado. | Servicio de fotos personal documentado, respaldado, monitorizado y recuperable. | Sprints 2, 3 y 4 | Planned |
| 6 — Hardening | Reducir riesgos. | Revisión de acceso, secretos, actualizaciones y configuración segura. | Sprints 2 a 5 | Planned |
| 7 — CI/CD | Automatizar validaciones y entregas seguras. | Validaciones y proceso de despliegue automatizado y documentado. | Sprint 6 | Planned |
| 8 — Servicios | Ampliar de forma controlada. | Nuevos servicios evaluados con los mismos criterios de calidad. | Sprints 3, 4 y 6 | Planned |

## Criterios transversales

Antes de cerrar un sprint, los cambios deben estar documentados, no contener secretos y actualizar estado, changelog, runbooks y ADR cuando corresponda. Los sprints de la fase 0 no autorizan implementación de infraestructura.
