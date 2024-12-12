package user

import (
	"testing"

	"github.com/zlx2019/ecom/types"
)

// mock 数据
type mockUserStore struct {
}

func (*mockUserStore) GetUserByUsername(username string) (*types.User, error) {
	return nil, nil
}

func (*mockUserStore) GetUserByID(id int) (*types.User, error) {
	return nil, nil
}
func (*mockUserStore) CreateUser(user *types.User) error {
	return nil
}

// UserHandler unit testing
func TestUserServiceHandlers(t *testing.T) {
	userStore := &mockUserStore{}
	_ = NewHandler(userStore)

}
