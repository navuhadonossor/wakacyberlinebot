package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"log"
	"os"
)

func insertUser(tgId int, tgName string) error {
	checkFile()
	users, err := getUserList()
	for _, user := range users {
		if user.TelegramId == tgId {
			return errors.New("user already exist")
		}
	}
	id, err := uuid.NewUUID()
	newUser := User{
		Id:               id.String(),
		TelegramId:       tgId,
		TelegramName:     tgName,
		WakatimeName:     "",
		WakatimeApiToken: "",
	}
	users = append(users, newUser)
	content, err := json.Marshal(users)
	fmt.Print("REGISTER USER")
	fmt.Println(string(content))
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
		if user.TelegramId == tgId {
			users[i].WakatimeName = wakatimeName
			users[i].WakatimeApiToken = apiToken
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
	fmt.Println("%+v\n", users)
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
