package api_func

import (
	"fmt"
	"net/http"

	"github.com/Kavanaugh-Frank/linkbush/aws_func"
	"github.com/Kavanaugh-Frank/linkbush/hash"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/gin-gonic/gin"
)

func Delete(c *gin.Context, client *dynamodb.Client) {
	// Get user_id from the URL parameter
	userID := c.Param("user_id")

	// Parse the password from the request body
	var requestData struct {
		Password string `json:"password"`
	}
	if err := c.ShouldBindJSON(&requestData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Retrieve the existing item from DynamoDB
	existingItem, err := aws_func.GetItem(client, userID)
	if err != nil {
		fmt.Println("Error retrieving item from DynamoDB:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve existing data"})
		return
	}

	// Compare the provided password with the stored hashed password
	if !hash.CheckPasswordHash(requestData.Password, existingItem.Password) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Incorrect password"})
		return
	}

	// Delete item from DynamoDB
	err = aws_func.DeleteItem(client, userID)
	if err != nil {
		fmt.Println("Error deleting item from DynamoDB:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete page"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Page deleted successfully"})
}
