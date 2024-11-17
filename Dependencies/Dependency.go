package Dependencies

import (
	"go.mongodb.org/mongo-driver/mongo"
	"gorm.io/gorm"
)

type Dependency struct {
	DB    *gorm.DB
	Mongo *mongo.Client
}

func InitDependencies(options ...Option) *Dependency {
	deps := &Dependency{}
	for _, apply := range options {
		apply(deps)
	}

	return deps
}
