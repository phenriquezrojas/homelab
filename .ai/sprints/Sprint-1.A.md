# Sprint 1.A — Bootstrap del Host (Diseño)

## Objetivo

Diseñar el contrato técnico que deberá cumplir cualquier implementación futura del bootstrap del host. El contrato define las propiedades, validaciones y evidencias que Sprint 1.B deberá satisfacer para dar el bootstrap por completado.

## Fase

A — Diseño

> En Fase A se generan los artefactos de diseño; en Fase B se ejecuta lo aprobado.
> Ver [FEATURE_LIFECYCLE.md](../lifecycle/FEATURE_LIFECYCLE.md) para reglas de transición.

## Principios del Bootstrap

Toda implementación del bootstrap deberá cumplir estos principios. Si el script viola alguno, no cumple el contrato del Sprint.

1. **Determinista.** La misma entrada produce el mismo resultado.
2. **Idempotente.** Puede ejecutarse múltiples veces sin efectos secundarios (ver [ADR-009](../../adr/ADR-009-bootstrap-idempotency.md)).
3. **Observable.** Cada paso genera log con timestamp y nivel (`INFO`/`OK`/`WARN`/`ERROR`).
4. **Reversible cuando sea posible.** Cada acción documenta su reversión explícita.
5. **Seguro frente a reejecuciones.** Una segunda ejecución no modifica nada si el estado objetivo ya fue alcanzado.
6. **Sin intervención manual durante la ejecución.** El script no solicita input interactivo.

## Alcance

### Incluido

- Diseño del script `bootstrap/bootstrap.sh` (estructura, flujo y contratos, sin código ejecutable).
- Diseño de la estrategia de idempotencia (aprobada como ADR-009).
- Diseño del sistema de logging del script.
- Diseño de la estrategia de rollback.
- Definición de prerrequisitos verificables del host.
- Definición de criterios de aceptación medibles.
- Definición de evidencias esperadas, incluyendo prueba de idempotencia.
- Identificación de riesgos y mitigaciones.
- Creación de Tasks HL-0009 a HL-0012.

### Excluido

- Escribir código ejecutable en ningún script.
- Modificar el servidor.
- Instalar Docker, Docker Compose o cualquier otro software.
- Diseñar Caddy, Tailscale, servicios Docker o backups.
- Diseñar monitorización o CI/CD.
- Configurar redes (`home.arpa`, DNS).

### No Objetivos

- No optimizar tiempos de ejecución del bootstrap.
- No instalar software en ningún host.
- No endurecer la seguridad del sistema operativo (corresponde a Sprint 6).
- No configurar usuarios adicionales ni claves SSH.
- No desplegar servicios ni contenedores de aplicación.
- No diseñar pipelines de CI/CD para el bootstrap.

## Entregables

### Fase A — Diseño

- [x] Sprint Specification completada (este documento).
- [x] Implementation Plan aprobado (`Sprint-1-Plan.md`).
- [x] Tasks con checklist verificable creadas (HL-0009, HL-0010, HL-0011, HL-0012).
- [x] Criterios de aceptación definidos en cada tarea.
- [x] Checklist de revisión preparado en el plan.
- [x] ADR-009 creado para la estrategia de idempotencia del bootstrap.
- [x] Evidencias de idempotencia definidas como criterio de aceptación para Sprint 1.B.

### Fase B

La Fase B ejecutará este contrato. Los entregables de implementación se documentarán en `Sprint-1.B.md`.

## Evidencia de idempotencia requerida para Sprint 1.B

Sprint 1.B deberá demostrar las siguientes ejecuciones secuenciales:

| Ejecución | Resultado esperado |
|---|---|
| Primera ejecución (host limpio) | Instala todo, exit code 0, log completo |
| Segunda ejecución (inmediata) | No modifica nada, exit code 0, log indica "ya instalado" |
| Tercera ejecución (tras reinicio) | No modifica nada, exit code 0 |

## Criterios de cierre del Sprint 1.A

Este Sprint solo podrá cerrarse cuando:

- [x] `Sprint-1-Plan.md` en estado `Approved`.
- [x] Todas las Tasks (HL-0009 a HL-0012) en estado `Approved`.
- [x] Ninguna pregunta abierta pendiente.
- [x] Todos los riesgos documentados en el Plan.
- [x] Evidencias de idempotencia definidas.
- [x] Rollback definido para cada fase del bootstrap.
- [x] Criterios de aceptación completos y verificables.
- [x] ADR-009 aprobado.

## Archivos modificados

- `.ai/sprints/Sprint-1.A.md`
- `.ai/implementation/Sprint-1-Plan.md`
- `.ai/tasks/HL-0009.md`, `HL-0010.md`, `HL-0011.md`, `HL-0012.md`
- `adr/ADR-009-bootstrap-idempotency.md`

## Trazabilidad

| Tipo | Referencia |
|---|---|
| ADR | ADR-001, ADR-002, ADR-003, ADR-009 |
| Tasks | HL-0009, HL-0010, HL-0011, HL-0012 |
| Sprint anterior | Sprint 0.B |
| Sprint siguiente | Sprint 1.B |
| Phase | Phase 1 — Bootstrap |

## Estado

Approved

## Lecciones aprendidas

Del Sprint 0: La separación A/B es efectiva. El diseño detallado en Fase A permite que Fase B opere sin ambigüedad. El bootstrap manual del Sprint 0 demostró que la estructura `/srv` correcta es condición necesaria y debe garantizarse por el propio script, no asumirse como prerrequisito externo.

## Deuda técnica

Ninguna generada en esta fase.

## Próximo Sprint

Sprint 1.B (Implementación del Bootstrap)
