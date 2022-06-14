package controllers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/moviesalreadytaken/recommendation_service/internal/models"
	"github.com/moviesalreadytaken/recommendation_service/internal/models/rest"
	"github.com/moviesalreadytaken/recommendation_service/internal/serivces"
)

type MovieRecommendationController struct {
	recomService *serivces.MovieReccomendationService
}

func NewMovieRecommendationService(
	recomService *serivces.MovieReccomendationService) *MovieRecommendationController {
	return &MovieRecommendationController{
		recomService: recomService,
	}
}

func (c *MovieRecommendationController) NewRecommendation(g *gin.Context) {
	usrId := bindUserId(g)
	if g.IsAborted() {
		return
	}
	exists, err := c.recomService.UserExists(*usrId)
	if err != nil {
		log.Printf("error while getting user info = %s", err.Error())
		g.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	if !exists {
		g.AbortWithStatusJSON(http.StatusBadRequest,
			rest.Result{Msg: "user not found"})
		return
	}
	movies, err := c.recomService.GenRecommndation(*usrId)
	if err != nil {
		log.Printf("error while recommendation recomendation = %s", err.Error())
		g.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	g.JSON(http.StatusOK, models.DbMoviesToRest(movies))
}

func (c *MovieRecommendationController) GetAllRecommendations(g *gin.Context) {
	recoms, err := c.recomService.GetAllRecommendations()
	if err != nil {
		log.Printf("error while getting recommendations info = %s", err.Error())
		g.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	g.JSON(http.StatusOK, models.DbRecommendationsToRest(recoms))
}

func bindUserId(g *gin.Context) *uuid.UUID {
	idStr := g.Param("userId")
	if idStr == "" {
		g.AbortWithStatusJSON(http.StatusBadRequest,
			rest.Result{Msg: "user id is empty!"})
		return nil
	}
	id, err := uuid.Parse(idStr)
	if err != nil {
		g.AbortWithStatusJSON(http.StatusBadRequest,
			rest.Result{Msg: err.Error()})
	}
	return &id
}
