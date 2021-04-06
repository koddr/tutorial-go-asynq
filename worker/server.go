package main

import (
	"log"

	"tutorial-go-asynq/tasks"

	"github.com/hibiken/asynq"
)

func main() {
	// Create and configuring Redis connection.
	redisConnection := asynq.RedisClientOpt{
		Addr: "localhost:6379", // Redis server address
	}

	// Create and configuring Asynq worker server.
	worker := asynq.NewServer(redisConnection, asynq.Config{
		// Specify how many concurrent workers to use.
		Concurrency: 10,
		// Specify multiple queues with different priority.
		Queues: map[string]int{
			"critical": 6, // processed 60% of the time
			"default":  3, // processed 30% of the time
			"low":      1, // processed 10% of the time
		},
	})

	// Create a new task's mux instance.
	mux := asynq.NewServeMux()

	// Define a task handler for the welcome email task.
	mux.HandleFunc(
		tasks.TypeWelcomeEmail,       // task type
		tasks.HandleWelcomeEmailTask, // handler function
	)

	// Define a task handler for the reminder email task.
	mux.HandleFunc(
		tasks.TypeReminderEmail,       // task type
		tasks.HandleReminderEmailTask, // handler function
	)

	// Run worker server.
	if err := worker.Run(mux); err != nil {
		log.Fatal(err)
	}
}
