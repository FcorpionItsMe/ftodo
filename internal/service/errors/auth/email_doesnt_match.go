package auth

type EmailDoesntMatch struct {
	Email string
}

func (e EmailDoesntMatch) Error() string {
	return "Email doesnt match!"
}
