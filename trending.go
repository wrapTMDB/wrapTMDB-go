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

type trending struct{}

var Trending trending

/********************
 * 1.GET /trending/{media_type}/{time_window}
 * @description Get the daily or weekly trending items.
 * The daily trending list tracks items over the period of a day while items have a 24 hour half life.
 * The weekly list tracks items over a 7 day period, with a 7 day half life.
 * @param {string} media_type ("all" || "movie" || "tv" || "person")
 * @param {string} time_window ("day" || "week")
 * @returns {any} JSON
 * @doc https://developers.themoviedb.org/3/trending/get-trending
 ********************/
func (t *trending) GetTrending(media_type string, time_window string) interface{} {
	var token = c_module.GetToken()
	var header = c_module.GetHeader()
	var targetURL string = baseURL +
		c_module.Route.TRENDING +
		`${media_type}/` +
		`${time_window}` +
		`?api_key=${token}`
	if token == "UnitTest_api_key" {
		return targetURL
	}
	var data, _ = http_methods_golang.GET(targetURL, header)
	return data
}

/*
1.GET Get Trending
*/
