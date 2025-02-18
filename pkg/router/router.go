package router

import (
	"log"
	"net/http"
	"square-pos/pkg/service/user"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func InitializeRoutes(db *gorm.DB) *mux.Router {
	router := mux.NewRouter()

	subRouter := router.PathPrefix("/api/v1").Subrouter()
	userHandler := user.NewHandler()

	userHandler.RegisterRoutes(subRouter)
	log.Println("Server started............")
	return router
}

func StartServer(db *gorm.DB) {
	router := InitializeRoutes(db)
	http.Handle("/", router)

	server := &http.Server{Addr: ":8080"}

	go func() {
		if err := server.ListenAndServe(); err != nil {
			panic(err)
		}
	}()
}
