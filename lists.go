package wraptmdb_go

import (
	"strconv"

	"github.com/kwangsing3/http_methods_golang"
)

/*
 * The MIT License (MIT)
 *
 * Copyright (c) kwangsing3
 *
 * https://github.com/wrapTMDB/wrapTMDB_go
 *
 */
type list struct{}

var Lists list

/********************
 * 1.GET /list/{list_id}
 * @description Get the details of a list.
 * @param { int|string} list_id
 * @param { int|string} language(optional)
 * @returns JSON
 * @doc https://developers.themoviedb.org/3/lists/get-list-details
 ********************/
func (l *list) GetDetails(list_id string, language string) interface{} {
	var token = c_module.GetToken()
	var header = c_module.GetHeader()
	var targetURL string = baseURL + c_module.Route.LIST + list_id + `?api_key=` + token
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
 * 2.GET /list/{list_id}/item_status
 * @description You can use this method to check if a movie has already been added to the list.
 * @param { int|string} list_id
 * @param { int|string} movie_id
 * @returns JSON
 * @doc https://developers.themoviedb.org/3/lists/check-item-status
 ********************/
func (l *list) GetCheckItemStatus(
	list_id string,
	movie_id string,
) interface{} {
	var token = c_module.GetToken()
	var header = c_module.GetHeader()
	var targetURL string = baseURL +
		c_module.Route.LIST +
		list_id + `/` +
		"item_status" +
		`?api_key=` + token
	if movie_id != "" {
		targetURL += `&movie_id=` + movie_id
	}
	if token == "UnitTest_api_key" {
		return targetURL
	}
	var data, _ = http_methods_golang.GET(targetURL, header)
	return data
}

/********************
 * 3.POST /list
 * @description Create a list.
 * @param {any} query
 * @param { int|string} session_id
 * @example query{
 *    "name" "This is my awesome test list.",
 *    "description" "Just an awesome list dawg.",
 *    "language" "en"
 * }
 * @returns JSON
 * @doc https://developers.themoviedb.org/3/lists/create-list
 ********************/
func (l *list) PostCreateList(query []byte, session_id string) interface{} {
	var token = c_module.GetToken()
	var header = c_module.GetHeader()
	var targetURL string = baseURL + c_module.Route.LIST + `?api_key=` + token
	if session_id != "" {
		targetURL += `&session_id=` + session_id
	}

	if token == "UnitTest_api_key" {
		return targetURL
	}
	var data, _ = http_methods_golang.POST(targetURL, header, query)
	return data
}

/********************
 * 4.POST /list/{list_id}/add_item
 * @description Add a movie to a list.
 * @param {any} query
 * @param { int|string} list_id
 * @param { int|string} session_id
 * @example query{
 *    "media_id" 18
 * }
 * @returns JSON
 * @doc https://developers.themoviedb.org/3/lists/add-movie
 ********************/
func (l *list) PostAddMovie(
	query []byte,
	list_id string,
	session_id string,
) interface{} {
	var token = c_module.GetToken()
	var header = c_module.GetHeader()
	var targetURL string = baseURL +
		c_module.Route.LIST +
		list_id + `/add_item` +
		`?api_key=` + token
	if session_id != "" {
		targetURL += `&session_id=` + session_id
	}

	if token == "UnitTest_api_key" {
		return targetURL
	}
	var data, _ = http_methods_golang.POST(targetURL, header, query)
	return data
}

/********************
 * 5.POST /list/{list_id}/remove_item
 * @description Remove a movie from a list.
 * @param {any} query
 * @param { int|string} list_id
 * @param { int|string} session_id
 * @example query{
 *    "media_id" 12
 * }
 * @returns JSON
 * @doc https://developers.themoviedb.org/3/lists/remove-movie
 ********************/
func (l *list) PostRemoveMovie(
	query []byte,
	list_id string,
	session_id string,
) interface{} {
	var token = c_module.GetToken()
	var header = c_module.GetHeader()
	var targetURL string = baseURL +
		c_module.Route.LIST +
		list_id + `/remove_item` +
		`?api_key=` + token
	if session_id != "" {
		targetURL += `&session_id=` + session_id
	}

	if token == "UnitTest_api_key" {
		return targetURL
	}
	var data, _ = http_methods_golang.POST(targetURL, header, query)
	return data
}

/********************
 * 6.POST /list/{list_id}/clear
 * @description Clear all of the items from a list.
 * @param { int|string} list_id
 * @param { int|string} session_id
 * @param {boolean} confirm
 * @returns JSON
 * @doc https://developers.themoviedb.org/3/lists/clear-list
 ********************/
func (l *list) PostClearList(
	list_id string,
	session_id string,
	confirm bool,
) interface{} {
	var token = c_module.GetToken()
	var header = c_module.GetHeader()
	var targetURL string = baseURL + c_module.Route.LIST + list_id + `/clear` + `?api_key=` + token
	if session_id != "" {
		targetURL += `&session_id=` + session_id
	}
	targetURL += `&confirm=` + strconv.FormatBool(confirm)
	if token == "UnitTest_api_key" {
		return targetURL
	}
	var data, _ = http_methods_golang.POST(targetURL, header, nil)
	return data
}

/********************
 * 7.DELETE /list/{list_id}
 * @description Delete a list.
 * @param { int|string} list_id
 * @param { int|string} session_id
 * @returns JSON
 * @doc https://developers.themoviedb.org/3/lists/delete-list
 ********************/
func (l *list) DeleteDeleteList(
	list_id string,
	session_id string,
) interface{} {
	var token = c_module.GetToken()
	var header = c_module.GetHeader()
	var targetURL string = baseURL + c_module.Route.LIST + list_id + `?api_key=` + token
	if session_id != "" {
		targetURL += `&session_id=` + session_id
	}
	if token == "UnitTest_api_key" {
		return targetURL
	}
	var data, _ = http_methods_golang.DELETE(targetURL, header, nil)
	return data
}

/*
  1.GET Get Details
  2.GET Check Item Status
  3.POST Create List
  4.POST Add Movie
  5.POST Remove Movie
  6.POST Clear List
  7.DELETE Delete List
*/
