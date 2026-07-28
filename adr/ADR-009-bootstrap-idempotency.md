# ADR-009 — Estrategia de idempotencia del Bootstrap

## Estado

Aceptada.

## Fecha

2026-07-28.

## Contexto

El proyecto requiere un script de bootstrap (`bootstrap/bootstrap.sh`) capaz de convertir un Ubuntu Server LTS recién instalado en una plataforma lista para servicios containerizados. Este script será ejecutado por agentes de IA y operadores humanos, posiblemente múltiples veces sobre el mismo host.

## Problema

Un script que solo instala y no valida puede dejar el sistema en un estado inconsistente si una ejecución anterior fue parcial o si el entorno ya fue configurado manualmente. Necesitamos garantizar que el bootstrap sea seguro de reejecutar en cualquier momento.

## Decisión

Se adopta un **enfoque mixto de idempotencia** basado en tres principios:

1. **El gestor de paquetes (`apt`) administra el estado de instalación.** Se confía en que `apt-get install` es idempotente por diseño.
2. **El Bootstrap valida el estado funcional esperado** después de cada operación mediante verificaciones explícitas.
3. **El Bootstrap nunca considera una operación exitosa únicamente porque el comando terminó sin error;** la operación solo se considera completada cuando las verificaciones posteriores demuestran que el estado objetivo fue alcanzado.

Este principio se implementa mediante tres niveles de funciones:

| Función | Responsabilidad |
|---|---|
| `ensure_package_installed <pkg>` | Delega en `apt`; instala si falta, no falla si ya existe |
| `ensure_service_running <svc>` | Verifica que el servicio está activo y habilitado; lo activa si no lo está |
| `validate_component <name>` | Verifica el estado funcional esperado del componente; aborta con error claro si falla |

## Alternativas consideradas

- **Solo confiar en apt:** Muy simple pero no valida que el sistema quedó como se esperaba. Una instalación incompleta podría pasar inadvertida.
- **Todo mediante verificaciones manuales en Bash:** Control absoluto pero demasiado código, fácil de introducir errores y duplica lógica que ya resuelve apt.

## Consecuencias

- El bootstrap puede ejecutarse múltiples veces de forma segura.
- Cada componente tiene una validación explícita que confirma su estado funcional.
- El patrón `ensure → validate` es reutilizable para futuros scripts (Caddy, Tailscale, Restic).

## Riesgos

- Las validaciones pueden dar falsos positivos si los comandos de verificación no son suficientemente estrictos.
- El mantenimiento de las funciones `validate_component` requiere actualización cuando cambian las versiones de software.

## Referencias relacionadas

- [ADR-002 — Docker First](ADR-002-docker-first.md)
- [ADR-003 — Ubuntu Server LTS](ADR-003-ubuntu-server.md)
- [Sprint-1-Plan.md](../.ai/implementation/Sprint-1-Plan.md)
