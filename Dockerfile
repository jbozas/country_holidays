FROM golang:latest

WORKDIR /country_holidays

COPY go.mod .
COPY go.sum .

# Descargar las dependencias
RUN go mod download

# Copiar el resto del código fuente
COPY . .

# Compilar la aplicación
RUN go build -o main .

# Exponer el puerto en el que se ejecutará la aplicación
EXPOSE 8080

# Comando para ejecutar la aplicación
CMD ["./main"]
