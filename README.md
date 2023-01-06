# Prueba de Concepto en Go Lang


Proyecto de servicio Back-End en Go utilizando Gin para los servicios HTTP y GORM como ORM para la conexión a la BD.


## Ejecución

Para ejecutar es necesario tener Go instalado. 
El proyecto se levanta con el siguiente comando:

`go run .` 

(`.`) este simbolo identifica el archivo main.go como entrypoint siempre y cuando se encuentre dentro de un módulo Go.

## Tecnologías

### Go
    - Gin Framework
    - GORM

### Docker

Este repositorio cuenta con una imagen docker para virtualizar una instancia de MySQL 8. La instancia crea un directorio llamado `db/` que contiene la data de la BD

Se puede ejecutar utilizando docker-compose 

`docker-compose -f mysql.yml up`

### Diseño

El proyecto está hecho bajo el patrón controller-service-repository-model

