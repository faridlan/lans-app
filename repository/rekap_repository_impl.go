package repository

import (
	"context"
	"errors"

	"github.com/faridlan/lans-app/model/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type RekapRepositoryImpl struct {
	DB *mongo.Collection
}

func NewRekapRepository(DB *mongo.Collection) RekapRepository {
	return &RekapRepositoryImpl{
		DB: DB,
	}
}

func (repository *RekapRepositoryImpl) CreateOne(ctx context.Context, rekap domain.Rekap) (*domain.Rekap, error) {
	result, err := repository.DB.InsertOne(ctx, rekap)
	if err != nil {
		return nil, err
	}

	rekap.Id = result.InsertedID.(primitive.ObjectID)

	return &rekap, nil
}

func (repository *RekapRepositoryImpl) UpdateOne(ctx context.Context, rekap domain.Rekap) (*domain.Rekap, error) {
	filter := bson.M{"_Id": rekap.Id}
	field := bson.M{"$set": rekap}
	result := repository.DB.FindOneAndUpdate(ctx, filter, field, options.FindOneAndUpdate().SetReturnDocument(1))
	err := result.Decode(&rekap)
	if err != nil {
		return nil, err
	}

	return &rekap, nil
}

func (repository *RekapRepositoryImpl) DeleteOne(ctx context.Context, rekap domain.Rekap) error {
	result, err := repository.DB.DeleteOne(ctx, bson.M{"_id": rekap.Id})
	if err != nil {
		return errors.New(err.Error())
	}

	if result.DeletedCount == 0 {
		return errors.New("no document deleted")
	}

	return nil
}

func (repository *RekapRepositoryImpl) FindOne(ctx context.Context, rekapId primitive.ObjectID) (*domain.Rekap, error) {

	filter := bson.D{{Key: "_id", Value: rekapId}}
	result := repository.DB.FindOne(ctx, filter)

	rekap := domain.Rekap{}

	err := result.Decode(&rekap)
	if err != nil {
		return nil, err
	}
	return &rekap, nil

	// cursor, err := repository.DB.Find(ctx, bson.M{"_id": rekapId})
	// if err != nil {
	// 	return nil, err
	// }
	// rekap := domain.Rekap{}
	// if cursor.Next(ctx) {
	// 	err := cursor.Decode(&rekap)
	// 	if err != nil {
	// 		return nil, err
	// 	}

	// 	return &rekap, nil
	// } else {
	// 	return nil, errors.New("rekap not found")
	// }
}

func (repository *RekapRepositoryImpl) FindMany(ctx context.Context) ([]domain.Rekap, error) {
	cursor, err := repository.DB.Find(ctx, bson.M{})
	if err != nil {
		return nil, errors.New(err.Error())
	}

	rekaps := []domain.Rekap{}
	for cursor.Next(ctx) {
		rekap := domain.Rekap{}
		err := cursor.Decode(&rekap)
		if err != nil {
			return nil, errors.New(err.Error())
		}

		rekaps = append(rekaps, rekap)

	}
	return rekaps, nil
}
