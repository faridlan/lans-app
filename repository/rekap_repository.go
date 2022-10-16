package repository

import (
	"context"

	"github.com/faridlan/lans-app/model/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type RekapRepository interface {
	CreateOne(ctx context.Context, rekap domain.Rekap) (*domain.Rekap, error)
	UpdateOne(ctx context.Context, rekap domain.Rekap) (*domain.Rekap, error)
	DeleteOne(ctx context.Context, rekap domain.Rekap) error
	FindOne(ctx context.Context, rekapId primitive.ObjectID) (*domain.Rekap, error)
	FindMany(ctx context.Context) ([]domain.Rekap, error)
}
