package router

import (
	"fmt"
	"log"
	"net/http"
	"zapys/admin"
	"zapys/consultas/cep"
	cpfv2 "zapys/consultas/cpfs/CPFV2"
	"zapys/consultas/placa"
	"zapys/src/config"
	"zapys/src/middleware"
	"github.com/rs/cors"
	"github.com/gorilla/mux"
)

func INITSERVER() {
	// mux
	r := mux.NewRouter()
	c := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET"},
		AllowedHeaders: []string{"*"},
		AllowCredentials: true,
	})
	handler := c.Handler(r)

	// rotas
	r.HandleFunc("/api/v2/cpf", middleware.TokenMiddleware(cpfv2.CPFV2)).Methods(http.MethodGet)
	
	r.HandleFunc("/api/v1/cep", middleware.TokenMiddleware(cep.CEP)).Methods(http.MethodGet)
	r.HandleFunc("/api/v1/placa", middleware.TokenMiddleware(placa.PLACAV1)).Methods(http.MethodGet)
	r.HandleFunc("/admin/criar", middleware.ADMMIDDLEWARE(admin.CriarUser)).Methods(http.MethodGet)
	r.HandleFunc("/admin/users", middleware.ADMMIDDLEWARE(admin.Users)).Methods(http.MethodGet)
	r.HandleFunc("/admin/delete", middleware.ADMMIDDLEWARE(admin.DeleteUser)).Methods(http.MethodGet)
	

	// configurar roteamento
	PORT := config.GetVarEnv("PORT")
	fmt.Printf("Servidor rodando http://localhost:%s", PORT)
	if err := http.ListenAndServe(":"+PORT, handler); err != nil {
		log.Fatal(err)
	}
}
