# Sprint 2.B — Plataforma Central (Implementación)

## Objetivo

Ejecutar el contrato técnico aprobado en el Sprint 2.A para convertir el host bootstrapped en una plataforma con acceso privado (Tailscale), proxy inverso (Caddy), resolución DNS interna (`home.arpa`) y red Docker compartida (`homelab-net`).

## Fase

Subfase B1 — Implementación Declarativa (Repositorio): Completada ✅  
Subfase B2 — Despliegue y Validación Operacional (Host): En Curso ⏳

> En Subfase B1 se versiona en Git el 100% de la configuración. En Subfase B2 se despliega en el host y se captura la evidencia empírica en `docs/evidence/Sprint-2.B2-validation.md`.

## Entregables Declarativos (Repositorio)

- [x] Declaración Compose para Caddy y red `homelab-net` en `services/caddy/docker-compose.yml`.
- [x] Caddyfile base con endpoint `health.home.arpa` en `services/caddy/Caddyfile`.
- [x] Documentación DNS (`home.arpa` MagicDNS + Split DNS) en `docs/DNS.md`.
- [x] Documentación operativa de Caddy en `docs/CADDY.md`.
- [x] Layout de persistencia en `/srv/homelab/app_data/caddy/` especificado y listo para despliegue.

## Trazabilidad de Tareas (Tasks HL-0015 a HL-0019)

| Task | Objetivo | Artefacto / Evidencia Declarativa | Estado Repo | Estado Host |
|---|---|---|---|---|
| **HL-0015** | Instalación y verificación de Tailscale | `docs/DNS.md` / Script de Host | PASS (Doc) | ⏳ Pendiente |
| **HL-0016** | Docker Compose y Caddyfile de Caddy | `services/caddy/docker-compose.yml` & `Caddyfile` | PASS | ⏳ Pendiente |
| **HL-0017** | Estrategia DNS `home.arpa` | `docs/DNS.md` (Desired State) | PASS | ⏳ Pendiente |
| **HL-0018** | Layout de persistencia bajo `/srv` | `docker-compose.yml` (`/srv/homelab/app_data/caddy/`) | PASS | ⏳ Pendiente |
| **HL-0019** | Validaciones End-to-End | `docs/CADDY.md` / Checklists de validación | PASS | ⏳ Pendiente |

## Matriz de Estado (Repositorio vs Host)

| Contrato / Componente | Implementado en Repo | Ejecutado / Validado en Host | Estado Consolidado |
|---|---|---|---|
| Compose Caddy (`caddy:2.7.6-alpine`) | ✔ `services/caddy/docker-compose.yml` | — | **PASS (Declarativo)** |
| Caddyfile (`health.home.arpa`) | ✔ `services/caddy/Caddyfile` | — | **PASS (Declarativo)** |
| Red Docker (`homelab-net`) | ✔ `services/caddy/docker-compose.yml` | — | **PASS (Declarativo)** |
| Desired State DNS (`home.arpa`) | ✔ `docs/DNS.md` | — | **PASS (Declarativo)** |
| Operación y Reversión | ✔ `docs/CADDY.md` | — | **PASS (Declarativo)** |
| `tailscale status` / IP 100.x | N/A (Host) | ⏳ Pendiente | **Pending (Host)** |
| `docker compose up -d` | N/A (Host) | ⏳ Pendiente | **Pending (Host)** |
| `docker exec caddy caddy validate` | N/A (Host) | ⏳ Pendiente | **Pending (Host)** |
| `dig health.home.arpa` (tailnet) | N/A (Host) | ⏳ Pendiente | **Pending (Host)** |
| `curl -f http://health.home.arpa` | N/A (Host) | ⏳ Pendiente | **Pending (Host)** |
| Idempotencia (`docker compose up`) | N/A (Host) | ⏳ Pendiente | **Pending (Host)** |

## Checklist de Evidencia Pendiente en Host para Cierre Definitivo

Antes de cambiar el estado del Sprint a `Completed` y proceder al merge del Gate de Cierre, se requiere recopilar y adjuntar las siguientes evidencias empíricas:

- [ ] `tailscale status` (Nodo online en la tailnet)
- [ ] `tailscale ip -4` (IP asignada 100.x.y.z)
- [ ] `docker compose up -d` (Despliegue exitoso de Caddy en host)
- [ ] `docker compose ps` (Servicio Caddy en estado healthy/running)
- [ ] `docker exec caddy caddy validate` (Configuración Caddyfile validada dentro del contenedor)
- [ ] `docker network inspect homelab-net` (Red bridge creada y activa)
- [ ] `ls -la /srv/homelab/app_data/caddy/data` (Estructura persistente creada en `/srv`)
- [ ] `dig health.home.arpa` (Resolución DNS verificada desde cliente tailnet)
- [ ] `curl -f http://health.home.arpa` (Respuesta `OK` recibida desde cliente tailnet)
- [ ] Segunda ejecución de `docker compose up -d` (Demostración de idempotencia sin cambios)

## Archivos Producidos / Modificados

- `services/caddy/docker-compose.yml`
- `services/caddy/Caddyfile`
- `docs/DNS.md`
- `docs/CADDY.md`
- `.ai/sprints/Sprint-2.B.md`

## Trazabilidad

| Tipo | Referencia |
|---|---|
| Plan | Sprint-2-Plan.md |
| Tasks | HL-0015 a HL-0019 |
| Sprint anterior | Sprint 2.A |
| Sprint siguiente | Sprint 3.A |
| Phase | Phase 2 (Storage), Phase 3 (Networking), Phase 5 (Platform Services) |

## Estado

In Review
