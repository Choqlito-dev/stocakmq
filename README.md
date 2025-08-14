USO DE LIBRERIA STOCAKMQ

# Instalacion
> # XD
```
go mod init tu-proyecto
go get https://github.com/Choqlito-dev/stocakmq.git
```

# Uso

Crea un archivo `main.go` y copia y pega este archivo.

```go
package main

import (
	"fmt"
	"log"

	rabbitmq "github.com/CHoqlito-dev/stocakmq"
)

func main() {
	rmq, err := rabbitmq.NuevaConexion("amqp://guest:guest@localhost:5672/")
	if err != nil {
		log.Fatal(err)
	}
	defer rmq.Cerrar()

	// publicar mensaje
	err = rmq.Publicar("test_queue", []byte(`{"msg": "Hola desde Go"}`))
	if err != nil {
		log.Fatal("Error publicando:", err)
	}

	// consumir mensajes
	rmq.Consumir("test_queue", func(msg amqp.Delivery) {
		fmt.Println("Mensaje recibido:", string(msg.Body))
	})

	select {} // mantener vivo
}
```
