package Config

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/redis/go-redis/v9"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {

	godotenv.Load()

	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbHost := os.Getenv("DB_HOST")
	dbName := os.Getenv("DB_NAME")
	dbPort := os.Getenv("DB_PORT")

	dsn := fmt.Sprintf("host=%v user=%v password=%v dbname=%v port=%v sslmode=disable", dbHost, dbUser, dbPass, dbName, dbPort)

	fmt.Println(dsn)
	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  dsn,
		PreferSimpleProtocol: true,
	}), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	DB = db
	fmt.Println("Database connected!")

}

func InitMongoDB() *mongo.Client {
	mongoHost := os.Getenv("MONGO_HOST")
	mongoPort := os.Getenv("MONGO_PORT")
	mongoUserName := os.Getenv("MONGO_USERNAME")
	mongoPassword := os.Getenv("MONGO_PASSWORD")

	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(fmt.Sprintf("mongodb://%s:%s@%s:%s", mongoUserName, mongoPassword, mongoHost, mongoPort)))
	if err != nil {
		log.Fatal(err)
	}

	if errPing := client.Ping(context.Background(), nil); errPing != nil {
		log.Fatal("Could not connect to MongoDB:", errPing)
	}
	fmt.Println("Connected to MongoDB!")

	return client
}

func InitRedis() *redis.Client {
	RedisClient := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_HOST") + ":" + os.Getenv("REDIS_PORT"),
		Password: os.Getenv("REDIS_PASS"),
		Username: os.Getenv("REDIS_UNAME"),
		DB:       0,
	})

	_, err := RedisClient.Ping(context.Background()).Result()
	if err != nil {
		log.Fatalf("log: failed to connect to redis: %v", err)
	}

	log.Println("log: Redis connected!")
	return RedisClient
}
