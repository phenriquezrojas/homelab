# ADR-010 — Runtime único y modular del Homelab

## Estado

Aceptada.

## Fecha

2026-08-05.

## Contexto

El proyecto Homelab adoptó desde sus inicios una estrategia implícita en la que cada Sprint producía artefactos de ejecución independientes. Bajo este modelo, cada fase de implementación generaba su propio mecanismo de despliegue sin que existiera un punto de entrada unificado capaz de construir el servidor completo desde cero.

Durante la validación operacional del Sprint 2.B2, el Owner planteó dos cuestionamientos que revelaron un vacío arquitectónico fundamental:

1. **Ausencia de un punto de entrada único:** No existía un mecanismo capaz de instalar dependencias, resolver conflictos con servicios heredados y levantar todos los servicios de forma automatizada e idempotente.
2. **Gestión de estado preexistente:** No se había definido cómo el proyecto convive, migra o desmantela configuración heredada del host.

El rechazo del Sprint 2.B2 (evidencia: `docs/evidence/Sprint-2.B2-validation.md`) demostró que la infraestructura declarativa (manifiestos Compose, documentación) era insuficiente sin un motor unificado de ejecución.

## Problema

El Homelab carece de un motor único y reproducible para desplegar, configurar, validar y reparar la plataforma completa. Los artefactos de ejecución existentes están vinculados a Sprints individuales y no componen un sistema coherente. Esto impide la reconstrucción desde cero y genera divergencia entre Sprints sucesivos.

## Decisión

Se adopta un **Runtime** como motor único de ejecución del Homelab. El Runtime no ejecuta acciones: converge el sistema hacia un estado deseado. Este ADR establece los contratos arquitectónicos que lo rigen.

### 1. Runtime único

El proyecto posee un único Runtime. Este Runtime es el único mecanismo oficial para desplegar, configurar, validar y reparar el Homelab. Ningún Sprint podrá crear un mecanismo de despliegue independiente como punto de entrada principal.

### 2. Capacidades

Una **capacidad** representa una función observable del sistema, independiente de la tecnología utilizada para implementarla.

Ejemplos:

- *Container Runtime* es una capacidad. Docker es una implementación de esa capacidad.
- *Reverse Proxy* es una capacidad. Caddy es una implementación de esa capacidad.
- *Private Network Access* es una capacidad. Tailscale es una implementación de esa capacidad.

Las capacidades son estables. Las implementaciones pueden cambiar. El Runtime razona sobre capacidades, no sobre productos.

El Runtime razona sobre capacidades y estados, no sobre tecnologías específicas. La sustitución de una implementación por otra (por ejemplo Docker por Podman, o Caddy por Traefik) no modifica el modelo del Runtime; únicamente cambia la implementación de la capacidad correspondiente.

### 3. Componentes

Un **componente** es el elemento primario del Runtime. Cada componente representa una capacidad del sistema y expone exactamente cuatro operaciones:

| Operación | Responsabilidad | Restricción |
|---|---|---|
| **install** | Proveer las dependencias necesarias para que la capacidad exista en el sistema | No configura, no valida, no repara |
| **configure** | Llevar la capacidad de instalada a operativa según su configuración deseada | No instala, no valida, no repara |
| **validate** | Observar el estado actual de la capacidad y proporcionar al Runtime la información necesaria para determinar el Current State | No modifica estado bajo ninguna circunstancia. No decide qué acciones ejecutar; esa responsabilidad pertenece al motor de planificación |
| **repair** | Converger la capacidad nuevamente al estado deseado cuando se ha detectado divergencia | No instala capacidades nuevas; opera sobre lo que ya debería existir |

Cada operación es idempotente: ejecutarla múltiples veces produce siempre el mismo estado final.

Las dependencias se declaran entre componentes, no entre operaciones individuales. Si el componente *Reverse Proxy* depende de *Container Runtime*, las cuatro operaciones del primero asumen que el segundo ya fue satisfecho.

