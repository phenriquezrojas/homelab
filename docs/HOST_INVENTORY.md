# Inventario del host

## Objetivo

`scripts/inventory.sh` genera una instantánea Markdown de solo lectura del host
objetivo. Su propósito es aportar evidencia para Sprint -0.5 y preparar el
diseño de Sprint 0 sin instalar software, cambiar configuración ni iniciar o
detener servicios.

El informe cubre sistema operativo, hardware, almacenamiento, red, acceso,
servicios, herramientas y hallazgos heurísticos. Los valores de variables de
entorno no se incluyen: solo se registran nombres de variables no sensibles
cuando están definidos.

## Cómo ejecutar

Ejecute el script desde la raíz del repositorio:

```bash
./scripts/inventory.sh
```

El resultado se crea en
`reports/HOST_INVENTORY_<hostname>_<fecha>.md`. El directorio `reports/` está
excluido de Git porque los informes pueden contener detalles de infraestructura
local. Para obtener datos de hardware, discos y servicios con mejor cobertura,
ejecútelo con los permisos de lectura apropiados para el host; el script no
eleva privilegios por sí mismo.

Si se ejecuta sin privilegios elevados, el informe muestra una advertencia al
inicio. La ejecución continúa, pero algunas secciones pueden quedar incompletas
o indicar `Permission Denied`.

## Limitaciones

- No instala dependencias. Cuando puede identificar la causa, el script informa
  `Command Not Installed`, `Permission Denied` o `Command Failed`. `Not
  Available` indica que no hay datos detectables para una sección.
- La detección de GPU depende de `lspci`; SMART depende de `smartctl`.
- Los estados de firewall y SSH reflejan las herramientas detectables en el
  momento de la ejecución y requieren revisión humana.
- El informe no sustituye auditoría de seguridad, pruebas de restauración ni
  monitorización continua.
- El script está orientado a Ubuntu Server, aunque degrada de forma segura en
  otros sistemas con comandos o rutas ausentes.

## Interpretación del reporte

Use el informe para establecer una línea base y decidir el trabajo de Sprint 0:

- Empiece por el **Resumen ejecutivo**. Es una vista `Best Effort`: sintetiza
  hostname, fecha, sistema, kernel, CPU, RAM, disco principal y espacio
  disponible, pero no reemplaza las secciones detalladas.
- Confirme versión LTS, kernel, arquitectura y capacidad disponible.
- Revise discos, particiones, UUID y filesystem antes de definir `/srv`.
- Valide rutas de red, DNS, SSH, firewall y puertos antes de instalar servicios.
- Trate los riesgos detectados como señales heurísticas, no como diagnóstico
  definitivo; valide cada hallazgo con el operador.
- Conserve el informe fuera del control de versiones y compártalo solo con
  personas autorizadas, pues describe el host.
