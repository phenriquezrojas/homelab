# Operación de Caddy (Proxy Inverso)

Este documento describe la arquitectura, despliegue y operación del proxy inverso centralizado (Caddy) en la plataforma homelab.

## Arquitectura

Caddy actúa como el punto de entrada principal para todo el tráfico HTTP/HTTPS dentro de la red privada (Tailscale).

- **Modo de ejecución:** Contenedor Docker (`caddy:2.7.6-alpine`).
- **Ubicación del código:** `services/caddy/docker-compose.yml` y `services/caddy/Caddyfile`.
- **Puertos expuestos:** `80:80` y `443:443` (bind a `0.0.0.0`, aislamiento garantizado por la falta de port forwarding público en el router).
- **Red Docker:** `homelab-net` (propietario inicial durante Sprint 2.B; a partir de Sprint 3 otros servicios se conectan vía `external: true`).

## Persistencia

Siguiendo el **ADR-001**, los datos de Caddy se persisten en el host:

| Ruta en el Host | Ruta en el Contenedor | Propósito |
|---|---|---|
| `/srv/homelab/app_data/caddy/data` | `/data` | Certificados TLS e identificadores internos |
| `/srv/homelab/app_data/caddy/config` | `/config` | Estado de configuración de Caddy |

## Comandos Operativos

### 1. Despliegue / Inicio
```bash
cd services/caddy
docker compose up -d
```

### 2. Validación de Estado y Configuración
```bash
docker compose ps
docker exec caddy caddy validate
```

### 3. Prueba de Salud (Health Check)
Desde un cliente conectado a la tailnet:
```bash
curl -f http://health.home.arpa
```

### 4. Recarga de Configuración (Sin Downtime)
```bash
docker exec caddy caddy reload --config /etc/caddy/Caddyfile
```

### 5. Reversión / Desinstalación
```bash
cd services/caddy
docker compose down
docker network rm homelab-net
rm -rf /srv/homelab/app_data/caddy
```

> **Advertencia de Reversión:** La eliminación de la red `homelab-net` y de `/srv/homelab/app_data/caddy` solo debe ejecutarse mientras Caddy sea el único consumidor de la red y no existan otros contenedores dependientes.
