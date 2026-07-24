# Skills

Las skills son roles de ingeniería que los agentes (Antigravity, Cursor, Codex, etc.) pueden asumir.

## Estructura Unificada

Para garantizar que todos los agentes descubran e inyecten correctamente las reglas sin depender de archivos ocultos duplicados, todo skill debe seguir estrictamente este formato:

1. **Directorio:** Una carpeta con el nombre del skill (ej. `architecture/`).
2. **Archivo:** Dentro de la carpeta, un archivo llamado exactamente `SKILL.md`.
3. **Metadatos:** El archivo debe comenzar con un bloque YAML Frontmatter indicando `name` y `description`.

```yaml
---
name: nombre-del-skill
description: Descripción breve de su propósito.
---
```

## Reglas de Mantenimiento

- **Añadir un skill:** Crea la carpeta y el `SKILL.md` con el YAML. Actualiza la tabla inferior.
- **Modificar un skill:** Edita directamente el `SKILL.md`. Todos los agentes leerán la versión actualizada.
- **Borrar un skill:** Elimina la carpeta completa y quítalo de la tabla inferior.

## Roles Disponibles

| Skill | Propósito |
| --- | --- |
| [Architecture](architecture/SKILL.md) | Proteger coherencia arquitectónica y ADR. |
| [Implementation](implementation/SKILL.md) | Convertir planes aprobados en cambios acotados y verificables. |
| [Review](review/SKILL.md) | Detectar riesgos, inconsistencias y deuda antes de aprobar. |
| [Operations](operations/SKILL.md) | Mantener evidencia operativa, runbooks y recuperación. |
