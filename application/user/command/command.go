package command

import (
	"go-ddd-layered/pkg/utils"
	"io"
	"time"
)

type (
	RegisterUser struct {
		Name     string
		Email    string
		Birthday time.Time
	}
)

func NewRegisterUser(body io.Reader) (cmd RegisterUser, err error) {
	err = utils.UnmarshalFromReader(body, &cmd)
	if err != nil {
		return cmd, nil
	}

	return cmd, nil
}
