package cockroach

import "github.com/samulastech/cockroach/internal/entities"

type Create interface {
	DataProcessing(in *entities.CreateCockroachDTO) error
}
