package models

import (
	"fmt"

	"github.com/wmd/utils"

	"github.com/go-redis/redis"
	"golang.org/x/crypto/bcrypt"
)

// User struct
type User struct {
	id int64
}

// NewUser func
func NewUser(prof string, mail string, username string, hash []byte) (*User, error) {
	exists, err := client.HExists("user:by-username", username).Result()
	if exists {
		return nil, utils.ErrUsernameTaken
	}
	id, err := client.Incr("user:next-id").Result()
	if err != nil {
		return nil, err
	}

	key := fmt.Sprintf("user:%d", id)
	pipe := client.Pipeline()
	pipe.HSet(key, "id", id)
	pipe.HSet(key, "username", username)
	pipe.HSet(key, "mail", mail)
	pipe.HSet(key, "hash", hash)
	pipe.HSet(key, "prof", prof)
	pipe.HSet("user:by-username", username, id)
	_, err2 := pipe.Exec()
	if err2 != nil {
		return nil, err2
	}
	return &User{id}, nil
}

// GetUsername method
func (user *User) GetUsername() (string, error) {
	key := fmt.Sprintf("user:%d", user.id)
	return client.HGet(key, "username").Result()
}

// GetID method
func (user *User) GetID() (int64, error) {
	return user.id, nil
}

// GetHash method
func (user *User) GetHash() ([]byte, error) {
	key := fmt.Sprintf("user:%d", user.id)
	return client.HGet(key, "hash").Bytes()
}

// Authentificate method
func (user *User) Authentificate(password string) error {
	hash, err := user.GetHash()
	if err != nil {
		return err
	}
	err = bcrypt.CompareHashAndPassword(hash, []byte(password))
	if err == bcrypt.ErrMismatchedHashAndPassword {
		return utils.ErrInvalidLogin
	}
	return err
}

// GetUserByID func
func GetUserByID(id int64) (*User, error) {
	return &User{id}, nil
}

// GetUserByUsername func
func GetUserByUsername(username string) (*User, error) {
	id, err := client.HGet("user:by-username", username).Int64()
	if err == redis.Nil {
		return nil, utils.ErrUserNotFound
	} else if err != nil {
		return nil, err
	}
	return GetUserByID(id)
}

// AuthentificateUser func
func AuthentificateUser(username, password string) (*User, error) {
	user, err := GetUserByUsername(username)
	if err != nil {
		return nil, err
	}
	return user, user.Authentificate(password)
}

// RegisterUser func
func RegisterUser(prof, mail, username, password string) error {
	cost := bcrypt.DefaultCost
	hash, err := bcrypt.GenerateFromPassword([]byte(password), cost)
	if err != nil {
		return err
	}
	_, err = NewUser(prof, mail, username, hash)
	return err
}

// ReturnUsername func
func ReturnUsername(id int64) (string, error) {
	key := fmt.Sprintf("user:%d", id)
	return client.HGet(key, "username").Result()
}

// CheckProf func
func CheckProf(id int64) bool {
	key := fmt.Sprintf("user:%d", id)
	s, err := client.HGet(key, "prof").Result()
	if err != nil || s == "" {
		return false
	}
	return true
}
