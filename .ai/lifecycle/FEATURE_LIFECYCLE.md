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
  ↓  (aprobación explícita del diseño)
Sprint x.B1 — Implementación Declarativa (Repositorio)
  ↓  (código y manifiestos versionados en Git)
Sprint x.B2 — Despliegue y Validación Operacional (Host)
  ↓  (evidencia empírica en docs/evidence/)
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

### Fase A — Diseño (Sprint x.A)

El humano y el agente colaboran para producir todos los artefactos necesarios. **No se escribe código, configuración ni se modifica infraestructura.**

Artefactos producidos:

- Sprint Specification (`.ai/sprints/Sprint-x.A.md` completado).
- Implementation Plan (enlazado desde la tarea).
- Tasks con checklist verificable (`.ai/tasks/`).
- Criterios de aceptación.
- Checklist de revisión.
- ADR (si la decisión es arquitectónica o duradera).

### Fase B — Implementación (Sprint x.B)

El agente ejecuta exactamente lo aprobado en la Fase A. Se estructura explícitamente en dos subfases secuenciales:

#### Subfase B1 — Implementación Declarativa (Repositorio)
- **Alcance:** Creación y versionado de código, manifiestos docker-compose, archivos de configuración y documentación canónica en Git.
- **Criterio de salida:** Finaliza cuando todos los artefactos del contrato de diseño están implementados y versionados en el repositorio.

#### Subfase B2 — Despliegue y Validación Operacional (Host)
- **Alcance:** Aplicación runtime en el servidor objetivo, ejecución de comandos de salud y verificación de funcionalidad.
- **Criterio de salida:** Finaliza cuando existen evidencias empíricas comprobables del despliegue registradas en el artefacto obligatorio de evidencia.

### Artefacto Obligatorio de Evidencia Operacional (`docs/evidence/`)

Toda implementación que incluya la subfase B2 debe generar un artefacto de evidencia versionado bajo la carpeta canónica `docs/evidence/` (ej. `docs/evidence/Sprint-x.B2-validation.md`). 

Este artefacto es **obligatorio para cerrar el Sprint** y debe registrar como mínimo:
- Fecha
- Sprint / Subfase
- Host objetivo
- Comandos ejecutados
- Resultado obtenido
- Evidencias relevantes (logs, salidas de comandos)
- Conclusión

Si durante la Fase B (B1 o B2) se descubre que se necesitan cambios de diseño, se retorna a la Fase A para actualizar los artefactos mediante el **Ciclo de Re-planificación (Replanning / Revision Cycle)** descrito a continuación.

## Replanning / Revision Cycle

Cuando se produce un hallazgo arquitectónico, o en general cualquier cambio material (nuevo requisito, bug de diseño, cambio de alcance/dependencia) que invalida el *baseline* actual, se activa el ciclo de re-planificación formal para preservar el contexto sin perder trazabilidad:

1. **Separación de Estados:** El ciclo distingue estrictamente entre `Document Status` (ej. `ACTIVE`, `SUPERSEDED`) y `Execution Status` (ej. `PLANNED`, `IN PROGRESS`, `ABORTED`, `COMPLETED`).
2. **Design Baseline:** Toda fase de implementación (B) debe declarar explícitamente en su metadata qué diseño y revisión la gobierna de manera inequívoca (ej. `Design Baseline: Artifact: Sprint-2-Runtime-Design.md, Revision: 1`).
3. **Supersession:** Todo documento superado pasa a `Document Status: SUPERSEDED`. Se preserva su valor histórico (indicando `Superseded By` y `Reason`, y opcionalmente `Design Role`).
4. **Revisiones:** Cuando un Sprint cambia o fracasa por un hallazgo que requiere rediseño, el intento antiguo conserva su `Execution Status: ABORTED`, pero su documento pasa a `Document Status: SUPERSEDED`. La revisión antigua se archiva físicamente en un directorio `history/` (ej. `history/Sprint-x.B-r1.md`). Se genera entonces una nueva revisión con `Document Status: ACTIVE` y `Execution Status: PLANNED`, apoyada en el nuevo `Design Baseline`.

## Reglas de transición

- Idea a discusión: definir problema, valor y límites.
- Discusión a ADR: crear ADR cuando la decisión sea arquitectónica o duradera.
- ADR a plan: enlazar decisiones, riesgos y reversión.
- Plan a Sprint x.A: crear Sprint Specification con criterios verificables.
- **Sprint x.A a Sprint x.B1 (Aprobación de Diseño):** Todos los artefactos de diseño deben estar aprobados explícitamente por el humano. Este es el Gate de Diseño.
- **Sprint x.B1 a Sprint x.B2 (Implementación Declarativa Completa):** Los artefactos de código, manifiestos y documentación están completos y versionados en Git.
- **Sprint x.B2 a Review (Validación Operacional Completa):** La solución fue desplegada en el host y se generó el artefacto obligatorio de evidencia empírica en `docs/evidence/`.
- Review a merge: resolver observaciones o registrarlas como deuda.
- Operación a evolución: registrar evidencia que justifique el cambio siguiente.

Plantillas: [ADR](../templates/ADR.md), [plan](../templates/IMPLEMENTATION_PLAN.md), [sprint](../templates/SPRINT.md), [task](../templates/TASK.md), [review](../templates/REVIEW.md) y [runbook](../templates/RUNBOOK.md).
