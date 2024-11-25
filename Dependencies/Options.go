package Dependencies

import "2024_akutansi_project/Config"

type Option func(provider *Dependency)

func WithDB() Option {
	return func(deps *Dependency) {
		Config.Connect()
		db := Config.DB
		if db == nil {
			panic("Failed to connect to database!")
		}
		deps.DB = db
	}
}

func WithMongo() Option {
	return func(deps *Dependency) {
		deps.Mongo = Config.InitMongoDB()
	}
}

func WithRedis() Option {
	return func(deps *Dependency) {
		deps.Redis = Config.InitRedis()
	}
}
