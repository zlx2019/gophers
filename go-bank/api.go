package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

// Web Server
type APIServer struct {
	listenAddr string
	store      *PostgresStore
}

func NewAPIServer(listenAddr string, store *PostgresStore) *APIServer {
	return &APIServer{
		listenAddr,
		store,
	}
}

// HTTP 请求处理函数
type iHandleFunc func(http.ResponseWriter, *http.Request) error

// WriteJson Response Json
func WriteJson(w http.ResponseWriter, httpStatus int, value any) error {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(httpStatus)
	return json.NewEncoder(w).Encode(value)
}
func StatusOk(w http.ResponseWriter, value any) error {
	return WriteJson(w, http.StatusOK, value)
}

// WriteText Response Text
func WriteText(w http.ResponseWriter, httpStatus int, value string) error {
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(httpStatus)
	_, err := w.Write([]byte(value))
	return err
}

type ApiError struct {
	message string
}

// 执行 iHandlerFunc，进行错误处理，以及返回包装后的 HandlerFunc
func wrapHTTPHandleFunc(f iHandleFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := f(w, r); err != nil {
			// handle the error
			log.Println(err)
			_ = WriteJson(w, http.StatusInternalServerError, ApiError{message: err.Error()})
		}
	}
}

// 服务运行.
func (s *APIServer) Startup() {
	// 创建路由器
	router := mux.NewRouter()
	// 注册路由
	router.HandleFunc("/", wrapHTTPHandleFunc(s.ping))
	router.HandleFunc("/account", wrapHTTPHandleFunc(s.handleAccount))
	router.HandleFunc("/account/{id}", wrapHTTPHandleFunc(s.GetAccountByID))
	log.Println("Server running on port: ", s.listenAddr)
	// 启动服务
	if err := http.ListenAndServe(s.listenAddr, router); err != nil {
		panic(err)
	}
}

// 账户处理
func (s *APIServer) handleAccount(w http.ResponseWriter, r *http.Request) error {
	switch strings.ToUpper(r.Method) {
	case "GET":
		return s.GetAccounts(w, r)
	case "POST":
		return s.CreateAccount(w, r)
	case "DELETE":
		return s.DeleteAccount(w, r)
	default:
		return fmt.Errorf("method not support %s", r.Method)
	}
}

func (s *APIServer) ping(w http.ResponseWriter, r *http.Request) error {
	return WriteText(w, http.StatusOK, "API Server on running...")
}

// 根据ID 获取账户信息
func (s *APIServer) GetAccountByID(w http.ResponseWriter, r *http.Request) error {
	// 读取 /account/{id} 参数
	id := mux.Vars(r)["id"]

	fmt.Println(id)
	return nil
}

// 获取所有账户信息
func (s *APIServer) GetAccounts(w http.ResponseWriter, r *http.Request) error {
	accounts, err := s.store.GetAccounts()
	if err != nil {
		return err
	}
	return StatusOk(w, accounts)
}

// 创建账户
func (s *APIServer) CreateAccount(w http.ResponseWriter, r *http.Request) error {
	req := new(CreateAccountRequest)
	// 解析请求体，映射为实体
	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		return err
	}
	account := NewAccount(req.FirstName, req.LastName)
	id, err := s.store.CreateAccount(account)
	if err != nil {
		return err
	}
	account.ID = id
	return StatusOk(w, account)
}

// 删除账户
func (s *APIServer) DeleteAccount(w http.ResponseWriter, r *http.Request) error {
	return nil
}

// 账户转账
func (s *APIServer) TransferAccount(w http.ResponseWriter, r *http.Request) error {
	return nil
}
