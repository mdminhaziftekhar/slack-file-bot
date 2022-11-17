package main

import (
	"fmt"
	"os"

	"github.com/slack-go/slack"
)

func main() {

	os.Setenv("SLACK_BOT_TOKEN", "xoxb-4391442328641-4382398975477-4TWX7Ifdi9nZAHEpfUO9rCia")
	os.Setenv("CHANNEL_ID", "C04B1R6T0UD")

	api := slack.New(os.Getenv("SLACK_BOT_TOKEN"))
	channelArr := []string{os.Getenv("CHANNEL_ID")}
	fileArr := []string{"file1.jpg", "file2.jpg", "file3.jpg", "file4.jpg", "file5.jpg"}

	for i := 0; i < len(fileArr); i++ {
		params := slack.FileUploadParameters{
			Channels: channelArr,
			File:     fileArr[i],
		}

		file, err := api.UploadFile(params)

		if err != nil {
			fmt.Printf("%s\n", err)
		}

		fmt.Printf("Name : %s, URL : %s\n", file.Name, file.URL)
	}

}
