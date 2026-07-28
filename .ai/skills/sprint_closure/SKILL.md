---
name: sprint_closure
description: Procedimiento obligatorio para cerrar un Sprint o Fase en el entorno EOS.
---

# Procedimiento de Cierre de Sprint

Este proyecto utiliza un modelo estricto de Engineering Operating System (EOS).

Cuando se te pida cerrar un Sprint (ya sea Fase A o Fase B) o cuando crees el artefacto final de un Sprint (ej. `Sprint-0.B.md`), **DEBES OBLIGATORIAMENTE** actualizar los siguientes 5 archivos en el mismo commit para mantener el estado global alineado:

1. `CURRENT_STATE.md` (Actualizar la fase y estado general)
2. `ROADMAP.md` (Marcar el sprint actual como `Completed` y el siguiente como `Planned`)
3. `docs/SPRINT_LOG.md` (Añadir la entrada histórica narrativa del sprint)
4. `CHANGELOG.md` (Registrar los cambios técnicos o documentales logrados)
5. `.ai/implementation/MASTER_PLAN.md` (Marcar los checks y cambiar el estado de la Fase si corresponde)

Si omites alguno de estos archivos al intentar subir el código (`git commit`), el hook de `pre-commit` implementado en este repositorio bloqueará mecánicamente tu intento. Evita errores leyendo este documento y preparando bien tus commits.
