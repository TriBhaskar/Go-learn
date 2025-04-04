package main

import (
	"fmt"
	"go_project1/internal/handlers"
	"net/http"

	"github.com/go-chi/chi"
	log "github.com/sirupsen/logrus"
)

func printBanner() {
    banner := `
   ______  ____      ___    ____  ____
  / ____/ / __ \    /   |  / __ \/  _/
 / / __  / / / /   / /| | / /_/ // /  
/ /_/ / / /_/ /   / ___ |/ ____// /   
\____/  \____/   /_/  |_/_/   /___/   
                                    
`
    fmt.Println(banner)
}

func main(){
	log.SetReportCaller(true)
	var r *chi.Mux = chi.NewRouter()
	handlers.Handle(r)

	fmt.Println("Starting GO API Service")
	fmt.Println()
	printBanner()

	err :=http.ListenAndServe("localhost:8080", r)
	if err != nil {
		log.Error(err)
	}
}