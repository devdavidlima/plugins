package crudDynamodb

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/devdavidlima/plugins/utils"
)

// Setting local test
type AwsConfig struct {
	DBEndpoint string
	DBRegion   string
}

// Generic model to represent a database data
type Model struct {
	TableName  string
	PrimaryKey string
	Svc        *dynamodb.DynamoDB
}

// -> NewModel: Create db models using a struct. It`s like a "constructor of my interface Model"
func NewModel(awscfg AwsConfig, tableName, primaryKey string) *Model {
	sess, err := session.NewSession(&aws.Config{
		Endpoint: aws.String(awscfg.DBEndpoint),
		Region:   aws.String(awscfg.DBRegion),
	})
	utils.CheckErrAbortProgram(err, "Unable to create a db model")

	return &Model{
		TableName:  tableName,
		PrimaryKey: primaryKey,
		Svc:        dynamodb.New(sess),
	}
}

// -> CreateItem: insert a new item respecting the PK
func (m *Model) CreateItem(data interface{}) error {
	input := &dynamodb.PutItemInput{
		TableName: aws.String(m.TableName),
		Item: map[string]*dynamodb.AttributeValue{
			m.PrimaryKey: {S: aws.String(fmt.Sprintf("%v", data.(map[string]interface{})[m.PrimaryKey]))},
		},
	}

	_, err := m.Svc.PutItem(input)
	utils.CheckErr(err, "")

	return nil
}

// -> ReadItem: Get item by PK
func (m *Model) ReadItem(id interface{}) (map[string]interface{}, error) {
	input := &dynamodb.GetItemInput{
		TableName: aws.String(m.TableName),
		Key: map[string]*dynamodb.AttributeValue{
			m.PrimaryKey: {S: aws.String(fmt.Sprintf("%v", id))},
		},
	}

	result, err := m.Svc.GetItem(input)
	utils.CheckErr(err, "Unable to get item")

	item := make(map[string]interface{})
	for k, v := range result.Item {
		item[k] = v.String()
	}

	return item, nil
}

// EditItem: Edit an item by PK
func (m *Model) EditItem(id interface{}, data interface{}) error {
	key := map[string]*dynamodb.AttributeValue{
		m.PrimaryKey: {S: aws.String(fmt.Sprintf("%v", id))},
	}

	updateExpression := "SET "
	expressionAttributeValues := map[string]*dynamodb.AttributeValue{}
	i := 0

	// Construct the update expression and expression attribute values
	for k, v := range data.(map[string]interface{}) {
		i++
		attributeName := fmt.Sprintf(":val%d", i)
		updateExpression += fmt.Sprintf("%s = %s, ", k, attributeName)
		expressionAttributeValues[attributeName] = &dynamodb.AttributeValue{S: aws.String(fmt.Sprintf("%v", v))}
	}

	// Remove the trailing comma and space from the update expression
	updateExpression = updateExpression[:len(updateExpression)-2]

	input := &dynamodb.UpdateItemInput{
		TableName:                 aws.String(m.TableName),
		Key:                       key,
		ExpressionAttributeValues: expressionAttributeValues,
		UpdateExpression:          aws.String(updateExpression),
	}

	_, err := m.Svc.UpdateItem(input)
	utils.CheckErr(err, "Unable to update item")

	return nil
}

// -> DelItem: delete an item by PK
func (m *Model) DelItem(id interface{}) error {
	key := map[string]*dynamodb.AttributeValue{
		m.PrimaryKey: {S: aws.String(fmt.Sprintf("%v", id))},
	}

	input := &dynamodb.DeleteItemInput{
		TableName: aws.String(m.TableName),
		Key:       key,
	}

	_, err := m.Svc.DeleteItem(input)
	utils.CheckErr(err, "")

	return nil
}
