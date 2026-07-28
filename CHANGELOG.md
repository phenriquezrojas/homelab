# Changelog

Todos los cambios relevantes se documentan siguiendo [Keep a Changelog](https://keepachangelog.com/es-ES/1.1.0/) y versionado semántico.

## [Unreleased]

### Added
- Cierre del Sprint 1.A (Diseño). Generados artefactos de diseño para el bootstrap del host, incluyendo el contrato de idempotencia (ADR-009) y las tareas de validación (HL-0009 a HL-0012).
- Refactorización del EOS (Sprint 0.C). Creación del `STATE_INDEX.md` para abstraer los documentos de estado global. Modificación del hook `pre-commit` para validación cruzada de tareas y actualización de Skills.
- Cierre del Sprint 0.B (Implementación). Ejecutada la estructura física en el host `/srv/homelab` y documentados los mecanismos de red (`docs/NETWORK.md`) y secretos (`docs/SECRETS.md`).
- Cierre del Sprint 0.A (Diseño). Generados artefactos de diseño para la persistencia, secretos y red (tareas [HL-0006](.ai/tasks/HL-0006.md), [HL-0007](.ai/tasks/HL-0007.md), [HL-0008](.ai/tasks/HL-0008.md) y plan `.ai/implementation/Sprint-0-Plan.md`).

### Added
- Metodología de Sprint bifásico (Fase A: Diseño / Fase B: Implementación) incorporada al EOS del proyecto. Aplica desde el Sprint 0.
- Principio permanente nº 13 en la Constitución del proyecto estableciendo la ejecución bifásica obligatoria.
- Actualización de `FEATURE_LIFECYCLE.md`, `AI_FRAMEWORK.md`, plantilla de Sprint, `ROADMAP.md` y flujo de revisiones para reflejar el ciclo bifásico.

## [0.4.0] - 2026-07-24

### Added
- Creación de la especificación formal del AI Development Framework (`.ai/context/AI_FRAMEWORK.md`) para el Sprint -0.6 (tarea [HL-0004](.ai/tasks/HL-0004.md)).

### Changed
- Cierre formal del Sprint -0.5.1 (tarea [HL-0003](.ai/tasks/HL-0003.md)) tras validar y conciliar la metadata de trazabilidad del script de inventario.
- Actualización de los índices globales (`CURRENT_STATE.md`, `ROADMAP.md`, `docs/SPRINT_LOG.md`) marcando la Phase 0 completamente funcional a nivel documental.

## [0.3.0] - 2026-07-22

### Added
- Creación de la estructura del Engineering Operating System (EOS) bajo la carpeta `.ai/`.
- Constitución del proyecto (`.ai/PROJECT_CONSTITUTION.md`) que establece principios operativos máximos.
- Documentos de contexto, plan maestro (MASTER_PLAN.md), ciclo de vida, roles (skills), revisiones y plantillas en `.ai/`.
- Tareas `HL-0001` a `HL-0005` y sprints del EOS (`.ai/sprints/`).

### Changed
- Regularización de toda la documentación raíz (`README.md`, `CURRENT_STATE.md`, `ROADMAP.md`) y operativa (`docs/SPRINT_LOG.md`, `docs/CONVENTIONS.md`) para reflejar los estados y dependencias reales del proyecto.

## [0.2.0] - 2026-07-22

### Added
- Script de recopilación de inventario del host Ubuntu Server (`scripts/inventory.sh`).
- Guía operativa de inventario (`docs/HOST_INVENTORY.md`).

## [0.1.0] - 2026-07-22

### Added
- Estructura inicial del repositorio Homelab.
- Visión, charter, hoja de ruta e índice de decisiones iniciales (Sprint -1).
- Decisiones arquitectónicas registradas de ADR-001 a ADR-008.
- Convenciones de contribución, versionado semántico, nombres y estilo Markdown.
- Reglas de exclusión para archivos locales, secretos y reportes.
