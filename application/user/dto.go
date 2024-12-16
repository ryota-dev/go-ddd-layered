package user

type (
	UserDTO struct {
		Id        string `json:"id"`
		Name      string `json:"name"`
		Email     string `json:"email"`
		Birthday  string `json:"birthday"`
		CreatedAt string `json:"createdAt"`
		UpdatedAt string `json:"updatedAt"`
	}
)
