# 🚀 Markitos SVC Boilerplates REST

¡Bienvenido a **Markitos SVC Boilerplates**! Este proyecto es un conjunto de plantillas y herramientas diseñadas para acelerar el desarrollo de servicios backend con las mejores prácticas de DevSecOps. 🥷

---

## 📋 Índice

1. 🎯 [¿Qué es esto?](#qué-es-esto)
2. 🛠️ [Tecnologías utilizadas](#tecnologías-utilizadas)
3. 🚀 [Cómo empezar](#cómo-empezar)
4. ⚙️ [Configuración: Variables de entorno al poder](#configuración-variables-de-entorno-al-poder)
5. 📂 [Estructura del proyecto](#estructura-del-proyecto)
6. ☸️ [Pipelines de CI/CD](#pipelines-de-cicd)
7. 🤝 [Contribuciones](#contribuciones)
8. 🧪 [Testing](#testing)
9. 🛡️ [Seguridad](#seguridad)
10. 📜 [Licencia](#licencia)
11. 🧙‍♂️ [Autor](#autor)

---

## 🎯 ¿Qué es esto?

Este repositorio contiene boilerplates para servicios backend que incluyen:

- Integración con **PostgreSQL** usando GORM.
- Servidor **REST** con Gin.
- Scripts para automatizar tareas comunes.
- Hooks de pre-commit para seguridad y calidad del código.
- ¡Y mucho más! 🚀

---

## 🛠️ Tecnologías utilizadas

- **Go** (1.24.1)
- **PostgreSQL** con GORM
- **Viper** para configuración
- **Gin** para APIs REST
- **Gitleaks** para detección de secretos
- **Snyk** para análisis de vulnerabilidades

---

## 🚀 Cómo empezar

1. Clona este repositorio:

    ```bash
    git clone https://markitos-svc-boilerplates-rest.git
    cd markitos-svc-boilerplates-rest
    ```

2. Levanta la base de datos con Docker:

    ```bash
    make postgres
    ```

3. Crea la base de datos:

    ```bash
    make createdb
    ```

4. Ejecuta los tests:

    ```bash
    make test
    ```

    *   **Por qué:** Ejecuta las pruebas unitarias para verificar la integridad del código.

    ```bash
    make run
    ```

    *   **Por qué:** Ejecuta la aplicación localmente con la configuración predeterminada o las variables de entorno definidas.

    ```bash
    make image VERSION=1.2.3
    ```

    *   **Por qué:** Construye la imagen Docker con una versión específica.

    ```bash
    make testv
    ```

    *   **Por qué:** Ejecuta las pruebas en modo verboso, mostrando más detalles.

    ```bash
    make postgres
    ```

    *   **Por qué:** Levanta una instancia de PostgreSQL en Docker para desarrollo local.

    ```bash
    make createdb
    ```

    *   **Por qué:** Crea la base de datos y el usuario necesarios para la aplicación.

    ```bash
    make dropdb
    ```

    *   **Por qué:** Elimina la base de datos.

    ```bash
    make proto
    ```

    *   **Por qué:** Construye la imagen Docker con una versión específica.

    ```bash
    make caas SERVICE_NAME=nuevo-servicio ENTITY_NAME=entidad
    ```

    *   **Por qué:** Crea un nuevo servicio y entidad en el proyecto.

5. ¡Listo para desarrollar! 🎉

---

## ⚙️ Configuración: Variables de entorno al poder

Este proyecto está diseñado para ser configurado principalmente a través de variables de entorno. Aunque se puede utilizar un archivo `app.env` para desarrollo local, **la configuración final y autoritativa siempre provendrá de las variables de entorno**.

### ¿Por qué solo variables de entorno? 🤔

- **Facilidad de despliegue:** Los entornos de producción modernos (Kubernetes, Cloud Functions, etc.) favorecen la configuración a través de variables de entorno.
- **Seguridad:** Evita la inclusión de secretos (contraseñas, tokens) en el código fuente o archivos de configuración versionados.
- **Consistencia:** Asegura que la configuración sea la misma en todos los entornos, reduciendo errores y sorpresas.

### ¿Cómo configurar mi entorno? ⚙️

1. **Variables obligatorias:**

    - `DATABASE_DSN`: Cadena de conexión a la base de datos PostgreSQL. Ejemplo:

        ```text
        host=localhost user=admin password=admin dbname=markitos-svc-boilerplates-rest sslmode=disable
        ```

    - `HTTP_SERVER_ADDRESS`: Dirección del servidor REST. Ejemplo: `:3000`

2. **Variables opcionales:**

    - Si no se especifica `HTTP_SERVER_ADDRESS`, se usarán los valores por defecto (`:3000`).

### Prioridad de configuración 🥇

1. Variables de entorno (prioridad máxima).
2. Archivo `app.env` (si existe, se sobreescribe con las variables de entorno).
3. Valores por defecto (si no se encuentra ni variable de entorno ni valor en `app.env`).

### Ejemplo de uso en Docker Compose 🐳

```yaml
version: "3.8"
services:
  app:
    image: markitos-svc-boilerplates-rest:latest
    ports:
      - "3000:3000"
    environment:
      DATABASE_DSN: "host=db user=admin password=admin dbname=markitos-svc-boilerplates-rest sslmode=disable"
      HTTP_SERVER_ADDRESS: ":3000"
```

### Recomendaciones ✅

- Utiliza herramientas como dotenv para gestionar variables de entorno en desarrollo local.
- En entornos de producción, configura las variables de entorno directamente en el sistema operativo o plataforma de despliegue (Kubernetes Secrets, AWS Secrets Manager, etc.).
- Evita incluir información sensible directamente en los archivos de definición de contenedores (Dockerfiles, docker-compose.yml).

---

## ☸️ Pipelines de CI/CD

Este proyecto utiliza GitHub Actions para automatizar el proceso de Continuous Integration y Continuous Deployment (CI/CD). Los pipelines se definen en el directorio [.github/workflows](.github/workflows).

### Pipelines activos ⚙️

1. **Publish Docker Image (publish.yml)**

    - **Disparador:** Creación de un tag (ej: `1.2.3`).
    - **Acciones:**
        - Autentica con GitHub Packages.
        - Construye la imagen Docker.
        - Publica la imagen en GitHub Container Registry (ghcr.io).
        - Verifica la imagen publicada.

    - **Por qué:** Automatiza la publicación de nuevas versiones de la aplicación en el registro de contenedores.

### Otros Pipelines (pueden activarse cambiando el branch) ⚙️

1. **[Nombre del Pipeline]**

    - **Disparador:** [Evento que dispara el pipeline].
    - **Acciones:**
        - [Lista de acciones que realiza el pipeline].

    - **Por qué:** [Explicación del propósito del pipeline].

### Recomendaciones ✅

*   Para ejecutar la aplicación en modo de desarrollo (con logs detallados):

    ```bash
    make run
    ```

*   Para ejecutar la aplicación en modo de producción (sin logs detallados):

    ```bash
    GIN_MODE=release make run
    ```

    *   **Por qué:** Reduce la cantidad de información en los logs, útil para entornos de producción.

*   Adapta los pipelines a tus necesidades específicas.
*   Añade más pruebas y análisis de seguridad a los pipelines.
*   Utiliza variables de entorno y secretos para configurar los pipelines de forma segura.

---

## 📂 Estructura del proyecto

```plaintext
├── cmd/                # Punto de entrada de la aplicación
├── infrastructure/     # Configuración, base de datos, gRPC, API REST, etc.
├── bin/                # Scripts útiles
├── etc/                # Configuración adicional (hooks, etc.)
├── go.mod              # Dependencias del proyecto
└── README.md           # Este archivo 😎
```

---

## 🤝 Contribuciones

¡Las contribuciones son bienvenidas! Si tienes ideas, mejoras o encuentras bugs, no dudes en abrir un issue o un pull request. 💡

---

## 🧪 Testing

Este proyecto incluye un conjunto de pruebas automatizadas para garantizar la calidad del código y la funcionalidad. Antes de realizar un commit, asegúrate de ejecutar las pruebas para evitar introducir errores.

### Ejecutar pruebas

Para ejecutar las pruebas, utiliza el siguiente comando:

```bash
make test
```

### Cobertura de pruebas

Puedes generar un informe de cobertura de pruebas ejecutando:

```bash
go test ./... -coverprofile=coverage.out
go tool cover -html=coverage.out
```

Esto abrirá un informe en tu navegador mostrando qué partes del código están cubiertas por las pruebas.

---

## 🛡️ Seguridad

La seguridad es una prioridad en este proyecto. Aquí hay algunas herramientas y prácticas que se utilizan:

1. **Gitleaks**: Para detectar secretos en el código fuente.
    - Ejecuta `gitleaks detect` para analizar el repositorio.
2. **Snyk**: Para identificar vulnerabilidades en las dependencias.
    - Ejecuta `snyk test` para realizar un análisis de seguridad.

Asegúrate de revisar y solucionar cualquier problema reportado antes de desplegar a producción.

---

## 📜 Licencia

Este proyecto está bajo la licencia MIT.  
¡Haz lo que quieras con él, pero no olvides dar crédito! 😉

---

## 🧙‍♂️ Autor

Creado con ❤️ por Marco Antonio (alias Markitos).

- 💼 [LinkedIn](#)
- 🐦 [Twitter](#)
