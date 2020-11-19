package main

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Employee Data format
type Employee struct {
	Name       string
	Department string
	Salary     int
}

// Error validation call
func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {

	fmt.Println("Starting MongoDB connection program.")

	// Set client options
	clientOptions := options.Client().ApplyURI("mongodb://root:rootpassword@localhost:27017")

	// Connect to Mongo
	client, err := mongo.Connect(context.TODO(), clientOptions)
	checkErr(err)

	// Test the connection
	err = client.Ping(context.TODO(), nil)
	checkErr(err)

	fmt.Println("Connection to Mongo successful.")

	// Fetch handle for collection
	collection := client.Database("test").Collection("employees")

	// Data instanciation
	employee1 := Employee{"Steve", "IT", 75000}
	employee2 := Employee{"Marie", "HR", 50000}
	employee3 := Employee{"John", "Accounting", 56000}

	// Append a single employee document to collection
	insertResult, err := collection.InsertOne(context.TODO(), employee1)
	checkErr(err)
	fmt.Println("Inserted an employee entry: ", insertResult.InsertedID)

	// Append multiple employee documents to collection
	employees := []interface{}{employee2, employee3}
	insertManyResult, err := collection.InsertMany(context.TODO(), employees)
	checkErr(err)
	fmt.Println("Inserted many employee entries:", insertManyResult.InsertedIDs)

	// Modify an employees salary (Steve gets a raise, congrats Steve)
	filter := bson.D{{"name", "Steve"}}
	update := bson.D{
		primitive.E{"$set", bson.D{
			{"salary", 80000},
		}},
	}
	updateResult, err := collection.UpdateOne(context.TODO(), filter, update)
	checkErr(err)
	fmt.Printf("Matched %v employees and updated %v documents.\n", updateResult.MatchedCount, updateResult.ModifiedCount)

	// Fetch Steves entry to verify that his raise took place
	var result Employee
	err = collection.FindOne(context.TODO(), filter).Decode(&result)
	checkErr(err)
	fmt.Printf("Found an employee record: %+v\n", result)

	// Fire an employee (sorry Steve you werent worth the raise)
	deleteResult, err := collection.DeleteOne(context.TODO(), bson.D{{"name", "Steve"}})
	checkErr(err)
	fmt.Printf("Deleted %v employee record.\n", deleteResult.DeletedCount)

	// Find options to get all employees of the company
	findOptions := options.Find()
	findOptions.SetLimit(3)
	var results []Employee

	// Match the remaining documents
	cursor, err := collection.Find(context.TODO(), bson.D{{}}, findOptions)
	checkErr(err)

	// Iterate over the provided cursor for decoding
	for cursor.Next(context.TODO()) {
		var elem Employee
		err := cursor.Decode(&elem)
		checkErr(err)
		results = append(results, elem)
	}
	checkErr(cursor.Err())
	cursor.Close(context.TODO())
	fmt.Printf("Found several employee records: %v\n", results)

	// Disconnect the client
	err = client.Disconnect(context.TODO())
	checkErr(err)
	fmt.Println("Connection to Mongo closed.")
}
