# Mercado Libre Mutans

instale el repositorio

`git clone https://github.com/jeffleon/Mercado_Libre_Mutans.git`

## API REST
para que la api rest funcione por favor ubicate en la carpeta Mercado_Libre_Mutans

`cd Mercado_Libre_Mutans/`
* utiliza el comando go get para instalar las dependencias
  `go get ./...`
    
__NOTA:__ es importante que de antemano crearas una base de datos en Cloud SQL (GCP)
* corre el cloud_sql_proxy
  
`./cloud_sql_proxy -instances=Connection name=tcp:5432`
* dentro de la carpeta database agrega tus credenciales de GCP

`dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s", "user", "password", "127.0.0.1:5432", "dbname")`
* correr la api
  `go run main.go`

* la api corre en el puerto 8080

### EndPoints

* __GET__  _/api/v1/stats_
* __GET__  _/api/v1/stats/:id_
* __POST__  _/api/v1/mutant_

### Depenpendencias usadas

si por alguna razon la api no corre se utilizaron las siguientes dependencias

* `go get github.com/gofiber/fiber/v2`

* `go get github.com/jinzhu/gorm`

* `go get github.com/jinzhu/gorm/dialects/sqlite`

* `go get database/sql`

* `go get github.com/GoogleCloudPlatform/cloudsql-proxy/proxy/dialers/mysql`