package aws_func

import (
	"context"
	"fmt"

	"github.com/Kavanaugh-Frank/linkbush/template"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

func PutItem(client *dynamodb.Client, data template.PageData) error {
	// Marshal struct to DynamoDB attribute map
	item, err := attributevalue.MarshalMap(data)
	if err != nil {
		return fmt.Errorf("error marshalling struct: %w", err)
	}

	// Put item into DynamoDB
	_, err = client.PutItem(context.TODO(), &dynamodb.PutItemInput{
		TableName: aws.String("linkbush"),
		Item:      item,
	})
	if err != nil {
		return fmt.Errorf("error writing to DynamoDB: %w", err)
	}

	fmt.Println("Item written to DynamoDB.")
	return nil
}

func GetItem(client *dynamodb.Client, userID string) (template.PageData, error) {
	// Define the key
	key := map[string]types.AttributeValue{
		"UserID": &types.AttributeValueMemberS{Value: userID},
	}

	// Fetch the item
	result, err := client.GetItem(context.TODO(), &dynamodb.GetItemInput{
		TableName: aws.String("linkbush"), // replace with your table name
		Key:       key,
	})
	if err != nil {
		return template.PageData{}, fmt.Errorf("error getting item: %w", err)
	}

	if result.Item == nil {
		return template.PageData{}, fmt.Errorf("no item found")
	}

	// Unmarshal the item
	var pageData template.PageData
	err = attributevalue.UnmarshalMap(result.Item, &pageData)
	if err != nil {
		return template.PageData{}, fmt.Errorf("error unmarshalling item: %w", err)
	}

	return pageData, nil
}

func DeleteItem(client *dynamodb.Client, userID string) error {
	// Define the table name and key
	tableName := "linkbush" // Replace with your actual table name
	key := map[string]types.AttributeValue{
		"UserID": &types.AttributeValueMemberS{Value: userID},
	}

	// Create the DeleteItemInput
	input := &dynamodb.DeleteItemInput{
		TableName: &tableName,
		Key:       key,
	}

	// Call the DeleteItem API
	_, err := client.DeleteItem(context.TODO(), input)
	if err != nil {
		return err
	}

	return nil
}
