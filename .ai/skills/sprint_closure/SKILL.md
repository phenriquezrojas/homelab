---
name: sprint_closure
description: Procedimiento obligatorio para cerrar un Sprint o Fase en el entorno EOS.
---

# Procedimiento de Cierre de Sprint

Este proyecto utiliza un modelo estricto de Engineering Operating System (EOS).

Cuando se te pida cerrar un Sprint (ya sea Fase A o Fase B) o cuando crees el artefacto final de un Sprint (ej. `Sprint-0.B.md`), **DEBES OBLIGATORIAMENTE** consultar el archivo `.ai/eos/STATE_INDEX.md`.

Ese índice contiene la lista oficial de documentos de estado global del proyecto. Debes abrir todos los documentos listados en él, actualizarlos de acuerdo al avance del proyecto, e incluirlos en el mismo commit que cierra el Sprint.

Si omites alguno de estos archivos al intentar subir el código (`git commit`), el hook de `pre-commit` implementado en este repositorio bloqueará mecánicamente tu intento. Evita errores leyendo el `STATE_INDEX.md` y preparando bien tus commits.
