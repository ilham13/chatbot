package services

import (
	"encoding/json"
	"net/http"
)

var baseURL = "http://dev.to/api/"

type Article struct {
	TypeOf              string `json:"type_of"`
	ID                  int    `json:"id"`
	Title               string `json:"title"`
	Description         string `json:"description"`
	CoverImage          string `json:"cover_image"`
	ReadablePublishDate string `json:"readable_publish_date"`
	SocialImage         string `json:"social_image"`
	PublishedAt         string `json:"published_at"`
	// Slug                     string
	// Path                     string
	// Url                      string
	// CanonicalUrl            string
	// comments_count           string
	// positive_reactions_count string
	// collection_id            string
	// created_at               string
	// edited_at                string
	// crossposted_at           string
	// published_at             string
	// last_comment_at          string
	// published_timestamp      string
	// tag_list                 struct {
	// 	discuss string
	// 	design  string
	// }
	// tags string
	// user struct {
	// 	name             string
	// 	username         string
	// 	twitter_username string
	// 	github_username  string
	// 	website_url      string
	// 	profile_image    string
	// 	profile_image_90 string
	// }
	// flare_tag struct {
	// 	name           string
	// 	bg_color_hex   string
	// 	text_color_hex string
	// }
}

func (article *Article) getArticles() ([]Article, error) {
	var err error
	var client = &http.Client{}
	var data []Article

	request, err := http.NewRequest("GET", baseURL+"articles?page=1", nil)

	if err != nil {
		return nil, err
	}

	response, err := client.Do(request)

	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	err = json.NewDecoder(response.Body).Decode(&data)
	// err = json.Unmarshal([]byte(jsonString), &data)

	if err != nil {
		return nil, err
	}
	return data, nil
}
