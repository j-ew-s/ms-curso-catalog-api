package userService

type User struct {
	Id      string   `json:"id,omitempty"`
	Name    string   `json:"name,omitempty"`
	Email   string   `json:"email,omitempty"`
	Address *Address `json:"address,omitempty"`
	Account *Account `json:"account,omitempty"`
}

type Address struct {
	Street  string `json:"street,omitempty"`
	ZipCode string `json:"zipCode,omitempty"`
}

type Account struct {
	User     string `json:"user,omitempty"`
	Password string `json:"password,omitempty"`
}

type Token struct {
	Token string `json:"token,omitempty"`
}

type TokenValidation struct {
	Status  bool   `json:"status,omitempty"`
	Message string `json:"message,omitempty"`
}
