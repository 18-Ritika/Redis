package students

import (
	"Training/Redis/Redis/internal/models"
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis"
	"strconv"
)

type Stores struct {
}

func New() *Stores {
	return &Stores{}
}

func redisConnection() *redis.Client {
	Client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	return Client
}

func (st *Stores) Post(ctx context.Context, std models.Student) (int, error) {
	Client := redisConnection()
	json, err := json.Marshal(std)
	if err != nil {
		return 0, err
	}
	err = Client.Set(std.ID, json, 0).Err()
	fmt.Println(err)
	if err != nil {
		return 0, err
	}
	id, _ := strconv.Atoi(std.ID)
	return id, nil
}

func (st *Stores) Get(ctx context.Context, ID string) (string, error) {
	Client := redisConnection()
	resp, err := Client.Get(ID).Result()
	if err != nil {
		return "", err
	}
	return resp, nil
}

func (st *Stores) Delete(ctx context.Context, ID string) error {
	Client := redisConnection()
	err := Client.Del(ID).Err()
	if err != nil {
		return err
	}
	return nil
}
