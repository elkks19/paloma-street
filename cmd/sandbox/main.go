package main

import (
	"log"
	"path/filepath"
)

func main() {
	log.Print(filepath.Join("/imagenes", "prueba", "test.txt"))
}

