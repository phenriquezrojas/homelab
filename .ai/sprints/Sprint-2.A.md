# Sprint 2.A — Diseño del Runtime y Plataforma Central

## 1. Objetivo del Sprint

Diseñar el **Runtime del Homelab** como el motor único de ejecución del proyecto (según ADR-010). El principal entregable del Sprint no es el diseño de capacidades específicas, sino la especificación del Runtime como plataforma permanente del proyecto. 

Las capacidades de la Plataforma Central pasan a ser el primer dominio del Runtime que servirá para validar el diseño. Al finalizar esta fase, todos los artefactos de diseño técnico deben estar aprobados y ser suficientes para que Sprint 2.B implemente el código sin ambigüedades arquitectónicas ni funcionales.

## 2. Alcance

El alcance del Sprint se divide en dos áreas principales:

### Diseño del Runtime
- Definir las interfaces lógicas, contratos e interacciones entre los subsistemas del motor de ejecución.
- Establecer las reglas lógicas para la resolución del modelo de convergencia (Current State vs Desired State).
- *Excluido:* No se diseñará la implementación física del Runtime. No se decidirá el lenguaje de programación, la estructura física de directorios en el repositorio ni el formato sintáctico del registro. (Esas decisiones corresponden a la implementación en Fase B).

### Primer Dominio del Runtime
- Modelar las capacidades existentes de la plataforma central dentro del modelo de componentes del Runtime.
- Diseñar las operaciones lógicas (`install`, `configure`, `validate`, `repair`) y dependencias para las capacidades fundacionales.

## 3. Preguntas Arquitectónicas que este Sprint debe responder

El diseño técnico del Runtime debe dar respuesta explícita a las siguientes interrogantes:

- ¿Cómo descubre el Runtime los componentes?
- ¿Cómo representa el Current State?
- ¿Cómo genera un Plan de Ejecución?
- ¿Cómo resuelve las dependencias entre componentes?
- ¿Cómo informa los resultados de las operaciones?
- ¿Cómo detecta los fallos y qué acciones toma al respecto?
- ¿Cómo incorpora nuevas capacidades en el futuro sin modificar su núcleo?
- ¿Cómo garantiza la convergencia del sistema?

## 4. Arquitectura del Runtime

El diseño del Runtime debe adherirse estrictamente a los principios establecidos en el [ADR-010](../../adr/ADR-010-homelab-runtime.md). El objetivo en Fase A es especificar cómo se materializan lógicamente esos principios para que la Fase B pueda codificarlos.

Sprint 2.B deberá implementar los siguientes subsistemas lógicos, los cuales deben quedar completamente especificados en este Sprint 2.A:

- **Contrato de Componentes:** Definición lógica de los inputs, outputs y expectativas de las 4 operaciones (`install`, `configure`, `validate`, `repair`), así como la identidad única del componente.
- **Registro:** El mecanismo lógico para el descubrimiento de componentes y la declaración de dependencias para construir el DAG.
- **Planner (Planificador):** El algoritmo lógico que evalúa el Current State (vía `validate`), lo compara con el Desired State, resuelve el DAG y genera el Plan de Ejecución.
- **Executor (Ejecutor):** El mecanismo lógico que recorre el plan, invoca las operaciones de los componentes, gestiona las transiciones de estado y aplica la política estricta de fallos.
- **Reporter (Reporte/Estado):** El mecanismo lógico para registrar los resultados y reportar los estados consolidados del sistema al finalizar o fallar un plan.

## 5. Primer Dominio del Runtime

Las siguientes capacidades demostrarán la viabilidad del Runtime y serán las primeras en ser modeladas:

### Container Runtime
- **Implementación vigente:** Ver ADR-002.
- **Responsabilidad:** Habilitar la ejecución aislada de componentes que no requieran acceso directo al host.

### Private Network
- **Implementación vigente:** Ver ADR-006.
- **Responsabilidad:** Proveer acceso privado, seguro y exclusivo a los servicios de la plataforma.

### Reverse Proxy
- **Implementación vigente:** Ver ADR-007.
- **Responsabilidad:** Exponer servicios internos de forma segura exclusivamente dentro de la red privada.

### Internal Name Resolution
- **Implementación vigente:** Ver ADR-008.
- **Responsabilidad:** Permitir resolución de nombres de servicios hacia la IP de la red privada.

## 6. Entregables

