package main

import (
	"context"
	"fmt"
)

// Disco struct
type Disco struct {
	ctx context.Context
}

func NewDisco() *Disco {
	return &Disco{}
}

func PruebaDisco() {
	fmt.Println("siuuu")
}
