# Homelab Runtime Engine (v1.0)

El **Homelab Runtime Engine** es el único mecanismo de ejecución oficial del proyecto, responsable de converger el sistema hacia su estado deseado. Se implementa como un CLI (`homelab`) escrito en Go.

## 1. Arquitectura y Conceptos Base

El Runtime no conoce la implementación de ninguna tecnología, sino que coordina **capacidades** a través de **componentes** utilizando un modelo estricto basado en dependencias (DAG) y observación de estado (Observer). 
* Ver [Sprint-2-Runtime-Design.md](../.ai/implementation/Sprint-2-Runtime-Design.md) para el documento de diseño completo.
* Ver [ADR-010](../adr/ADR-010-homelab-runtime.md) para el origen de esta decisión arquitectónica.

## 2. Compilación e Instalación

La compilación y aprovisionamiento inicial se realiza automáticamente mediante `bootstrap/bootstrap.sh`.

```bash
sudo ./bootstrap/bootstrap.sh
```

El script se encargará de instalar Go, compilar el código fuente de `runtime/` e instalar el ejecutable en `/usr/local/bin/homelab`.

## 3. Comandos de la CLI

El comando principal que orquesta todo el proceso de observación, planificación y ejecución es `converge`.

### `homelab converge`
Evalúa el estado actual de todas las capacidades registradas, calcula el plan topológico de ejecución y ejecuta las operaciones necesarias (`install`, `configure`, `repair`) para que todas alcancen el estado `HEALTHY`. Si una operación falla, la convergencia se detiene de inmediato (Halt-on-Fail).

### `homelab plan`
Realiza el descubrimiento y la observación del estado actual, y muestra por pantalla las operaciones que se ejecutarían en un `converge`, pero sin mutar el sistema.

### `homelab validate`
Ejecuta la operación `validate` de todos los componentes y muestra un listado con el estado empírico de cada capacidad (ej. `ABSENT`, `INSTALLED`, `CONFIGURED`, `HEALTHY`, `FAILED`).

## 4. Registro y Componentes

El registro que dicta qué componentes deben ejecutarse y cuáles son sus dependencias se encuentra en `runtime/registry.yaml`.

Cada componente está alojado en `runtime/components/<provider_id>/` y debe contener exactamente cuatro scripts ejecutables:
* `install.sh`
* `configure.sh`
* `validate.sh`
* `repair.sh`

## 5. Casos de Intervención Humana (Fricciones)

El contrato impone que la ejecución de los componentes debe ser automatizable e inatendida, sin embargo existen excepciones documentadas:

- **Tailscale Authentication:** El primer despliegue de Tailscale requiere autenticación manual. El script `configure.sh` detendrá el plan (`homelab converge`) devolviendo un exit code > 0 si requiere intervención. El usuario debe autenticar manualmente (`sudo tailscale up`) y re-ejecutar `homelab converge`.
- **MagicDNS Secrets:** MagicDNS requiere que existan credenciales/sesiones activas u operadores manuales en el admin console de Tailscale para configurar los registros DNS.
