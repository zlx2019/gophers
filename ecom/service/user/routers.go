package user

import (
	"fmt"
	"net/http"
	"time"

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
	req, err := utils.ParseJSON[types.RegisterUserRequest](r)
	if err != nil {
		utils.Fail(w, err.Error())
		return
	}
	// check if the user exists?
	user, err := h.store.GetUserByUsername(req.Username)
	if err == nil {
		utils.Fail(w, fmt.Sprintf("user with %s already exists", req.Username))
		return
	}
	// hash password
	hashedPassword, err := utils.HashPassword(req.Password)
	if err != nil {
		utils.Fail(w, err.Error())
		return
	}
	err = h.store.CreateUser(&types.User{
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Username:  req.Username,
		Password:  hashedPassword,
		CreatedAt: time.Now(),
	})
	if err != nil {
		utils.Fail(w, err.Error())
		return
	}
	utils.Ok(w, user)
}
