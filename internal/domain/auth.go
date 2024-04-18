package domain

type User struct {
	ID         string //uuid
	Email      string
	FirstName  string
	LastName   string
	Password   string
	IsActive   bool
	IsVerified bool
	IsStaff    bool
}

type SignUpInput struct {
	Email     *string `json:"email" validate:"required,email"`
	FirstName *string `json:"firstName" validate:"required,gte=2"`
	LastName  *string `json:"lastName" validate:"required,gte=2"`
	Password  *string `json:"password" validate:"required,gte=6"`
}

func (i SignUpInput) Validate() error {
	return validate.Struct(i)
}

type SignInInput struct {
	Email    *string `json:"email" validate:"required"`
	Password *string `json:"password" validate:"required"`
}

func (i SignInInput) Validate() error {
	return validate.Struct(i)
}
