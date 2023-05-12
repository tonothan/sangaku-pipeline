package controllers

import (
	"context"
	"net/http"
	"pipeline/configs"
	"pipeline/models"
	"pipeline/responses"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var sangakuCollection *mongo.Collection = configs.GetCollection(configs.DB, "sangakus")
var sangakuValidate = validator.New()

func Ping() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(200, gin.H{
			"data": "Hello world",
		})
	}
}

func CreateSangaku() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var sangaku models.Sangaku
		defer cancel()

		//validate the request body
		if err := c.BindJSON(&sangaku); err != nil {
			c.JSON(http.StatusBadRequest, responses.SangakuResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		//use the validator library to validate required fields
		if validationErr := sangakuValidate.Struct(&sangaku); validationErr != nil {
			c.JSON(http.StatusBadRequest, responses.SangakuResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": validationErr.Error()}})
			return
		}

		newSangaku := models.Sangaku{
			Id: primitive.NewObjectID(),

			// Time-Space metadata
			Temple_en:   sangaku.Temple_en,
			Temple_jp:   sangaku.Temple_jp,
			Location_en: sangaku.Location_en,
			Location_jp: sangaku.Location_jp,
			Year_from:   sangaku.Year_from,
			Year_to:     sangaku.Year_to,

			// Content data
			Transcription_en: sangaku.Transcription_en,
			Transcription_jp: sangaku.Transcription_jp,
			Problem_en:       sangaku.Problem_en,
			Problem_jp:       sangaku.Problem_jp,
			Formula_en:       sangaku.Formula_en,
			Formula_jp:       sangaku.Formula_jp,
			Solution_en:      sangaku.Solution_en,
			Solution_jp:      sangaku.Solution_jp,
			School_en:        sangaku.School_en,
			School_jp:        sangaku.School_jp,
			Author_en:        sangaku.Author_en,
			Author_jp:        sangaku.Author_jp,

			// Medium metadata
			Dimension:   sangaku.Dimension,
			Material_en: sangaku.Material_en,
			Material_jp: sangaku.Material_jp,
		}

		result, err := sangakuCollection.InsertOne(ctx, newSangaku)
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.SangakuResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		c.JSON(http.StatusCreated, responses.SangakuResponse{Status: http.StatusCreated, Message: "success", Data: map[string]interface{}{"data": result}})
	}
}

func UploadImages() gin.HandlerFunc {
	return func(c *gin.Context) {

		file, _ := c.FormFile("data")

		print(file)

	}
}
