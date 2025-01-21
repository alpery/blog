package user

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

type BlogUser struct {
	ID       primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Username string             `bson:"username" json:"username"`
	Password string             `bson:"password" json:"-"` // The "-" prevents the password from being included in JSON responses
	Email    string             `bson:"email" json:"email"`
	Created  time.Time          `bson:"created" json:"created"`
	Updated  time.Time          `bson:"updated" json:"updated"`
}

func (u *BlogUser) HashPassword() error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(hashedPassword)
	return nil
}

// ComparePassword compares the provided password with the hashed password
func (u *BlogUser) ComparePassword(password string) error {
	return bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
}
