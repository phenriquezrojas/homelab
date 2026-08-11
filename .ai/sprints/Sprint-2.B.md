# Sprint 2.B — Implementación del Runtime Engine

```yaml
Sprint: 2.B
Revision: 2
Document Status: ACTIVE
Execution Status: COMPLETED
Design Baseline: 
  Artifact: Sprint-2-Runtime-Design.md
  Revision: 1
```

## Objetivo

Desarrollar y validar el **Homelab Runtime Engine** (Motor de Convergencia) v1.0, y utilizarlo para desplegar el primer dominio de capacidades (Docker, Tailscale, Caddy y MagicDNS) tal como está especificado en el *Design Baseline*.

## Fase

A — Diseño | B1 — Implementación Declarativa (Repositorio) | B2 — Despliegue y Validación Operacional (Host)

> En Fase A se generan los artefactos de diseño.
> En Subfase B1 se implementan y versionan en Git los manifiestos, código y documentación.
> En Subfase B2 se despliega en el host objetivo y se adjunta el artefacto obligatorio de evidencia (`docs/evidence/`).
> Ver [FEATURE_LIFECYCLE.md](../lifecycle/FEATURE_LIFECYCLE.md) para reglas de transición.

## Alcance

**Incluido:**
- Desarrollo del motor (Discovery, Observer, Planner, Transition Resolver, Executor, Reporter).
- Implementación de los 4 componentes iniciales según el contrato del Runtime.
- Despliegue en el host usando el motor.

**Excluido:**
- Servicios adicionales fuera del primer dominio.
- Implementación de estado persistente complejo más allá de los directorios en `/srv`.

## Entregables

### Subfase B1 — Implementación Declarativa (Repositorio)

- [x] Código fuente del motor versionado en el repositorio (`runtime/`).
- [x] Scripts de los componentes de Docker, Tailscale, Caddy y MagicDNS (`runtime/components/`).
- [x] Registro (`registry.yaml`) con las capacidades y proveedores configurados.

### Subfase B2 — Despliegue y Validación Operacional (Host)

- [x] Motor ejecutado en el host objetivo logrando la convergencia de `ABSENT` a `HEALTHY` para todas las capacidades.
- [x] Artefacto obligatorio de evidencia operacional creado en `docs/evidence/Sprint-2.B2-validation.md` documentando la traza del Executor.

## Estado de Ejecución

**COMPLETED**

## Tasks Asociadas (Fase B1)

- [HL-0020] Compilar motor del Runtime en Go (`runtime/cmd/homelab/main.go`).
- [HL-0021] Escribir contratos para `docker-engine`, `tailscale`, `caddy`, `magic-dns`.
- [HL-0022] Crear registro declarativo `registry.yaml`.
- [HL-0023] Documentar el runtime y adaptar bootstrap.

## Lecciones aprendidas

*A llenar post-ejecución*

## Deuda técnica

- **Subcomando `homelab status`:** Definido en el plan como herramienta de diagnóstico ("Muestra el último estado conocido") pero no implementado en B1. No bloquea B2 pero debe añadirse en un sprint posterior.
- **`go.sum` generado manualmente:** El archivo fue creado con hashes conocidos. Debe regenerarse con `go mod tidy` al compilar en el servidor para garantizar integridad criptográfica.
- **Reporter solo genera texto:** El plan mencionaba "texto/JSON". La serialización JSON puede añadirse como mejora futura.

## Próximo Sprint

Sprint 3.A (Respaldo - Diseño)
