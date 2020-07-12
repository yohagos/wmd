package models

import (
	"fmt"
	"github.com/wmd/utils"
	"strconv"
	"strings"
	"time"
)

// Upload struct
type Upload struct {
	id int64
}

// NewUpload func
func NewUpload(user_id int64, itemName string, description string, scientificType string, categories string, path string) (*Upload, error) {
	id, err := client.Incr("upload:next-id").Result()
	if err != nil {
		return nil, err
	}
	key := fmt.Sprintf("upload:%d", id)
	pipe := client.Pipeline()
	pipe.HSet(key, "id", id)
	pipe.HSet(key, "user_id", user_id)
	pipe.HSet(key, "item", itemName)
	userName, _ := GetUserForUpload(user_id)
	pipe.HSet(key, "username", userName)
	createdAt := time.Now()
	pipe.HSet(key, "createdAt", createdAt)
	pipe.HSet(key, "description", description)
	pipe.HSet(key, "type", scientificType)
	pipe.HSet(key, "categories", categories)
	pipe.HSet(key, "path", path)
	pipe.LPush("uploads", id)
	pipe.LPush(fmt.Sprintf("user:%d:uploads", user_id), id)
	_, err = pipe.Exec()
	if err != nil {
		return nil, err
	}
	return &Upload{id}, nil
}

func (up Upload) GetItemname() (string, error) {
	key := fmt.Sprintf("upload:%d", up.id)
	str, err := client.HGet(key, "item").Result()
	if err != nil {
		utils.LoggingErrorFile(err.Error())
		return "", err
	}
	return str, err
}

func (up Upload) GetTimestamp() (string, error) {
	key := fmt.Sprintf("upload:%d", up.id)
	return client.HGet(key, "createdAt").Result()
}

func (up Upload) GetDescription() (string, error) {
	key := fmt.Sprintf("upload:%d", up.id)
	return client.HGet(key, "description").Result()
}

func (up Upload) GetType() (string, error) {
	key := fmt.Sprintf("upload:%d", up.id)
	return client.HGet(key, "type").Result()
}

func (up Upload) GetCategories() (string, error) {
	key := fmt.Sprintf("upload:%d", up.id)
	return client.HGet(key, "categories").Result()
}

func (up Upload) GetPath() (string, error) {
	key := fmt.Sprintf("upload:%d", up.id)
	return client.HGet(key, "path").Result()
}

func (up *Upload) GetUser() (*User, error) {
	key := fmt.Sprintf("upload:%d", up.id)
	userID, err := client.HGet(key, "user_id").Int64()
	if err != nil {
		return nil, err
	}
	return GetUserByID(userID)
}

func (up Upload) GetUserID() (int64, error) {
	key := fmt.Sprintf("upload:%d", up.id)
	userID, err := client.HGet(key, "user_id").Int64()
	if err != nil {
		return -1, err
	}
	return userID, nil
}

func queryUploads(key string) ([]*Upload, error) {
	uploadIDs, err := client.LRange(key, 0, 10).Result()
	if err != nil {
		return nil, err
	}
	uploads := make([]*Upload, len(uploadIDs))
	for i, upID := range uploadIDs {
		id, err := strconv.Atoi(upID)
		if err != nil {
			return nil, err
		}
		uploads[i] = &Upload{int64(id)}
	}
	return uploads, err
}

// UploadsPath func
func UploadsPath(key string) (string, error) {
	uploads, err := GetAllUploads()
	if err != nil {
		utils.LoggingErrorFile(err.Error())
		return "", err
	}
	var p string
	for _, up := range uploads {
		str, _ := up.GetItemname()
		if strings.EqualFold(str, key) {
			p, _ = up.GetPath()
			return p, nil
		}
	}

	return "", utils.ErrNothingFound
}

// GetAllUploads func
func GetAllUploads() ([]*Upload, error) {
	return queryUploads("uploads")
}

// GetUserForUpload func
func GetUserForUpload(id int64) (string, error) {
	key := fmt.Sprintf("user:%d", id)
	return client.HGet(key, "username").Result()
}
