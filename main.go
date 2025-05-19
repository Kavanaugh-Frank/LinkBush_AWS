package main

import (
	"context"
	"fmt"
	htmlTemplate "html/template"

	"github.com/Kavanaugh-Frank/linkbush/api_func"
	"github.com/Kavanaugh-Frank/linkbush/template"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	ginadapter "github.com/awslabs/aws-lambda-go-api-proxy/gin"
	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()

	// router.Use(cors.New(cors.Config{
	// 	AllowOrigins:     []string{"*"}, // or use specific origins like http://localhost:3000
	// 	AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
	// 	AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
	// 	ExposeHeaders:    []string{"Content-Length"},
	// 	AllowCredentials: true,
	// 	MaxAge:           12 * time.Hour,
	// }))

	// Load AWS config
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		fmt.Println("Error loading AWS config:", err)
		return
	}

	// Create DynamoDB client
	dynamoDBclient := dynamodb.NewFromConfig(cfg)
	// create the s3 client
	s3client := s3.NewFromConfig(cfg)

	tmpl := htmlTemplate.Must(htmlTemplate.New("index").Parse(template.UserHTMTL))
	router.SetHTMLTemplate(tmpl)

	router.POST("/new_page", func(c *gin.Context) {
		api_func.New(c, dynamoDBclient, s3client)
	})

	router.GET("/healthz", func(c *gin.Context) {
		c.String(200, "ok")
	})

	router.POST("/edit_page/:user_id", func(c *gin.Context) {
		api_func.Edit(c, dynamoDBclient, s3client)
	})

	router.GET("/:user_id", func(c *gin.Context) {
		api_func.Get_User_Page(c, dynamoDBclient)
	})

	router.POST("/delete_page/:user_id", func(c *gin.Context) {
		api_func.Delete(c, dynamoDBclient)
	})

	lambdaHandler := ginadapter.NewV2(router)

	// start
	lambda.Start(lambdaHandler.ProxyWithContext)
}
