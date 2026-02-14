package data

import (
	"context"
	"time"

	"taskmanagement/config"
	"taskmanagement/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

func getTaskCollection() *mongo.Collection {
	if config.Client == nil {
		panic("MongoDB client not initialized - call config.ConnectDB() first")
	}
	return config.Client.Database(config.DBName).Collection("tasks")
}

func GetTasks() ([]models.Task, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cursor, err := getTaskCollection().Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var tasks []models.Task
	if err = cursor.All(ctx, &tasks); err != nil {
		return nil, err
	}
	return tasks, nil
}

func GetTaskByID(id string) (models.Task, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	objID, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return models.Task{}, err
	}
	var task models.Task
	err = getTaskCollection().FindOne(ctx, bson.M{"_id": objID}).Decode(&task)
	return task, err

}

func CreateTask(task models.Task) (*mongo.InsertOneResult, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	return getTaskCollection().InsertOne(ctx, task)
}

func UpdateTask(id string, update bson.M) (*mongo.UpdateResult, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	return getTaskCollection().UpdateOne(ctx, bson.M{"_id": objID}, bson.M{"$set": update})
}

func DeleteTask(id string) (*mongo.DeleteResult, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	return getTaskCollection().DeleteOne(ctx, bson.M{"_id": objID})
}
