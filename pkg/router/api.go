package router

/*
* Dependency injection helpes to maintain a single database connections
* so router or services no need to create the database connections. when we need to scale the application
* no need to create many datbase connections

* helps to maintain modularity of the code

 */
import (
	"log"
	"net/http"
	"square-pos/pkg/service/pos"
	"square-pos/pkg/service/user"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func InitializeRoutes(db *gorm.DB) *mux.Router {
	router := mux.NewRouter()

	subRouter := router.PathPrefix("/api/v1").Subrouter()

	// Register routes for authentiaction & authorizations
	userStore := user.NewStore(db)
	userHandler := user.NewHandler(userStore)

	userHandler.RegisterRoutes(subRouter)

	/// register routes for pos operations
	posStore := pos.NewPosStore(db)
	posHandler := pos.NewPosHandler(posStore)

	posHandler.RegisterRoutes(subRouter)
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
