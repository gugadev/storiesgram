package helpers

import (
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/gugadev/storiesgram/models"
)

// Files bunch of methods of IO ops
type Files struct{}

/*
Write fetch and image and save it in disk
*/
func (m *Files) Write(story models.Story, out string) {
	parts := strings.Split(story.Source, ".")
	ext := parts[len(parts)-1]
	file, filerr := os.Create(out + "/" + strconv.Itoa(story.PK) + "." + ext)
	if filerr != nil {
		log.Fatal(filerr)
	}
	defer file.Close()

	response, reqerr := http.Get(story.Source)
	if reqerr != nil {
		log.Fatal(reqerr)
	}
	defer response.Body.Close()

	_, copyerr := io.Copy(file, response.Body)
	// err2 := response.DownloadToFile(out + "/" + strconv.Itoa(story.PK) + "." + ext)
	if copyerr != nil {
		log.Fatal(copyerr)
	}
}
