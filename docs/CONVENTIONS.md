# Convenciones del repositorio

Estas convenciones reducen ambigüedad y preservan trazabilidad. Se aplican desde Sprint -1.

## Conventional Commits

Formato: `tipo(alcance opcional): descripción breve en imperativo`.

Tipos permitidos: `feat`, `fix`, `docs`, `chore`, `refactor`, `test`, `build`, `ci` y `security`. El alcance identifica un área, por ejemplo `docs`, `adr`, `backup` o `caddy`. La descripción comienza en minúscula, no termina con punto y describe un cambio.

Ejemplos:

```text
docs(adr): documenta la persistencia bajo /srv
chore(repo): añade reglas para archivos locales
feat(backup): incorpora configuración de restic
```

Los cambios incompatibles se marcan con `!` tras tipo o alcance y se explican en el cuerpo. Los commits deben ser atómicos, revisables y no incluir secretos ni datos generados.

## Semantic Versioning

Se usa `MAJOR.MINOR.PATCH`:

- **MAJOR:** cambio incompatible en contratos, procedimientos o estructura pública.
- **MINOR:** capacidad nueva compatible, como un servicio terminado.
- **PATCH:** corrección compatible de configuración, documentación o proceso.

Antes de `1.0.0`, `0.y.z` comunica evolución. Cada versión publicada debe etiquetarse y resumirse en `CHANGELOG.md`.

## Convenciones de nombres

- Directorios y archivos de configuración: minúsculas y `kebab-case`.
- Documentos de raíz y ADR: nombres establecidos en mayúsculas cuando corresponda; los restantes, `kebab-case.md`.
- ADR: `ADR-NNN-descripcion-corta.md`, con tres dígitos y descripción técnica concisa en inglés para estabilidad de rutas.
- Servicios: un directorio por servicio en `kebab-case`.
- Variables de entorno: `UPPER_SNAKE_CASE`; secretos fuera de Git.
- Dominios internos: subdominios en minúscula bajo `home.arpa`.

## Estilo Markdown

- Use ATX headings y un único H1.
- Deje líneas en blanco alrededor de títulos, listas, tablas y bloques de código.
- Use español para documentación operativa; tecnología, rutas, comandos y variables pueden permanecer en inglés.
- Use enlaces relativos para contenido interno y texto descriptivo para enlaces externos.
- Apunte a líneas de aproximadamente 100 caracteres cuando sea práctico.
- No incluya secretos, tokens, IPs privadas reales ni configuraciones generadas sin revisión.

## Organización de carpetas

| Ruta | Propósito |
| --- | --- |
| `adr/` | Decisiones de arquitectura inmutables. |
| `assets/` | Recursos estáticos de documentación. |
| `backup/` | Configuración declarativa y documentación de respaldos. |
| `bootstrap/` | Material declarativo para preparar el host. |
| `compose/` | Definiciones Docker Compose por stack. |
| `diagrams/` | Diagramas fuente y exportados. |
| `docs/` | Estándares y documentación transversal. |
| `restore/` | Material de restauración y validación. |
| `runbooks/` | Procedimientos de operación e incidencias. |
| `scripts/` | Automatizaciones versionadas y revisadas. |
| `services/` | Configuración específica de cada servicio. |
| `tests/` | Validaciones automatizadas o manuales reproducibles. |

La raíz contiene documentos de gobierno, metadatos y puntos de entrada explícitos. Cada incorporación debe ir en la carpeta más específica; evite duplicar configuración.
