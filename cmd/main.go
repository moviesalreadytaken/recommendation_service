package main

import (
	"log"

	ginPkg "github.com/gin-gonic/gin"

	"github.com/moviesalreadytaken/recommendation_service/internal/controllers"
	"github.com/moviesalreadytaken/recommendation_service/internal/repos"
	"github.com/moviesalreadytaken/recommendation_service/internal/serivces"
	"github.com/moviesalreadytaken/recommendation_service/internal/utils"
)

func main() {
	gin := ginPkg.New()
	cnf := utils.LoadCnfFromEnv()

	usersClient, err := serivces.NewUsersServiceClient(cnf)
	if err != nil {
		log.Fatalf("error while initializing users serivce client = %s", err.Error())
	}
	moviesClient, err := serivces.NewMovieServiceClient(cnf)
	if err != nil {
		log.Fatalf("error while initializing movies serivce client = %s", err.Error())
	}
	recomRepo := repos.NewInMemoryRecommendationRepo()
	recomService := serivces.NewMovieRecommendationService(recomRepo, usersClient, moviesClient)
	recomController := controllers.NewMovieRecommendationService(recomService)

	api := gin.Group("/")
	controllers.AddMovieRecommendationControllerRoutes(recomController, api)

	log.Fatal(gin.Run(":10003"))
}
