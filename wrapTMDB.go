package wraptmdb

import "encoding/json"

var Router = struct {
	MOVIE          string
	TV             string
	ACCOUNT        string
	CERTIFICATION  string
	AUTHENTICATION string
	COLLECTION     string
	COMPANY        string
	CONFIGURATION  string
	CREDIT         string
	DISCOVER       string
	FIND           string
	GENRE          string
	GUESTSESSION   string
	KEYWORD        string
	LIST           string
	NETWORK        string
	TRENDING       string
	PEOPLE         string
	REVIEWS        string
	SEARCH         string
	EPISODE        string
	EPISODEGROUPS  string
	WATCHPROVIDERS string
}{
	MOVIE:          "movie/",
	TV:             "tv/",
	ACCOUNT:        "account/",
	CERTIFICATION:  "certification/",
	AUTHENTICATION: "Authentication/",
	COLLECTION:     "colleciton/",
	COMPANY:        "company/",
	CONFIGURATION:  "Configuration/",
	CREDIT:         "credit/",
	DISCOVER:       "discover/",
	FIND:           "find/",
	GENRE:          "genre/",
	GUESTSESSION:   "guest_session/",
	KEYWORD:        "keyword/",
	LIST:           "list/",
	NETWORK:        "network/",
	TRENDING:       "trending/",
	PEOPLE:         "person/",
	REVIEWS:        "reviews/",
	SEARCH:         "search/",
	EPISODE:        "episode/",
	EPISODEGROUPS:  "episode_group/",
	WATCHPROVIDERS: "watch/providers/",
}

var header = map[string]string{
	"User-Agent": "",
	"Referer":    "",
}
var common = map[string]interface{}{
	"TOKEN":    "",
	"BASE_URL": "https://api.themoviedb.org/3/",
	"HEADER":   header,
}

//#region function

//header
func SetHeader(user_agent string, refer string) {
	header["User_Agent"] = user_agent
	header["Referer"] = refer
}
func GetHeader() ([]byte, error) {
	return json.Marshal(header)
}

//token
func Init(input string) {
	common["TOKEN"] = input
}

func GetToken() string {
	return common["TOKEN"].(string)
}

//URL
func GetURL() string {
	return common["BASE_URL"].(string)
}

//#endregion
