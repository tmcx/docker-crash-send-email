package main

import (
	"bytes"
	"os/exec"
	"strings"

	"github.com/jasonlvhit/gocron"
)

const (
	OFF = "OFF"
	ON  = "ON"
)

func startCheckDockerStatusJob() {
	gocron.Every(config.VerificationInterval).Minute().Do(checkDockerStatus)
	appendLog("[Job State]: Started cron job")
	<-gocron.Start()
}

func checkDockerStatus() {
	dockerState := OFF

	if dockerIsRunning() {
		emailSent = false
		dockerState = ON
	}
	appendLog("[Docker State]: " + dockerState)

	if dockerState == OFF && emailSent == false || lastEmailNotWasSent {
		sendEmailReportingDockerFailure()
	}
}

func dockerIsRunning() bool {
	cmd := exec.Command("docker", "ps", "-a")
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()

	if err != nil {
		return false
	}
	return strings.Contains(out.String(), "IMAGE") && strings.Contains(out.String(), "NAME")
}
