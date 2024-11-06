package auth

type UserAlreadyExistError struct {
}

func (e UserAlreadyExistError) Error() string {
	return "User Already Exist!"
}
