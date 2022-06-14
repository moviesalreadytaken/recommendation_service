package controllers

import "github.com/gin-gonic/gin"

func AddMovieRecommendationControllerRoutes(c *MovieRecommendationController, api *gin.RouterGroup) {
	api.GET("/recoms/:userId", c.NewRecommendation)
	api.GET("/recoms", c.GetAllRecommendations)
}	
