package wraptmdb

import (
	"github.com/kwangsing3/http_methods_golang"
)

/*
 * The MIT License (MIT)
 *
 * Copyright (c) kwangsing3
 *
 * https://github.com/wrapTMDB/wraptmdb_go
 *
 */

type discover struct{}

var Discover discover

/********************
 * 1./discover/movie
 * @description Discover movies by different types of data like average rating,
 * number of votes, genres and certifications.
 * You can get a valid list of certifications from the "certifications list" method.
 * @param {any} query
 * @example "query" are more than 32 available properties.... see more in discover.ts
 * @returns JSON
 * @doc https://developers.themoviedb.org/3/discover/movie-discover
 ********************/
func (d *discover) GetMovieDiscover(query map[string]string) interface{} {
	var token = c_module.GetToken()
	var header = c_module.GetHeader()
	var targetURL string = baseURL +
		c_module.Route.DISCOVER +
		c_module.Route.MOVIE +
		` api_key=` + token
	var defauquery = map[string]string{
		"language":                      "en-US",
		"sort_by":                       "popularity.desc",
		"page":                          "1",
		"include_adult":                 "false",
		"include_video":                 "false",
		"with_watch_monetization_types": "flatrate",
	}
	if query == nil {
		targetURL += stringByQuery(query)
	} else {
		targetURL += stringByQuery(defauquery)
	}

	if token == "UnitTest_api_key" {
		return targetURL
	}
	var data, _ = http_methods_golang.GET(targetURL, header)
	return data
}

/*
   @example movieQuery{
       "language"  {string},
       "region"  {string},
       "sort_by"  {string},
       "certification_country"  {string},
       "certification"  {string},
       "certification.lte"  {string},
       "certification.gte"  {string},
       "include_adult"  {bool},
       "include_video"  {bool},
       "page"  {number},
       "primary_release_year"  {string},
       "primary_release_date.gte"  {string},
       "primary_release_date.lte"  {string},
       "release_date.gte"  {string},
       "release_date.lte"  {string},
       "with_release_type"  {string},
       "year"  {string},
       "vote_count.gte"  {string},
       "vote_count.lte"  {string},
       "vote_average.gte"  {string},
       "vote_average.lte"  {string},
       "with_cast"  {string},
       "with_crew"  {string},
       "with_people"  {string},
       "with_companies"  {string},
       "with_genres"  {string},
       "without_genres"  {string},
       "with_keywords"  {string},
       "without_keywords"  {string},
       "with_runtime.gte"  {string},
       "with_runtime.lte"  {string},
       "with_original_language"  {string},
       "with_watch_providers"  {string},
       "watch_region"  {string},
       "with_watch_monetization_types"  {string} ,
       "without_companies"  {string}
   }
   ....more available properties plz visit
   @doc https://developers.themoviedb.org/3/discover/movie-discover
*/
/********************
 * 2./discover/tv
 * @description Discover TV shows by different types of data like average rating,
 * number of votes, genres, the network they aired on and air dates.
 * @param {any} query
 * @example "query" are more than 24 available properties.... see more in discover.ts
 * @returns JSON
 * @doc https://developers.themoviedb.org/3/discover/tv-discover
 ********************/
func (d *discover) GetTVDiscover(query map[string]string) interface{} {
	var token = c_module.GetToken()
	var header = c_module.GetHeader()
	var targetURL string = baseURL +
		c_module.Route.DISCOVER +
		c_module.Route.TV +
		` api_key=` + token
	var defauquery = map[string]string{
		"language":                      "en-US",
		"sort_by":                       "popularity.desc",
		"page":                          "1",
		"timezone":                      "America/New_York",
		"with_watch_monetization_types": "flatrate",
	}
	if query == nil {
		targetURL += stringByQuery(query)
	} else {
		targetURL += stringByQuery(defauquery)
	}
	if token == "UnitTest_api_key" {
		return targetURL
	}
	var data, _ = http_methods_golang.GET(targetURL, header)
	return data
}

/*
   @example tvQuery{
       "language"  {string},
       "sort_by"  {string},
       "air_date.gte"  {string},
       "air_date.lte"  {string},
       "first_air_date.gte"  {string},
       "first_air_date.lte"  {string},
       "first_air_date_year"  {string},
       "page"  {number},
       "timezone"  {string},
       "vote_average.gte"  {string},
       "vote_count.gte"  {string},
       "with_genres"  {string},
       "with_networks"  {string},
       "without_genres"  {string},
       "with_runtime.gte"  {string},
       "with_runtime.lte"  {string},
       "include_null_first_air_dates"  {bool},
       "with_original_language"  {string},
       "without_keywords"  {string},
       "screened_theatrically"  {bool},
       "with_companies"  {string},
       "with_keywords"  {string},
       "with_watch_providers"  {string},
       "watch_region"  {string},
       "with_watch_monetization_types"  {string} ,
       "with_status"  {number},
       "with_type"  {number},
       "without_companies"  {string}
   }
   ....more available properties plz visit
   @doc https://developers.themoviedb.org/3/discover/tv-discover
*/

func stringByQuery(query map[string]string) string {
	var result = ""
	for key, value := range query {
		result += `&` + key + `=` + value
	}
	return result
}

/*
  1.GET Movie Discover
  2.GET TV Discover
*/
