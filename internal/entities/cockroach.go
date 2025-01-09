package entities

import "time"

type Cockroach struct {
	Id        int       `json:"id"`
	Amount    int       `json:"amount"`
	CreatedAt time.Time `json:"createdAt"`
}

type InsertCockroachDTO struct {
	Id        int       `json:"id"`
	Amount    int       `json:"amount"`
	CreatedAt time.Time `json:"createdAt"`
}

type CockroachPushNotificationDTO struct {
	Title        string `json:"title"`
	Amount       int    `json:"amount"`
	ReportedTime string `json:"createdAt"`
}

type CreateCockroachDTO struct {
	Amount int `json:"amount"`
}
