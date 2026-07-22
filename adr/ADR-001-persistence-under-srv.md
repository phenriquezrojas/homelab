# ADR-001 — Persistencia bajo `/srv`

## Estado

Aceptada.

## Fecha

2026-07-22.

## Contexto

Los servicios Docker requieren datos que sobrevivan a la recreación de contenedores y que entren en procedimientos de respaldo y restauración.

## Problema

Sin una ubicación única, los datos se dispersan entre perfiles de usuario, volúmenes no documentados y rutas implícitas.

## Decisión

Todo dato persistente gestionado por el homelab se alojará bajo `/srv`. La jerarquía detallada se definirá durante Sprint 0 y se separará del contenido versionado.

## Alternativas consideradas

- Directorios bajo `/home`: mezcla datos de servicio con perfiles de usuario.
- Rutas por defecto de Docker: ofrecen menor visibilidad y portabilidad operacional.
- Un volumen sin ruta estándar: podría complementar la convención lógica, pero no sustituirla.

## Consecuencias

- Runbooks, permisos, respaldos y restauraciones tendrán una raíz común.
- Sprint 0 debe definir jerarquía, propiedad y capacidad.

## Riesgos

- Permisos incorrectos pueden exponer o bloquear datos.
- Capacidad insuficiente o un único disco son riesgos que deberán monitorizarse y respaldarse.

## Referencias relacionadas

- [ADR-002 — Docker First](ADR-002-docker-first.md)
- [ADR-005 — Restic + Backblaze B2](ADR-005-restic-backblaze.md)
- [Project Charter](../PROJECT_CHARTER.md)
