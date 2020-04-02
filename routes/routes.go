package routes

import (
	"net/http"
	"github.com/brendonhps/controllers"
)

// CarregaRotas : 
func CarregaRotas()  {
	http.HandleFunc("/", controllers.Index)
	http.HandleFunc("/new", controllers.Store)
	http.HandleFunc("/insert", controllers.Insert)
	http.HandleFunc("/delete", controllers.Delete)
	http.HandleFunc("/edit", controllers.Edit)
	http.HandleFunc("/update", controllers.Update)

}