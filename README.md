
# Entrega Evaluación - Desarrollo de taller Microservicios

## Integrantes

- Cristian Simon
- Juan Jara
- Julio Cáceres
- Sebastian Hechenleitner

##  Levantar proyecto

1. __Clonar__ o realizar __fork__ del repositorio presente. Para clonar el repositorio, ejecutar el siguiente comando:

```shell
cd <directorio de trabajo>
git clone https://github.com/malarcon-79/microservices-lab.git
cd microservices-lab
```

2. Iniciar la base de datos y otros requisitos. Para esto, se debe ingresar a la carpeta __01-kubernetes__ y ejecutar scripts `00-create-cluster.sh` y `01-create-backend.sh`

```shell
cd 01-kubernetes

./00-create-cluster.sh
./01-create-backend.sh

cd ..
```

3. Generar imagen de los servicios de backend, para esto debemos dirigirnos a a la carpeta `02-servicios-backend` y ejecutar el comando de docker build en las siguientes rutas:

```shell
cd 02-servicios-backend

cd billing-service
docker build -t billing-server:latest .

cd ..

cd custody-service
docker build -t custody-server:latest .

cd ../..
```

4. Generar imagen de los servicios de frontend, para esto debemos dirigirnos a la carpeta `03-servicios-frontend` y ejecutar el comando de docker build en las siguientes rutas:


```shell
cd 03-servicios-frontend

cd frontend-api
docker build -t frontend-api:latest .

cd ..

cd frontend-app
docker build -t frontend-app:latest .

cd ../..
```

5. Exportar las imagenes de docker recientemente creadas  al cluster de kubernetes:

```shell

k3d image import billing-service:latest -c dev

k3d image import custody-service:latest -c dev

k3d image import frontend-api:latest -c dev

k3d image import frontend-app:latest -c dev

```

6. Cargar configuraciones a nuestro cluster de kubernetes:

```shell
cd k8s

kubectl apply -f liberacion.yaml

```

7. Con los pasos realizados, ya podremos hacer uso de nuestro proyecto.

---
