---
Estado: In Review
Versión: 1.0
Autor: AI Agent
Relacionado: Sprint 2.A, Sprint 2.B, ADR-001, ADR-002, ADR-006, ADR-007, ADR-008
---

# Implementation Plan — Plataforma Central (Sprint 2)

## Objetivo

Definir el contrato técnico completo que Sprint 2.B deberá ejecutar para habilitar acceso privado (Tailscale), proxy inverso (Caddy), resolución DNS interna (`home.arpa`) y red Docker compartida sobre el host bootstrapped en Sprint 1.

## Contexto y ADR

Este plan se fundamenta en los siguientes ADR aprobados:

- [ADR-001](../../adr/ADR-001-persistence-under-srv.md) — Persistencia bajo `/srv`. Los datos operativos de Caddy (certificados internos, estado) se persisten en `/srv/homelab/app_data/caddy`.
- [ADR-002](../../adr/ADR-002-docker-first.md) — Docker First. Caddy se ejecuta como contenedor. Tailscale es la excepción documentada: requiere acceso directo al kernel.
- [ADR-006](../../adr/ADR-006-tailscale.md) — Tailscale para acceso privado. Todo el tráfico remoto pasa por la tailnet. No se exponen puertos públicos.
- [ADR-007](../../adr/ADR-007-caddy.md) — Caddy como proxy inverso. Centraliza enrutamiento HTTP/S, TLS interno y políticas comunes.
- [ADR-008](../../adr/ADR-008-home-arpa.md) — `home.arpa` como dominio interno. Los servicios se acceden mediante nombres estables bajo este dominio reservado (RFC 8375).

## Prerrequisitos del Host

Antes de ejecutar Sprint 2.B, el host debe cumplir:

1. Sprint 1.B completado (bootstrap exitoso, Docker funcional).
2. Estructura `/srv/homelab` existente y validada.
3. Acceso SSH confirmado desde la máquina de gestión.
4. Cuenta de Tailscale creada y accesible.

## Alcance

### Incluido en Sprint 2.B

1. **Tailscale en el host:** Instalación idempotente, autenticación y validación.
2. **Red Docker compartida:** Creación de la red `homelab-net` (bridge).
3. **Caddy como contenedor:** Docker Compose, Caddyfile base y persistencia en `/srv`.
4. **Resolución DNS `home.arpa`:** Configuración del mecanismo elegido.
5. **Validación end-to-end:** Un endpoint de health check accesible vía `health.home.arpa` a través de la tailnet.

### Excluido de Sprint 2.B

