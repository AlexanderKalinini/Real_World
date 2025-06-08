package user

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"rwa/internal/model/user"
)

func (c *UsersController) Create(w http.ResponseWriter, r *http.Request) error {
	ctx := r.Context()
	body := r.Body
	w.Header().Set("Content-Type", "application/json")
	defer func(body io.ReadCloser) {
		err := body.Close()
		if err != nil {
			return
		}
	}(body)
	type UsersWrapper struct {
		Users user.User `json:"users"`
	}
	bodyByte, err := io.ReadAll(body)
	if err != nil {
		return err
	}
	var usersWrapper UsersWrapper
	err = json.Unmarshal(bodyByte, &usersWrapper)
	if err != nil {
		fmt.Println(err)
		return err
	}

	err = c.UserRepo.Create(ctx, &usersWrapper.Users)

	if err != nil {
		http.Error(w, err.Error(), http.StatusUnprocessableEntity)
		return err
	}

	fmt.Println(string(bodyByte))
	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(bodyByte)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}
