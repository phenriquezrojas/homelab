# Visión

Construir una plataforma que preserve la memoria digital de una familia durante décadas, independientemente del hardware, del operador o de la infraestructura utilizada. posible de reconstruir tras un fallo y seguro sin complejidad innecesaria.

## Visión de producto

El homelab usará Ubuntu Server LTS y Docker. Tailscale habilitará acceso privado; Caddy ordenará el acceso HTTP(S) bajo `home.arpa`; los datos persistentes se mantendrán bajo `/srv`; Restic con Backblaze B2 protegerá la información ante pérdida del host.

El operador debe poder responder con claridad qué servicios existen, dónde viven sus datos, cómo se accede, cómo se respaldan y cómo se recuperan.

## Principios de experiencia operativa

- La documentación acompaña a la configuración.
- La simplicidad operativa prevalece sobre añadir tecnología.
- Cada servicio tendrá propósito, propietario de datos y ruta de recuperación.
- El acceso remoto será privado por defecto.
- El sistema crecerá por etapas y conservará evidencia de decisiones.

## Horizonte

La hoja de ruta avanza desde la base documentada hacia bootstrap, plataforma central, respaldo, monitorización, servicios, endurecimiento y CI/CD. Consulte [ROADMAP.md](ROADMAP.md).
