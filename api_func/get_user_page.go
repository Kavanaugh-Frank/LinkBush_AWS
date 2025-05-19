package api_func

import (
	"fmt"
	"net/http"

	"github.com/Kavanaugh-Frank/linkbush/aws_func"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/gin-gonic/gin"
)

func Get_User_Page(c *gin.Context, client *dynamodb.Client) {
	// Get user_id from the URL parameter
	userID := c.Param("user_id")

	// Get item from DynamoDB
	response, err := aws_func.GetItem(client, userID)
	if err != nil {
		fmt.Println("Error getting item:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch data"})
		return
	}

	c.HTML(http.StatusOK, "index", response)
}
