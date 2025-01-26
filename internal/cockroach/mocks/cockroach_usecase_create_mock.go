package mocks

import (
	"github.com/samulastech/cockroach/internal/entities"
	"github.com/stretchr/testify/mock"
)

type CockroachUsecaseCreateMock struct {
	mock.Mock
}

func NewCockroachUsecaseCreateMock() *CockroachUsecaseCreateMock {
	return &CockroachUsecaseCreateMock{}
}

func (c *CockroachUsecaseCreateMock) DataProcessing(in *entities.CreateCockroachDTO) error {
	args := c.Called(in)
	return args.Error(0)
}
