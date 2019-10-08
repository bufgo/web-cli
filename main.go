package main

import (
	"bufgo/server"
)

func main() {
	// Load router
	r := server.NewRouter()
	r.Run()
}
