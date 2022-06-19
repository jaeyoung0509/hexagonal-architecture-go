package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	mr "github.com/jaeyoung0509/hexagonal-go/serializer/repository/mongo"
	rr "github.com/jaeyoung0509/hexagonal-go/serializer/repository/redis"
	"github.com/jaeyoung0509/hexagonal-go/shortner"
)

func main() {

}

// repo <- service -> serializer -> http
func httpPort() string {
	port := "8000"
	if os.Getenv("PORT") != "" {
		port = os.Getenv("PORT")
	}
	return fmt.Sprintf(":%s", port)
}

func chooseRepo() shortner.RedirectRepository {
	switch os.Getenv("URL_DB") {
	case "redis":
		redisURL := os.Getenv("REDIS_URL")
		repo, err := rr.NewRedisRepository(redisURL)
		if err != nil {
			log.Fatal(err)
		}
		return repo
	case "mongo":
		mongoURL := os.Getenv("MONGO_URL")
		mongodb := os.Getenv("MONGO_DB")
		mongoTimeout, _ := strconv.Atoi(os.Getenv("MONGO_TIMEOUT"))
		repo, err := mr.NewMongoRepository(mongoURL, mongodb, mongoTimeout)
		if err != nil {
			log.Fatal(err)
		}
		return repo
	}
	return nil
}
