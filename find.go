package wraptmdb

import "github.com/kwangsing3/http_methods_golang"

/*
 * The MIT License (MIT)
 *
 * Copyright (c) kwangsing3
 *
 * https://github.com/wrapTMDB/wrapTMDB-ts
 *
 */
type find struct{}

var Find find

/********************
 * 1.GET /find/{external_id}
 * @description The find method makes it easy to search for objects in TMDB database by an external id.
 * @param {int|string} external_id
 * @param {string} external_source (default imdb_id)
 * @param {string} language (optional)
 * @returns JSON
 * @doc https://developers.themoviedb.org/3/find/find-by-id
 ********************/
func (f *find) GetFindByID(
	external_id string,
	language string,
	external_source string,
) interface{} {
	var token = c_module.GetToken()
	var header = c_module.GetHeader()
	var targetURL string = baseURL + c_module.Route.FIND + external_id + `?api_key=` + token
	if language != "" {
		targetURL += `&language=` + language
	}
	if external_source != "" {
		targetURL += `&external_source=` + external_source
	}
	if token == "UnitTest_api_key" {
		return targetURL
	}
	var data, _ = http_methods_golang.GET(targetURL, header)
	return data
}

/*
  1.GET Find by ID
*/
