# Configuración DNS — `home.arpa`

Este documento define el **desired state** (configuración esperada) de la resolución DNS interna para la plataforma homelab bajo el dominio reservado `home.arpa` (RFC 8375).

## Arquitectura de Resolución

La resolución de nombres dentro de la red privada utiliza **Tailscale MagicDNS + Split DNS**.

- **Desired State (Repositorio):** Declarado en este documento.
- **Runtime State (Consola de Tailscale):** La consola de administración de Tailscale implementa la resolución asignando la IP Tailscale del host a las peticiones del dominio `home.arpa`.

## Registros Requeridos

| Dominio / Subdominio | Destino (IP) | Propósito |
|---|---|---|
| `health.home.arpa` | IP Tailscale del host (`100.x.y.z`) | Endpoint de salud y validación de la plataforma central |

## Procedimiento de Configuración (Consola Tailscale Admin)

1. **Habilitar MagicDNS:**
   - Navegar a `DNS` en el Panel de Administración de Tailscale.
   - Activar la opción `MagicDNS`.

2. **Configurar Split DNS / Extra Records:**
   - En la sección `Custom DNS` / `Search Domains`, añadir el dominio `home.arpa`.
   - Agregar el registro A/AAAA o la regla de Split DNS apuntando hacia la IP de Tailscale del host servidor `homelab` (ej. `100.x.y.z`).

## Validación

Desde cualquier cliente autenticado en la tailnet:

```bash
dig health.home.arpa
```

Debe retornar la IP Tailscale (`100.x.y.z`) del servidor `homelab`.

## Reversión

Para revertir la configuración DNS:
1. Eliminar la entrada Split DNS para `home.arpa` en el panel de Tailscale Admin.
2. Eliminar los registros custom de `home.arpa`.
