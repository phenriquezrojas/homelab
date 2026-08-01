# Sprint 2.A — Plataforma Central (Diseño)

## Objetivo

Diseñar el contrato técnico que Sprint 2.B deberá ejecutar para convertir el host bootstrapped en una plataforma con acceso privado (Tailscale), proxy inverso (Caddy), resolución DNS interna (`home.arpa`) y red Docker compartida. Al finalizar esta fase, todos los artefactos de diseño deben estar aprobados y ser suficientes para que Sprint 2.B opere sin ambigüedad.

## Fase

A — Diseño

> En Fase A se generan los artefactos de diseño; en Fase B se ejecuta lo aprobado.
> Ver [FEATURE_LIFECYCLE.md](../lifecycle/FEATURE_LIFECYCLE.md) para reglas de transición.

## Principios de la Plataforma Central

Toda implementación de la plataforma central deberá cumplir estos principios. Si la implementación viola alguno, no cumple el contrato del Sprint.

1. **Docker First.** Todos los componentes de plataforma se ejecutan como contenedores Docker (ADR-002). Tailscale es la única excepción justificada (requiere acceso directo al kernel para crear la interfaz de red).
2. **Configuración Declarativa Versionada.** Toda configuración (Caddyfile, docker-compose, Tailscale ACL) se versiona en el repositorio (Constitución §1).
3. **Persistencia bajo `/srv`.** Los datos operativos de Caddy (certificados, estado) se persisten bajo `/srv/homelab/app_data/caddy` (ADR-001).
4. **Acceso Privado Exclusivo.** No se expone ningún puerto al enrutador físico ni a Internet pública. Todo el tráfico pasa por la tailnet (ADR-006).
5. **Reversible.** Cada componente define su procedimiento de reversión antes de implementarse (Constitución §12).
6. **Independencia de Servicios.** La plataforma central no depende de ningún servicio de aplicación. Los servicios de aplicación dependen de ella.

## Alcance

### Incluido

- Diseño de la instalación y configuración de Tailscale en el host.
- Diseño del Caddyfile base con un servicio de prueba mínimo (health check).
- Diseño del docker-compose para Caddy como proxy inverso.
- Diseño de la estrategia DNS `home.arpa` (mecanismo de resolución).
- Diseño de la red Docker compartida (`homelab-net`) para comunicación entre Caddy y servicios futuros.
- Diseño del layout de persistencia bajo `/srv/homelab/app_data/caddy`.
- Definición de criterios de aceptación medibles para Sprint 2.B.
- Definición de evidencias de validación esperadas.
- Identificación de riesgos y mitigaciones.
- Creación de Tasks HL-0015 a HL-0019.

### Excluido

- Escribir código ejecutable, configuración operativa o modificar el host.
- Instalar Tailscale, Caddy o cualquier software.
- Crear contenedores, redes o volúmenes Docker.
- Configurar PostgreSQL, Redis u otros servicios de datos.
- Desplegar Immich o cualquier servicio de aplicación.
- Configurar respaldos o monitorización.
- Configurar el enrutador físico, firewall o reglas de NAT.
- Diseñar TLS con certificados públicos (fuera del alcance para red interna).

### No Objetivos

- No se diseña HA (alta disponibilidad) ni failover.
- No se implementa CI/CD para despliegue de configuración.
- No se diseñan runbooks de operación (corresponde a Sprint 4+).
- No se endurece la seguridad del host (corresponde a Sprint 6).

## Componentes de Diseño

### 1. Tailscale — Acceso Privado

**ADR:** [ADR-006](../../adr/ADR-006-tailscale.md)

Tailscale se instala directamente en el host (no como contenedor) porque necesita acceso directo al stack de red del kernel para crear y operar la interfaz `tailscale0`.

El diseño debe definir:
- Procedimiento de instalación idempotente.
- Autenticación del nodo (auth key preconfigurada vs. interactiva).
- ACL mínima esperada (definida en la consola Tailscale, documentada en el repo).
- Validación: `tailscale status` confirma que el nodo está conectado a la tailnet.
- Reversión: `tailscale down`, desinstalación y eliminación de estado.

### 2. Caddy — Proxy Inverso

**ADR:** [ADR-007](../../adr/ADR-007-caddy.md)

Caddy corre como contenedor Docker en la red `homelab-net`. Escucha en los puertos 80 y 443 del host. Su configuración se declara en un `Caddyfile` versionado.

El diseño debe definir:
- Docker Compose para Caddy (`services/caddy/docker-compose.yml`).
- Caddyfile base con al menos un endpoint de health check (`health.home.arpa`).
- Volúmenes: configuración montada desde el repo, datos persistentes en `/srv/homelab/app_data/caddy`.
- Red Docker: nombre `homelab-net`, tipo `bridge`.
- Validación: `curl -f http://health.home.arpa` retorna respuesta desde el proxy.
- Reversión: `docker compose down`, eliminar volúmenes y red.

### 3. DNS `home.arpa` — Resolución Interna

**ADR:** [ADR-008](../../adr/ADR-008-home-arpa.md)

El mecanismo de resolución debe permitir que los clientes de la tailnet resuelvan `*.home.arpa` hacia la IP Tailscale del servidor.

El diseño debe evaluar y recomendar una de estas estrategias:

