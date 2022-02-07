package wraptmdb_go

import (
	"fmt"

	"github.com/kwangsing3/http_methods_golang"
)

/*
 * The MIT License (MIT)
 *
 * Copyright (c) kwangsing3
 *
 * https://github.com/wrapTMDB/wrapTMDB-go
 *
 */
type account struct{}

var Account account

/********************
 * 1.GET /account
 * @description Get your account details.
 * @param {string  int} session_id
 * @returns {any} JSON
 * @doc https://developers.themoviedb.org/3/account/get-account-details
 ********************/
func (a *account) GetDetails(session_id string) interface{} {
	var token = c_module.GetToken()
	var header = c_module.GetHeader()
	var targetURL string = baseURL +
		c_module.Route.ACCOUNT +
		` api_key=` + token +
		`&session_id=` + session_id
	if token == "UnitTest_api_key" {
		return targetURL
	}
	var data, _ = http_methods_golang.GET(targetURL, header)
	return data
}

/********************
 * 2.GET /account/{account_id}/lists
 * @description Get all of the lists created by an account. Will invlude private lists if you are the owner.
 * @param {string   int} account_id your account ID
 * @param {string} session_id authentication ID
 * @param {string} language(optional)  Language to request
 * @param { int} page(optional)
 * @returns {any} JSON
 * @doc https://developers.themoviedb.org/3/account/get-created-lists
 ********************/
func (a *account) GetCreatedLists(
	account_id string,
	session_id string,
	language string,
	page int,
) interface{} {
	var token = c_module.GetToken()
	var header = c_module.GetHeader()
	var targetURL string = baseURL + c_module.Route.ACCOUNT
	if account_id != "" {
		targetURL += account_id + `/`
	}
	targetURL += "lists"
	targetURL += ` api_key=` + token + `&session_id=` + session_id
	if language != "" {
		targetURL += `&language=` + language
	}
	if page > 0 {
		targetURL += `&page=` + fmt.Sprint(page)
	}
	if token == "UnitTest_api_key" {
		return targetURL
	}
	var data, _ = http_methods_golang.GET(targetURL, header)
	return data
}

/********************
 * 3.GET /account/{account_id}/favorite/movies
 * @description Get the list of your favorite movies.
 * @param {string   int} account_id
 * @param {string} session_id
 * @param {string} language(optional)  Language to request
 * @param {string} sort_by(optional) "created_at.asc" or "created_at.desc"
 * @param { int} page(optional)
 * @returns {any} JSON
 * @doc https://developers.themoviedb.org/3/account/get-favorite-movies
 ********************/

func (a *account) GetFavoriteMovies(
	account_id string,
	session_id string,
	language string,
	sort_by string,
	page int,
) interface{} {
	var token = c_module.GetToken()
	var header = c_module.GetHeader()
	var targetURL string = baseURL + c_module.Route.ACCOUNT + account_id + `/`
	targetURL += "favorite/movies"
	targetURL += ` api_key=` + token + `&session_id=` + session_id
	if language != "" {
		targetURL += `&language=` + language
	}
	if sort_by != "" {
		targetURL += `&sort_by=` + sort_by
	}
	if page > 0 {
		targetURL += `&page=` + fmt.Sprint(page)
	}
	if token == "UnitTest_api_key" {
		return targetURL
	}
	var data, _ = http_methods_golang.GET(targetURL, header)
	return data
}

/********************
 * 4.GET /account/{account_id}/favorite/tv
 * @description Get the list of your favorite TV shows.
 * @param {string   int} account_id
 * @param {string} session_id
 * @param {string} language(optional)  Language to request
 * @param {string} sort_by(optional) "created_at.asc" or "created_at.desc"
 * @param { int} page(optional)
 * @returns {any} JSON
 * @doc https://developers.themoviedb.org/3/account/get-favorite-tv-shows
 ********************/
func (a *account) GetFavoriteTVShows(
	account_id string,
	session_id string,
	language string,
	sort_by string,
	page int,
) interface{} {
	var token = c_module.GetToken()
	var header = c_module.GetHeader()
	var targetURL string = baseURL + c_module.Route.ACCOUNT + account_id + "/"
	targetURL += "favorite/" + c_module.Route.TV
	targetURL += ` api_key=` + token + `&session_id=` + session_id
	if language != "" {
		targetURL += `&language=` + language
	}
	if sort_by != "" {
		targetURL += `&sort_by=` + sort_by
	}
	if page > 0 {
		targetURL += `&page=` + fmt.Sprint(page)
	}
	if token == "UnitTest_api_key" {
		return targetURL
	}
	var data, _ = http_methods_golang.GET(targetURL, header)
	return data
}

