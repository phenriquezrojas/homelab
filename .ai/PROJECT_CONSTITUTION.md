# Project Constitution

## Autoridad

Este documento es la autoridad operativa máxima del repositorio. Todo agente,
mantenedor o automatización debe respetarlo junto con los ADR aceptados. Si una
instrucción local contradice esta Constitución o un ADR, se debe detener el
cambio y solicitar una decisión explícita mediante ADR.

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
13. Todo Sprint se ejecuta en dos fases secuenciales: Diseño (A) e Implementación (B). No se inicia la Fase B sin aprobación explícita de los artefactos de la Fase A. Aplica desde el Sprint 0.

## Reglas de trabajo

- Antes de actuar, leer `.ai/context/`, el plan maestro y las tareas relacionadas.
- Respetar los límites del sprint activo; no adelantar infraestructura ni servicios.
- Preferir referencias a documentos canónicos antes que duplicar contenido.
- Mantener secretos, datos operativos y reportes sensibles fuera de Git.
- Actualizar estado, tarea, sprint, plan y changelog cuando el cambio lo requiera. Al cerrar un sprint o fase, es OBLIGATORIO incluir en el commit todos los archivos referenciados en `.ai/eos/STATE_INDEX.md`.
- Tratar los ADR aceptados como registro histórico: se sustituyen con un ADR nuevo, no se reescriben.
- En Fase A (Diseño), no se escribe código, configuración ni se modifica infraestructura; solo se producen artefactos de diseño.
- En Fase B (Implementación), se ejecuta estrictamente lo aprobado en la Fase A. Si el diseño resulta insuficiente, se retorna a Fase A antes de continuar.

## Navegación mínima

Un agente que se incorpora debe leer, en orden:

1. `.ai/PROJECT_CONSTITUTION.md`
2. `.ai/context/`
3. `.ai/skills/`
4. `.ai/implementation/MASTER_PLAN.md`
5. `.ai/tasks/`
