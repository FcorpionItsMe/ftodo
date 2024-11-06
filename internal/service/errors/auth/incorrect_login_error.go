package auth

type IncorrectLoginError struct {
	AdditionalInfo string
}

func (e IncorrectLoginError) Error() string {
	return "Incorect Login! " + e.AdditionalInfo
}
