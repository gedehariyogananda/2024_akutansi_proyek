package Dependencies

import (
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type Dependency struct {
	DB    *gorm.DB
	Redis *redis.Client
}

func InitDepedencies(options ...Option) *Dependency {
	deps := &Dependency{}
	for _, apply := range options {
		apply(deps)
	}

	return deps

}
