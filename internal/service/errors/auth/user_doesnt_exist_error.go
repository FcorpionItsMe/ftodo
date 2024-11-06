package auth

type UserDoesntExistError struct {
	Login string
}

func (e UserDoesntExistError) Error() string {
	return "User Doesnt Exist!"
}
