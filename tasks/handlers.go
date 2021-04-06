package tasks

import (
	"context"
	"fmt"

	"github.com/hibiken/asynq"
)

// HandleWelcomeEmailTask handler for welcome email task.
func HandleWelcomeEmailTask(c context.Context, t *asynq.Task) error {
	// Get user ID from given task.
	id, err := t.Payload.GetInt("user_id")
	if err != nil {
		return err
	}

	// Dummy message to the worker's output.
	fmt.Printf("Send Welcome Email to User ID %d\n", id)

	return nil
}

// HandleReminderEmailTask for reminder email task.
func HandleReminderEmailTask(c context.Context, t *asynq.Task) error {
	// Get int with the user ID from the given task.
	id, err := t.Payload.GetInt("user_id")
	if err != nil {
		return err
	}

	// Get string with the sent time from the given task.
	time, err := t.Payload.GetString("sent_in")
	if err != nil {
		return err
	}

	// Dummy message to the worker's output.
	fmt.Printf("Send Reminder Email to User ID %d\n", id)
	fmt.Printf("Reason: time is up (%v)\n", time)

	return nil
}
