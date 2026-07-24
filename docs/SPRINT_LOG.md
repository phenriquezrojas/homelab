# Sprint Log

Registro histórico de sprints aprobados. Cada nueva entrada debe conservar este
formato para facilitar la trazabilidad del proyecto.

## Sprint -1 — Fundaciones documentales y organizacionales

- **Fecha:** 2026-07-22
- **Objetivo:** establecer la base documental, organizacional y de gobierno del
  Homelab sin implementar funcionalidad técnica.
- **Entregables:** README, visión, project charter, roadmap, estado actual,
  índice de decisiones, changelog, convenciones y ADR-001 a ADR-008 completos.
- **Decisiones importantes:** se documentaron como aceptadas Ubuntu Server LTS,
  Docker First, persistencia bajo `/srv`, `home.arpa`, Tailscale, Caddy,
  PostgreSQL, Redis, Restic + Backblaze B2 y un único repositorio.
- **Riesgos abiertos:** todavía no existe evidencia del host objetivo ni de su
  capacidad, red, almacenamiento o controles de acceso.
- **Deuda técnica:** inicializar Git, completar el texto de licencia MIT y
  definir inventario, secretos y estructura de `/srv` en Sprint 0.
- **Estado:** Approved

## Sprint -0.5 — Host Inventory

- **Fecha:** 2026-07-22
- **Objetivo:** incorporar una herramienta de inventario del host en modo solo
  lectura para recoger evidencia previa a Sprint 0.
- **Entregables:** `scripts/inventory.sh`, documentación de uso y reportes
  Markdown ignorados bajo `reports/`.
- **Decisiones importantes:** el inventario no instala dependencias, no cambia
  configuración y omite valores de variables de entorno para evitar secretos.
- **Riesgos abiertos:** la cobertura de hardware, SMART, firewall y servicios
  depende de comandos instalados y de los permisos de ejecución.
- **Deuda técnica:** ejecutar y revisar el inventario en el host Ubuntu
  definitivo; definir una política de dependencia para herramientas como
  `smartctl`; validar los hallazgos antes del bootstrap.
- **Estado:** Approved

## Sprint -0.5.1 — Pull Request Review Corrections

- **Fecha:** 2026-07-24
- **Objetivo:** incorporar las observaciones de revisión del inventario y dejar la base documental lista para Sprint 0 sin añadir capacidades de host.
- **Entregables:** advertencia de privilegios, clasificación de disponibilidad de comandos, resumen ejecutivo, nombre de reporte con hostname y actualización de la guía de inventario.
- **Decisiones importantes:** el reporte conserva el enfoque solo lectura y clasifica errores sin ejecutar expresiones o pipelines arbitrarios.
- **Riesgos abiertos:** los datos siguen siendo `Best Effort`; la cobertura depende de permisos y herramientas existentes en el host.
- **Deuda técnica:** ninguna.
- **Estado:** Approved

## Sprint -0.6 — AI Development Framework

- **Fecha:** 2026-07-24
- **Objetivo:** definir el marco agnóstico de modelo para trabajo asistido por agentes.
- **Entregables:** `.ai/context/AI_FRAMEWORK.md` definiendo directrices, jerarquía de autoridad y reglas de operación para agentes de IA.
- **Decisiones importantes:** el marco de desarrollo de IA debe ser agnóstico y basarse únicamente en estándares del repositorio.
- **Riesgos abiertos:** ninguno.
- **Deuda técnica:** ninguna.
- **Estado:** Approved

## Sprint -0.7 — Engineering Operating System (EOS)

- **Fecha:** 2026-07-22
- **Objetivo:** crear la capa de conocimiento permanente del repositorio (EOS) y regularizar la narrativa documental de sprints y estados.
- **Entregables:** Constitución del proyecto, contexto resumido, plan maestro por fases, tareas, sprints, ciclo de vida, roles y plantillas transversales bajo `.ai/`.
- **Decisiones importantes:** la Constitución en `.ai/PROJECT_CONSTITUTION.md` se establece como la máxima autoridad de gobierno técnico del repositorio.
- **Riesgos abiertos:** desalineación documental de tareas si no se mantiene la consistencia cruzada en los commits.
- **Deuda técnica:** resolver la revisión pendiente de inventario (Sprint -0.5.1) y el marco de IA (Sprint -0.6) en fases posteriores.
- **Estado:** Approved
