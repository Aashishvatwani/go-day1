package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
	"github.com/slack-go/slack"
)

func main() {
	// Load .env file if present
	_ = godotenv.Load()

	api := slack.New(os.Getenv("SLACK_BOT_TOKEN"))
	channelArr := []string{os.Getenv("SLACK_CHANNEL_ID")}
	fileArr := []string{"Proj_Synopsis_format.docx"}

	for _, channelID := range channelArr {
		for i := range fileArr {
			fp := fileArr[i]
			// Ensure file exists and is non-empty
			f, err := os.Open(fp)
			if err != nil {
				fmt.Printf("cannot open file %q: %v\n", fp, err)
				continue
			}
			fi, err := f.Stat()
			if err != nil {
				f.Close()
				fmt.Printf("cannot stat file %q: %v\n", fp, err)
				continue
			}
			if fi.Size() == 0 {
				f.Close()
				fmt.Printf("file %q is empty; nothing to upload\n", fp)
				continue
			}

			params := slack.UploadFileV2Parameters{
				Channel:  channelID,
				Filename: filepath.Base(fp),
				Title:    filepath.Base(fp),
				Reader:   f,
				// Provide size so Slack doesn't treat it as 0
				FileSize: int(fi.Size()),
			}
			uploaded, err := api.UploadFileV2(params)
			// Close after upload attempt
			f.Close()
			if err != nil {
				fmt.Printf("failed to upload file: %v\n", err)
				continue
			}
			fmt.Printf("Uploaded: %s (id: %s)\n", uploaded.Title, uploaded.ID)
		}
	}
}
