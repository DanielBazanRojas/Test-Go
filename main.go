package main

import (
	"fmt"

	"github.com/google/uuid"
)

func main() {
	fmt.Println("UUID Generada: ")
	fmt.Println(uuid.New().String())
}
