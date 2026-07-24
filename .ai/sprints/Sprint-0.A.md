# Sprint 0.A — Diseño del Bootstrap Base

## Objetivo

Definir e interconectar conceptualmente la estructura de persistencia, el modelo de gestión de secretos y la topología de red para preparar el host.

## Fase

A — Diseño

> En Fase A se generan los artefactos de diseño; en Fase B se ejecuta lo aprobado.
> Ver [FEATURE_LIFECYCLE.md](../lifecycle/FEATURE_LIFECYCLE.md) para reglas de transición.

## Alcance

- Diseñar la persistencia bajo `/srv`.
- Definir formato y ubicación de archivos con variables de entorno (secretos).
- Establecer enrutamiento y nombres internos.
- **Excluido:** Modificación del host, Docker Compose, o ejecución de scripts.

## Entregables

### Fase A — Diseño

- [x] Sprint Specification completada (este documento).
- [x] Implementation Plan aprobado (`Sprint-0-Plan.md`).
- [x] Tasks con checklist verificable creadas (HL-0006, HL-0007, HL-0008).
- [x] Criterios de aceptación definidos en las tareas.
- [x] Checklist de revisión preparado en el plan.
- [x] ADR creado (si aplica) - No se requiere ADR nuevo, se basan en ADR 001, 006, 008.

### Fase B — Implementación (A ejecutar en Sprint 0.B)

- [ ] Creada la estructura `/srv` con permisos seguros.
- [ ] Documentación paralela actualizada (si aplica).
- [ ] Evidencia de validación de exclusión en `.gitignore`.

## Archivos modificados

- `.ai/sprints/Sprint-0.A.md`
- `.ai/implementation/Sprint-0-Plan.md`
- `.ai/tasks/HL-0006.md`, `HL-0007.md`, `HL-0008.md`

## Estado

Completed

## Lecciones aprendidas

La separación física entre la planificación de la Fase A y la ejecución técnica de la Fase B garantiza un diseño atómico sin contaminación cruzada ni asunciones falsas.

## Deuda técnica

Ninguna.

## Próximo Sprint

Sprint 0.B (Implementación)
