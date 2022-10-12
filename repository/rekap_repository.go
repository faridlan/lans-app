package repository

import (
	"context"

	"github.com/faridlan/lans-app/model/domain"
	"go.mongodb.org/mongo-driver/mongo"
)

type RekapRepository interface {
	CreateOne(ctx context.Context, db *mongo.Database, rekap domain.Rekap) (*domain.Rekap, error)
	UpdateOne(ctx context.Context, db *mongo.Database, rekap domain.Rekap) (domain.Rekap, error)
	DeleteOne(ctx context.Context, db *mongo.Database, rekap domain.Rekap) error
	FindOne(ctx context.Context, db *mongo.Database, rekapId string) (domain.Rekap, error)
	FindMany(ctx context.Context, db *mongo.Database) ([]domain.Rekap, error)
}
