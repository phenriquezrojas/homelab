---
name: repository_consistency
description: Principios transversales de consistencia documental en el EOS.
---

# Consistencia del Repositorio (EOS)

El repositorio es un sistema interconectado. Cuando realices cambios técnicos o documentales, aplica siempre el principio de consistencia transversal:

1. **Tareas (Tasks):** Si completas una tarea (ej. pasas el estado de `HL-XXXX.md` a `Completed`), verifica si eso te permite avanzar o cerrar el Sprint asociado.
2. **Sprints:** Nunca marques un Sprint como `Completed` si las tareas referenciadas dentro de él aún están abiertas o marcadas como `Planned`.
3. **ADRs (Architecture Decision Records):** Si implementas un cambio que altera la arquitectura, no solo cambies el código. Redacta un ADR, solicita aprobación, y luego asegúrate de que el `.ai/implementation/MASTER_PLAN.md` refleje ese nuevo hito.
4. **Pruebas y Evidencia:** Toda tarea completada debe llevar en su archivo Markdown una sección de `## Evidence` con la salida real de los comandos que prueban que fue completada exitosamente.

Actúa siempre como un Arquitecto/Mantenedor global, no solo como un ejecutor de scripts.
