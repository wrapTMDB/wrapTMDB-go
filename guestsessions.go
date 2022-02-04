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

type guestsessions struct{}

var GuestSessions guestsessions

/********************
 * 1.GET /guest_session/{guest_session_id}/rated/movies
 * @description Get the rated movies for a guest session.
 * @param {number|string} guest_session_id
 * @param {string} language (optional)
 * @param {string} sort_by (optional) "created_at.asc" or "created_at.desc"
 * @returns JSON
 * @doc https://developers.themoviedb.org/3/guest-sessions/get-guest-session-rated-movies
 ********************/
func (g *guestsessions) GetRatedMovies(
	guest_session_id string,
	language string,
	sort_by string,
) interface{} {
	var token = c_module.GetToken()
	var header = c_module.GetHeader()
	var targetURL string = baseURL +
		c_module.Route.GUESTSESSION +
		guest_session_id + `/` +
		"rated/movies" +
		`?api_key=` + token
	if language != "" {
		targetURL += `&language=` + language
	}
	if sort_by != "" {
		targetURL += `&sort_by=` + sort_by
	}
	if token == "UnitTest_api_key" {
		return targetURL
	}
	var data, _ = http_methods_golang.GET(targetURL, header)
	return data
}

/********************
 * 2.GET /guest_session/{guest_session_id}/rated/tv
 * @description Get the rated TV shows for a guest session.
 * @param {number|string} guest_session_id
 * @param {string} language (optional)
 * @param {string} sort_by (optional) "created_at.asc" or "created_at.desc"
 * @returns JSON
 * @doc https://developers.themoviedb.org/3/guest-sessions/get-guest-session-rated-tv-shows
 ********************/
func (g *guestsessions) GetRatedTVShows(
	guest_session_id string,
	language string,
	sort_by string,
) interface{} {
	var token = c_module.GetToken()
	var header = c_module.GetHeader()
	var targetURL string = baseURL +
		c_module.Route.GUESTSESSION +
		guest_session_id + "/" +
		"rated/tv" +
		`?api_key=` + token
	if language != "" {
		targetURL += `&language=` + language
	}
	if sort_by != "" {
		targetURL += `&sort_by=` + sort_by
	}
	if token == "UnitTest_api_key" {
		return targetURL
	}
	var data, _ = http_methods_golang.GET(targetURL, header)
	return data
}

/********************
 * 3.GET /guest_session/{guest_session_id}/rated/tv/episodes
 * @descriptionGet the rated TV episodes for a guest session.
 * @param {number|string} guest_session_id
 * @param {string} language (optional)
 * @param {string} sort_by (optional) "created_at.asc" or "created_at.desc"
 * @returns JSON
 * @doc https://developers.themoviedb.org/3/guest-sessions/get-guest-session-rated-tv-shows
 ********************/
func (g *guestsessions) GetRatedTVEpisodes(
	guest_session_id string,
	language string,
	sort_by string,
) interface{} {
	var token = c_module.GetToken()
	var header = c_module.GetHeader()
	var targetURL string = baseURL +
		c_module.Route.GUESTSESSION +
		guest_session_id + "/" +
		"rated/tv/episodes" +
		`?api_key=` + token
	if language != "" {
		targetURL += `&language=` + language
	}
	if sort_by != "" {
		targetURL += `&sort_by=` + sort_by
	}

	if token == "UnitTest_api_key" {
		return targetURL
	}
	var data, _ = http_methods_golang.GET(targetURL, header)
	return data
}

/*
1.GET Get Rated Movies
2.GET Get Rated TV Shows
3.GET Get Rated TV Episodes
*/