La Fase A debe producir los siguientes artefactos de diseño técnico:

- [x] **Architecture Decision Log:** Registro ligero de decisiones abiertas, descartadas y pendientes durante el diseño para evitar discusiones cíclicas futuras.
- [x] **Runtime Architecture Specification:** Documento de diseño que especifica el flujo lógico del Planner, Executor y Reporter.
- [x] **Component Contract Specification:** Definición técnica de las interfaces de entrada/salida para el contrato de los componentes.
- [x] **Registry Specification:** Definición lógica de cómo se estructuran las relaciones y dependencias.
- [x] **Execution Plan Specification:** Definición conceptual de las fases y contenido del plan que genera el Planner.
- [x] **Primer Dominio de Capacidades:** Especificación de las operaciones lógicas, estados y dependencias para las capacidades del dominio inicial bajo el paradigma del Runtime.

## 7. Criterios de Aceptación (Gate Review)

Para que el diseño de Fase A sea aprobado, debe cumplir con:

### Del Diseño del Runtime
- [x] Contrato de componente definido de forma independiente a la tecnología.
- [x] Registro lógicamente definido (mecanismo de descubrimiento).
- [x] Modelo de estados y flujo de convergencia completamente especificados.
- [x] Planificador definido (incluyendo detección de ciclos e invalidación de dependencias circulares).
- [x] Política de fallos definida de acuerdo al principio Halt-on-fail.
- [x] Interfaces lógicas entre subsistemas (Registro, Planner, Executor, Reporter) documentadas.
- [x] **Vacíos Arquitectónicos:** Cualquier decisión arquitectónica faltante requerida para que Sprint 2.B pueda implementarse sin ambigüedad debe registrarse y resolverse explícitamente en el diseño.

### De las Capacidades Iniciales (Derivados)
- [x] Se documentan los criterios lógicos para validar el estado de la capacidad *Container Runtime*.
- [x] Se documentan los criterios lógicos para validar el estado de la capacidad *Private Network*.
- [x] Se documentan los criterios lógicos para validar el estado de la capacidad *Reverse Proxy*.
- [x] Se documentan los criterios lógicos para validar el estado de la capacidad *Internal Name Resolution*.
- [x] Se especifican las dependencias (DAG) entre estas capacidades.

## 8. Archivos a Producir

El diseño técnico no debe producir artefactos ejecutables. Los documentos técnicos a producir son:

- `.ai/sprints/Sprint-2.A.md` (Este documento).
- `.ai/implementation/Sprint-2-Runtime-Design.md` (Documento principal consolidado). Debe construirse y aprobarse secuencialmente en este orden:
  1. Modelo Conceptual
  2. Contrato del Componente
  3. Registro (Descubrimiento y Dependencias)
  4. Modelo de Estado
  5. Planner (Planificador)
  6. Plan de Ejecución
  7. Executor (Ejecutor)
  8. Reporter
  9. Casos de Error y Fallos
  10. Primer Dominio de Capacidades
- Actualización de documentación de diseño en `docs/` (adaptando diseños al modelo de capacidades del Runtime).

## 9. Trazabilidad

| Tipo | Referencia |
|---|---|
| ADR | ADR-001, ADR-002, ADR-006, ADR-007, ADR-008, ADR-010 |
| Sprint anterior | Sprint 1.B |
| Sprint siguiente | Sprint 2.B |
| Phase | Phase 2 (Storage), Phase 3 (Networking), Phase 5 (Platform Services) |

## 10. Estado

Completed

## 11. Lecciones Aprendidas

- **Del Sprint 1:** La separación A/B ha demostrado su valor. El contrato de diseño detallado en la Fase A permitió que la Fase B se ejecutara de forma autónoma, con revisiones enfocadas en el cumplimiento del contrato. La regla de Gate Review añadida a la Constitución protege contra la ampliación involuntaria del alcance.
- **Reapertura desde Sprint 2.B:** El Sprint 2.B falló en la Subfase B2 porque la infraestructura declarativa carecía de un motor unificado de ejecución. Esto derivó en el ADR-010 que instituye un Runtime único y modular. El Sprint 2.A se reabre porque el ADR-010 modifica el contrato arquitectónico del Runtime y deja incompleto el diseño aprobado originalmente.

## 12. Próximo Sprint

Sprint 2.B: Implementación del Runtime y de las primeras capacidades de la Plataforma Central utilizando dicho Runtime.
