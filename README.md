Leal Rewards - Configuración y Ejecución
Este documento describe cómo configurar y ejecutar la base de datos y la aplicación Go para el proyecto Leal Rewards. Puedes elegir entre ejecutar los servicios en local o utilizar Docker Compose.

Opción 1: Ejecutar en Local

Paso 1: Configuración de la Base de Datos
Navega al directorio de configuración de la base de datos:

cd database

Construye la imagen de Docker para la base de datos PostgreSQL:
docker build -t leal_db .

Paso 2: Ejecución de la Base de Datos
Crea y ejecuta un contenedor a partir de la imagen de la base de datos:

docker run -d -p 54321:5432 --name leal_container leal_db

Asegúrate de que en el archivo .env la variable DATABASE_URL esté apuntando a la instancia local de PostgreSQL:
DATABASE_URL=postgres://postgres:postgres@localhost:54321/postgres?sslmode=disable

Paso 3: Ejecución de la Aplicación Go
Regresa al directorio raíz y luego navega al directorio cmd:

cd .. && cd cmd
Ejecuta la aplicación Go:

go run main.go


Opción 2: Ejecutar con Docker Compose

Paso 1: Configuración de la Base de Datos
Navega al directorio de configuración de la base de datos:

cd database
Construye la imagen de Docker para la base de datos PostgreSQL:

docker build -t leal_db .

Paso 2: Ejecución con Docker Compose
Desde el directorio raíz del proyecto, ejecuta Docker Compose:

cd ..
docker-compose up --build

Este comando levantará dos contenedores:

leal_container: para la aplicación Go.
psql_container: para la base de datos PostgreSQL.