package api

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/zlx2019/ecom/service/user"

	"github.com/gorilla/mux"
)

type APIServer struct {
	listenAddr string
	store      *sql.DB
}

func NewAPIServer(addr string, db *sql.DB) *APIServer {
	return &APIServer{addr, db}
}

func (s *APIServer) Startup() error {
	// 创建并且注册路由
	router := mux.NewRouter()
	sr := router.PathPrefix("/api/v1").Subrouter()

	// 注册 user 模块
	userStore := user.NewStore(s.store)
	userHandle := user.NewHandler(userStore)
	userHandle.RegisterRoutes(sr)

	log.Println("Listen and running on", s.listenAddr)
	return http.ListenAndServe(s.listenAddr, router)
}
