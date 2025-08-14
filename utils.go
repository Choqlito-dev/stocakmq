package rabbitmq

import (
	"encoding/json"
	"log"
)

// FailOnError imprime un error y termina la ejecución.
func FailOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

// ToJSON convierte una estructura a JSON.
func ToJSON(v interface{}) []byte {
	b, err := json.Marshal(v)
	if err != nil {
		log.Printf("Error al convertir a JSON: %v", err)
		return nil
	}
	return b
}

// FromJSON convierte JSON a una estructura.
func FromJSON(data []byte, v interface{}) error {
	return json.Unmarshal(data, v)
}
