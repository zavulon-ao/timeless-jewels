package main

import (
	"fmt"

	"github.com/zavulon-ao/timeless-jewels/wasm/exposition"
)

func main() {
	exposition.Expose()
	fmt.Println("Calculator Initialized")
	select {}
}
