package main

import (
	"assignment3/router"
	"net/http"
)

func main() {
	rtr := router.Start()

	http.ListenAndServe(":8080", rtr)
}
