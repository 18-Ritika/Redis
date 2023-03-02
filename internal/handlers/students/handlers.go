package students

import (
	"Training/Redis/Redis/internal/models"
	"Training/Redis/Redis/internal/services"
	"context"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io"
	"net/http"
)

type Handler struct {
	svc services.StudentServices
}

func New(svc services.StudentServices) *Handler {
	return &Handler{svc: svc}
}

func (h *Handler) Get(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	v := mux.Vars(r)
	id := v["id"]

	resp, err := h.svc.Get(context.TODO(), id)
	if err != nil {
		json.NewEncoder(w).Encode(ErrorResp{
			Code:    http.StatusInternalServerError,
			Status:  "ERROR",
			Message: "Internal Server Error",
		})
		return
	}

	json.NewEncoder(w).Encode(SuccessResp{
		Code:   http.StatusOK,
		Status: "SUCCESS",
		Data:   resp,
	})
}

func (h *Handler) Post(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	req, err := io.ReadAll(r.Body)
	if err != nil {
		json.NewEncoder(w).Encode(ErrorResp{
			Code:    http.StatusBadRequest,
			Status:  "ERROR",
			Message: "ID is not of Proper DataType",
		})
		return
	}

	var student models.Student
	err = json.Unmarshal(req, &student)
	if err != nil {
		fmt.Println(err)
		return
	}

	res, err := h.svc.Post(context.Background(), student)
	if err != nil {
		json.NewEncoder(w).Encode(ErrorResp{
			Code:    http.StatusInternalServerError,
			Status:  "ERROR",
			Message: "Internal Server Error",
		})
		return
	}

	json.NewEncoder(w).Encode(SuccessResp{
		Code:   http.StatusOK,
		Status: "SUCCESS",
		Data:   res,
	})
}

func (h *Handler) Delete(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	v := mux.Vars(r)
	id := v["id"]
	err := h.svc.Delete(context.Background(), id)
	if err != nil {
		json.NewEncoder(w).Encode(SuccessResp{
			Code:   http.StatusInternalServerError,
			Status: "ERROR",
			Data:   "Internal Server Error",
		})
		return
	}
	json.NewEncoder(w).Encode(SuccessResp{
		Code:   http.StatusNoContent,
		Status: "SUCCESS",
		Data:   "Movie Deleted Successfully",
	})
}
