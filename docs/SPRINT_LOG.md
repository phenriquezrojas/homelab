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

## Sprint 0.A — Diseño del Bootstrap Base

- **Fecha:** 2026-07-24
- **Objetivo:** Definir e interconectar conceptualmente la estructura de persistencia, secretos y red.
- **Entregables:** Sprint Specification, Implementation Plan, Tasks (HL-0006, HL-0007, HL-0008).
- **Decisiones importantes:** Se preparan decisiones basadas en ADR 001, 006, 008.
- **Lecciones aprendidas:** La separación física Fase A/B garantiza diseño atómico.
- **Estado:** Approved

## Sprint 0.B — Implementación del Bootstrap Base

- **Fecha:** 2026-07-24
- **Objetivo:** Materializar decisiones de Fase A en servidor.
- **Infraestructura creada:** Directorios base `/srv/homelab` y permisos.
- **Documentación creada:** `docs/SECRETS.md`, `docs/NETWORK.md`.
- **Repositorio clonado:** En `homelab`.
- **Estado:** Approved

## Sprint 1.A — Bootstrap del Host (Diseño)

- **Fecha:** 2026-07-28
- **Objetivo:** Diseñar el contrato técnico que deberá cumplir cualquier implementación futura del bootstrap del host.
- **Entregables:** Sprint Specification, Implementation Plan, Tasks (HL-0009 a HL-0012) y ADR-009.
- **Decisiones importantes:** Se adoptó un enfoque mixto de idempotencia donde el gestor de paquetes (apt) administra el estado y el script valida el estado funcional esperado (ADR-009).
- **Estado:** Approved

## Sprint 1.B — Bootstrap del Host (Implementación)

- **Fecha:** 2026-07-28
- **Objetivo:** Ejecutar el diseño del Sprint 1.A y entregar un script validado en el host que prepare Ubuntu Server LTS con Docker y los directorios persistentes de forma determinista e idempotente.
- **Entregables:** Script `bootstrap/bootstrap.sh`, tareas HL-0013 y HL-0014 completadas, y evidencias de ejecución en el host.
- **Decisiones importantes:** Se implementó el script con 6 fases y 3 niveles de idempotencia exigidos.
- **Lecciones aprendidas:** Ejecutar vía SSH (SCP) un script idempotente permite validar la automatización remotamente sin necesidad de ensuciar el host con código no versionado o clonaciones prematuras. El script sobrevivió y demostró idempotencia en 3 ejecuciones consecutivas.
- **Estado:** Approved

## Sprint 2.A — Plataforma Central y Diseño del Runtime (Diseño)

- **Fecha:** 2026-08-05
- **Objetivo:** Diseñar el contrato técnico que Sprint 2.B deberá ejecutar para convertir el host bootstrapped en una plataforma con acceso privado (Tailscale), proxy inverso (Caddy), resolución DNS interna (`home.arpa`) y red Docker compartida. Además, debido al rechazo original de Sprint 2.B, diseñar el Homelab Runtime (v1.0) como motor de convergencia universal.
- **Entregables:** Sprint Specification, Implementation Plan, Tasks (HL-0015 a HL-0019) y la Especificación Técnica del Runtime v1.0 (`Sprint-2-Runtime-Design.md`).
- **Decisiones importantes:** Se definió Tailscale MagicDNS + Split DNS para resolución. Además, se estructuró la arquitectura del Runtime con fronteras estrictas de 6 subsistemas: Discovery, Observer, Planner, Transition Resolver, Executor y Reporter (conforme a ADR-010).
- **Lecciones aprendidas:** Ejecutar una prueba de estrés teórica sobre una arquitectura (Capítulo 10 del diseño) con capacidades reales antes de implementar código es la mejor forma de detectar fricciones tempranas sin generar deuda técnica.
- **Estado:** Approved

## Sprint 2.B — Plataforma Central (Implementación del Runtime Engine)

- **Fecha:** 2026-08-11
- **Objetivo:** Desarrollar y validar el Homelab Runtime Engine v1.0, y utilizarlo para desplegar el primer dominio de capacidades (Docker, Tailscale, Caddy y MagicDNS).
- **Entregables:** Código Go del Runtime (`runtime/`), scripts de componentes, registro `registry.yaml` y evidencias en `docs/evidence/`.
- **Decisiones importantes:** Implementar un motor puro de convergencia basado en DAG de dependencias (Topological Sort) y política Halt-on-Fail (ADR-010).
- **Estado:** Approved





