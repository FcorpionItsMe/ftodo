package auth

type HashingError struct {
}

func (e HashingError) Error() string {
	return "Hashing Error!"
}
