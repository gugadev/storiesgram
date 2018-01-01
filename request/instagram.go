package request

import (
	"net/http"
	"strings"

	"github.com/gugadev/storiesgram/models"
	"github.com/levigross/grequests"
)

const uri = "https://i.instagram.com/api/v1/feed/user/"

/*
GetStories get stories of an user
@param {string} userid Instragram ID of target
@return raw response
*/
func GetStories(userid, sessionid string) []models.Story {
	// var result models.Story
	var result models.Raw
	var stories []models.Story
	options := &grequests.RequestOptions{
		Headers: map[string]string{
			"CacheControl":    "no-cache",
			"ContentLanguage": "en",
			"Vary":            "Cookie, Accept-Language",
			"AcceptEncoding":  "gzip, deflate",
			"Content-Type":    "application/json",
		},
		Cookies: []*http.Cookie{
			{
				Name:  "ds_user_id",
				Value: userid,
			},
			{
				Name:  "sessionid",
				Value: sessionid,
			},
		},
		UserAgent: "Instagram 10.3.2 (iPhone7,2; iPhone OS 9_3_3; en_US; en-US; scale=2.00; 750x1334) AppleWebkit/420+",
	}

	response, err := grequests.Get(uri+userid+"/reel_media/", options)
	if err != nil {
		panic(err)
	}
	response.JSON(&result)
	items := result.Items
	for _, value := range items {
		var image models.Image
		var story models.Story
		image = value.Images.Candidates[0]
		story.PK = value.PK
		story.Source = strings.Split(image.URL, "?")[0]
		if value.MediaType == 1 {
			story.Type = "image"
		} else {
			story.Type = "video"
		}
		stories = append(stories, story)
	}

	return stories
}
