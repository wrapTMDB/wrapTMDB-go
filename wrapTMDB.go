package wraptmdb_go

type routepath struct {
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
}
type common struct {
	TOKEN    string
	BASE_URL string
	HEADER   map[string]string
	Route    routepath
}

var c_module = common{
	TOKEN:    "",
	BASE_URL: "https://api.themoviedb.org/3/",
	HEADER: map[string]string{
		"User-Agent": "",
		"Referer":    "",
	},
	Route: routepath{
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
	},
}
var baseURL = c_module.GetURL()

//#region function
func Init(input string) {
	c_module.TOKEN = input
}

//header
func SetHeader(input map[string]string) {
	c_module.HEADER = input
}
func (c *common) GetHeader() map[string]string {
	return c.HEADER
}

//token
func (c *common) Init(input string) {
	c.TOKEN = input
}

func (c *common) GetToken() string {
	return c.TOKEN
}

//URL
func (c *common) GetURL() string {
	return c.BASE_URL
}

//#endregion
