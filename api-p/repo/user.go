package repo

type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

var UserList []User

func FindUserByEmail(email string) *User {
	for _, user := range UserList {
		if user.Email == email {
			return &user
		}
	}
	return nil
}
