package repository

import (
	"context"
	"errors"

	"github.com/faridlan/lans-app/model/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type RekapRepositoryImpl struct {
	// DB *mongo.Collection
}

func NewRekapRepository() RekapRepository {
	return &RekapRepositoryImpl{
		// DB: &mongo.Collection{},
	}
}

func (repository *RekapRepositoryImpl) CreateOne(ctx context.Context, DB *mongo.Collection, rekap domain.Rekap) (*domain.Rekap, error) {
	result, err := DB.InsertOne(ctx, rekap)
	if err != nil {
		return nil, errors.New(err.Error())
	}

	id := result.InsertedID
	rekap.Id = id.(primitive.ObjectID)

	return &rekap, nil
}

func (repository *RekapRepositoryImpl) UpdateOne(ctx context.Context, DB *mongo.Collection, rekap domain.Rekap) (*domain.Rekap, error) {
	filter := bson.M{"_Id": rekap.Id}
	field := bson.M{"$set": rekap}
	_, err := DB.UpdateOne(ctx, filter, field)
	if err != nil {
		return nil, errors.New(err.Error())
	}

	return &rekap, nil
}

func (repository *RekapRepositoryImpl) DeleteOne(ctx context.Context, DB *mongo.Collection, rekap domain.Rekap) error {
	_, err := DB.DeleteOne(ctx, bson.M{"_id": rekap.Id})
	if err != nil {
		return errors.New(err.Error())
	}

	return nil
}

func (repository *RekapRepositoryImpl) FindOne(ctx context.Context, DB *mongo.Collection, rekapId primitive.ObjectID) (*domain.Rekap, error) {
	cursor, err := DB.Find(ctx, bson.M{"_id": rekapId})
	if err != nil {
		return nil, errors.New(err.Error())
	}

	rekap := domain.Rekap{}
	if cursor.Next(ctx) {
		err := cursor.Decode(&rekap)
		if err != nil {
			return nil, errors.New(err.Error())
		}

		return &rekap, nil
	} else {
		return nil, errors.New("rekap not found")
	}
}

func (repository *RekapRepositoryImpl) FindMany(ctx context.Context, DB *mongo.Collection) ([]*domain.Rekap, error) {
	cursor, err := DB.Find(ctx, bson.M{})
	if err != nil {
		return nil, errors.New(err.Error())
	}

	rekaps := []*domain.Rekap{}
	for cursor.Next(ctx) {
		rekap := &domain.Rekap{}
		err := cursor.Decode(&rekap)
		if err != nil {
			return nil, errors.New(err.Error())
		}

		rekaps = append(rekaps, rekap)

	}
	return rekaps, nil
}
