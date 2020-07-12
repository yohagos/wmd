package models

import (
	"fmt"
	"github.com/wmd/utils"
	"strconv"
)

// User struct
type Verifying struct {
	id int64
}

// NewUser func
func NewVerification(mail string, username string, password string, rand []byte) (*Verifying, error) {
	exists, err := client.HExists("user:by-username", username).Result()
	if exists {
		return nil, utils.ErrUsernameTaken
	}
	id, err := client.Incr("verif:next-id").Result()
	if err != nil {
		return nil, err
	}

	key := fmt.Sprintf("verif:%d", id)
	pipe := client.Pipeline()
	pipe.HSet(key, "id", id)
	pipe.HSet(key, "username", username)
	pipe.HSet(key, "password", password)
	pipe.HSet(key, "mail", mail)
	pipe.HSet(key, "rand", rand)
	//pipe.HSet("user:by-username", username, id)
	pipe.LPush("verifs", id)
	_, err2 := pipe.Exec()
	if err2 != nil {
		return nil, err2
	}
	return &Verifying{id}, nil
}

// GetUsername method
func (verif *Verifying) GetUsername() (string, error) {
	key := fmt.Sprintf("verif:%d", verif.id)
	return client.HGet(key, "username").Result()
}

// GetID method
func (verif *Verifying) GetID() (int64, error) {
	return verif.id, nil
}

// GetHash method
func (verif *Verifying) GetRand() ([]byte, error) {
	key := fmt.Sprintf("verif:%d", verif.id)
	return client.HGet(key, "rand").Bytes()
}

func (verif *Verifying) GetPassword() (string, error) {
	key := fmt.Sprintf("verif:%d", verif.id)
	return client.HGet(key, "password").Result()
}

func (verif *Verifying) GetMail() (string, error) {
	key := fmt.Sprintf("verif:%d", verif.id)
	return client.HGet(key, "mail").Result()
}

func (verif *Verifying) DelVerif() {
	client.Del(fmt.Sprintf("verif:%d", verif.id))
	client.Decr(fmt.Sprintf("verif:next-id"))
}

func queryVerifs(key string) ([]*Verifying, error) {
	verifIDs, err := client.LRange(key, 0, -1).Result()
	if err != nil {
		return nil, err
	}
	verifs := make([]*Verifying, len(verifIDs))
	for i, upID := range verifIDs {
		id, err := strconv.Atoi(upID)
		if err != nil {
			return nil, err
		}
		verifs[i] = &Verifying{int64(id)}
	}
	return verifs, err
}

// GetAllUploads func
func GetAllVerif() ([]*Verifying, error) {
	return queryVerifs("verifs")
}
