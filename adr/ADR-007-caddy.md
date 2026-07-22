# ADR-007 — Caddy como proxy inverso

## Estado

Aceptada.

## Fecha

2026-07-22.

## Contexto

Los servicios HTTP necesitan un punto de entrada uniforme para enrutamiento, cabeceras, TLS cuando corresponda y control de acceso.

## Problema

Exponer puertos por servicio dispersa configuración y dificulta políticas coherentes.

## Decisión

Caddy será el proxy inverso estándar. Su configuración declarativa se versionará e integrará con `home.arpa` y Tailscale.

## Alternativas consideradas

- Nginx: se mantiene el estándar aprobado de Caddy.
- Traefik: no es la tecnología seleccionada.
- Exposición directa de puertos: fragmenta acceso y enrutamiento.

## Consecuencias

- Los servicios HTTP declararán su integración con Caddy.
- Caddy centralizará políticas comunes de HTTP.

## Riesgos

- Una configuración incorrecta puede afectar a varios servicios.
- Certificados, DNS y acceso requieren validación antes de exponer servicios.

## Referencias relacionadas

- [ADR-006 — Tailscale](ADR-006-tailscale.md)
- [ADR-008 — `home.arpa`](ADR-008-home-arpa.md)
- [ADR-002 — Docker First](ADR-002-docker-first.md)
