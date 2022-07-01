package services

import (
	"context"
	"errors"

	"github.com/cmani097/go-mongo/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserService interface {
	CreateUser(*models.User) error
	GetUser(*string) (*models.User, error)
	GetAll() ([]*models.User, error)
	UpdateUser(*models.User) error
	DeleteUser(*string) error
}
type UserServiceImpl struct {
	usercollection *mongo.Collection
	ctx            context.Context
}

func NewUserService(usercollection *mongo.Collection, ctx context.Context) UserService {
	return &UserServiceImpl{
		usercollection: usercollection,
		ctx:            ctx,
	}
}

func (userservice *UserServiceImpl) CreateUser(user *models.User) error {
	_, err := userservice.usercollection.InsertOne(userservice.ctx, user)
	return err
}

func (userservice *UserServiceImpl) GetUser(name *string) (*models.User, error) {
	var user *models.User
	query := bson.D{bson.E{Key: "user_name", Value: name}}
	err := userservice.usercollection.FindOne(userservice.ctx, query).Decode(&user)
	return user, err
}

func (userservice *UserServiceImpl) GetAll() ([]*models.User, error) {
	var users []*models.User
	cursor, err := userservice.usercollection.Find(userservice.ctx, bson.D{{}})
	if err != nil {
		return nil, err
	}
	for cursor.Next(userservice.ctx) {
		var user models.User
		err := cursor.Decode(&user)
		if err != nil {
			return nil, err
		}
		users = append(users, &user)
	}
	if err := cursor.Err(); err != nil {
		return nil, err
	}

	cursor.Close(userservice.ctx)

	if len(users) == 0 {
		return nil, errors.New("Documents not found")
	}
	return users, nil
}

func (userservice *UserServiceImpl) UpdateUser(user *models.User) error {
	filter := bson.D{bson.E{Key: "user_name", Value: user.Name}}
	update := bson.D{bson.E{Key: "$set", Value: bson.D{bson.E{Key: "user_name", Value: user.Name}, bson.E{Key: "user_age", Value: user.Age}, bson.E{Key: "user_address", Value: user.Address}}}}
	result, _ := userservice.usercollection.UpdateOne(userservice.ctx, filter, update)
	if result.MatchedCount != 1 {
		return errors.New("No matched document found for update")
	}
	return nil
}

func (userservice *UserServiceImpl) DeleteUser(name *string) error {
	filter := bson.D{bson.E{Key: "user_name", Value: name}}
	result, _ := userservice.usercollection.DeleteOne(userservice.ctx, filter)
	if result.DeletedCount != 1 {
		return errors.New("No matched document found for delete")
	}
	return nil
}
