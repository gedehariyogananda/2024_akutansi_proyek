package Dependencies

import (
	"github.com/redis/go-redis/v9"
	"go.mongodb.org/mongo-driver/mongo"
	"gorm.io/gorm"
)

type Dependency struct {
	DB    *gorm.DB
	Mongo *mongo.Client
	Redis *redis.Client
}

func InitDependencies(options ...Option) *Dependency {
	deps := &Dependency{}
	for _, apply := range options {
		apply(deps)
	}

	return deps
}