- PostgreSQL, Redis u otros servicios de datos.
- Cualquier servicio de aplicación (Immich, etc.).
- Respaldos, monitorización o alertas.
- Hardening del host o configuración de firewall.
- Certificados TLS públicos (Let's Encrypt).
- CI/CD o pipelines de despliegue.

## Arquitectura de la Plataforma

```
┌─────────────────────────────────────────────────────────────┐
│  CLIENTE (laptop, móvil)                                    │
│  ┌──────────────┐                                           │
│  │ Tailscale    │                                           │
│  │ Client       │                                           │
│  └──────┬───────┘                                           │
└─────────┼───────────────────────────────────────────────────┘
          │ WireGuard tunnel (tailnet)
          │
┌─────────┼───────────────────────────────────────────────────┐
│  HOST HOMELAB (Ubuntu Server LTS)                           │
│         │                                                   │
│  ┌──────┴───────┐                                           │
│  │ Tailscale    │ ← Instalación nativa (acceso a kernel)    │
│  │ (tailscale0) │                                           │
│  └──────┬───────┘                                           │
│         │ 100.x.y.z (IP Tailscale)                          │
│         │                                                   │
│  ┌──────┴───────┐                                           │
│  │ Caddy        │ ← Contenedor Docker                       │
│  │ :80 / :443   │   Caddyfile versionado                    │
│  └──────┬───────┘                                           │
│         │ homelab-net (Docker bridge)                        │
│         │                                                   │
│  ┌──────┴───────┐  ┌──────────────┐  ┌──────────────┐      │
│  │ Servicio A   │  │ Servicio B   │  │ Servicio N   │      │
│  │ (futuro)     │  │ (futuro)     │  │ (futuro)     │      │
│  └──────────────┘  └──────────────┘  └──────────────┘      │
│                                                             │
│  /srv/homelab/app_data/caddy/ ← Persistencia de estado      │
└─────────────────────────────────────────────────────────────┘
```

## Pasos del Sprint 2.B

### PASO 1: Instalación de Tailscale

**Procedimiento idempotente:**

```
1. Verificar si tailscale está instalado (command -v tailscale).
2. Si no está instalado:
   a. Descargar e instalar via script oficial (curl -fsSL https://tailscale.com/install.sh).
   b. El script oficial es idempotente y detecta la distribución.
3. Verificar el servicio tailscaled:
   a. systemctl is-active tailscaled
   b. systemctl is-enabled tailscaled
   c. Si no está activo/habilitado: systemctl enable --now tailscaled
4. Autenticar el nodo:
   a. Verificar si ya está conectado: tailscale status
   b. Si no está conectado: tailscale up (requiere auth key o login interactivo).
   c. Documentar: el primer tailscale up requiere intervención humana (abrir URL y autorizar).
5. Validar: tailscale status muestra el nodo como "online".
```

**Autenticación:**

La primera ejecución de `tailscale up` requiere una acción humana (autorizar el nodo en la consola Tailscale). Esto es una excepción aceptada al principio de "sin intervención manual" porque:
- Solo ocurre en la primera ejecución.
- Es un requisito de seguridad de Tailscale (autenticar identidad).
- Las ejecuciones subsecuentes son automáticas.

Se documentará como prerrequisito de Sprint 2.B.

**Reversión:**

```bash
tailscale down
sudo apt-get remove --purge tailscale
sudo rm -rf /var/lib/tailscale
```

### PASO 2: Estructura de persistencia para Caddy

**Directorios requeridos:**

| Directorio | Permisos | Propietario | Propósito |
|---|---|---|---|
| `/srv/homelab/app_data/caddy` | 755 | root:root | Raíz de persistencia de Caddy |
| `/srv/homelab/app_data/caddy/data` | 755 | root:root | Certificados y estado TLS |
| `/srv/homelab/app_data/caddy/config` | 755 | root:root | Configuración persistente generada |

Estos directorios se crean usando `ensure_directory` del bootstrap o un script equivalente. Si ya existen, no se modifican.

**Reversión:**

```bash
rm -rf /srv/homelab/app_data/caddy
```

### PASO 3: Red Docker compartida

**Creación:**

La red `homelab-net` se declara en el `docker-compose.yml` de Caddy. Docker Compose la crea automáticamente si no existe.

```yaml
networks:
  homelab-net:
    name: homelab-net
    driver: bridge
```

Los servicios futuros (a partir de Sprint 3) se unirán a esta red considerándola infraestructura existente, declarando:

```yaml
networks:
  homelab-net:
    external: true
```

**Validación:**

```bash
docker network inspect homelab-net
```

**Reversión:**

```bash
docker network rm homelab-net
```

### PASO 4: Configuración y despliegue de Caddy

**Estructura de archivos en el repositorio:**

```
services/caddy/
├── docker-compose.yml
└── Caddyfile
```

**docker-compose.yml (diseño):**

```yaml
services:
  caddy:
    image: caddy:2.7.6-alpine # TODO (Deuda Técnica): Fijar a sha256 específico en producción
    container_name: caddy
    restart: unless-stopped
    ports:
      - "80:80"
      - "443:443"
    volumes:
      - ./Caddyfile:/etc/caddy/Caddyfile:ro
      - /srv/homelab/app_data/caddy/data:/data
      - /srv/homelab/app_data/caddy/config:/config
    networks:
      - homelab-net

networks:
  homelab-net:
    name: homelab-net
    driver: bridge
    # El compose de Caddy es el propietario inicial de la red durante Sprint 2.B. 
    # A partir de Sprint 3 la red pasa a considerarse infraestructura existente.
```

**Decisión explícita sobre el bind de puertos:** Caddy expone los puertos `80:80` y `443:443` (bind a `0.0.0.0` por defecto). La restricción de "acceso privado exclusivo" (ADR-006) se cumple mediante la ausencia de port forwarding en el enrutador físico. No se fuerza el bind a la IP de Tailscale para evitar problemas de orden de arranque (Docker intentando hacer bind antes de que la interfaz `tailscale0` esté lista).

**Caddyfile base (diseño):**

```caddyfile
health.home.arpa {
    respond "OK" 200
}
```

Este Caddyfile mínimo:
- Define un servicio de health check accesible en `health.home.arpa`.
- No configura TLS porque la red interna (tailnet) ya provee cifrado end-to-end.
- Actúa como plantilla base que Sprint 2.B+ ampliará con servicios reales.

**Validación:**

La implementación deberá validar mediante:

```bash
docker compose -f services/caddy/docker-compose.yml ps   # Caddy "running"
docker exec caddy caddy validate                          # Valida el Caddyfile
curl -f http://health.home.arpa                           # Retorna "OK" (desde un cliente perteneciente a la tailnet)
```

**Reversión:**

```bash
cd services/caddy && docker compose down
docker network rm homelab-net
rm -rf /srv/homelab/app_data/caddy
```

> **Nota de reversión:** La eliminación de la red `homelab-net` y del almacenamiento persistente `/srv/homelab/app_data/caddy` sólo aplica mientras Caddy sea el único consumidor y no existan otros contenedores utilizando esta infraestructura.

### PASO 5: Resolución DNS `home.arpa`

**Evaluación de estrategias:**

| Estrategia | Complejidad | Mantenibilidad | Escalabilidad | Riesgo |
|---|---|---|---|---|
| Tailscale MagicDNS + Split DNS | Baja | Alta (gestionada) | Buena | Dependencia de Tailscale |
| CoreDNS como contenedor | Media | Media | Muy buena | Componente adicional que operar |
| `/etc/hosts` en clientes | Muy baja | Muy baja | Mala | No escala, no versionable |

**Recomendación: Tailscale MagicDNS + Split DNS**

Justificación:
1. **Simplicidad (Constitución §11):** No introduce componentes adicionales que operar.
2. **Integración natural:** Tailscale ya está instalado como plano de acceso; MagicDNS extiende esa capacidad sin software adicional.
3. **Configuración declarativa:** Los registros DNS se definen en la consola de Tailscale (Extra Records / Split DNS) y se documentan en el repositorio.
4. **Alcance controlado:** Solo los dispositivos de la tailnet resuelven `*.home.arpa`, alineado con el acceso privado exclusivo.

**Configuración esperada en Tailscale Admin Console:**

- Habilitar MagicDNS.
- Configurar Split DNS: dominio `home.arpa` resuelve hacia la IP Tailscale del servidor homelab (100.x.y.z).
- **Registro requerido explícito:** `health.home.arpa` → IP Tailscale del host homelab.
- Documentar la configuración en `docs/DNS.md` como fuente de verdad versionada (Constitución §1). La consola de Tailscale implementa el desired state definido en `docs/DNS.md`.

**Nota:** Esta configuración no requiere ADR adicional porque utiliza una capacidad nativa del componente ya aprobado (ADR-006). La configuración en la consola de Tailscale es *runtime state* (estado operativo gestionado externamente), mientras que `docs/DNS.md` contiene el *desired state* (configuración esperada documentada en el repositorio). Esta distinción es común en infraestructura administrada (Cloudflare, GCP, AWS) y no constituye una excepción a la Constitución §1.

**Reversión:**

- Deshabilitar Split DNS en la consola Tailscale.
- Eliminar registros de `home.arpa`.

## Validaciones End-to-End

| # | Validación | Comando | Resultado esperado |
|---|---|---|---|
| 1 | Tailscale activo | `tailscale status` | Nodo online |
| 2 | IP Tailscale asignada | `tailscale ip -4` | Retorna 100.x.y.z |
| 3 | Red Docker creada | `docker network inspect homelab-net` | Red tipo bridge |
| 4 | Caddy corriendo | `docker compose -f services/caddy/docker-compose.yml ps` | Estado "running" |
| 4.1 | Caddyfile válido | `docker exec caddy caddy validate` | Retorna validación exitosa |
| 5 | Persistencia Caddy | `ls /srv/homelab/app_data/caddy/data` | Directorio existe |
| 6 | DNS resuelve | `dig health.home.arpa` (desde un cliente perteneciente a la tailnet) | Retorna IP Tailscale del host |
| 7 | Health check OK | `curl -f http://health.home.arpa` (desde un cliente perteneciente a la tailnet) | Retorna "OK" |
| 8 | Idempotencia | `docker compose up -d` segunda vez | Sin cambios |

## Criterios de Éxito (Success Criteria)

Este plan se considera ejecutado exitosamente cuando Sprint 2.B demuestre:

- [ ] Tailscale instalado y nodo conectado a la tailnet.
- [ ] Red Docker `homelab-net` creada y operativa.
- [ ] Caddy ejecutándose como contenedor Docker con datos en `/srv/homelab/app_data/caddy`.
- [ ] `health.home.arpa` resoluble (desde un cliente perteneciente a la tailnet).
- [ ] `curl -f http://health.home.arpa` retorna "OK" (desde un cliente perteneciente a la tailnet).
- [ ] Segunda ejecución de `docker compose up -d` sin cambios (idempotencia).
- [ ] Documentación DNS en `docs/DNS.md`.
- [ ] Documentación Caddy en `docs/CADDY.md`.

## Riesgos y Mitigaciones

| Riesgo | Probabilidad | Impacto | Mitigación |
|---|---|---|---|
| Tailscale requiere intervención humana para autenticar el primer nodo | Alta | Bajo | Documentar como prerrequisito; no es un fallo del script |
| MagicDNS no soporta Split DNS para `home.arpa` | Baja | Alto | Fallback: agregar registros en Tailscale Extra Records; evaluar CoreDNS si falla |
| Puerto 80/443 ya ocupado en el host | Baja | Medio | FASE 0 de Sprint 2.B verifica puertos libres antes de desplegar |
| Conflictos de red Docker con la subred de la tailnet | Baja | Alto | Especificar subnet para `homelab-net` si hay conflicto |
| Caddy no arranca por Caddyfile inválido | Media | Medio | Validar Caddyfile con `caddy validate` antes de desplegar |
| IP Tailscale cambia y rompe DNS | Baja | Medio | Tailscale asigna IPs estables; documentar cómo actualizarlas si cambian |
| Tailscale fuera de servicio | Baja | Crítico | Mantener acceso SSH directo en la LAN como mecanismo de recuperación fallback |

## Trazabilidad

- **Phases:** Phase 2 (Storage — persistencia Caddy), Phase 3 (Networking — Tailscale, DNS, Caddy), Phase 5 (Platform Services — red compartida)
- **Sprint:** Sprint 2
- **Tareas:** HL-0015, HL-0016, HL-0017, HL-0018, HL-0019
- **ADR relacionados:** ADR-001, ADR-002, ADR-006, ADR-007, ADR-008
