package api_func

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Kavanaugh-Frank/linkbush/aws_func"
	"github.com/Kavanaugh-Frank/linkbush/hash"
	"github.com/Kavanaugh-Frank/linkbush/template"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/gin-gonic/gin"
)

func Edit(c *gin.Context, client *dynamodb.Client, s3Client *s3.Client) {
	var data template.PageData

	// Parse form data (with file + JSON fields)
	if err := c.Request.ParseMultipartForm(10 << 20); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to parse form data"})
		return
	}

	// Manually get fields from the multipart form
	data.UserID = c.PostForm("UserID")
	data.Title = c.PostForm("Title")
	data.Password = c.PostForm("Password")

	// You may also want to parse JSON string manually if Links come in JSON format
	linksJson := c.PostForm("Links")
	if linksJson != "" {
		if err := json.Unmarshal([]byte(linksJson), &data.Links); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to parse links JSON"})
			return
		}
	}

	// Check if user exists
	existing, err := aws_func.GetItem(client, data.UserID)
	if err != nil {
		c.JSON(http.StatusConflict, gin.H{"error": "Failed to get user information"})
		return
	}

	// Check if the stored password matches
	if !hash.CheckPasswordHash(data.Password, existing.Password) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Incorrect Password"})
		return
	}

	// File upload
	file, _, err := c.Request.FormFile("ProfileImage")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to retrieve file"})
		return
	}
	defer file.Close()

	// Upload to S3
	s3Key := "profile-images/" + data.UserID + ".jpg"
	err = aws_func.UploadToS3(file, s3Key, s3Client)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to upload image to S3"})
		return
	}

	// Set profile image URL
	// https://linkbush-images.s3.us-east-1.amazonaws.com/profile-images/{image_name}
	data.ProfileImageURL = fmt.Sprintf("https://%s.s3.us-east-1.amazonaws.com/%s", aws_func.S3BucketName, s3Key)

	// Hash password
	hashedPassword, err := hash.HashPassword(data.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
		return
	}
	data.Password = hashedPassword

	// Put new item in DynamoDB
	err = aws_func.PutItem(client, data)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create new page"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Page created successfully"})
}
