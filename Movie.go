package wraptmdb

import (
	"encoding/json"
	"log"

	"github.com/kwangsing3/http_methods_golang"
)

var Movies movie

type movie struct {
}

/*

 */

func (r *movie) GetMsg() string {
	return "Hello"
}

/********************
 * 1.GET /movie/{movie_id}
 * @description Get the primary information about a movie.
 * @param {int|string} movie_id  Movie ID in TMDB
 * @param {string} language(optional)  Language to request
 * @returns {string} JSON
 * @doc https://developers.themoviedb.org/3/movies/get-movie-details
 ********************/
func (r *movie) GetDetail(movie_id string, language string) interface{} {
	//624860
	var token = GetToken()
	//var header, _ = GetHeader()
	var baseURL = GetURL()
	var targetURL string = baseURL + Router.MOVIE + movie_id + `?api_key=` + token

	if language != "" {
		targetURL += `&language=` + language
	}

	if token == "UnitTest_api_key" {
		return targetURL
	}
	var data, err = http_methods_golang.GET(targetURL)
	if err != nil {
		log.Print(err)
		return ""
	}
	var result GetDetailType
	json.Unmarshal(data, &result)
	return result
}

//#region query_struct
type GetDetailType struct {
	Adult               bool   `json:"adult"`
	BackdropPath        string `json:"backdrop_path"`
	BelongsToCollection struct {
		ID           int    `json:"id"`
		Name         string `json:"name"`
		PosterPath   string `json:"poster_path"`
		BackdropPath string `json:"backdrop_path"`
	} `json:"belongs_to_collection"`
	Budget int `json:"budget"`
	Genres []struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	} `json:"genres"`
	Homepage            string  `json:"homepage"`
	ID                  int     `json:"id"`
	ImdbID              string  `json:"imdb_id"`
	OriginalLanguage    string  `json:"original_language"`
	OriginalTitle       string  `json:"original_title"`
	Overview            string  `json:"overview"`
	Popularity          float64 `json:"popularity"`
	PosterPath          string  `json:"poster_path"`
	ProductionCompanies []struct {
		ID            int    `json:"id"`
		LogoPath      string `json:"logo_path"`
		Name          string `json:"name"`
		OriginCountry string `json:"origin_country"`
	} `json:"production_companies"`
	ProductionCountries []struct {
		Iso31661 string `json:"iso_3166_1"`
		Name     string `json:"name"`
	} `json:"production_countries"`
	ReleaseDate     string `json:"release_date"`
	Revenue         int    `json:"revenue"`
	Runtime         int    `json:"runtime"`
	SpokenLanguages []struct {
		EnglishName string `json:"english_name"`
		Iso6391     string `json:"iso_639_1"`
		Name        string `json:"name"`
	} `json:"spoken_languages"`
	Status      string  `json:"status"`
	Tagline     string  `json:"tagline"`
	Title       string  `json:"title"`
	Video       bool    `json:"video"`
	VoteAverage float64 `json:"vote_average"`
	VoteCount   int     `json:"vote_count"`
}

//#endregion
