package cockroach

import (
	"github.com/samulastech/cockroach/internal/entities"
	"log"
	"os"
	"time"
)

type CockroachUsecaseCreate struct {
	cockroachRepository CockroachRepository
	cockroachMessaging  CockroachMessaging
	log                 *log.Logger
}

func NewCockroachUsecaseCreate(
	cockroachRepository CockroachRepository,
	cockroachMessaging CockroachMessaging,
) *CockroachUsecaseCreate {
	return &CockroachUsecaseCreate{
		cockroachRepository: cockroachRepository,
		cockroachMessaging:  cockroachMessaging,
		log:                 log.New(os.Stdout, "[cockroach-usecase-create] ", log.LstdFlags),
	}
}

func (u *CockroachUsecaseCreate) DataProcessing(in *entities.CreateCockroachDTO) error {
	data := &entities.InsertCockroachDTO{
		Amount: in.Amount,
	}

	if err := u.cockroachRepository.InsertCockroach(data); err != nil {
		return err
	}

	if err := u.cockroachMessaging.PushNotification(&entities.CockroachPushNotificationDTO{
		Title:        "Some cockroaches are being created... I don't know why 🪳",
		Amount:       in.Amount,
		ReportedTime: time.Now().Local().Format("2006-01-02 15:04:05"),
	}); err != nil {
		return err
	}

	u.log.Println("[message: create usecase finished successfully]")
	return nil
}
