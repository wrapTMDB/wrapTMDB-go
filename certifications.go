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
type certifications struct{}

var Certifications certifications

/********************
 * 1.GET /certification/movie/list
 * @description Get an up to date list of the officially supported movie certifications on TMDB.
 * @returns {any} JSON
 * @doc https://developers.themoviedb.org/3/certifications/get-movie-certifications
 ********************/
func (c *certifications) GetMovieCertifications() interface{} {
	var token = c_module.GetToken()
	var header = c_module.GetHeader()
	var targetURL string = baseURL +
		c_module.Route.CERTIFICATION +
		c_module.Route.MOVIE +
		"list" +
		`?api_key=` + token
	if token == "UnitTest_api_key" {
		return targetURL
	}
	var data, _ = http_methods_golang.GET(targetURL, header)
	return data
}

/********************
 * 2.GET /certification/tv/list
 * @description Get an up to date list of the officially supported TV show certifications on TMDB.
 * @returns {any} JSON
 * @doc https://developers.themoviedb.org/3/certifications/get-tv-certifications
 ********************/
func (c *certifications) GetTVCertifications() interface{} {
	var token = c_module.GetToken()
	var header = c_module.GetHeader()
	var targetURL string = baseURL +
		c_module.Route.CERTIFICATION +
		c_module.Route.TV +
		"list" +
		`?api_key=` + token
	if token == "UnitTest_api_key" {
		return targetURL
	}
	var data, _ = http_methods_golang.GET(targetURL, header)
	return data
}

/*
  1.GET Get Movie Certifications
  2.GET Get TV Certifications
*/
