package wraptmdb

import "github.com/kwangsing3/http_methods_golang"

/*
 * The MIT License (MIT)
 *
 * Copyright (c) kwangsing3
 *
 * https://github.com/wrapTMDB/wraptmdb-go
 *
 */

type genres struct{}

var Genres genres

/********************
 * 1.GET /genre/movie/list
 * @description Get the list of official genres for movies.
 * @param {string} language (optional)
 * @returns JSON
 * @doc https://developers.themoviedb.org/3/genres/get-movie-list
 ********************/
func (g *genres) GetMovieList(language string) interface{} {
	var token = c_module.GetToken()
	var header = c_module.GetHeader()
	var targetURL string = baseURL +
		c_module.Route.GENRE +
		c_module.Route.MOVIE +
		"list" +
		`?api_key=` + token
	if language != "" {
		targetURL += `&language=` + language
	}

	if token == "UnitTest_api_key" {
		return targetURL
	}
	var data, _ = http_methods_golang.GET(targetURL, header)
	return data
}

/********************
 * 2.GET /genre/tv/list
 * @description Get the list of official genres for TV shows.
 * @param {string} language (optional)
 * @returns JSON
 * @doc https://developers.themoviedb.org/3/genres/get-tv-list
 ********************/
func (g *genres) GetTVList(language string) interface{} {
	var token = c_module.GetToken()
	var header = c_module.GetHeader()
	var targetURL string = baseURL +
		c_module.Route.GENRE +
		c_module.Route.TV +
		"list" +
		`?api_key=` + token
	if language != "" {
		targetURL += `&language=` + language
	}

	if token == "UnitTest_api_key" {
		return targetURL
	}
	var data, _ = http_methods_golang.GET(targetURL, header)
	return data
}
