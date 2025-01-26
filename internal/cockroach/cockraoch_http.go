package cockroach

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/samulastech/cockroach/internal/entities"
	"github.com/samulastech/cockroach/pkg/common"
)

type CockroachHTTPHandler struct {
	cockroachUsecaseCreate Create
	log                    *log.Logger
}

func NewCockroachHTTPHandler(cockroachUsecaseCreate Create) *CockroachHTTPHandler {
	return &CockroachHTTPHandler{
		cockroachUsecaseCreate: cockroachUsecaseCreate,
		log:                    log.New(os.Stdout, "[cockroach-httphandler] ", log.LstdFlags),
	}
}

func (h *CockroachHTTPHandler) CreateCockroach(w http.ResponseWriter, r *http.Request) {
	var dto entities.CreateCockroachDTO
	if err := json.NewDecoder(r.Body).Decode(&dto); err != nil {
		h.log.Println(fmt.Sprintf("[message: error while trying to parse the body][err: %s]", err.Error()))
		common.SendJSON(w, http.StatusBadRequest, common.Response{
			Message: "invalid body",
			Data:    nil,
		})
		return
	}

	if err := h.cockroachUsecaseCreate.DataProcessing(&dto); err != nil {
		h.log.Println(fmt.Sprintf("[message: unknow error was thrown][err: %s]", err.Error()))
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
