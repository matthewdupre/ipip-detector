package main

import (
	"fmt"
	"os"
)

func is_ipip_enabled() bool {
	return false
}

func ping_someone() bool {
	return false
}

func set_ipip(enabled bool) {
	return
}

func main() {
	enabled := is_ipip_enabled()

	if enabled {
		fmt.Println("Trying without IPIP")
		set_ipip(false)
	}
	connectivityOk := ping_someone()

	if !connectivityOk {
		fmt.Println("Trying with IPIP")
		set_ipip(true)
	}
	connectivityOk = ping_someone()

	if connectivityOk {
		fmt.Println("Connectivity OK!: IPIP required:", enabled)
	} else {
		fmt.Println("Failed to establish connectivity")
		os.Exit(1)
	}
}
