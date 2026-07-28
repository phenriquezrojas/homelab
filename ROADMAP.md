# Hoja de ruta

La hoja de ruta se organiza por resultados y dependencias, no por fechas. Un sprint termina cuando sus criterios de salida se cumplen y su documentación queda actualizada.

A partir del Sprint 0, cada Sprint se divide en dos fases: **Fase A (Diseño)** y **Fase B (Implementación)**. La tabla refleja esta separación con sub-filas por fase.

| Sprint | Objetivo | Resultado esperado | Dependencias | Estado |
| --- | --- | --- | --- | --- |
| -1 — Fundaciones | Definir base documental y de gobierno. | Charter, visión, ADR, convenciones, estado y roadmap completos. | Ninguna | Completed |
| -0.5 — Host Inventory | Recoger evidencia del host. | Script de inventario `scripts/inventory.sh` y reporte local en Markdown. | Sprint -1 | Completed |
| -0.5.1 — PR Corrections | Atender observaciones del PR de inventario. | Corrección de privilegios y clasificación de comandos en script. | Sprint -0.5 | Completed |
| -0.6 — AI Framework | Definir marco de desarrollo para agentes de IA. | Criterios y directrices de desarrollo asistido por agentes (`.ai/context/AI_FRAMEWORK.md`). | Sprint -1 | Completed |
| -0.7 — EOS | Crear el Engineering Operating System (EOS). | Constitución, plan maestro, tareas y plantillas transversales en `.ai/`. | Sprints -1, -0.5.1, -0.6 | Completed |
| 0.A — Preparación (Diseño) | Diseñar modelo de secretos, plan de red y estructura de `/srv`. | Sprint Specification, Implementation Plan, Tasks y criterios de aceptación aprobados. | Sprints -1, -0.7 | Completed |
| 0.B — Preparación (Implementación) | Ejecutar y validar los artefactos de diseño del Sprint 0. | Modelo de secretos, plan de red y estructura de `/srv` aprobados. | Sprint 0.A | Completed |
| 1.A — Bootstrap (Diseño) | Diseñar proceso reproducible de base del host. | Sprint Specification, Implementation Plan, Tasks y criterios de aceptación aprobados. | Sprint 0.B | Planned |
| 1.B — Bootstrap (Implementación) | Reproducir base del host. | Proceso validado para Ubuntu Server LTS, Docker y directorios persistentes. | Sprint 1.A | Planned |
| 2.A — Plataforma central (Diseño) | Diseñar componentes compartidos. | Sprint Specification, Implementation Plan, Tasks y criterios de aceptación aprobados. | Sprint 1.B | Planned |
| 2.B — Plataforma central (Implementación) | Establecer componentes compartidos. | Tailscale, Caddy, `home.arpa` y dependencias de datos operables. | Sprint 2.A | Planned |
| 3.A — Respaldo (Diseño) | Diseñar protección y recuperación de datos. | Sprint Specification, Implementation Plan, Tasks y criterios de aceptación aprobados. | Sprint 2.B | Planned |
| 3.B — Respaldo (Implementación) | Proteger y recuperar datos. | Restic a Backblaze B2 y restauración probada. | Sprint 3.A | Planned |
| 4.A — Monitorización (Diseño) | Diseñar observabilidad del sistema. | Sprint Specification, Implementation Plan, Tasks y criterios de aceptación aprobados. | Sprints 2.B y 3.B | Planned |
| 4.B — Monitorización (Implementación) | Obtener visibilidad. | Métricas, logs, alertas y runbooks. | Sprint 4.A | Planned |
| 5.A — Immich (Diseño) | Diseñar incorporación del primer servicio priorizado. | Sprint Specification, Implementation Plan, Tasks y criterios de aceptación aprobados. | Sprints 2.B, 3.B y 4.B | Planned |
| 5.B — Immich (Implementación) | Incorporar el primer servicio priorizado. | Servicio de fotos personal documentado, respaldado, monitorizado y recuperable. | Sprint 5.A | Planned |
| 6.A — Hardening (Diseño) | Diseñar reducción de riesgos. | Sprint Specification, Implementation Plan, Tasks y criterios de aceptación aprobados. | Sprints 2.B a 5.B | Planned |
| 6.B — Hardening (Implementación) | Reducir riesgos. | Revisión de acceso, secretos, actualizaciones y configuración segura. | Sprint 6.A | Planned |
| 7.A — CI/CD (Diseño) | Diseñar automatización de validaciones y entregas. | Sprint Specification, Implementation Plan, Tasks y criterios de aceptación aprobados. | Sprint 6.B | Planned |
| 7.B — CI/CD (Implementación) | Automatizar validaciones y entregas seguras. | Validaciones y proceso de despliegue automatizado y documentado. | Sprint 7.A | Planned |
| 8.A — Servicios (Diseño) | Diseñar ampliación controlada de servicios. | Sprint Specification, Implementation Plan, Tasks y criterios de aceptación aprobados. | Sprints 3.B, 4.B y 6.B | Planned |
| 8.B — Servicios (Implementación) | Ampliar de forma controlada. | Nuevos servicios evaluados con los mismos criterios de calidad. | Sprint 8.A | Planned |

## Criterios transversales

Antes de cerrar un sprint, los cambios deben estar documentados, no contener secretos y actualizar estado, changelog, runbooks y ADR cuando corresponda. Los sprints de la fase 0 no autorizan implementación de infraestructura.

A partir del Sprint 0, cada sprint se divide en Fase A (Diseño) y Fase B (Implementación). La Fase A produce los artefactos de diseño (Sprint Specification, Implementation Plan, Tasks y criterios de aceptación). La Fase B ejecuta exactamente esos artefactos. No se inicia la Fase B sin aprobación explícita de la Fase A.
