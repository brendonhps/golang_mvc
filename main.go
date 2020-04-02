package main

import (
	"net/http"
	"github.com/brendonhps/routes"
)

func main() {
	routes.CarregaRotas()
	http.ListenAndServe(":8000", nil)
}
