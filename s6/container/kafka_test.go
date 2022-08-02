package container

import (
	"context"
	"strings"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/segmentio/kafka-go"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/testcontainers/testcontainers-go"
)

func TestWithKafka(t *testing.T) {
	kafkaContainer := testcontainers.NewLocalDockerCompose(
		[]string{"testdata/docker-compose.yml"},
		strings.ToLower(uuid.New().String()),
	)
	execError := kafkaContainer.WithCommand([]string{"up", "-d"}).Invoke()
	require.NoError(t, execError.Error)

	// kafka starts ver slow
	time.Sleep(time.Minute)
	defer destroyKafka(kafkaContainer)

	// test write
	w := &kafka.Writer{
		Addr:     kafka.TCP("localhost:9092"),
		Topic:    "test-topic",
		Balancer: &kafka.LeastBytes{},
	}

	err := w.WriteMessages(context.Background(),
		kafka.Message{
			Key:   []byte("Key-A"),
			Value: []byte("Hello World!"),
		},
		kafka.Message{
			Key:   []byte("Key-B"),
			Value: []byte("One!"),
		},
		kafka.Message{
			Key:   []byte("Key-C"),
			Value: []byte("Two!"),
		},
	)
	assert.NoError(t, err)
	err = w.Close()
	assert.NoError(t, err)

	// test read
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers:   []string{"localhost:9092"},
		Topic:     "test-topic",
		Partition: 0,
		MinBytes:  10e3, // 10KB
		MaxBytes:  10e6, // 10MB
	})

	m, err := r.ReadMessage(context.Background())
	assert.NoError(t, err)
	assert.NotEmpty(t, m)

}

func destroyKafka(compose *testcontainers.LocalDockerCompose) {
	compose.Down()
	time.Sleep(1 * time.Second)
}
