package main

import (
	"flag"
	"log"
	"os"

	"github.com/gugadev/storiesgram/helpers"
	"github.com/gugadev/storiesgram/request"
)

func main() {
	var files helpers.Files

	userid := flag.String("u", "353545498924", "Target User ID")
	sessionid := flag.String("s", "ri9485fggg345444%$3KD", "Your session ID")
	output := flag.String("o", "/home/johndoe/stories", "A folder to write data")
	flag.Parse()

	if userid == nil {
		log.Fatal("Provide the target user id")
		os.Exit(-1)
	}
	if sessionid == nil {
		log.Fatal("Provide your session id")
		os.Exit(-1)
	}
	if output == nil {
		log.Fatal("Provide the output directory")
		os.Exit(-1)
	}

	// 1976714778
	// 6097358349%3A4uOkWWhPs8fmQC%3A4
	// /Users/gugadev/stories
	stories := request.GetStories(*userid, *sessionid)
	for _, story := range stories {
		files.Write(story, *output)
	}
}
