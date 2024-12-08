package Dependencies

import (
	firebase "firebase.google.com/go"
	"firebase.google.com/go/messaging"
	"github.com/redis/go-redis/v9"
	"go.mongodb.org/mongo-driver/mongo"
	"gorm.io/gorm"
)

type Dependency struct {
	DB        *gorm.DB
	Mongo     *mongo.Client
	Redis     *redis.Client
	Firebase  *firebase.App
	Messaging *messaging.Client
}

func InitDependencies(options ...Option) *Dependency {
	deps := &Dependency{}
	for _, apply := range options {
		apply(deps)
	}

	return deps
}
