package main

import (
	"fmt"

	"github.com/cmackenzie1/go-uuid"
)

func main() {
	v4, _ := uuid.NewV4()
	fmt.Printf("UUIDv4: %s\n", v4) // c07526de-40e5-418f-93d1-73ba20d2ac2c

	v7, _ := uuid.NewV7()
	fmt.Printf("UUIDv7: %s\n", v7) // 0185e1af-a3c1-704f-80f5-6fd2f8387f09
}
