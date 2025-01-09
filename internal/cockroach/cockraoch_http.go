package cockroach

import (
	"encoding/json"
	"fmt"
	"github.com/samulastech/cockroach/internal/entities"
	"github.com/samulastech/cockroach/pkg/common"
	"log"
	"net/http"
	"os"
)

type CockroachHTTPHandler struct {
	cockroachUsecaseCreate *CockroachUsecaseCreate
	log                    *log.Logger
}

func NewCockroachHTTPHandler(cockroachUsecaseCreate *CockroachUsecaseCreate) *CockroachHTTPHandler {
	return &CockroachHTTPHandler{
		cockroachUsecaseCreate: cockroachUsecaseCreate,
		log:                    log.New(os.Stdout, "[cockroach-httphandler] ", log.LstdFlags),
	}
}

func (h *CockroachHTTPHandler) CreateCockroach(w http.ResponseWriter, r *http.Request) {
	var dto entities.CreateCockroachDTO
	if err := json.NewDecoder(r.Body).Decode(&dto); err != nil {
		h.log.Println(fmt.Sprintf("[err: %s]", err.Error()))
		common.SendJSON(w, http.StatusBadRequest, common.Response{
			Message: "invalid body",
			Data:    nil,
		})
		return
	}

	if err := h.cockroachUsecaseCreate.DataProcessing(&dto); err != nil {
		h.log.Println(fmt.Sprintf("[err: %s]", err.Error()))
		common.SendJSON(w, http.StatusInternalServerError, common.Response{
			Message: "internal server error",
			Data:    nil,
		})
		return
	}

	common.SendJSON(w, http.StatusOK, common.Response{
		Message: "success",
		Data:    "🪳",
	})
	return
}
