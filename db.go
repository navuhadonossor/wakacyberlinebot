package main

import (
	"encoding/json"
	"errors"
	"github.com/google/uuid"
	"log"
	"os"
)

func insertUser(tgId int, tgName string) error {
	checkFile()
	users, err := getUserList()
	for _, user := range users {
		if user.telegramId == tgId {
			return errors.New("user already exist")
		}
	}
	id, err := uuid.NewUUID()
	newUser := &User{
		id:               id.String(),
		telegramId:       tgId,
		telegramName:     tgName,
		wakatimeName:     "",
		wakatimeApiToken: "",
	}
	users = append(users, *newUser)
	content, err := json.Marshal(users)
	err = os.WriteFile(JsonFilepath, content, 0644)
	if err != nil {
		log.Println("Insert user failed: " + err.Error())
	}
	return nil
}

func updateUser(tgId int, apiToken string, wakatimeName string) error {
	checkFile()
	users, err := getUserList()
	for i, user := range users {
		if user.telegramId == tgId {
			users[i].wakatimeName = wakatimeName
			users[i].wakatimeApiToken = apiToken
		}
	}
	content, err := json.Marshal(users)
	err = os.WriteFile(JsonFilepath, content, 0644)
	if err != nil {
		log.Println("Update user failed: " + err.Error())
	}
	return nil
}

func getUserList() ([]User, error) {
	var users []User
	content, err := os.ReadFile(JsonFilepath)
	err = json.Unmarshal(content, &users)
	log.Print("JSON CONTENTS")
	log.Println(users)
	if err != nil {
		return []User{}, errors.New("Cannot read users: " + err.Error())
	}
	return users, nil
}

func checkFile() error {
	_, err := os.Stat(JsonFilepath)
	if os.IsNotExist(err) {
		_, err = os.Create(JsonFilepath)
		if err != nil {
			log.Println("Cannot create file: " + err.Error())
			return err
		}
	}
	return nil
}