### 4. Contrato de un componente

Todo componente cumple el siguiente contrato:

- **Identidad:** Todo componente posee un identificador único y estable dentro del Runtime. Ese identificador es la referencia utilizada por el registro, las dependencias, los planes de ejecución y las evidencias. La forma concreta de representarlo es una decisión de implementación.
- **Nombre:** Identificador de la capacidad que provee.
- **Clasificación:** Según su naturaleza (ver Sección 5).
- **Dependencias:** Lista explícita de capacidades que requiere. Se expresan como nombres de capacidad, nunca como referencias a Sprints, módulos ni productos.
- **Operaciones:** Las cuatro operaciones obligatorias: install, configure, validate, repair.
- **Estado de salida:** Cada operación informa de manera inequívoca si terminó con éxito o con error, incluyendo la razón del resultado.

Un componente que no cumpla este contrato no puede registrarse en el Runtime.

### 5. Clasificación de componentes

El Runtime distingue componentes según su naturaleza:

| Clasificación | Descripción | Ejemplo conceptual |
|---|---|---|
| **Infraestructura** | Capacidades de base que habilitan la plataforma | Container Runtime, Network |
| **Plataforma** | Servicios compartidos consumidos por múltiples aplicaciones | Base de datos relacional, Caché |
| **Aplicación** | Servicios de usuario final que consumen infraestructura y plataforma | Gestión de fotos |
| **Operación** | Capacidades transversales de mantenimiento | Respaldo, Monitorización |

La clasificación informa prioridades dentro del grafo de dependencias y aísla capas de responsabilidad.

### 6. Registro

El Runtime descubre los componentes disponibles mediante un **registro**. El registro es la única fuente de verdad para el descubrimiento de componentes. Un componente que no esté registrado no existe para el Runtime.

El registro contiene, como mínimo, la identidad de cada componente, su clasificación y sus dependencias declaradas. El formato y la ubicación del registro son decisiones de implementación que se definirán en el Sprint correspondiente.

### 7. Dependencias por capacidad

Los componentes declaran dependencias sobre capacidades, no sobre Sprints, productos ni otros componentes directamente. El Runtime construye un **grafo acíclico dirigido (DAG)** a partir de las dependencias declaradas en el registro.

**Resolución del grafo:**

- El Runtime resuelve el orden de ejecución a partir del DAG. El orden histórico de los Sprints que originaron cada componente es irrelevante.
- **Las dependencias circulares invalidan el plan de ejecución y constituyen un error arquitectónico.** El Runtime rechaza cualquier plan que contenga ciclos en el grafo de dependencias.

### 8. Modelo de estado

El Runtime mantiene un modelo de estado por componente. Un componente se encuentra en exactamente uno de los siguientes estados:

| Estado | Significado |
|---|---|
| **Absent** | La capacidad no existe en el sistema |
| **Installed** | La capacidad fue instalada pero no configurada |
| **Configured** | La capacidad fue configurada pero no validada como operativa |
| **Healthy** | La capacidad fue validada y opera según lo esperado |
| **Failed** | La capacidad no cumple el estado deseado; requiere intervención |
| **Unknown** | El estado no ha sido determinado; requiere validación |

Las operaciones producen transiciones de estado:

- **install** transiciona de *Absent* a *Installed*.
- **configure** transiciona de *Installed* a *Configured*.
- **validate** transiciona de *Configured* a *Healthy* o a *Failed*.
- **repair** transiciona de *Failed* a *Configured* (y requiere validación posterior).

Este modelo permite que el Runtime determine qué operaciones son necesarias sin reejecutar operaciones que ya alcanzaron su estado objetivo.

### 9. Desired State y convergencia

El Runtime opera bajo un modelo de convergencia. No ejecuta acciones arbitrarias: compara el estado actual de cada componente contra un **Desired State** y ejecuta únicamente las operaciones necesarias para cerrar la brecha.

