package command

import (
	"go-ddd-layered/pkg/utils"
	"io"
)

type (
	RegisterUser struct {
		Name  string `json:"name"`
		Email string `json:"email"`
	}

	ListUser struct {
		Page    *int32 `json:"page"`
		PerPage *int32 `json:"perPage"`
	}
)

func NewRegisterUser(body io.Reader) (cmd RegisterUser, err error) {
	err = utils.UnmarshalFromReader(body, &cmd)
	if err != nil {
		return cmd, nil
	}

	return cmd, nil
}
