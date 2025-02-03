package entities

type User struct {
	Id       int
	Username string
	Password string
	Role     int
}

func NewUser(id int, username string, password string, role int) *User {
	return &User{Id: id, Username: username, Password: password, Role: role}
}

func (u *User) GetUsername() string {
	return u.Username
}

func (u *User) SetUsername(username string) {
	u.Username = username
}

func (u *User) SetRole(role int) {
	u.Role = role
}