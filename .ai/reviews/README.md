# Technical Review Flow

Toda revisión comprueba primero la Constitución, los ADR aplicables, la tarea y la fase del Master Plan. Después revisa alcance, seguridad, reversión, documentación, pruebas y deuda.

A partir del Sprint 0, el flujo de revisión distingue dos puntos de control correspondientes a las fases del Sprint bifásico:

## Gate Review (Fase A → Fase B)

Revisión de los artefactos de diseño antes de autorizar la implementación. Verifica que:

1. La Sprint Specification tiene objetivo verificable, alcance definido y entregables claros.
2. El Implementation Plan enlaza ADR, riesgos y estrategia de reversión.
3. Las Tasks tienen criterios de aceptación verificables y checklist completo.
4. No hay ambigüedades que obligarían a improvisar durante la Fase B.
5. El alcance no excede los límites del sprint ni adelanta infraestructura.

**Decisión:** Approved (se inicia Fase B) | Changes Requested (se itera en Fase A).

## Completion Review (Fase B → Merge)

Revisión de la implementación contra los criterios definidos en la Fase A. Sigue el flujo estándar:

1. Identificar tarea, sprint, fase y ADR relacionados.
2. Comprobar que el cambio se limita al alcance aprobado en la Fase A.
3. Validar enlaces, convenciones y ausencia de secretos.
4. Ejecutar validaciones proporcionales al riesgo.
5. Clasificar hallazgos como Critical, Major o Minor.
6. Resolver, aceptar explícitamente o registrar deuda antes de aprobar.

**Decisión:** Approved | Approved with debt | Changes requested.

Use la [plantilla de review](../templates/REVIEW.md). Las observaciones históricas pertenecen al sprint o tarea que las originó; no se reescribe el registro de decisiones.
