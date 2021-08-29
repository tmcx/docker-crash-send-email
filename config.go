package main

import (
	"encoding/json"
	"io/ioutil"
	"os"

	"github.com/go-playground/validator/v10"
)

const (
	configFilePath = "./configuration.json"
)

var (
	config   Config
	validate = validator.New()
)

func loadConfiguration() error {
	jsonFile, err := os.Open(configFilePath)
	if err != nil {
		appendLog("[Load configuration]: Error: " + err.Error())
		return err
	}
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	json.Unmarshal(byteValue, &config)

	err = validate.Struct(config)
	if err != nil {
		appendLog("[Load configuration]: Error: " + err.Error())
		return err
	}

	appendLog("[Load configuration]: Success")
	return nil
}

type Config struct {
	VerificationInterval uint64      `json:"verification_interval" validate:"required,number,min=1"`
	Credentials          Credentials `json:"credentials" validate:"required"`
	Email                Email       `json:"email" validate:"required"`
	Log                  Log         `json:"log" validate:"required"`
}

type Log struct {
	MaxFileSize int64  `json:"max_file_size" validate:"required,min=1"`
	Format      string `json:"format" validate:"required"`
}

type Email struct {
	To      []string `json:"to" validate:"required,dive,email"`
	Body    string   `json:"html_body" validate:"required"`
	Subject string   `json:"subject" validate:"required"`
}

type Credentials struct {
	SmtpPort int    `json:"smtp_port" validate:"required,number,min=1"`
	Email    string `json:"email" validate:"required,email"`
	SmtpHost string `json:"smtp_host" validate:"required"`
	Password string `json:"password" validate:"required"`
}