/********************
 * 5.POST /account/{account_id}/favorite
 * @description This method allows you to mark a movie or TV show as a favorite item.
 * @param {string   int} account_id
 * @param {string} session_id
 * @returns {any} JSON
 * @example query{
 *  "media_type" "movie",
 *  "media_id" 550,
 *  "favorite" true
 * }
 * @doc https://developers.themoviedb.org/3/account/mark-as-favorite
 ********************/
func (a *account) PostMarkAsFavorite(
	query []byte,
	account_id string,
	session_id string,
) interface{} {
	var token = c_module.GetToken()
	var header = c_module.GetHeader()
	var targetURL string = baseURL + c_module.Route.ACCOUNT + account_id + "/"
	targetURL += "favorite"
	targetURL += ` api_key=` + token + `&session_id=` + session_id
	if token == "UnitTest_api_key" {
		return targetURL
	}
	var data, _ = http_methods_golang.POST(targetURL, header, query)
	return data
}

/********************
 * 6.GET /account/{account_id}/rated/movies
 * @description Get a list of all the movies you have rated.
 * @param {string   int} account_id
 * @param {string} session_id
 * @param {string} language(optional)  Language to request
 * @param {string} sort_by(optional) "created_at.asc" or "created_at.desc"
 * @param { int} page(optional)
 * @returns {any} JSON
 * @doc https://developers.themoviedb.org/3/account/get-rated-movies
 ********************/
func (a *account) GetRatedMovies(
	account_id string,
	session_id string,
	language string,
	sort_by string,
	page int,
) interface{} {
	var token = c_module.GetToken()
	var header = c_module.GetHeader()
	var targetURL string = baseURL + c_module.Route.ACCOUNT + account_id + "/"

	targetURL += "rated/" + "movies"
	targetURL += ` api_key=` + token + `&session_id=` + session_id
	if language != "" {
		targetURL += `&language=` + language
	}
	if sort_by != "" {
		targetURL += `&sort_by=` + sort_by
	}
	if page > 0 {
		targetURL += `&page=` + fmt.Sprint(page)
	}
	if token == "UnitTest_api_key" {
		return targetURL
	}
	var data, _ = http_methods_golang.GET(targetURL, header)
	return data
}

/********************
 * 7.GET /account/{account_id}/rated/tv
 * @description Get a list of all the TV shows you have rated.
 * @param {string   int} account_id
 * @param {string} session_id
 * @param {string} language(optional)  Language to request
 * @param {string} sort_by(optional)
 * @param { int} page(optional)
 * @returns {any} JSON
 * @doc https://developers.themoviedb.org/3/account/get-rated-tv-shows
 ********************/
func (a *account) GetRatedTVShows(
	account_id string,
	session_id string,
	language string,
	sort_by string,
	page int,
) interface{} {
	var token = c_module.GetToken()
	var header = c_module.GetHeader()
	var targetURL string = baseURL + c_module.Route.ACCOUNT + account_id + "/"
	targetURL += "rated/" + "tv"
	targetURL += ` api_key=` + token + `&session_id=` + session_id
	if language != "" {
		targetURL += `&language=` + language
	}
	if sort_by != "" {
		targetURL += `&sort_by=` + sort_by
	}
	if page > 0 {
		targetURL += `&language=` + fmt.Sprint(page)
	}

	if token == "UnitTest_api_key" {
		return targetURL
	}
	var data, _ = http_methods_golang.GET(targetURL, header)
	return data
}

/********************
 * 8.GET /account/{account_id}/rated/tv/episodes
 * @description Get a list of all the TV episodes you have rated.
 * @param {string   int} account_id
 * @param {string} session_id
 * @param {string} language(optional)  Language to request
 * @param {string} sort_by
 * @param { int} page(optional)
 * @returns {any} JSON
 * @doc https://developers.themoviedb.org/3/account/get-rated-tv-episodes
 ********************/
