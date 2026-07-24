---
Estado: Approved
Versión: 1.0
Autor: AI Agent
Relacionado: Sprint 0.A, Sprint 0.B, ADR-001, ADR-006, ADR-008
---

# Implementation Plan — Bootstrap Base (Sprint 0)

## Objetivo

Materializar las decisiones aprobadas durante Sprint 0.A para preparar la plataforma antes de desplegar servicios.

## Prerrequisitos

Antes de ejecutar este plan en la Fase B, se debe verificar:
- [ ] Ubuntu Server instalado.
- [ ] Repositorio clonado en el host.
- [ ] Usuario con privilegios de `sudo`.
- [ ] Conexión a Internet disponible.
- [ ] Sprint 0.A (Diseño) formalmente aprobado.
- [ ] ADR vigentes (001, 006, 008) leídos y comprendidos.

## Alcance

- Crear la estructura de directorios en el host.
- Definir el procedimiento de paso de secretos vía archivos planos fuera del repositorio.
- Documentar la arquitectura de red (Caddy + Tailscale).
- **Excluido:** Instalar servicios funcionales finales (ej. Immich, PostgreSQL).

## Resultados Esperados

### 1. Estructura de Persistencia
- **Resultado esperado:** Estructura persistente base creada bajo `/srv`.
- **Criterio de aceptación:** Los directorios `app_data`, `secrets` y `backups` existen bajo `/srv/homelab`. Los permisos cumplen la política estricta.
- **Implementación:** Ver [HL-0006](../tasks/HL-0006.md).

### 2. Gestión de Secretos
- **Resultado esperado:** Convención de secretos locales protegidos y excluidos del control de versiones.
- **Criterio de aceptación:** Archivos `.env.*` ignorados globalmente en Git. Documentación sobre inyección en Compose elaborada.
- **Implementación:** Ver [HL-0007](../tasks/HL-0007.md).

### 3. Arquitectura de Red
- **Resultado esperado:** Topología de red interna documentada y lista para configurar.
- **Criterio de aceptación:** Flujo de peticiones (Tailscale → Caddy → Contenedor) y resolución DNS (`home.arpa`) detallados.
- **Implementación:** Ver [HL-0008](../tasks/HL-0008.md).

## Validación

| Validación | Evidencia |
| --- | --- |
| Directorios creados | Salida de `tree /srv/homelab` |
| Permisos de seguridad | Salida de `ls -ld /srv/homelab/secrets` |
| Archivos `.env` ignorados | Salida limpia de `git status` tras crear un `.env` de prueba |
| Red documentada | Revisión documental aprobada |

## Criterios de Éxito (Success Criteria)

Este Plan se considera ejecutado exitosamente cuando:
- [ ] Existe `/srv/homelab`.
- [ ] Existe `/srv/homelab/app_data`.
- [ ] Existe `/srv/homelab/backups`.
- [ ] Existe `/srv/homelab/secrets`.
- [ ] `/srv/homelab/secrets` tiene permisos correctos (solo lectura/escritura para root/admin).
- [ ] Git ignora sistemáticamente los archivos `.env`.
- [ ] La topología de red base está documentada teóricamente.
- [ ] No se instalaron servicios funcionales.

## Riesgos

- **Permisos incorrectos:** Exposición de secretos locales a usuarios no privilegiados.
- **Error de montaje:** Contenedores escribiendo en rutas efímeras del host si `/srv` no se mapea bien.
- **Confusión entre repo y persistencia:** Guardar archivos `docker-compose.yml` en `/srv` en lugar del repo versionado.
- **Rollback incompleto:** Dejar credenciales u orfandad de carpetas tras revertir un cambio.

## Rollback

En caso de fallo crítico en la implementación:
- **Qué eliminar:** El directorio completo creado. Ejecutar `sudo rm -rf /srv/homelab`.
- **Qué conservar:** El clon del repositorio en el directorio del usuario.
- **Cómo validar limpieza:** La ejecución de `ls /srv/homelab` debe retornar error de archivo inexistente.

## Trazabilidad

- **Fase:** Phase 1 — Bootstrap
- **Sprint:** Sprint 0
- **Tareas:** HL-0006, HL-0007, HL-0008