- El **Desired State** es la declaración de qué capacidades deben existir y en qué estado deben encontrarse.
- El **Current State** es el estado observado de cada componente, determinado por la operación validate.
- El Runtime converge el sistema del Current State al Desired State.

### 10. Plan de ejecución

El Runtime no ejecuta componentes de manera secuencial ni indiscriminada. Ante una intención declarada, el Runtime construye un **plan de ejecución**. El plan es el artefacto intermedio entre la intención del operador y la ejecución de operaciones. El Runtime nunca opera sin un plan.

Un plan de ejecución contiene, conceptualmente:

- **Intención:** Qué se desea lograr (despliegue completo, operación parcial, reparación).
- **Capacidades objetivo:** Qué componentes deben alcanzar el Desired State.
- **Dependencias resueltas:** El subgrafo del DAG necesario para satisfacer las capacidades objetivo.
- **Estado actual:** El Current State de cada componente involucrado.
- **Operaciones necesarias:** Las transiciones de estado que cierran la brecha entre Current State y Desired State.
- **Orden de ejecución:** La secuencia que respeta las dependencias del grafo.

Ejemplos de intenciones:

- **Despliegue completo:** Todas las capacidades registradas deben alcanzar el estado *Healthy*. El Runtime resuelve el grafo completo.
- **Operación parcial:** Una capacidad específica debe alcanzar un estado específico. El Runtime resuelve únicamente el subgrafo necesario.

El plan de ejecución constituye una representación efímera del proceso de convergencia. No forma parte del estado persistente del sistema ni reemplaza el registro de componentes. Cada ejecución del Runtime genera un nuevo plan a partir del estado observado y del estado deseado.

### 11. Política de fallos

Cuando una operación falla durante la ejecución de un plan:

- **El Runtime detiene la ejecución del plan.** No continúa con operaciones que dependen del componente fallido.
- **Los componentes sin dependencia sobre el componente fallido no se ven afectados** si su ejecución ya fue completada o si no estaban incluidos en el plan.
- **El componente fallido transiciona al estado *Failed*.**
- **El Runtime informa qué componente falló, en qué operación, la razón del error y qué capacidades del plan quedaron sin satisfacer.**
- **El Runtime no realiza rollback automático.** La reversión es una decisión del operador, no del motor. Este principio se alinea con el Principio 12 de la Constitución (reversibilidad explícita).

### 12. Alcance de repair

La operación repair tiene límites definidos:

- Repair converge un componente en estado *Failed* de vuelta a *Configured*, con la expectativa de que una validación posterior lo lleve a *Healthy*.
- Repair **no reinstala** un componente. Si la capacidad se encuentra *Absent*, la operación correcta es install, no repair.
- Repair **no destruye datos** de usuario. Puede reiniciar servicios, regenerar configuración derivada o restablecer estado operativo, pero no elimina datos persistentes.
- Repair **no modifica otros componentes.** Opera exclusivamente sobre el componente al que pertenece.

### 13. Acumulación incremental

Cada Sprint contribuye componentes al Runtime o amplía componentes existentes. Ningún Sprint reemplaza el Runtime completo. El Runtime permanece estable; los componentes crecen con el proyecto.

### 14. Reinstalación completa

El Runtime debe permitir reconstruir un servidor vacío (Ubuntu Server LTS recién instalado) ejecutando un despliegue completo. No debe depender del orden histórico de Sprints ni de conocimiento externo al repositorio. La combinación del repositorio, el registro y los secretos gestionados externamente es suficiente para reconstruir la plataforma.

### 15. Extensibilidad

Agregar un nuevo servicio al Homelab implica únicamente:

1. Crear un componente que cumpla el contrato (Sección 4).
2. Registrarlo con sus dependencias (Sección 6).

Nunca debe ser necesario modificar la arquitectura del Runtime para incorporar una capacidad nueva.

### 16. Independencia del framework EOS

El Runtime y el framework EOS son sistemas independientes:

