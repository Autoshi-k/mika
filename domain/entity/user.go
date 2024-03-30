package entity

type User struct {
	id          string
	phoneNumber string
	email       string
	foreName    string
	surname     string
}

func (u User) IsEmpty() bool {
	return u.id != ""
}
