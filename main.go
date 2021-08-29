package main

func main() {
	appendLog("[Service]: Started")
	err := loadConfiguration()
	if err != nil {
		return
	}

	startCheckDockerStatusJob()
}