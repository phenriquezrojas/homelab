# Estado actual

**Fase:** Sprint -0.7 — Engineering Operating System (EOS)  
**Fecha de actualización:** 2026-07-22  
**Estado general:** Base documental y la capa de conocimiento permanente (EOS) completados; herramientas de inventario en modo solo lectura validadas. Implementación técnica de infraestructura no iniciada.

## Completado

- **Sprint -1 (Fundaciones):** Visión, charter, roadmap, convenciones y ADR-001 a ADR-008 aprobados.
- **Sprint -0.5 (Host Inventory):** Herramienta de inventario solo lectura `scripts/inventory.sh` implementada y validada en modo local.
- **Sprint -0.5.1 (Correcciones del PR):** Observaciones del Pull Request de inventario verificadas e integradas (tarea [HL-0003](.ai/tasks/HL-0003.md)).
- **Sprint -0.6 (AI Development Framework):** Diseñado el marco agnóstico de modelo para agentes en `.ai/context/AI_FRAMEWORK.md` (tarea [HL-0004](.ai/tasks/HL-0004.md)).
- **Sprint -0.7 (EOS):** Estructura del Engineering Operating System (`.ai/`) creada con principios, contexto, plan maestro, tareas y plantillas transversales.

## Planificado / Pendiente

- **Sprint 0 (Preparación de Bootstrap):** Definir modelo de secretos, topología de red, estructura física de `/srv` y criterios de acceptance del bootstrap.

## Restricciones vigentes

No se añadirán secretos, servicios funcionales, scripts de automatización destructivos, Docker Compose ni GitHub Actions durante esta fase. `bootstrap.sh` y `restore.sh` siguen siendo marcadores de posición.

## Próximo hito

Iniciar la preparación del **Sprint 0** para avanzar hacia la configuración e infraestructura base del servidor.
