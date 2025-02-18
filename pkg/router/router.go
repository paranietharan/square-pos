package router

import (
	"net/http"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func InitializeRoutes(db *gorm.DB) *mux.Router {
	router := mux.NewRouter()

	// router.HandleFunc("/users", handler.CreateNewUser(db)).Methods("POST")

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
