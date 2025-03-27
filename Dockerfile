# ==================================================================================
# 🔥 INSTRUCCIONES DE USO 🔥
# ==================================================================================
# 
# 🛠️ CONSTRUCCIÓN (sin Makefile):
# --------------------------------
# Para construir la imagen manualmente:
#    docker build -t markitos-svc-boilerplates-rest:1.0.0 .
#
# 🚀 EJECUCIÓN:
# ------------
# Para ejecutar con variables de entorno:
#    docker run -p 3000:3000 -p 30000:30000 \
#      -e DATABASE_DSN="host=db user=admin password=admin dbname=markitos-svc-boilerplates-rest sslmode=disable" \
#      -e HTTP_SERVER_ADDRESS=":3000" \
#      markitos-svc-boilerplates-rest:1.0.0
#
# 🌐 CONEXIÓN:
# -----------
# - API HTTP disponible en: http://localhost:3000
# - API gRPC disponible en: localhost:30000
# ==================================================================================

#:[.'.]:> Fase de construcción - ¡Usamos Alpine porque es ligero y rápido! Perfecto para compilar nuestra app
FROM golang:1.24-alpine AS builder

#:[.'.]:> Definimos nuestro espacio de trabajo, ¡como tener la mesa limpia antes de cocinar!
WORKDIR /app

#:[.'.]:> Primero copiamos las dependencias para aprovechar la caché de Docker
#:[.'.]:> ¡Truquito ninja para builds más rápidos! 🚀
COPY go.mod go.sum ./
RUN go mod download

#:[.'.]:> Ahora copiamos todo el código fuente - ¡Los ingredientes para nuestra receta!
COPY . .

#:[.'.]:> ¡Compilamos nuestra aplicación con esteroides! 💪
#:[.'.]:> CGO_ENABLED=0 para un binario sin dependencias externas
#:[.'.]:> ¡Así tenemos una app autosuficiente y lista para rockear!
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o /app/bin/server ./cmd/main.go

#:[.'.]:> Fase final con imagen distroless de Google - ¡Mínima y súper segura!
#:[.'.]:> Sin shell, sin paquetes innecesarios, ¡puro músculo y nada de grasa!
FROM gcr.io/distroless/static-debian12

#:[.'.]:> Definimos dónde vivirá nuestra app - ¡Hogar dulce hogar!
WORKDIR /

#:[.'.]:> Solo copiamos el binario compilado - ¡Nada de basura extra!
#:[.'.]:> ¡Es como mudarte solo con lo esencial y dejar los trastos viejos!
COPY --from=builder /app/bin/server /server

#:[.'.]:> Declaramos los puertos que usaremos - ¡Las puertas de entrada a nuestra app!
#:[.'.]:> Recuerda que esto es solo documentación, no abre realmente los puertos 😉
EXPOSE 3000 30000

#:[.'.]:> Variables de entorno para configurar nuestra app - ¡Personalización a la carta!
#:[.'.]:> ¡Cambia estos valores para adaptar la app a tus necesidades!
#:[.'.]:> Estas variables de entorno son valores por defecto.
#:[.'.]:> Si defines estas variables al ejecutar el contenedor, tendrán prioridad sobre estos valores.
ENV DATABASE_DSN="host=localhost user=admin password=admin dbname=markitos-svc-boilerplates-rest sslmode=disable"
ENV HTTP_SERVER_ADDRESS=":3000"

#:[.'.]:> ¡La orden para arrancar nuestra máquina! Con formato exec para mejor gestión de señales
#:[.'.]:> ¡A darle caña! 🔥
CMD ["/server"]