package controllers

import (
	"context"
	"gin-mongo-api/configs"
	"gin-mongo-api/models"
	"gin-mongo-api/responses"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var manifestCollection *mongo.Collection = configs.GetCollection(configs.DB, "manifests")
var manifestValidate = validator.New()

func CreateManifest() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var manifest models.Manifest
		defer cancel()

		//validate the request body
		if err := c.BindJSON(&manifest); err != nil {
			c.JSON(http.StatusBadRequest, responses.ManifestResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		//use the validator library to validate required fields
		if validationErr := manifestValidate.Struct(&manifest); validationErr != nil {
			c.JSON(http.StatusBadRequest, responses.ManifestResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": validationErr.Error()}})
			return
		}

		newManifest := models.Manifest{
			Id:       primitive.NewObjectID(),
			Image:    manifest.Image,
			Location: manifest.Location,
			Date:     manifest.Date,
		}

		// The image should be sent to the cantaloupe server

		result, err := manifestCollection.InsertOne(ctx, newManifest)
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.ManifestResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		c.JSON(http.StatusCreated, responses.ManifestResponse{Status: http.StatusCreated, Message: "success", Data: map[string]interface{}{"data": result}})
	}
}

func GetAllManifests() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var manifests []models.Manifest
		defer cancel()

		results, err := manifestCollection.Find(ctx, bson.M{})

		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.ManifestResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		//reading from the db in an optimal way
		defer results.Close(ctx)
		for results.Next(ctx) {
			var singleManifest models.Manifest
			if err = results.Decode(&singleManifest); err != nil {
				c.JSON(http.StatusInternalServerError, responses.ManifestResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			}

			manifests = append(manifests, singleManifest)
		}

		c.JSON(http.StatusOK,
			responses.ManifestResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": manifests}},
		)
	}
}
