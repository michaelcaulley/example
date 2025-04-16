package model

import (
	"context"
	"example/internal/ent"
)

type Remindable interface {
	ent.Noder
	Reminders(
		ctx context.Context,
	) ([]*ent.Reminder, error)
	IsRemindable()
}
