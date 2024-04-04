package model

type User struct {
	ID          int    `db:"id"`
	Name        string `db:"name"`
	Email       string `db:"email"`
	Password    string `db:"password"`
	DefaultCity string `db:"default_city"`
}

func (c *User) TableName() string {
	return "users"
}
