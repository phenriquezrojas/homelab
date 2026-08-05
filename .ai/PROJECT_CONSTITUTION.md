# Project Constitution

## Autoridad

Este documento es la autoridad operativa máxima del repositorio. Todo agente,
mantenedor o automatización debe respetarlo junto con los ADR aceptados. Si una
instrucción local contradice esta Constitución o un ADR, se debe detener el
cambio y solicitar una decisión explícita mediante ADR.

## Framework EOS (Evidence-based Operating System)

Este proyecto está gobernado por el **Evidence-based Operating System (EOS)**, un entorno diseñado para gestionar el progreso basado en evidencias empíricas y documentales.

El marco EOS se apoya en los siguientes pilares:
- **Gestión por Evidencia**: Toda transición de estado requiere evidencia explícita versionada.
- **EOS Runtime Engine**: El agente actúa de manera proactiva infiriendo el siguiente paso del ciclo de vida (`FEATURE_LIFECYCLE.md`) a partir del estado documental, eliminando la necesidad de que el usuario dirija el flujo constantemente.
- **Comandos del Motor**: El entorno responde a `EOS STATUS` (evaluación de estado actual basada en evidencia) y `EOS CONTINUE` (determinación de la siguiente transición válida obligatoria).

## Principios permanentes

1. El repositorio es la única fuente de verdad para intención, configuración, procedimientos y decisiones.
2. El hardware es reemplazable; los datos son el activo principal e irremplazable.
3. Todo debe poder reconstruirse desde cero con el repositorio y secretos gestionados externamente.
4. Docker First es el modelo de ejecución predeterminado.
5. La persistencia se mantiene bajo `/srv` y fuera del contenido versionado.
6. Ningún cambio manual permanente sustituye configuración, procedimiento o evidencia versionados.
7. Toda decisión arquitectónica importante requiere un ADR; ningún agente cambia arquitectura sin uno.
8. Disaster Recovery es obligatorio: respaldo y restauración se diseñan, documentan y verifican juntos.
9. Documentación y código tienen el mismo valor de ingeniería y se actualizan en el mismo cambio.
10. Todo sprint termina en un estado funcional para su alcance, con estado, deuda y siguiente paso explícitos. Es OBLIGATORIO consultar y actualizar todos los documentos listados en el índice oficial `.ai/eos/STATE_INDEX.md` en el mismo commit que cierra el sprint.
11. Se mantiene la simplicidad y se evitan dependencias innecesarias.
12. Todo cambio debe ser reversible o describir claramente su reversión antes de ejecutarse.
13. Todo Sprint se ejecuta en dos fases secuenciales: Diseño (A) e Implementación (B). La Fase B se estructura explícitamente en dos subfases: Implementación Declarativa en Repositorio (B1) y Despliegue/Validación Operacional en Host (B2). No se inicia la Fase B sin aprobación explícita de los artefactos de la Fase A. Aplica desde el Sprint 0.
14. Ningún agente posee autoridad para aprobar un Sprint. Todo Sprint implementado pasa automáticamente al estado: In Review, hasta recibir aprobación explícita del Owner.
    Mientras un Sprint esté In Review está prohibido:
    - marcar Completed
    - actualizar CURRENT_STATE
    - actualizar ROADMAP
    - actualizar CHANGELOG
    - registrar el Sprint como Approved
    - ejecutar el commit de cierre
15. Ningún Sprint podrá implementar capacidades del Homelab sin que exista previamente la capacidad correspondiente en el Runtime para instalarlas, configurarlas, validarlas y repararlas. El Runtime crece primero.

## Reglas de trabajo

- Antes de actuar, leer `.ai/context/`, el plan maestro y las tareas relacionadas.
- Respetar los límites del sprint activo; no adelantar infraestructura ni servicios.
- Preferir referencias a documentos canónicos antes que duplicar contenido.
- Mantener secretos, datos operativos y reportes sensibles fuera de Git.
- Actualizar estado, tarea, sprint, plan y changelog cuando el cambio lo requiera. Al cerrar un sprint o fase, es OBLIGATORIO incluir en el commit todos los archivos referenciados en `.ai/eos/STATE_INDEX.md`.
- `CURRENT_STATE.md` especifica la subfase exacta en ejecución (ej. `Sprint 2.A`, `Sprint 2.B1`, `Sprint 2.B2`).
- Toda validación operacional (subfase B2) exige registrar un artefacto obligatorio de evidencia versionado bajo `docs/evidence/` (incluyendo fecha, sprint, host objetivo, comandos ejecutados, resultado obtenido, evidencias relevantes y conclusión).
- Tratar los ADR aceptados como registro histórico: se sustituyen con un ADR nuevo, no se reescriben.
- En Fase A (Diseño), no se escribe código, configuración ni se modifica infraestructura; solo se producen artefactos de diseño.
- En Fase B1 (Implementación Declarativa), se crea y versiona en Git la solución aprobada.
- En Fase B2 (Despliegue y Validación Operacional), se despliega en el host objetivo y se certifica la evidencia empírica. Si el diseño resulta insuficiente, se retorna a Fase A antes de continuar.

## Navegación mínima

Un agente que se incorpora debe leer, en orden:

1. `.ai/PROJECT_CONSTITUTION.md`
2. `.ai/context/`
3. `.ai/skills/`
4. `.ai/implementation/MASTER_PLAN.md`
5. `.ai/tasks/`

## Regla de Gate Review

La revisión de un Sprint tiene como objetivo verificar el cumplimiento del contrato aprobado durante la Fase A.

Durante el Gate Review:

- No se amplía el alcance del Sprint.
- No se introducen nuevos requisitos.
- No se realizan refactorizaciones por preferencia.
- Las mejoras que no incumplan el contrato aprobado deberán registrarse como Deuda Técnica, ADR o backlog.

Únicamente los incumplimientos del contrato aprobado podrán bloquear el cierre del Sprint.

## Comportamiento de Gate

Todo agente que ejecute un Gate (revisión, auditoría o cierre) debe seguir este protocolo:

1. **Detectar** el hallazgo.
2. **Verificar** contra el contrato aprobado y la matriz de Tareas (HL-XXXX).
3. **Diferenciar** explícitamente entre:
   - **Cumplimiento Declarativo (Repositorio):** Archivos, compose, Caddyfile, docs y scripts de la solución.
   - **Cumplimiento Operativo (Host):** Evidencia empírica de ejecución runtime en el servidor.
4. **Clasificar** según estas categorías:

| Categoría | Significado | Efecto |
|---|---|---|
| **Incumplimiento del contrato** | El artefacto viola una condición del contrato aprobado en Fase A. | Bloquea el Gate. Debe corregirse. |
| **Ambigüedad del diseño** | El contrato no especifica una condición y el diseño admite más de una interpretación. | No bloquea. Requiere aclaración del Owner. |
| **Recomendación de mejora** | El artefacto cumple el contrato pero podría mejorarse. | No bloquea. Se registra como Deuda Técnica. |

El agente **no propone arquitectura** durante un Gate. Todo informe de Gate debe incluir trazabilidad explícita de Tasks, una matriz Repo vs Host y la lista de evidencias pendientes para el cierre definitivo.
