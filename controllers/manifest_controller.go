package controllers

import (
	"context"
	"net/http"
	"path/filepath"
	"time"
	"tonothan/sangaku-pipeline-server/configs"
	"tonothan/sangaku-pipeline-server/models"
	"tonothan/sangaku-pipeline-server/responses"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var manifestCollection *mongo.Collection = configs.GetCollection(configs.DB, "manifests")
var manifestValidate = validator.New()

func Ping() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"data": "pong",
		})
	}
}

func CreateManifestMetadata() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var manifest models.Manifest
		defer cancel()

		// Validate the request
		if err := c.BindJSON(&manifest); err != nil {
			c.JSON(http.StatusBadRequest, responses.ManifestResponse{Status: http.StatusBadRequest, Message: "Error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		// Validate the required fields
		if validationErr := manifestValidate.Struct(&manifest); validationErr != nil {
			c.JSON(http.StatusBadRequest, responses.ManifestResponse{Status: http.StatusBadRequest, Message: "Error", Data: map[string]interface{}{"data": validationErr.Error()}})
			return
		}

		manifestId := uuid.NewString()

		// Create the new manifest
		newManifest := models.Manifest{
			UUID:  manifestId,
			Label: manifest.Label,
		}

		// Insert into db
		result, err := manifestCollection.InsertOne(ctx, newManifest)
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.ManifestResponse{Status: http.StatusInternalServerError, Message: "Error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		c.JSON(http.StatusCreated, responses.ManifestResponse{Status: http.StatusCreated, Message: "Success!", Data: map[string]interface{}{"mongodb-id": result.InsertedID, "uuid": manifestId}})
	}
}

func UploadImage() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Source
		file, err := c.FormFile("file")
		if err != nil {
			c.JSON(http.StatusBadRequest, responses.ManifestResponse{Status: http.StatusBadRequest, Message: "Error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		filename := filepath.Base(file.Filename)
		if err := c.SaveUploadedFile(file, configs.EnvImageStorePath()+filename); err != nil {
			c.JSON(http.StatusBadRequest, responses.ManifestResponse{Status: http.StatusBadRequest, Message: "Error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		c.JSON(http.StatusCreated, responses.ManifestResponse{Status: http.StatusCreated, Message: "Success", Data: map[string]interface{}{"data": filename}})
	}
}

func GetManifestMetadata() gin.HandlerFunc {
	return func(c *gin.Context) {

		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		manifestId := c.Param("manifestId")

		var manifest models.Manifest
		defer cancel()

		err := manifestCollection.FindOne(ctx, bson.M{"uuid": manifestId}).Decode(&manifest)
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.ManifestResponse{Status: http.StatusInternalServerError, Message: "Error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		c.JSON(http.StatusOK, responses.ManifestResponse{Status: http.StatusOK, Message: "Success", Data: map[string]interface{}{"data": manifest}})
	}
}
