# Reporte de Evidencia Operacional — Sprint 2.B2

## Metadata Obligatoria EOS

- **Fecha:** 2026-08-01
- **Sprint / Subfase:** Sprint 2.B2 — Despliegue y Validación Operacional (Host)
- **Host Objetivo:** `homelab` (Ubuntu Server LTS, 100.x.y.z Tailscale IP)
- **Ejecutor:** AI Agent / Mantenedor
- **Estado de Validación:** Pendiente de Ejecución Empírica en Host

---

## 1. Plan de Comandos a Ejecutar en Host

| # | Prueba / Validación | Comando a Ejecutar | Resultado Esperado | Estado |
|---|---|---|---|---|
| 1 | Tailscale activo | `tailscale status` | Nodo online | ⏳ Pendiente |
| 2 | IP Tailscale asignada | `tailscale ip -4` | IP 100.x.y.z asignada | ⏳ Pendiente |
| 3 | Red Docker `homelab-net` | `docker network inspect homelab-net` | Red bridge activa | ⏳ Pendiente |
| 4 | Caddy desplegado | `docker compose -f services/caddy/docker-compose.yml ps` | Servicio `caddy` healthy/running | ⏳ Pendiente |
| 5 | Validación Caddyfile | `docker exec caddy caddy validate` | Caddyfile válido | ⏳ Pendiente |
| 6 | Persistencia en `/srv` | `ls -la /srv/homelab/app_data/caddy/data` | Directorios existentes | ⏳ Pendiente |
| 7 | Resolución DNS | `dig health.home.arpa` (desde cliente tailnet) | Resuelve IP Tailscale | ⏳ Pendiente |
| 8 | Health Check HTTP | `curl -f http://health.home.arpa` (desde cliente tailnet) | Responde `OK` | ⏳ Pendiente |
| 9 | Prueba de Idempotencia | Segunda ejecución `docker compose up -d` | Estado `Unchanged` | ⏳ Pendiente |

---

## 2. Registro de Evidencias Relevantes (Logs y Salidas)

```text
[Salidas de comandos capturadas durante la ejecución real en el host]
```

---

## 3. Conclusión Operacional

- **Resultado Consolidado:** *[A completar tras la ejecución de los comandos]*
- **Veredicto de Cierre B2:** *[A completar tras verificar 100% de los criterios]*
