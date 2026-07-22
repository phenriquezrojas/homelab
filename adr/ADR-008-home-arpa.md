# ADR-008 — `home.arpa` como dominio interno

## Estado

Aceptada.

## Fecha

2026-07-22.

## Contexto

Los servicios internos necesitan nombres estables y legibles, sin depender de puertos o IPs memorizadas.

## Problema

Usar IPs o dominios no reservados complica operación y puede provocar conflictos de resolución.

## Decisión

`home.arpa` será el dominio interno. Los subdominios estarán en minúscula; la resolución se diseñará con la red interna y Tailscale.

## Alternativas consideradas

- IPs como acceso: menor legibilidad y flexibilidad.
- Dominios públicos propios: coste administrativo innecesario para el espacio interno.
- Sufijos locales no estandarizados: riesgo de conflicto y resolución inconsistente.

## Consecuencias

- Documentación y configuración usarán nombres bajo `home.arpa`.
- Se debe definir y validar DNS interno antes de depender de esos nombres.

## Riesgos

- DNS incompleto puede impedir acceso aunque el servicio esté disponible.
- Las convenciones de subdominios deben evitar colisiones.

## Referencias relacionadas

- [ADR-006 — Tailscale](ADR-006-tailscale.md)
- [ADR-007 — Caddy](ADR-007-caddy.md)
- [Convenciones](../docs/CONVENTIONS.md)
