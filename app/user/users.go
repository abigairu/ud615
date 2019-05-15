package user

type User struct {
	Username     string
	PasswordHash string
	Email        string
}

type Users map[string]User

var DB = Users{
	"user": User{
		Username: "user",
		// bcrypt hash for "password"
		PasswordHash: "$2a$10$KgFhp4HAaBCRAYbFp5XYUOKrbO90yrpUQte4eyafk4Tu6mnZcNWiK",
		Email: "user@example.com",
	},
	"user2": User{
		Username: "user2",
		// bcrypt hash for "password2"
		PasswordHash: "$2y$10$KagMZRoOywODjEpmI1NYJ.hhBVQxyEJU6zaDRiHLGv6QQ18DmhJDa",
		Email: "user2@example.com",
	},
}
