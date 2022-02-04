package wraptmdb

import "github.com/kwangsing3/http_methods_golang"

/*
 * The MIT License (MIT)
 *
 * Copyright (c) kwangsing3
 *
 * https //github.com/wrapTMDB/wrapTMDB-ts
 *
 */

type search struct{}

var Search search

/********************
 * 1.GET /search/company
 * @description Search for companies.
 * @param {string} query keywords
 * @param {int} page (optional)
 * @returns JSON
 * @doc https //developers.themoviedb.org/3/search/search-companies
 ********************/
func (s *search) GetSearchCompanies(query string, page int) interface{} {
	var token = c_module.GetToken()
	var header = c_module.GetHeader()
	var targetURL string = baseURL +
		c_module.Route.SEARCH +
		c_module.Route.COMPANY +
		`?api_key=${token}`
	if query != "" {
		targetURL += `&query=${query}`
	}
	if page > 0 {
		targetURL += `&page=${page}`
	}
	if token == "UnitTest_api_key" {
		return targetURL
	}
	var data, _ = http_methods_golang.GET(targetURL, header)
	return data
}

/********************
 * 2.GET /search/collection
 * @description Search for collections.
 * @param {string} language(optional)
 * @param {string} query keywords
 * @param {int} page (optional)
 * @returns JSON
 * @doc https //developers.themoviedb.org/3/search/search-collections
 ********************/
func (s *search) GetSearchCollections(
	query string,
	language string,
	page int,
) interface{} {
	var token = c_module.GetToken()
	var header = c_module.GetHeader()
	var targetURL string = baseURL +
		c_module.Route.SEARCH +
		c_module.Route.COLLECTION +
		`?api_key=${token}`
	if query != "" {
		targetURL += `&query=${query}`
	}
	if language != "" {
		targetURL += `&language=${language}`
	}
	if page > 0 {
		targetURL += `&page=${page}`
	}
	if token == "UnitTest_api_key" {
		return targetURL
	}
	var data, _ = http_methods_golang.GET(targetURL, header)
	return data
}

/********************
 * 3.GET /search/keyword
 * @description Search for keywords.
 * @param {string} query keywords
 * @param {int} page (optional)
 * @returns JSON
 * @doc https //developers.themoviedb.org/3/search/search-keywords
 ********************/
func (s *search) GetSearchKeywords(query string, page int) interface{} {
	var token = c_module.GetToken()
	var header = c_module.GetHeader()
	var targetURL string = baseURL +
		c_module.Route.SEARCH +
		c_module.Route.KEYWORD +
		`?api_key=${token}`
	if query != "" {
		targetURL += `&query=${query}`
	}
	if page > 0 {
		targetURL += `&page=${page}`
	}
	if token == "UnitTest_api_key" {
		return targetURL
	}
	var data, _ = http_methods_golang.GET(targetURL, header)
	return data
}

/********************
 * 4.GET /search/movie
 * @description Search for movies.
 * @param {string} language(optional)
 * @param {string} query
 * @param {bool} include_adult (optional)
 * @param {string} region (optional)
 * @param {int} year (optional)
 * @param {int} primary_release_year (optional)
 * @returns JSON
 * @doc https //developers.themoviedb.org/3/search/search-movies
 ********************/
func (s *search) GetSearchMovies(
	query string,
	language string,
	page int,
	include_adult bool,
	region string,
	year int,
	primary_release_year int,
) interface{} {
	var token = c_module.GetToken()
	var header = c_module.GetHeader()
	var targetURL string = baseURL +
		c_module.Route.SEARCH +
		c_module.Route.MOVIE +
		`?api_key=${token}`
	if query != "" {
		targetURL += `&query=${query}`
	}
	if language != "" {
		targetURL += `&language=${language}`
	}
	if page > 0 {
		targetURL += `&page=${page}`
	}
	targetURL += `&include_adult=${include_adult}`
	if region != "" {
		targetURL += `&region=${region}`
	}
	if year > 0 {
		targetURL += `&year=${year}`
	}
	if region != "" {
		targetURL += `&primary_release_year=${primary_release_year}`
	}
	if token == "UnitTest_api_key" {
		return targetURL
	}
	var data, _ = http_methods_golang.GET(targetURL, header)
	return data
}