| Estrategia | Ventaja | Riesgo |
|---|---|---|
| Tailscale MagicDNS | Integrado, cero infraestructura adicional | Depende de la consola Tailscale |
| Pi-hole / CoreDNS como contenedor | Control total de registros | Añade un componente más que operar |
| Archivo `/etc/hosts` en clientes | Simplicidad extrema | No escala, no se versiona |

La estrategia elegida se documentará con justificación y, si introduce un componente nuevo, requerirá un ADR.

### 4. Red Docker Compartida

**ADR:** [ADR-002](../../adr/ADR-002-docker-first.md)

Una red Docker bridge nombrada (`homelab-net`) permite que Caddy enrute tráfico hacia contenedores de servicio sin exponer puertos al host.

El diseño debe definir:
- Nombre de la red: `homelab-net`.
- Tipo: `bridge`.
- Creación (Sprint 2.B): declarada en el compose de Caddy como red `external: false` (se crea si no existe).
- Consumo futuro (Sprint 3+): a partir del Sprint 3, esta red pasa a considerarse infraestructura existente y todos los servicios deberán referenciarla obligatoriamente como `external: true`.

## Entregables

### Fase A — Diseño

- [x] Sprint Specification completada (este documento).
- [x] Implementation Plan aprobado (`Sprint-2-Plan.md`).
- [x] Tasks con checklist verificable creadas (HL-0015 a HL-0019).
- [x] Criterios de aceptación definidos en cada tarea.
- [x] Checklist de revisión preparado en el plan.
- [x] Estrategia DNS evaluada y recomendada.
- [x] ADR no requerido (MagicDNS es capacidad nativa de ADR-006, no introduce componente nuevo).

### Fase B — Implementación

La Fase B ejecutará este contrato. Los entregables de implementación se documentarán en `Sprint-2.B.md`.

## Evidencia de validación requerida para Sprint 2.B

Sprint 2.B deberá demostrar las siguientes validaciones:

| Validación | Comando/Evidencia |
|---|---|
| Tailscale conectado a la tailnet | `tailscale status` muestra el nodo online |
| Caddy corriendo como contenedor | `docker compose ps` muestra servicio healthy |
| Red Docker `homelab-net` creada | `docker network inspect homelab-net` |
| DNS `home.arpa` resuelve | `curl -f http://health.home.arpa` (desde un nodo en la tailnet) retorna OK |
| Persistencia de Caddy en `/srv` | `ls -la /srv/homelab/app_data/caddy/` |
| Idempotencia del despliegue | Segunda ejecución de `docker compose up -d` sin cambios |

## Criterios de cierre del Sprint 2.A

Para solicitar el cierre del Gate, deben cumplirse las siguientes condiciones documentales:

- [x] Todos los entregables de Fase A están completados.
- [x] `Sprint-2-Plan.md` en estado In Review.
- [x] Todas las Tasks (HL-0015 a HL-0019) en estado In Review.
- [x] Estrategia DNS definida y justificada.
- [x] Criterios de aceptación completos y verificables.

El cierre oficial ocurre cuando el Owner aprueba el Gate.

## Archivos a producir

- `.ai/sprints/Sprint-2.A.md`
- `.ai/implementation/Sprint-2-Plan.md`
- `.ai/tasks/HL-0015.md` a `HL-0019.md`
- `docs/DNS.md` (diseño del contenido esperado para Sprint 2.B)
- `docs/CADDY.md` (diseño del contenido esperado para Sprint 2.B)

## Trazabilidad

| Tipo | Referencia |
|---|---|
| ADR | ADR-001, ADR-002, ADR-006, ADR-007, ADR-008 |
| Tasks | HL-0015, HL-0016, HL-0017, HL-0018, HL-0019 |
| Sprint anterior | Sprint 1.B |
| Sprint siguiente | Sprint 2.B |
| Phase | Phase 2 (Storage), Phase 3 (Networking), Phase 5 (Platform Services) |

## Estado

Completed

## Lecciones aprendidas

Del Sprint 1: La separación A/B ha demostrado su valor. El contrato de diseño detallado en la Fase A (6 fases, 3 niveles de idempotencia, criterios de aceptación medibles) permitió que la Fase B se ejecutara de forma autónoma, con revisiones enfocadas en el cumplimiento del contrato y no en decisiones de diseño ad hoc. La regla de Gate Review añadida a la Constitución protege contra la ampliación involuntaria del alcance durante las revisiones.

## Deuda técnica heredada

Del Sprint 1.B:
- Separación Ensure → Validate → Evidence en el motor de convergencia del bootstrap.
- Parametrización CLI del bootstrap (`--dry-run`, `--verbose`).
- Refactorización de `validate_component` en funciones atómicas.
- Encapsulamiento del flujo principal en `main()`.
- Rotación de logs con logrotate.

Del Gate Sprint 2.A:
- **Redacción de Contratos (Plan):** Separar explícitamente la formulación de diseño ("La implementación deberá validar mediante...") de la ejecución pura ("Ejecutar curl..."), evitando mezclar el qué (Fase A) con el cómo (Fase B).
- **Inmutabilidad de imágenes:** Anclar las imágenes de contenedores a hashes `sha256` específicos para garantizar reproducibilidad estricta.

## Próximo Sprint

Sprint 2.B (Implementación de la Plataforma Central)
