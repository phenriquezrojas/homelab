# Feature Lifecycle

Toda capacidad o cambio duradero sigue este ciclo. Se puede detener en cualquier etapa si la información no permite avanzar sin riesgo.

```text
Idea
  ↓
Discusión
  ↓
ADR (si cambia arquitectura)
  ↓
Plan
  ↓
Sprint x.A — Diseño
  ↓  (aprobación explícita de artefactos)
Sprint x.B — Implementación
  ↓
Review
  ↓
Merge
  ↓
Operación
  ↓
Mantenimiento
  ↓
Evolución
```

## Fases del Sprint

A partir del Sprint 0, todo sprint se ejecuta en dos fases secuenciales:

### Fase A — Diseño

El humano y el agente colaboran para producir todos los artefactos necesarios. **No se escribe código, configuración ni se modifica infraestructura.**

Artefactos producidos:

- Sprint Specification (`.ai/sprints/Sprint-x.md` completado).
- Implementation Plan (enlazado desde la tarea).
- Tasks con checklist verificable (`.ai/tasks/`).
- Criterios de aceptación.
- Checklist de revisión.
- ADR (si la decisión es arquitectónica o duradera).

### Fase B — Implementación

El agente ejecuta exactamente lo aprobado en la Fase A. Al finalizar se revisa el resultado contra los criterios definidos y, si cumple, se cierra el Sprint.

Artefactos producidos:

- Código, configuración o infraestructura.
- Documentación paralela actualizada.
- Evidencia de validación y pruebas.

Si durante la Fase B se descubre que se necesitan cambios de diseño, se retorna a la Fase A para actualizar los artefactos antes de continuar.

## Reglas de transición

- Idea a discusión: definir problema, valor y límites.
- Discusión a ADR: crear ADR cuando la decisión sea arquitectónica o duradera.
- ADR a plan: enlazar decisiones, riesgos y reversión.
- Plan a Sprint x.A: crear Sprint Specification con criterios verificables.
- **Sprint x.A a Sprint x.B: todos los artefactos de diseño deben estar aprobados explícitamente por el humano. Este es el gate de diseño.**
- Sprint x.B a review: implementación completada según los artefactos de la Fase A; documentación, pruebas y runbooks actualizados.
- Review a merge: resolver observaciones o registrarlas como deuda.
- Operación a evolución: registrar evidencia que justifique el cambio siguiente.

Plantillas: [ADR](../templates/ADR.md), [plan](../templates/IMPLEMENTATION_PLAN.md), [sprint](../templates/SPRINT.md), [task](../templates/TASK.md), [review](../templates/REVIEW.md) y [runbook](../templates/RUNBOOK.md).
