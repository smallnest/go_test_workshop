package gnomock_test

import (
	"fmt"
	"testing"

	redisclient "github.com/go-redis/redis/v7"
	"github.com/orlangure/gnomock"
	"github.com/orlangure/gnomock/preset/redis"
)

func TestRedis(t *testing.T) {
	vs := make(map[string]interface{})

	vs["a"] = "foo"
	vs["b"] = 42
	vs["c"] = true

	p := redis.Preset(redis.WithValues(vs))
	container, _ := gnomock.Start(p,
		gnomock.WithDebugMode(),
		gnomock.WithUseLocalImagesFirst(),
	)

	defer func() { _ = gnomock.Stop(container) }()

	addr := container.DefaultAddress()
	client := redisclient.NewClient(&redisclient.Options{Addr: addr})

	fmt.Println(client.Get("a").Result())

	var number int

	err := client.Get("b").Scan(&number)
	fmt.Println(number, err)

	var flag bool

	err = client.Get("c").Scan(&flag)
	fmt.Println(flag, err)
}
