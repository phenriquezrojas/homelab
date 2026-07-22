# ADR-002 — Docker First

## Estado

Aceptada.

## Fecha

2026-07-22.

## Contexto

El homelab integrará servicios con dependencias y ciclos de vida distintos, por lo que necesita un modelo uniforme de ejecución.

## Problema

La instalación directa en el host genera deriva de configuración, dependencias acopladas y menor reproducibilidad.

## Decisión

Docker será el mecanismo por defecto para servicios y dependencias. La configuración se declarará en este repositorio y los datos persistentes seguirán `/srv`.

## Alternativas consideradas

- Instalación nativa: acopla servicios y sistema operativo.
- Máquinas virtuales por servicio: añaden sobrecarga operativa innecesaria.
- Kubernetes: complejidad desproporcionada para esta fase.

## Consecuencias

- Los servicios compartirán convenciones de imágenes, redes, volúmenes y documentación.
- Se deberán gestionar actualizaciones, recursos y seguridad de contenedores.
- Native installation becomes an explicit exception that must be justified by a future ADR.

## Riesgos

- Los contenedores no eliminan la responsabilidad de asegurar el host.
- Imágenes no controladas y privilegios excesivos afectan reproducibilidad y seguridad.

## Referencias relacionadas

- [ADR-001 — Persistencia bajo `/srv`](ADR-001-persistence-under-srv.md)
- [ADR-003 — Ubuntu Server LTS](ADR-003-ubuntu-server.md)
- [ADR-007 — Caddy](ADR-007-caddy.md)
