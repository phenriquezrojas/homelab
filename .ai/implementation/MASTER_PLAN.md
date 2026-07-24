# Master Plan

Documento vivo de ejecución técnica. Se organiza por fases, no por sprints. Las tareas y sprints enlazados dan trazabilidad; el estado real está en las tareas.

## Phase 0 — Foundation

- **Objetivo:** establecer gobierno, trazabilidad y contexto suficiente para continuar sin historial conversacional.
- **Descripción:** documentación base, ADR, inventario de host y EOS.
- **Dependencias:** ninguna.
- **Estado:** Completed.
- **Riesgos:** estado documental puede divergir si no se actualizan índices y tareas.
- **Checklist:**
  - [x] Charter, visión, convenciones y ADR aprobados.
  - [x] Inventario de host de solo lectura disponible.
  - [x] EOS con contexto, tareas, sprints, plantillas y flujo de revisión.
- **Sprints relacionados:** -1, -0.5, -0.5.1, -0.6, -0.7.
- **ADR relacionados:** ADR-001 a ADR-008.
- **Tareas relacionadas:** HL-0001 a HL-0005.

## Phase 1 — Bootstrap

- **Objetivo:** preparar de forma reproducible el host Ubuntu Server LTS.
- **Descripción:** definir y ejecutar el proceso aprobado para host, Docker y estructura inicial de `/srv`.
- **Dependencias:** Phase 0; inventario validado y modelo de secretos definido.
- **Estado:** Planned.
- **Riesgos:** diferencias entre host inventariado y host definitivo; cambios manuales no documentados.
- **Checklist:**
  - [ ] Validar host y versión LTS.
  - [ ] Definir estructura y permisos de `/srv`.
  - [ ] Aprobar gestión de secretos.
- **Sprints relacionados:** Sprint 0 y Sprint 1.
- **ADR relacionados:** ADR-001, ADR-002, ADR-003.
- **Tareas relacionadas:** por planificar.

## Phase 2 — Storage

- **Objetivo:** definir almacenamiento persistente y su operación.
- **Descripción:** capacidad, filesystem, permisos y propiedad de datos bajo `/srv`.
- **Dependencias:** Phase 1.
- **Estado:** Planned.
- **Riesgos:** capacidad insuficiente, permisos erróneos y dependencia de un único disco.
- **Checklist:**
  - [ ] Diseñar layout de persistencia.
  - [ ] Documentar capacidad y crecimiento.
  - [ ] Definir validaciones de filesystem.
- **Sprints relacionados:** Sprint 0 y Sprint 1.
- **ADR relacionados:** ADR-001.
- **Tareas relacionadas:** por planificar.

## Phase 3 — Networking

- **Objetivo:** habilitar conectividad privada y nombres internos coherentes.
- **Descripción:** Tailscale, DNS interno `home.arpa` y entrada HTTP(S) con Caddy.
- **Dependencias:** Phases 1 y 2.
- **Estado:** Planned.
- **Riesgos:** DNS incompleto, ACL excesivas y exposición accidental.
- **Checklist:**
  - [ ] Diseñar DNS y subdominios.
  - [ ] Definir acceso Tailscale y revocación.
  - [ ] Diseñar integración Caddy.
- **Sprints relacionados:** Sprint 2.
- **ADR relacionados:** ADR-006, ADR-007, ADR-008.
- **Tareas relacionadas:** por planificar.

## Phase 4 — Identity & Security

- **Objetivo:** aplicar identidad, secretos, acceso mínimo y endurecimiento.
- **Descripción:** modelo de secretos, SSH, actualizaciones, permisos y revisión de exposición.
- **Dependencias:** Phases 1 y 3.
- **Estado:** Planned.
- **Riesgos:** secretos fuera de control, privilegios excesivos y host sin actualizar.
- **Checklist:**
  - [ ] Aprobar modelo de secretos.
  - [ ] Revisar acceso administrativo.
  - [ ] Documentar hardening y reversión.
- **Sprints relacionados:** Sprint 0 y Sprint 6.
- **ADR relacionados:** ADR-003, ADR-006.
- **Tareas relacionadas:** por planificar.

## Phase 5 — Platform Services

- **Objetivo:** habilitar dependencias compartidas para servicios.
- **Descripción:** proxy, PostgreSQL y Redis bajo las decisiones aprobadas.
- **Dependencias:** Phases 1 a 4.
- **Estado:** Planned.
- **Riesgos:** dependencias compartidas sin operación ni recuperación definida.
- **Checklist:**
  - [ ] Definir contratos de red y persistencia.
  - [ ] Documentar operación de datos compartidos.
  - [ ] Añadir pruebas y runbooks.
- **Sprints relacionados:** Sprint 2.
- **ADR relacionados:** ADR-002, ADR-007.
- **Tareas relacionadas:** por planificar.

## Phase 6 — Applications

- **Objetivo:** incorporar servicios de usuario con criterios operativos completos.
- **Descripción:** evaluación, despliegue y operación de aplicaciones; Immich es el primer servicio priorizado.
- **Dependencias:** Phases 2 a 5 y capacidad de recuperación validada.
- **Estado:** Planned.
- **Riesgos:** añadir aplicaciones sin respaldo, monitorización o runbook.
- **Checklist:**
  - [ ] Evaluar servicio contra principios de Constitución.
  - [ ] Definir datos, acceso y recuperación.
  - [ ] Completar revisión técnica.
- **Sprints relacionados:** Sprint 5 y Sprint 8.
- **ADR relacionados:** ADR-001, ADR-002, ADR-005, ADR-007.
- **Tareas relacionadas:** por planificar.

## Phase 7 — Backup & Disaster Recovery

- **Objetivo:** proteger y restaurar datos y configuración.
- **Descripción:** Restic hacia Backblaze B2, retención, cifrado y ejercicios de restauración.
- **Dependencias:** Phases 1, 2 y 5.
- **Estado:** Planned.
- **Riesgos:** respaldos no verificables, pérdida de claves y costes no controlados.
- **Checklist:**
  - [ ] Definir alcance y retención.
  - [ ] Configurar secretos fuera de Git.
  - [ ] Probar restauración documentada.
- **Sprints relacionados:** Sprint 3.
- **ADR relacionados:** ADR-005.
- **Tareas relacionadas:** por planificar.

## Phase 8 — Operations

- **Objetivo:** operar la plataforma con visibilidad, procedimientos y evolución controlada.
- **Descripción:** monitorización, runbooks, mantenimiento, CI/CD y gestión de cambios.
- **Dependencias:** Phases 3 a 7.
- **Estado:** Planned.
- **Riesgos:** operación reactiva sin evidencia ni procedimientos.
- **Checklist:**
  - [ ] Definir observabilidad y alertas.
  - [ ] Completar runbooks.
  - [ ] Establecer validaciones de repositorio.
- **Sprints relacionados:** Sprint 4, Sprint 6 y Sprint 7.
- **ADR relacionados:** ADR-002 a ADR-008 según alcance.
- **Tareas relacionadas:** por planificar.
