package cockroach

import "github.com/samulastech/cockroach/internal/entities"

type CockroachFCMMessaging struct {
}

func NewCockroachMessaging() *CockroachFCMMessaging {
	return &CockroachFCMMessaging{}
}

func (r *CockroachFCMMessaging) PushNotification(message *entities.CockroachPushNotificationDTO) error {
	return nil
}
