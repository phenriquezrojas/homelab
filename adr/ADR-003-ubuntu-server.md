# ADR-003 — Ubuntu Server LTS como sistema operativo base

## Estado

Aceptada.

## Fecha

2026-07-22.

## Contexto

La plataforma requiere un sistema operativo de servidor con soporte previsible, documentación amplia y compatibilidad con Docker.

## Problema

Un sistema sin soporte prolongado incrementa el riesgo operativo y el coste de mantenimiento.

## Decisión

El host objetivo utilizará Ubuntu Server en una versión LTS soportada. La versión exacta y calendario de actualización se registrarán en Sprint 0 y Bootstrap.

## Alternativas consideradas

- Ubuntu no LTS: ciclo de soporte más corto.
- Debian: se mantiene la línea base ya aprobada.
- Distribuciones inmutables: introducen un modelo operativo adicional sin necesidad actual.

## Consecuencias

- Los procedimientos se escribirán para Ubuntu Server LTS.
- Las actualizaciones de seguridad y fin de soporte deben seguirse en operación.

## Riesgos

- Posponer actualizaciones puede dejar componentes sin soporte.
- Las diferencias entre versiones LTS deben validarse antes de actualizar.

## Referencias relacionadas

- [ADR-002 — Docker First](ADR-002-docker-first.md)
- [Project Charter](../PROJECT_CHARTER.md)
- [Roadmap](../ROADMAP.md)
