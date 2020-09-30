package main

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

// ItemInfo info
type ItemInfo struct {
	Plot   string  `json:"plot"`
	Rating float64 `json:"rating"`
}

// Item - Create struct to hold info about new item
type Item struct {
	Year  int      `json:"year"`
	Title string   `json:"title"`
	Info  ItemInfo `json:"info"`
}

func main() {

	// Initialize a session that the SDK will use to load
	// credentials from the shared credentials file ~/.aws/credentials
	// and region from the shared configuration file ~/.aws/config.
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
		Config: aws.Config{
			Region:   aws.String("us-west-2"),
			Endpoint: aws.String("http://localhost:8000"),
		},
	}))

	// Create DynamoDB client
	svc := dynamodb.New(sess)

	tableName := "Movies"
	movieName := "The Big New Movie"
	movieYear := "2015"

	result, err := svc.GetItem(&dynamodb.GetItemInput{
		TableName: aws.String(tableName),
		Key: map[string]*dynamodb.AttributeValue{
			"year": {
				N: aws.String(movieYear),
			},
			"title": {
				S: aws.String(movieName),
			},
		},
	})
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	if result.Item == nil {
		panic("Could not find '" + movieName + "'")
		// msg := "Could not find '" + movieName + "'"
		// return nil, errors.New(msg)
	}

	item := Item{}

	err = dynamodbattribute.UnmarshalMap(result.Item, &item)
	if err != nil {
		panic(fmt.Sprintf("Failed to unmarshal Record, %v", err))
	}

	fmt.Println("Found item:")
	fmt.Println("Year:  ", item.Year)
	fmt.Println("Title: ", item.Title)
	fmt.Println("Plot:  ", item.Info.Plot)
	fmt.Println("Rating:", item.Info.Rating)

}
