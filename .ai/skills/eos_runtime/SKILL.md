---
name: eos_runtime
description: Implementa el EOS Runtime Engine que permite ejecutar los comandos "EOS STATUS" y "EOS CONTINUE". Este motor asegura que el agente guíe el progreso del proyecto Homelab basándose en la evidencia documental sin que el usuario tenga que indicarle el siguiente paso.
---

# EOS Runtime Engine

A partir de esta versión, EOS incorpora un Runtime Engine cuyo objetivo es eliminar la necesidad de que el usuario recuerde el flujo del framework.

El agente nunca debe pedir al usuario que indique cuál es el siguiente paso cuando pueda determinarlo a partir del estado del proyecto.

## Comando: EOS STATUS

Objetivo:
Mostrar el estado actual del proyecto utilizando exclusivamente la evidencia existente.

Antes de responder, el agente debe inspeccionar como mínimo:
- CURRENT_STATE.md
- STATE_INDEX.md
- PROJECT_CONSTITUTION.md
- FEATURE_LIFECYCLE.md
- Sprint actualmente activo
- Implementation Plan asociado
- Tasks relacionadas
- Último Gate Review existente
- Evidencias bajo docs/evidence/
- Último commit relevante (cuando corresponda)

La respuesta debe incluir únicamente:
• Sprint actual
• Fase/Subfase actual
• Estado del Sprint
• Último Gate ejecutado
• Evidencias registradas
• Trabajo pendiente
• Siguiente transición permitida por el Lifecycle (ej. Iniciar Gate Review, Crear Implementation Plan, etc.)

Regla estricta:
El comando EOS STATUS es exclusivamente de lectura. No modifica ningún archivo ni propone arquitecturas.

## Comando: EOS CONTINUE

Objetivo:
Determinar e informar al usuario cuál es el siguiente paso lógico y válido según el Lifecycle.

Flujo de ejecución:
1. Análisis del estado: Igual que en EOS STATUS, el agente reconstruye mentalmente el estado del proyecto.
2. Verificación de completitud: Verifica si la transición anterior fue realmente cerrada (ej. si requiere Gate y si el Gate fue aprobado).
3. Verificación de evidencia: Si se requiere evidencia operativa (ej. Sprint B2), verifica su existencia en docs/evidence/.
4. Verificación de inconsistencias: Si un documento dice estar en "Fase C" pero falta el "Gate de Fase B", el agente detecta la inconsistencia.
5. Determinación: Basado en FEATURE_LIFECYCLE.md, determina la única transición válida siguiente.
6. Propuesta: Informa al usuario la acción requerida.

Reglas estrictas para EOS CONTINUE:
- Nunca debe saltar fases del Lifecycle.
- Si detecta inconsistencias documentales, detiene el proceso y pide al usuario autorización para corregirlas.
- Si el siguiente paso es un Gate Review, propone iniciarlo.
- Si el siguiente paso es un Commit de cierre, dicta el formato del commit.

Formato de salida esperado de EOS CONTINUE:
Estado actual: [Resumen conciso de dónde estamos]
Siguiente transición EOS: [Acción específica que dicta el Lifecycle]
Acción requerida: [Pregunta binaria o instrucción directa para el usuario, ej. "¿Iniciamos el Gate Review ahora?"]
