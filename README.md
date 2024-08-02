# Leal Rewards - Configuración y Ejecución
Este documento describe cómo configurar y ejecutar la base de datos y la aplicación Go para el proyecto Leal Rewards. Puedes elegir entre ejecutar los servicios en local o utilizar Docker Compose.

## Opción 1: Ejecutar en Local

### Paso 1: Configuración de la Base de Datos
Navega al directorio de configuración de la base de datos:

```cd database```

Construye la imagen de Docker para la base de datos PostgreSQL:
```docker build -t leal_db .```

### Paso 2: Ejecución de la Base de Datos
Crea y ejecuta un contenedor a partir de la imagen de la base de datos:

```docker run -d -p 54321:5432 --name leal_container leal_db```

Asegúrate de que en el archivo .env la variable DATABASE_URL esté apuntando a la instancia local de PostgreSQL:
DATABASE_URL=postgres://postgres:postgres@localhost:54321/postgres?sslmode=disable

### Paso 3: Ejecución de la Aplicación Go
Regresa al directorio raíz y luego navega al directorio cmd:

```cd .. && cd cmd```
Ejecuta la aplicación Go:

```go run main.go```


## Opción 2: Ejecutar con Docker Compose

### Paso 1: Configuración de la Base de Datos
Navega al directorio de configuración de la base de datos:

```cd database```
Construye la imagen de Docker para la base de datos PostgreSQL:

```docker build -t leal_db .```

### Paso 2: Ejecución con Docker Compose
Desde el directorio raíz del proyecto, ejecuta Docker Compose:

```cd ..```
```docker-compose up --build```

Este comando levantará dos contenedores:

leal_container: para la aplicación Go.
psql_container: para la base de datos PostgreSQL.

# Funcionamiento y uso de endpoints

## En la raiz del proyecto van a encontrar  LEAL-REWARDS.postman_collection 

esta es la coleccion de postman que contiene los curls ya armados para que puedan probar el api

## connect api 
esta solicitud es de lectura y solo confirma que la api se levanto con exito 

### respuesta
{
    "message": "Welcome api-REST",
    "status": 200
}

## create user 
esta solicitud de escritura creea usuarios para que puedan consultar y hacer compras 
### cuerpo de la solicitud 
{
    "document":"1022437014",
    "password" : "cristiangar"
}

### respuesta de la solicitud  
{
    "id": "2k7XfNsUp76ngcCgIIYDcvitDJq",
    "document": "1022437014"
}

## login  
Con tus cdredenciales anteriores ( document y password ) puedes acceder a la info relacionada con cada usuario , si no haz realizado comprar tendra valores nulos 

### cuerpo de la solicitud 
{
    "document":"1022437014",
    "password" : "cristiangar"
}

### respuesta de la solicitud  
{
    "id": "2k7XfNsUp76ngcCgIIYDcvitDJq",
    "document": "1022437014",
    "points": [
        
    ],
    "cashback": [
       
    ],
    "Leal_Coins": 390,
    "trades": null
}

## create trade 
Para realizar una compra , primero debemos crear un comercio . pasamos el nombre del comercio y un User_id ( id de el primer usuario logueado en la api , por defecto este usuario va a ser el administrador  )

### cuerpo de la solicitud 

{
    "name" :"Adidas",
    "user_id" : "2k7XfNsUp76ngcCgIIYDcvitDJq"
}

**recuerda que el user_id es el id del usuario que se creo al inicio del recorrido**

### respuesta de la solicitud 

{
    "message": "Comercio Registrado con exito",
    "id": "2k7XvWj1ISzYyJQROvCQcIwE4cj",
    "trade_name": "Adidas",
    "user_id": "2k7XfNsUp76ngcCgIIYDcvitDJq"
}

## create transaction
Y por ultimo vamos a crear la transaccion , enviamos : documento , monto y nombre del comercio 

### cuerpo de la solicitud 
{
    "document" : "1022437014",
    "amount" : 30000,
    "trade_name" : "Adidas" 
}

### respuesta de la solicitud 

{
    "message": "Transaction successfully registered",
    "id": "2k7XwGTeE3b1JJIKiYSWC5veExJ"
}

**Si revisas te logueas de nuevo vas a ver una respuesta con tus datos completos y porsupuesto tus leal coins y puntos 🚀**

{
    "id": "2k7XfNsUp76ngcCgIIYDcvitDJq",
    "document": "1022437014",
    "points": [
        {
            "trade_id": "2k7XlDfwqLE30n8BnwerDq9zarg",
            "points": 360
        },
        {
            "trade_id": "2k7XvWj1ISzYyJQROvCQcIwE4cj",
            "points": 580
        }
    ],
    "cashback": [
        {
            "amount": 150
        },
        {
            "amount": 180
        },
        {
            "amount": 30
        },
        {
            "amount": 30
        },
        {
            "amount": 550
        }
    ],
    "Leal_Coins": 940,
    "trades": null
}