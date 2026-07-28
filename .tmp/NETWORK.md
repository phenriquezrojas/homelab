# Topología de Red Base

Este documento detalla la ruta de tráfico teórica hacia los contenedores del Homelab y el mecanismo de resolución DNS.

## Arquitectura de Red Interna

El acceso a los servicios de Homelab se realiza exclusivamente de forma privada mediante **Tailscale**. No existe exposición pública de puertos en el enrutador físico.

### Flujo de Peticiones

El flujo de tráfico desde el cliente hasta el servicio final se compone de las siguientes etapas:

1. **Cliente Tailscale:** El usuario se conecta a la red privada (tailnet) desde su dispositivo.
2. **Tailscale IP:** La petición se dirige a la IP asignada por Tailscale al servidor físico (el host Ubuntu).
3. **Caddy (Proxy Inverso):** El servicio Caddy, escuchando en los puertos estándar (80/443), intercepta la petición entrante. Caddy actúa como punto de terminación TLS (si aplica para la red interna) y enrutador de capa 7.
4. **Contenedor Destino:** Caddy redirige el tráfico hacia el contenedor Docker correspondiente (ej. `immich-app:3001`) a través de una red bridge interna de Docker.

## Resolución DNS Local (`home.arpa`)

Para evitar el uso de direcciones IP en la navegación, se utiliza el dominio interno estándar `home.arpa` (definido en RFC 8375). 

- **Nombres de dominio:** Los servicios se expondrán bajo subdominios lógicos, por ejemplo: `immich.home.arpa`.
- **Resolución:** El mecanismo primario de resolución se implementará en etapas posteriores. Caddy utilizará estos nombres de dominio para aplicar las reglas de enrutamiento definidas en el `Caddyfile`.

Esta topología prepara el terreno para la instalación de Caddy y Tailscale en los próximos Sprints.
