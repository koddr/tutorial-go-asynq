package tasks

import (
	"time"

	"github.com/hibiken/asynq"
)

const (
	// TypeWelcomeEmail is a name of the task type
	// for sending a welcome email.
	TypeWelcomeEmail = "email:welcome"

	// TypeReminderEmail is a name of the task type
	// for sending a reminder email.
	TypeReminderEmail = "email:reminder"
)

// NewWelcomeEmailTask task payload for a new welcome email.
func NewWelcomeEmailTask(id int) *asynq.Task {
	// Specify task payload.
	payload := map[string]interface{}{
		"user_id": id, // set user ID
	}

	// Return a new task with given type and payload.
	return asynq.NewTask(TypeWelcomeEmail, payload)
}

// NewReminderEmailTask task payload for a reminder email.
func NewReminderEmailTask(id int, ts time.Time) *asynq.Task {
	// Specify task payload.
	payload := map[string]interface{}{
		"user_id": id,          // set user ID
		"sent_in": ts.String(), // set time to sending
	}

	// Return a new task with given type and payload.
	return asynq.NewTask(TypeReminderEmail, payload)
}
