# Reporte de Evidencia Operacional — Sprint 2.B2

## Metadata Obligatoria EOS

- **Fecha:** 2026-08-11
- **Sprint / Subfase:** Sprint 2.B2 — Despliegue y Validación Operacional (Host)
- **Host Objetivo:** `homelab` (Ubuntu Server LTS, Tailscale IP 100.119.17.36)
- **Ejecutor:** Antigravity (AI Agent)
- **Estado de Validación:**  EFECTUADO Y EXITOSO (100% Convergido)

---

## 1. Plan de Comandos a Ejecutar en Host

| # | Prueba / Validación | Comando a Ejecutar | Resultado Esperado | Estado |
|---|---|---|---|---|
| 1 | Compilación y Bootstrap | `sudo ./bootstrap/bootstrap.sh` | Compilación exitosa y binario `homelab` funcional |  EXITOSO |
| 2 | Plan de convergencia inicial | `sudo homelab plan` | Muestra los cambios necesarios ordenados por el DAG |  EXITOSO |
| 3 | Ejecución de convergencia | `sudo homelab converge` | Aplica mutaciones y reporta éxito |  EXITOSO |
| 4 | Resolución de conflictos (Puerto 80) | `curl -s -I http://localhost:80` | Responde con cabecera `Server: Caddy` redireccionando a HTTPS |  EXITOSO |
| 5 | Validación de MagicDNS | `tailscale status --json` | Detecta `"MagicDNSEnabled": true` |  EXITOSO |
| 6 | Prueba de Idempotencia | Segunda ejecución `sudo homelab plan` | Retorna `No operations needed` |  EXITOSO |

---

## 2. Registro de Evidencias Relevantes (Logs y Salidas)

### Compilación y Bootstrap (Fase 7 de bootstrap.sh)
```text
[INFO] 2026-08-11 02:42:19 - FASE 7: Compilación e instalación del Homelab Runtime Engine...
[INFO] 2026-08-11 02:42:19 - Go disponible: go version go1.18.1 linux/amd64
[INFO] 2026-08-11 02:42:19 - Compilando Homelab Runtime...
[INFO] 2026-08-11 02:42:20 - Instalando binario homelab en /usr/local/bin...
[INFO] 2026-08-11 02:42:20 - Asignando permisos de ejecución a los componentes...
[OK]   2026-08-11 02:42:20 - Binario homelab instalado y funcional.
```

### Ejecución de la Convergencia Final
```text
peter@desktop:~/homelab$ sudo homelab converge
No operations needed. System is in the desired state.
```

### Verificación de Firmas e Idempotencia
```text
peter@desktop:~/homelab$ sudo homelab plan
Execution Plan:
  No operations needed. System is in the desired state.
```

### Comprobación de Caddy (Puerto 80)
```text
peter@desktop:~/homelab$ curl -s -v -H 'Host: health.home.arpa' http://localhost:80
*   Trying 127.0.0.1:80...
* Connected to localhost (127.0.0.1) port 80 (#0)
> GET / HTTP/1.1
> Host: health.home.arpa
...
< HTTP/1.1 308 Permanent Redirect
< Connection: close
< Location: https://health.home.arpa/
< Server: Caddy
```

---

## 3. Conclusión Operacional

- **Resultado Consolidado:**  **CONVERGIDO Y ESTABLE.**
- **Veredicto de Cierre B2:** **APROBADO PARA GATE REVIEW.**
- **Logros:**
  - El motor Go (`homelab`) coordinó de forma autónoma el DAG de dependencias.
  - La resolución del conflicto de puerto 80 con Apache fue exitosa desplazando Apache a 8080 y robusteciendo la firma de `validate.sh` de Caddy.
  - La idempotencia del sistema se ha verificado empíricamente con planes vacíos tras la convergencia.
  - La integración de Tailscale y MagicDNS fue exitosamente validada con la salida del agente nativo.
