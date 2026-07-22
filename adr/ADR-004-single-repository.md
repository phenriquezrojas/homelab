# ADR-004 — Un único repositorio

## Estado

Aceptada.

## Fecha

2026-07-22.

## Contexto

Infraestructura, procedimientos, decisiones y configuraciones comparten dependencias y deben evolucionar con trazabilidad.

## Problema

Separar prematuramente documentación, configuraciones y runbooks fragmenta el contexto y dificulta cambios coherentes.

## Decisión

Este repositorio será la fuente de verdad para documentación, ADR, configuraciones declarativas, automatizaciones y pruebas. Secretos y datos operativos quedan fuera de Git.

## Alternativas consideradas

- Un repositorio por servicio: sobrecarga de coordinación y pérdida de visión integral.
- Repositorios separados para documentación e infraestructura: rompe la trazabilidad.
- Sin control de versiones: elimina historial y revisión.

## Consecuencias

- Los cambios interdependientes podrán revisarse y versionarse juntos.
- La estructura de carpetas y convenciones son esenciales para mantener navegabilidad.

## Riesgos

- El repositorio puede crecer y requerir disciplina.
- Permisos excesivos amplían el alcance de cambios; se mitiga con revisión y secretos externos.

## Referencias relacionadas

- [Convenciones](../docs/CONVENTIONS.md)
- [Project Charter](../PROJECT_CHARTER.md)
- [ADR-001 — Persistencia bajo `/srv`](ADR-001-persistence-under-srv.md)
