package main

import (
	"fmt"
	"log"
	_"github.com/lib/pq"
	"github.com/bxcodec/faker/v4"
	"github.com/go-redis/redis/v9"
	"github.com/jmoiron/sqlx"
	"github.com/nurmuhammaddeveloper/mudium_user_service/config"
	"github.com/nurmuhammaddeveloper/mudium_user_service/storage"
	"github.com/nurmuhammaddeveloper/mudium_user_service/storage/repo"
)

func main() {
	cfg := config.Load(".")

	psqlUrl := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.Postgres.Host,
		cfg.Postgres.Port,
		cfg.Postgres.User,
		cfg.Postgres.Password,
		cfg.Postgres.Database,
	)

	psqlConn, err := sqlx.Connect("postgres", psqlUrl)
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}

	rdb := redis.NewClient(&redis.Options{
		Addr: cfg.Redis.Addr,
	})
	fmt.Println(rdb)

	memory := storage.NewStoragePg(psqlConn)
	data, err := memory.User().Create(&repo.User{
		FirstName:   faker.FirstName(),
		LastName:    faker.LastName(),
		PhoneNumber: faker.CCNumber(),
		Gender:      "male",
		Email:       faker.Email(),
		Password:    faker.Sentence(),
		UserName:    faker.Name(),
		Type:        "superadmin",
	})

	if err != nil {
		panic(err)
	}
	fmt.Println("user created")
	fmt.Println(data)
}