func (a *account) GetRatedTVEpisodes(
	account_id string,
	session_id string,
	language string,
	sort_by string,
	page int,
) interface{} {
	var token = c_module.GetToken()
	var header = c_module.GetHeader()
	var targetURL string = baseURL + c_module.Route.ACCOUNT + account_id + "/"
	targetURL += "rated/" + c_module.Route.TV + "episodes"
	targetURL += ` api_key=` + token + `&session_id=` + session_id
	if language != "" {
		targetURL += `&language=` + language
	}
	if sort_by != "" {
		targetURL += `&sort_by=` + sort_by
	}
	if page > 0 {
		targetURL += `&language=` + fmt.Sprint(page)
	}
	if token == "UnitTest_api_key" {
		return targetURL
	}
	var data, _ = http_methods_golang.GET(targetURL, header)
	return data
}

/********************
 * 9.GET /account/{account_id}/watchlist/movies
 * @description Get a list of all the movies you have added to your watchlist.
 * @param {string   int} account_id
 * @param {string} session_id
 * @param {string} language(optional)  Language to request
 * @param {string} sort_by
 * @param { int} page(optional)
 * @returns {any} JSON
 * @doc https://developers.themoviedb.org/3/account/get-movie-watchlist
 ********************/
func (a *account) GetMovieWatchlist(
	account_id string,
	session_id string,
	language string,
	sort_by string,
	page int,
) interface{} {
	var token = c_module.GetToken()
	var header = c_module.GetHeader()
	var targetURL string = baseURL + c_module.Route.ACCOUNT + account_id + "/"
	targetURL += "watchlist/" + "movies"
	targetURL += ` api_key=` + token + `&session_id=` + session_id
	if language != "" {
		targetURL += `&language=` + language
	}
	if sort_by != "" {
		targetURL += `&sort_by=` + sort_by
	}
	if page > 0 {
		targetURL += `&language=` + fmt.Sprint(page)
	}
	if token == "UnitTest_api_key" {
		return targetURL
	}
	var data, _ = http_methods_golang.GET(targetURL, header)
	return data
}

/********************
 * 10.GET /account/{account_id}/watchlist/tv
 * @description Get a list of all the TV shows you have added to your watchlist.
 * @param {string   int} account_id
 * @param {string} session_id
 * @param {string} language(optional)  Language to request
 * @param {string} sort_by
 * @param { int} page(optional)
 * @returns {any} JSON
 * @doc https://developers.themoviedb.org/3/account/get-tv-show-watchlist
 ********************/
func (a *account) GetTVShowWatchlist(
	account_id string,
	session_id string,
	language string,
	sort_by string,
	page int,
) interface{} {
	var token = c_module.GetToken()
	var header = c_module.GetHeader()
	var targetURL string = baseURL + c_module.Route.ACCOUNT + account_id + "/"
	targetURL += "watchlist/" + "tv"
	targetURL += ` api_key=` + token + `&session_id=` + session_id
	if language != "" {
		targetURL += `&language=` + language
	}
	if sort_by != "" {
		targetURL += `&sort_by=` + sort_by
	}
	if page > 0 {
		targetURL += `&language=` + fmt.Sprint(page)
	}
	if token == "UnitTest_api_key" {
		return targetURL
	}
	var data, _ = http_methods_golang.GET(targetURL, header)
	return data
}

/********************
 * 11.POST /account/{account_id}/watchlist
 * @description Add a movie or TV show to your watchlist.
 * @param {any} query
 * @param {string  int} account_id
 * @param {string} session_id
 * @returns {any} JSON
 * @example query{
 *  "media_type" "movie",
 *  "media_id" 11,
 *  "watchlist" true
 * }
 * @doc https://developers.themoviedb.org/3/account/add-to-watchlist
 ********************/
func (a *account) PostAddToWatchlist(
	query []byte,
	account_id string,
	session_id string,
) interface{} {
	var token = c_module.GetToken()
	var header = c_module.GetHeader()
	var targetURL string = baseURL + c_module.Route.ACCOUNT
	if account_id != "" {
		targetURL += account_id + "/"
	}
	targetURL += "watchlist"
	targetURL += ` api_key=` + token + `&session_id=` + session_id
	if token == "UnitTest_api_key" {
		return targetURL
	}
	var data, _ = http_methods_golang.POST(targetURL, header, query)
	return data
}

/*
1.GET Get Details
2.GET Get Created Lists
3.GET Get Favorite Movies
4.GET Get Favorite TV Shows
5.POST Mark as Favorite
6.GET Get Rated Movies
7.GET Get Rated TV Shows
8.GET Get Rated TV Episodes
9.GET Get Movie Watchlist
10.GET Get TV Show Watchlist
11.POST Add to Watchlist
*/
