package models

import (
	"log"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

// User is the basic entity of the app and represents the user
type User struct {
	ID        primitive.ObjectID `bson:"_id" json:"id"`
	Email     string             `bson:"email" json:"email"`
	Password  string
	CreatedAt primitive.DateTime
	UpdatedAt primitive.DateTime
}

// HashPassword hashes the Password of the current User struct. Does not check if the password has already been hashed
func (u *User) HashPassword() {
	hash, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Panicln(err)
	}

	u.Password = string(hash)
}

// ComparePasswords hashes the password passed to it and compares it with the hashed password
func (u *User) ComparePasswords(plainPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(plainPassword))

	if err != nil {
		return false
	}
	return true
}