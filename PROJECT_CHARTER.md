# Project Charter — Homelab

## Propósito

Establecer y evolucionar un homelab self-hosted comprensible, reproducible y recuperable. El repositorio será la fuente de verdad de arquitectura, configuración declarativa, procedimientos y decisiones.

## Objetivos

- Proveer una plataforma estable para servicios internos seleccionados.
- Reducir trabajo manual con configuraciones y procedimientos repetibles.
- Tratar datos, seguridad y recuperación como responsabilidades explícitas.
- Mantener trazabilidad de cambios y decisiones.
- Permitir que un operador competente reconstruya la plataforma desde el repositorio y secretos externos.
- Preservar el patrimonio digital familiar para las siguientes generaciones.

## Alcance

El proyecto cubre host Ubuntu Server LTS, Docker, acceso privado, proxy HTTP(S), persistencia, respaldos, observabilidad, endurecimiento y guías operativas.

Sprint -1 cubre exclusivamente documentación y organización. No incluye instalaciones, despliegues, scripts funcionales, Docker Compose, servicios, CI/CD ni GitHub Actions.

## Fuera de alcance

- Servicios públicos de Internet por defecto.
- Alta disponibilidad empresarial o soporte 24/7.
- Secretos, claves privadas, datos personales o respaldos en Git.
- Cambiar una decisión aceptada sin un ADR que la reemplace.

## Principios rectores

- **Documentar antes de automatizar.**
- **Infraestructura reproducible:** configuración en repositorio y secretos externos.
- **Estado explícito:** datos persistentes bajo `/srv`.
- **Seguridad por defecto:** acceso remoto privado mediante Tailscale.
- **Recuperación verificable:** respaldo y restauración son parte del diseño.
- **Evolución incremental:** cada sprint debe ser acotado y validable.

## Roles

| Rol | Responsabilidad |
| --- | --- |
| Propietario | Priorizar, aprobar decisiones y custodiar secretos. |
| Mantenedor | Mantener documentación, configuración y trazabilidad. |
| Operador | Ejecutar runbooks, verificar respaldos y responder incidencias. |
| Colaborador | Proponer cambios según convenciones y ADR vigentes. |

Una persona puede asumir todos los roles.

## Criterios de éxito

- Las decisiones tienen ADR con contexto y consecuencias.
- La hoja de ruta comunica alcance, dependencias y salida.
- El repositorio no contiene secretos ni datos operativos.
- Los siguientes sprints pueden empezar con estructura y convenciones inequívocas.

## Gobernanza

Los cambios de arquitectura se proponen mediante ADR. Las decisiones aceptadas son registro histórico; un cambio se expresa mediante un ADR nuevo que las sustituya o complemente.
