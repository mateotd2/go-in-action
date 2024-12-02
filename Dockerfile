# Usar una imagen base con Go
FROM golang:latest

# Configurar el directorio de trabajo dentro del contenedor
WORKDIR /app

# Copiar los archivos del proyecto al contenedor
COPY . .

# Descargar dependencias y compilar el binario
RUN go mod tidy && go build -o app

# Comando por defecto para ejecutar el binario
# CMD ["./app"]