/********************
 * 5.GET /search/multi
 * @description Search multiple models in a single request.
 * Multi search currently supports searching for movies, tv shows and people in a single request.
 * @param {string} query
 * @param {string} language(optional)
 * @param {int} page (optional)
 * @param {bool} include_adult (optional)
 * @param {string} region (optional)
 * @returns JSON
 * @doc https //developers.themoviedb.org/3/search/multi-search
 ********************/
func (s *search) GetMultiSearch(
	query string,
	language string,
	page int,
	include_adult bool,
	region string,
) interface{} {
	var token = c_module.GetToken()
	var header = c_module.GetHeader()
	var targetURL string = baseURL + c_module.Route.SEARCH + "multi" + `?api_key=${token}`
	if query != "" {
		targetURL += `&query=${query}`
	}
	if language != "" {
		targetURL += `&language=${language}`
	}
	if page > 0 {
		targetURL += `&page=${page}`
	}
	targetURL += `&include_adult=${include_adult}`
	if region != "" {
		targetURL += `&region=${region}`
	}
	if token == "UnitTest_api_key" {
		return targetURL
	}
	var data, _ = http_methods_golang.GET(targetURL, header)
	return data
}

/********************
 * 6.GET /search/person
 * @description Search for people.
 * @param {string} language(optional)
 * @param {string} query
 * @param {int} page (optional)
 * @param {bool} include_adult (optional)
 * @param {string} region (optional)
 * @returns JSON
 * @doc https //developers.themoviedb.org/3/search/search-people
 ********************/
func (s *search) GetSearchPeople(
	query string,
	language string,
	page int,
	include_adult bool,
	region string,
) interface{} {
	var token = c_module.GetToken()
	var header = c_module.GetHeader()
	var targetURL string = baseURL +
		c_module.Route.SEARCH +
		c_module.Route.PEOPLE +
		`?api_key=${token}`
	if query != "" {
		targetURL += `&query=${query}`
	}
	if language != "" {
		targetURL += `&language=${language}`
	}
	if page > 0 {
		targetURL += `&page=${page}`
	}
	targetURL += `&include_adult=${include_adult}`
	if region != "" {
		targetURL += `&region=${region}`
	}
	if token == "UnitTest_api_key" {
		return targetURL
	}
	var data, _ = http_methods_golang.GET(targetURL, header)
	return data
}

/********************
 * 7.GET /search/tv
 * @description Search for a TV show.
 * @param {string} language(optional)
 * @param {int} page (optional)
 * @param {string} query
 * @param {bool} include_adult (optional)
 * @param {int} first_air_date_year (optional)
 * @returns JSON
 * @doc https //developers.themoviedb.org/3/search/search-tv-shows
 ********************/
func (s *search) GetSearchTVShows(
	query string,
	language string,
	page int,
	include_adult bool,
	first_air_date_year int,
) interface{} {
	var token = c_module.GetToken()
	var header = c_module.GetHeader()
	var targetURL string = baseURL + c_module.Route.SEARCH + c_module.Route.TV + `?api_key=${token}`
	if query != "" {
		targetURL += `&query=${query}`
	}
	if language != "" {
		targetURL += `&language=${language}`
	}
	if page > 0 {
		targetURL += `&page=${page}`
	}
	targetURL += `&include_adult=${include_adult}`

	if first_air_date_year > 0 {
		targetURL += `&first_air_date_year=${first_air_date_year}`
	}
	if token == "UnitTest_api_key" {
		return targetURL
	}
	var data, _ = http_methods_golang.GET(targetURL, header)
	return data
}

/*
1.GET Search Companies
2.GET Search Collections
3.GET Search Keywords
4.GET Search Movies
5.GET Multi Search
6.GET Search People
7.GET Search TV Shows
*/
