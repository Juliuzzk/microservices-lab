

# Levantar backend de manera local

# Billing Services
dsn="host=localhost port=5432 user=postgres password=postgres dbname=lab sslmode=disable" go run  server.go   

# Custody Service
dsn="host=localhost port=5432 user=postgres password=postgres dbname=lab sslmode=disable" go run  server.go   


# Levantar front end-api

BILLING_BACKEND=localhost:5000 npx ts-node src/server.ts

# Ejecutar container en back end con variables de entorno para poder pasar la conexion
# Para ejecucion local

docker run -e dsn="host=localhost port=5432 user=postgres password=postgres dbname=lab sslmode=disable" -d billing-service:0.0.1

# Importar imagen al cluster
k3d image import <docker_image>:latest -c dev