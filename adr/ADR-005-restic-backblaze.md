# ADR-005 — Restic + Backblaze B2 para respaldos

## Estado

Aceptada.

## Fecha

2026-07-22.

## Contexto

Los datos persistentes requieren una copia externa, cifrada y verificable para recuperarse ante pérdida o corrupción del host.

## Problema

Los respaldos locales no cubren pérdida física del equipo y los procedimientos manuales no proporcionan confianza de recuperación.

## Decisión

Restic creará y gestionará respaldos cifrados con Backblaze B2 como destino externo. Retención, credenciales, alcance y pruebas se definirán en Sprint 3.

## Alternativas consideradas

- Copias manuales a discos locales: dependen de operación manual y no dan copia externa consistente.
- Rsync como mecanismo principal: no aporta por sí solo snapshots cifrados.
- Otros proveedores: se mantiene la decisión aprobada de Backblaze B2.

## Consecuencias

- Las credenciales vivirán fuera del repositorio y se controlarán costes.
- Los runbooks tratarán respaldo y restauración como un flujo completo.

## Riesgos

- Un respaldo sin restauración probada puede fallar cuando se necesite.
- Pérdida de claves o credenciales impide recuperación.
- Retención mal configurada puede aumentar costes.

## Referencias relacionadas

- [ADR-001 — Persistencia bajo `/srv`](ADR-001-persistence-under-srv.md)
- [Estado actual](../CURRENT_STATE.md)
- [Roadmap](../ROADMAP.md)
