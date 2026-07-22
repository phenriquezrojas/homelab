# ADR-006 — Tailscale para acceso remoto privado

## Estado

Aceptada.

## Fecha

2026-07-22.

## Contexto

El operador necesita acceso fuera de la LAN sin exponer servicios administrativos a Internet ni operar una VPN tradicional.

## Problema

La exposición directa de puertos aumenta la superficie de ataque y el acceso remoto ad hoc dificulta una política de identidad consistente.

## Decisión

Tailscale será el plano de acceso remoto privado. Dispositivos, identidades, ACL y revocación se documentarán antes de su uso operativo.

## Alternativas consideradas

- Reenvío de puertos público: aumenta exposición.
- VPN autogestionada: añade operación de red innecesaria.
- Solo LAN: no cubre acceso remoto privado.

## Consecuencias

- El acceso remoto dependerá de identidad y plano de control de Tailscale.
- Los runbooks incluirán alta, baja y revisión de dispositivos.

## Riesgos

- Dispositivos autorizados en exceso amplían acceso no deseado.
- ACL deficientes pueden vulnerar mínimo privilegio.

## Referencias relacionadas

- [ADR-007 — Caddy](ADR-007-caddy.md)
- [ADR-008 — `home.arpa`](ADR-008-home-arpa.md)
- [Project Charter](../PROJECT_CHARTER.md)
