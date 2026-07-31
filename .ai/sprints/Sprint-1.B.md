# Sprint 1.B — Bootstrap del Host (Implementación)

## Objetivo

Ejecutar el diseño del Sprint 1.A y entregar un script validado en el host que prepare Ubuntu Server LTS con Docker y los directorios persistentes de forma determinista e idempotente.

## Fase

B — Implementación

> En Fase B se ejecuta estrictamente lo aprobado en la Fase A (ver `Sprint-1-Plan.md`).

## Entregables

- [x] Script `bootstrap/bootstrap.sh` implementado según el contrato de diseño.
- [x] Documentación `docs/BOOTSTRAP.md` creada.
- [x] Script ejecutado exitosamente en el host `homelab`.
- [x] Evidencia de idempotencia (tres pasadas) recopilada.
- [x] Tareas HL-0013 y HL-0014 creadas y completadas.

## Archivos modificados

- `bootstrap/bootstrap.sh`
- `docs/BOOTSTRAP.md`
- `docs/evidence/Sprint-1.B-idempotency.log`
- `docs/evidence/Sprint-1.B-shellcheck.log`
- `docs/evidence/Sprint-1.B-bash-syntax.log`
- `.ai/sprints/Sprint-1.B.md`
- `.ai/tasks/HL-0013.md`, `HL-0014.md`

## Trazabilidad

| Tipo | Referencia |
|---|---|
| Plan | Sprint-1-Plan.md |
| Tasks | HL-0013, HL-0014 |
| Sprint anterior | Sprint 1.A |
| Sprint siguiente | Sprint 2.A |
| Phase | Phase 1 — Bootstrap |

## Deuda Técnica

**Evaluación de Motor de Convergencia (Ensure -> Validate -> Evidence):**
El script actual consolida la generación de logs, pero mezcla la ejecución con el registro de la evidencia en el mismo flujo estándar (`stdout`/`stderr`). Se ha evaluado que para alcanzar un modelo de *state-convergence* maduro, el EOS requeriría desacoplar estas fases arquitectónicamente de modo que el bootstrap (o una herramienta superior) genere la evidencia versionable de forma nativa por fase, separando el acto de asegurar (Ensure) del acto de certificar el estado (Validate -> Evidence). Esta mejora excede el alcance del Sprint 1.B y queda documentada aquí para ser abordada mediante un ADR futuro.

## Estado

Completed
