# Bootstrap del Host

El aprovisionamiento inicial del servidor homelab está completamente automatizado y validado mediante un script idempotente.

## Prerrequisitos

1. Ubuntu Server LTS instalado.
2. Acceso SSH o consola con privilegios de `sudo`.
3. Conectividad a Internet.

> [!NOTE]
> No es necesario crear los directorios persistentes manualmente. El propio script se encarga de estructurarlos bajo `/srv/homelab` si no existen.

## Uso

El script puede copiarse y ejecutarse vía SSH desde cualquier máquina de gestión:

```bash
scp bootstrap/bootstrap.sh usuario@homelab:/tmp/bootstrap.sh
ssh usuario@homelab 'sudo SUDO_USER=$USER /tmp/bootstrap.sh'
```

## Idempotencia

El script está diseñado bajo el **principio de idempotencia mixta** ([ADR-009](../adr/ADR-009-bootstrap-idempotency.md)). Puede ejecutarse tantas veces como se desee de forma segura. En pasadas sucesivas donde el estado objetivo ya esté alcanzado, el script simplemente valida que los componentes estén funcionales y finaliza con éxito sin generar cambios destructivos.

## Registro

Toda la salida del script, incluyendo errores o advertencias, se duplica automáticamente en el archivo `~/homelab-bootstrap.log` del usuario que lo ejecuta.

## Reversibilidad

Si se desea revertir el aprovisionamiento de este host para devolverlo a su estado de fábrica (o antes de instalar Docker), se deben ejecutar los siguientes pasos manuales en el servidor:

```bash
# 1. Detener y desinstalar Docker y sus dependencias
sudo systemctl stop docker.socket docker.service
sudo apt-get purge -y docker-ce docker-ce-cli containerd.io docker-buildx-plugin docker-compose-plugin docker-ce-rootless-extras
sudo apt-get autoremove -y

# 2. Eliminar el repositorio y clave GPG de Docker
sudo rm -f /etc/apt/sources.list.d/docker.list
sudo rm -f /etc/apt/keyrings/docker.asc
sudo apt-get update

# 3. Remover al usuario del grupo docker (opcional)
sudo gpasswd -d $USER docker

# 4. Eliminar directorios (Opcional, DESTRUCTIVO)
# ATENCIÓN: Esto eliminará todos los datos, backups y secretos almacenados
# sudo rm -rf /srv/homelab
```
