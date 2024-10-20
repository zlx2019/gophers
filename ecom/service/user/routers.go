package user

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/zlx2019/ecom/types"
	"github.com/zlx2019/ecom/utils"
)

type Handler struct {
	store types.UserStore
}

func NewHandler(store types.UserStore) *Handler {
	return &Handler{store}
}

// user api routes
func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/login", h.login).Methods("POST")
	router.HandleFunc("/register", h.register).Methods("POST")
}

// 登录
func (h *Handler) login(w http.ResponseWriter, r *http.Request) {
}

// 注册
func (h *Handler) register(w http.ResponseWriter, r *http.Request) {
	// get JSON payload
	_, err := utils.ParseJSON[types.RegisterUserRequest](r)
	if err != nil {
		utils.Fail(w, err.Error())
		return
	}
	// check if the user exists?
}
