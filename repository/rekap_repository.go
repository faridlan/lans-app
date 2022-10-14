package repository

import (
	"context"

	"github.com/faridlan/lans-app/model/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type RekapRepository interface {
	// CreateOne(ctx context.Context, rekap domain.Rekap) (*domain.Rekap, error)
	// UpdateOne(ctx context.Context, rekap domain.Rekap) (*domain.Rekap, error)
	// DeleteOne(ctx context.Context, rekap domain.Rekap) error
	// FindOne(ctx context.Context, rekapId primitive.ObjectID) (*domain.Rekap, error)
	// FindMany(ctx context.Context) ([]*domain.Rekap, error)
	CreateOne(ctx context.Context, DB *mongo.Collection, rekap domain.Rekap) (*domain.Rekap, error)
	UpdateOne(ctx context.Context, DB *mongo.Collection, rekap domain.Rekap) (*domain.Rekap, error)
	DeleteOne(ctx context.Context, DB *mongo.Collection, rekap domain.Rekap) error
	FindOne(ctx context.Context, DB *mongo.Collection, rekapId primitive.ObjectID) (*domain.Rekap, error)
	FindMany(ctx context.Context, DB *mongo.Collection) ([]*domain.Rekap, error)
}
