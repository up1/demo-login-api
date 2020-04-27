package user

type LoginRequest struct {
	UserName string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`
}

type User struct {
	ID        int    `json:"id,omitempty"`
	FirstName string `json:"firstname,omitempty"`
	LastName  string `json:"lastname,omitempty"`
	Email     string `json:"email,omitempty"`
}