| Sistema | Responsabilidad |
|---|---|
| **EOS** | Gobierna el proyecto: fases, sprints, evidencias, gates y trazabilidad documental |
| **Runtime** | Despliega el Homelab: converge el sistema hacia el Desired State |

- EOS jamás forma parte del Runtime.
- El Runtime jamás implementa lógica del framework EOS.
- Los artefactos de gobierno (`.ai/`) no son artefactos del Runtime.
- Los componentes del Runtime no referencian sprints, tareas ni documentos EOS.

## Alternativas consideradas

- **Mantener artefactos de ejecución independientes por Sprint:** Cada Sprint sigue produciendo su propio mecanismo de despliegue. Genera divergencia, duplicación e imposibilita la reconstrucción completa. Fue la causa directa del fallo del Sprint 2.B2.
- **Orquestador externo (Ansible, Terraform):** Introduce una dependencia significativa, curva de aprendizaje y complejidad que contradice el Principio 11 de la Constitución (simplicidad y mínimas dependencias). Puede reconsiderarse mediante un ADR separado.
- **Ejecución secuencial sin grafo:** Un motor que ejecuta todos los módulos en un orden fijo. Funciona para despliegues completos pero no permite operaciones parciales ni escala con interdependencias crecientes.
- **Módulos independientes sin componente contenedor:** Los módulos de install, configure, validate y repair existen como entidades separadas con sus propias dependencias. Multiplica la complejidad del grafo innecesariamente y dificulta razonar sobre el estado de una capacidad como unidad.

## Consecuencias

### Ventajas

- Reinstalación completa desde un servidor vacío mediante un único punto de entrada.
- El modelo de convergencia (Desired State vs Current State) evita reejecutar operaciones innecesarias.
- Las dependencias entre capacidades permiten operaciones parciales sin reejecutar todo el sistema.
- Menor deuda técnica: no se acumulan artefactos de ejecución huérfanos.
- Crecimiento modular: cada servicio nuevo es un componente; la arquitectura del Runtime permanece estable.
- El modelo de estado permite que el Runtime tome decisiones informadas sobre qué ejecutar.

### Desventajas

- Requiere definir e implementar un registro antes de que el Runtime sea funcional.
- Exige que todo componente declare sus dependencias de forma explícita; las dependencias implícitas rompen el grafo.
- El DAG de dependencias y el modelo de convergencia aumentan la complejidad del diseño inicial respecto a un script lineal.
- Exige disciplina para mantener la separación de responsabilidades entre operaciones (install/configure/validate/repair).
- La determinación del Current State requiere que las operaciones validate sean confiables; validaciones débiles degradan todo el modelo.

## Riesgos

- El contrato de componente puede ser insuficiente si no se valida con los primeros servicios reales durante el Sprint de implementación.
- La resolución del DAG puede volverse compleja conforme crecen los servicios; se deberá vigilar la sobreingeniería.
- Si la disciplina de separación de operaciones no se mantiene, los componentes pueden degradarse en monolitos que mezclan responsabilidades.
- La migración del bootstrap existente (`bootstrap/bootstrap.sh`) al modelo de Runtime requerirá un plan de transición explícito.
- Las operaciones validate deben ser suficientemente estrictas para que el Current State sea confiable; falsos positivos comprometen el modelo de convergencia.

## Referencias relacionadas

- [ADR-002 — Docker First](ADR-002-docker-first.md)
- [ADR-003 — Ubuntu Server LTS](ADR-003-ubuntu-server.md)
- [ADR-004 — Un único repositorio](ADR-004-single-repository.md)
- [ADR-009 — Estrategia de idempotencia del Bootstrap](ADR-009-bootstrap-idempotency.md)
- [Evidencia Sprint 2.B2](../docs/evidence/Sprint-2.B2-validation.md)
- [Sprint 2.B](../.ai/sprints/Sprint-2.B.md)
- [Constitución del Proyecto](../.ai/PROJECT_CONSTITUTION.md) — Principios 3, 4, 11, 12
