package users

type AddressSchema struct {
	Street string `json:"street"`
	City   string `json:"city"`
}

type UserSchema struct {
	FirstName string           `json:"fname"`
	LastName  string           `json:"lname"`
	Age       int              `json:"age" validate:"gte=0,lte=130"`
	Email     string           `json:"email" validate:"required,email"`
	Address   []*AddressSchema `json:"address"`
}
