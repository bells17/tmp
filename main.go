package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Printf("uid: %d\n", os.Geteuid())
}
