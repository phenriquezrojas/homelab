# Gestión de Secretos

Este documento describe el mecanismo operativo para la inyección de secretos (contraseñas, tokens, llaves) en los servicios de Homelab, asegurando que ninguna credencial se filtre al repositorio de código.

## Principio de Aislamiento

Todos los secretos se mantienen estrictamente fuera del sistema de control de versiones. El archivo `.gitignore` en la raíz del repositorio está configurado para bloquear cualquier archivo que coincida con el patrón `.env*`.

## Ubicación Física

Físicamente en el host, los secretos sensibles o globales pueden almacenarse bajo `/srv/homelab/secrets`, directorio que cuenta con permisos restrictivos (`700`, propietario `root`). Sin embargo, el mecanismo primario de inyección es a nivel de entorno de Docker Compose.

## Inyección en Docker Compose

Para inyectar secretos en un servicio (ej. `services/immich/docker-compose.yml`), se utilizará un archivo `.env` ubicado en el mismo directorio que el archivo Compose, el cual será consumido automáticamente por Docker.

**Procedimiento:**
1. Crear el archivo `.env` en la ruta del servicio (ej. `services/immich/.env`).
2. (Opcional pero recomendado) Proveer un archivo `.env.example` versionado con valores ficticios para documentar las variables necesarias.
3. Docker Compose cargará automáticamente el archivo `.env` durante el despliegue.

Este mecanismo simple cumple con los requisitos actuales de operación y recuperación. En caso de requerir mayor escalabilidad, se evaluará la adopción de un gestor de secretos avanzado (ej. SOPS, Bitwarden Secrets Manager).
