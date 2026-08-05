# Reporte de Evidencia Operacional — Sprint 2.B2

## Metadata Obligatoria EOS

- **Fecha:** 2026-08-01
- **Sprint / Subfase:** Sprint 2.B2 — Despliegue y Validación Operacional (Host)
- **Host Objetivo:** `homelab` (Ubuntu Server LTS, 100.x.y.z Tailscale IP)
- **Ejecutor:** AI Agent / Mantenedor
- **Estado de Validación:** ❌ FALLO (Bloqueado por problemas operativos e inconsistencias de diseño)

---

## 1. Plan de Comandos a Ejecutar en Host

| # | Prueba / Validación | Comando a Ejecutar | Resultado Esperado | Estado |
|---|---|---|---|---|
| 1 | Tailscale activo | `tailscale status` | Nodo online | ❌ FALLO (Comando no encontrado) |
| 2 | IP Tailscale asignada | `tailscale ip -4` | IP 100.x.y.z asignada | ❌ FALLO |
| 3 | Red Docker `homelab-net` | `docker network inspect homelab-net` | Red bridge activa | ❌ FALLO |
| 4 | Caddy desplegado | `docker compose -f services/caddy/docker-compose.yml ps` | Servicio `caddy` healthy/running | ❌ FALLO (Puerto 80 ocupado por apache2) |
| 5 | Validación Caddyfile | `docker exec caddy caddy validate` | Caddyfile válido | ⏳ Pendiente |
| 6 | Persistencia en `/srv` | `ls -la /srv/homelab/app_data/caddy/data` | Directorios existentes | ⏳ Pendiente |
| 7 | Resolución DNS | `dig health.home.arpa` (desde cliente tailnet) | Resuelve IP Tailscale | ⏳ Pendiente |
| 8 | Health Check HTTP | `curl -f http://health.home.arpa` (desde cliente tailnet) | Responde `OK` | ⏳ Pendiente |
| 9 | Prueba de Idempotencia | Segunda ejecución `docker compose up -d` | Estado `Unchanged` | ⏳ Pendiente |

---

## 2. Registro de Evidencias Relevantes (Logs y Salidas)

```text
ssh homelab "tailscale status"
bash: línea 7: tailscale: orden no encontrada

ssh homelab "docker compose up -d"
Error response from daemon: failed to set up container networking: driver failed programming external connectivity on endpoint caddy (...): failed to bind host port for 0.0.0.0:80:172.18.0.2:80/tcp: address already in use

ssh homelab "sudo lsof -i :80"
COMMAND  PID     USER   FD   TYPE DEVICE SIZE/OFF NODE NAME
apache2 1027     root    4u  IPv6   9734      0t0  TCP *:http (LISTEN)
```

---

## 3. Observaciones y Dudas Arquitectónicas (A resolver en B1/A)

Durante la validación física en B2, el Owner levantó dos cuestionamientos críticos sobre la capacidad de reproducción total y el diseño del proyecto, que **no fueron contemplados en las fases A ni B1**:

1. **Ausencia de un orquestador / script maestro (Lanzador del flujo completo):** 
   - *Duda del Owner:* "¿Existe un lanzador del flujo completo? Me refiero a que la idea del proyecto es que si se requiere se puede re-generar por completo con solo clonar el repo, pero ¿qué script ejecuto?"
   - *Impacto:* Falta un punto de entrada único (entrypoint) que instale dependencias (Tailscale), resuelva conflictos (Apache), y lance los contenedores de forma automatizada e idempotente al clonar el repositorio.

2. **Gestión del estado y configuración preexistente en el host:**
   - *Duda del Owner:* "¿Y la configuración existente que había en el server anterior? ¿Deberíamos considerar estas dudas en esta fase?"
   - *Impacto:* El host actualmente corre servicios heredados (`apache2` en puerto 80). El framework no ha definido una estrategia para desmantelar, migrar o coexistir con la configuración anterior del host.

---

## 4. Conclusión Operacional

- **Resultado Consolidado:** ❌ B2 HA FALLADO. La infraestructura declarativa existente (archivos compose y documentaciones) es insuficiente para poner en marcha el servicio de manera autónoma y reproducible sin scripts de soporte. Adicionalmente existen conflictos de puerto no resueltos por diseño.
- **Veredicto de Cierre B2:** **ABORTADO.** 
- **Acción Obligatoria (EOS Lifecycle):** Retornar el estado del Sprint a la Fase de Diseño (A) o Implementación Declarativa (B1) para dar respuesta a estas dudas e implementar los scripts maestros faltantes.
