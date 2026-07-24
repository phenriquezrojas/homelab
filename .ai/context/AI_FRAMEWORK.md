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

### A. Inspección Antes de Modificación
- Ningún agente presupone la existencia de archivos, código o configuraciones.
- Se debe inspeccionar la fuente de verdad mediante lectura directa antes de proponer o aplicar cambios.

### B. Mínimo Alcance y Modificación Reversible
- Cada cambio debe limitarse estrictamente al alcance del Sprint y Tarea asignados.
- Todo cambio debe ser atómico y reversible.

### C. Gestión de Secretos y Datos
- **Prohibición absoluta:** Ningún agente escribirá o incluirá credenciales, tokens, llaves privadas o datos sensibles en ningún archivo del repositorio.
- Las configuraciones deben usar variables de entorno externalizadas (`.env` no versionado) o Docker secrets.

### D. Documentación Paralela
- Código y documentación se actualizan en la misma unidad de cambio.
- Cada tarea completada debe reflejar su estado en su correspondiente archivo en `.ai/tasks/`.

## 4. Ciclo de Vida del Trabajo del Agente

1. **Lectura de Contexto:** Consultar `.ai/PROJECT_CONSTITUTION.md` y `.ai/context/`.
2. **Verificación de Tarea:** Confirmar que la tarea está en estado de ejecución autorizada.
3. **Ejecución:** Aplicar cambios con alcance mínimo.
4. **Validación:** Comprobar sintaxis, pruebas existentes o consistencia documental.
5. **Cierre:** Actualizar la metadata de la tarea, el sprint y el changelog.

---
*Este marco aplica a todos los agentes en el proyecto Homelab desde la finalización del Sprint -0.6.*
