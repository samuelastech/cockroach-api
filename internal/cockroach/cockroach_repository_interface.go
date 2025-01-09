package cockroach

import "github.com/samulastech/cockroach/internal/entities"

type CockroachRepository interface {
	InsertCockroach(in *entities.InsertCockroachDTO) error
}

type CockroachMessaging interface {
	PushNotification(message *entities.CockroachPushNotificationDTO) error
}
