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
Sprint
  ↓
Implementación
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

## Reglas de transición

- Idea a discusión: definir problema, valor y límites.
- Discusión a ADR: crear ADR cuando la decisión sea arquitectónica o duradera.
- ADR a plan: enlazar decisiones, riesgos y reversión.
- Plan a sprint: crear tarea con criterios verificables.
- Implementación a review: actualizar documentación, pruebas y runbooks aplicables.
- Review a merge: resolver observaciones o registrarlas como deuda.
- Operación a evolución: registrar evidencia que justifique el cambio siguiente.

Plantillas: [ADR](../templates/ADR.md), [plan](../templates/IMPLEMENTATION_PLAN.md), [task](../templates/TASK.md), [review](../templates/REVIEW.md) y [runbook](../templates/RUNBOOK.md).
