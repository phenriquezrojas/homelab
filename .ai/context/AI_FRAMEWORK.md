# AI Development Framework — Homelab

Este documento establece el marco operativo agnóstico de proveedor y modelo para cualquier agente de Inteligencia Artificial o herramienta automatizada que contribuya a este repositorio.

## 1. Principio Agnóstico

Cualquier agente (sin importar su proveedor, arquitectura o interfaz) debe operar bajo las mismas reglas y contratos definidos en este repositorio. Ningún agente asume instrucciones o sesgos propios de su modelo que contradigan el contenido del repositorio.

## 2. Jerarquía de Autoridad

1. **`PROJECT_CONSTITUTION.md`**: Regla máxima e inquebrantable.
2. **Architecture Decision Records (`adr/`)**: Decisiones técnicas aceptadas.
3. **Contexto del Proyecto (`.ai/context/`)**: Marco operativo, stack y convenciones.
4. **Master Plan y Tareas (`.ai/implementation/`, `.ai/tasks/`)**: Alcance y trazabilidad.

Si un agente detecta una ambigüedad o conflicto entre sus instrucciones y la jerarquía del repositorio, debe detener la ejecución y solicitar aclaración o proponer un ADR.

## 3. Reglas de Operación del Agente

### A. Roles y Skills (Obligatorio)
- El proyecto utiliza roles de ingeniería explícitos ubicados en `.ai/skills/`.
- Antes de iniciar una tarea o revisión, el agente DEBE identificar qué skill aplica (Architecture, Implementation, Operations, Review) y leer el archivo `SKILL.md` correspondiente.
- El agente debe actuar estrictamente dentro de las responsabilidades y límites de dicho rol.

### B. Inspección Antes de Modificación
- Ningún agente presupone la existencia de archivos, código o configuraciones.
- Se debe inspeccionar la fuente de verdad mediante lectura directa antes de proponer o aplicar cambios.

### C. Mínimo Alcance y Modificación Reversible
- Cada cambio debe limitarse estrictamente al alcance del Sprint y Tarea asignados.
- Todo cambio debe ser atómico y reversible.

### C. Gestión de Secretos y Datos
- **Prohibición absoluta:** Ningún agente escribirá o incluirá credenciales, tokens, llaves privadas o datos sensibles en ningún archivo del repositorio.
- Las configuraciones deben usar variables de entorno externalizadas (`.env` no versionado) o Docker secrets.

### D. Documentación Paralela
- Código y documentación se actualizan en la misma unidad de cambio.
- Cada tarea completada debe reflejar su estado en su correspondiente archivo en `.ai/tasks/`.

## 4. Ciclo de Vida del Trabajo del Agente

A partir del Sprint 0, todo sprint sigue un ciclo bifásico. El agente adapta su comportamiento según la fase activa:

### Fase A — Diseño (Sprint x.A)

1. **Lectura de Contexto:** Consultar `.ai/PROJECT_CONSTITUTION.md`, `.ai/context/` y el plan maestro.
2. **Colaboración en Diseño:** Trabajar con el humano para producir los artefactos de diseño: Sprint Specification, Implementation Plan, Tasks con criterios verificables, criterios de aceptación y checklist de revisión.
3. **Restricción:** El agente **no ejecuta cambios de código, configuración ni infraestructura** durante esta fase. Solo produce artefactos documentales de diseño.
4. **Gate de Diseño:** La Fase A concluye cuando el humano aprueba explícitamente todos los artefactos producidos.

### Fase B — Implementación (Sprint x.B)

1. **Verificación de Aprobación:** Confirmar que los artefactos de la Fase A están aprobados y la tarea está autorizada para ejecución.
2. **Ejecución:** Aplicar cambios con alcance mínimo, siguiendo estrictamente lo definido en los artefactos de la Fase A.
3. **Validación:** Comprobar sintaxis, pruebas existentes o consistencia documental.
4. **Cierre:** Actualizar la metadata de la tarea, el sprint y el changelog.

### Retorno a Fase A

Si durante la Fase B se descubre que el diseño es insuficiente o incorrecto, el agente debe detenerse y retornar a la Fase A para actualizar los artefactos antes de continuar. No se improvisan soluciones fuera del diseño aprobado.

---
*Este marco aplica a todos los agentes en el proyecto Homelab desde la finalización del Sprint -0.6. La metodología bifásica aplica a partir del Sprint 0.*
