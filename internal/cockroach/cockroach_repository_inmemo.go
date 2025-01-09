package cockroach

import (
	"fmt"
	"github.com/samulastech/cockroach/internal/entities"
	"log"
	"os"
)

type CockroachRepositoryInMemo struct {
	list   []*entities.Cockroach
	nextId int
	log    *log.Logger
}

func NewCockroachRepositoryInMemo() *CockroachRepositoryInMemo {
	return &CockroachRepositoryInMemo{
		list: make([]*entities.Cockroach, 0),
		log:  log.New(os.Stdout, "[cockroach-inmemo-repo] ", log.LstdFlags),
	}
}

func (r *CockroachRepositoryInMemo) InsertCockroach(in *entities.InsertCockroachDTO) error {
	r.list = append(r.list, &entities.Cockroach{Id: r.nextId, Amount: in.Amount})
	r.log.Println(fmt.Sprintf("[message: inserted cockroach(es)][id: %d][amount: %d]", r.nextId, in.Amount))
	r.nextId++
	return nil
}
