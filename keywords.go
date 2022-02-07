package wraptmdb

import "github.com/kwangsing3/http_methods_golang"

/*
 * The MIT License (MIT)
 *
 * Copyright (c) kwangsing3
 *
 * https://github.com/wrapTMDB/wrapTMDB-go
 *
 */

type keywords struct{}

var KeyWords keywords

/********************
 * 1.GET /keyword/{keyword_id}
 * @description
 * @param {number|string} keyword_id
 * @returns JSON
 * @doc https://developers.themoviedb.org/3/keywords/get-keyword-details
 ********************/
func (k *keywords) GetDetails(keyword_id string) interface{} {
	var token = c_module.GetToken()
	var header = c_module.GetHeader()
	var targetURL string = baseURL + c_module.Route.KEYWORD + `${keyword_id}` + `?api_key=${token}`

	if token == "UnitTest_api_key" {
		return targetURL
	}
	var data, _ = http_methods_golang.GET(targetURL, header)
	return data
}

/********************
 * 2.GET /keyword/{keyword_id}/movies
 * @description Get the movies that belong to a keyword.
 * We highly recommend using "movie discover" instead of this method as it is much more flexible.
 * @param {number|string} keyword_id
 * @param {number|string} language(optional)
 * @param {boolean} include_adult(optional)
 * @returns JSON
 * @doc https://developers.themoviedb.org/3/keywords/get-movies-by-keyword
 ********************/
func (k *keywords) GetMovies(
	keyword_id string,
	language string,
	include_adult bool,
) interface{} {
	var token = c_module.GetToken()
	var header = c_module.GetHeader()
	var targetURL string = baseURL +
		c_module.Route.KEYWORD +
		`${keyword_id}/` +
		"movies" +
		`?api_key=${token}`
	if language != "" {
		targetURL += `&language=${language}`
	}

	targetURL += `&include_adult=${include_adult}`
	if token == "UnitTest_api_key" {
		return targetURL
	}
	var data, _ = http_methods_golang.GET(targetURL, header)
	return data
}

/*
  1.GET Get Details
  2.GET Get Movies
*/
