package services

import (
	"context"
	"fmt"

	"github.com/slice-mani/postgres-jinzhu/models"
	"go.mongodb.org/mongo-driver/mongo"
)

type SampleService interface {
	TestService(*models.User) error
}
type SampleServiceImpl struct {
	usercollection *mongo.Collection
	ctx            context.Context
}

func NewSampleService(usercollection *mongo.Collection, ctx context.Context) SampleService {
	return &SampleServiceImpl{
		usercollection: usercollection,
		ctx:            ctx,
	}
}

func (sampleservice *SampleServiceImpl) TestService(user *models.User) error {
	fmt.Println("CreateUser of sampleservice called")
	return nil
}
