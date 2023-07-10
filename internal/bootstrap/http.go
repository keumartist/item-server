package bootstrap

import (
	"context"
	"log"
	"os"
	"time"

	handler "art-item/internal/handler/item"
	"art-item/internal/outbound"
	repository "art-item/internal/repository/item"
	service "art-item/internal/service/item"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func InitHTTPServer() error {
	mongodbURI := os.Getenv("MONGODB_URI")
	if mongodbURI == "" {
		log.Fatal("Database connection is not set")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongodbURI))
	if err != nil {
		return err
	}

	db := client.Database(os.Getenv("MONGODB_DATABASE"))
	itemRepo := repository.NewMongoDBItemRepository(db, "items")
	itemService := service.NewItemService(itemRepo)
	itemHandler := handler.NewItemHandler(itemService)

	authEndpoint := os.Getenv("AUTH_SERVICE_ENDPOINT") // ex) "auth-service.default.svc.cluster.local:50051"
	authClient := outbound.NewAuthServiceClient(authEndpoint)

	app := fiber.New()
	itemHandler.RegisterRoutes(app, authClient)

	err = app.Listen(":3000")
	if err != nil {
		return err
	}

	return nil
}
