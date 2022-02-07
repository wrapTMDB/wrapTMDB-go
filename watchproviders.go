package wraptmdb_go

import "github.com/kwangsing3/http_methods_golang"

/*
 * The MIT License (MIT)
 *
 * Copyright (c) kwangsing3
 *
 * https://github.com/wrapTMDB/wrapTMDB_go
 *
 */

type watchproviders struct{}

var Watchproviders watchproviders

/********************
 * 1.GET /watch/providers/regions
 * @description Returns a list of all of the countries we have watch provider (OTT/streaming) data for.
 * @param {string} language(optional)
 * @returns JSON
 * @doc https://developers.themoviedb.org/3/watch-providers/get-available-regions
 ********************/
func (w *watchproviders) GetAvailableRegions(language string) interface{} {
	var token = c_module.GetToken()
	var header = c_module.GetHeader()
	var targetURL string = baseURL + c_module.Route.WATCHPROVIDERS + "regions" + ` api_key=` + token
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
 * 2.GET /watch/providers/movie
 * @description Returns a list of the watch provider (OTT/streaming) data we have available for movies.
 * You can specify a watch_region param if you want to further filter the list by country.
 * @param {string} language(optional)
 * @param {string} watch_region(optional)
 * @returns JSON
 * @doc https://developers.themoviedb.org/3/watch-providers/get-movie-providers
 ********************/
func (w *watchproviders) GetMovieProviders(
	language string,
	watch_region string,
) interface{} {
	var token = c_module.GetToken()
	var header = c_module.GetHeader()
	var targetURL string = baseURL +
		c_module.Route.WATCHPROVIDERS +
		c_module.Route.MOVIE +
		` api_key=` + token
	if language != "" {
		targetURL += `&language=` + language
	}
	if watch_region != "" {
		targetURL += `&watch_region=` + watch_region
	}
	if token == "UnitTest_api_key" {
		return targetURL
	}
	var data, _ = http_methods_golang.GET(targetURL, header)
	return data
}

/********************
 * 3.GET /watch/providers/tv
 * @description Returns a list of the watch provider (OTT/streaming) data we have available for TV series.
 * You can specify a watch_region param if you want to further filter the list by country.
 * @param {string} language(optional)
 * @param {string} watch_region(optional)
 * @returns JSON
 * @doc https://developers.themoviedb.org/3/watch-providers/get-tv-providers
 ********************/
func (w *watchproviders) GetTVProviders(language string, watch_region string) interface{} {
	var token = c_module.GetToken()
	var header = c_module.GetHeader()
	var targetURL string = baseURL +
		c_module.Route.WATCHPROVIDERS +
		c_module.Route.TV +
		` api_key=` + token
	if language != "" {
		targetURL += `&language=` + language
	}
	if watch_region != "" {
		targetURL += `&watch_region=` + watch_region
	}
	if token == "UnitTest_api_key" {
		return targetURL
	}
	var data, _ = http_methods_golang.GET(targetURL, header)
	return data
}

/*
  1.GET GetAvailableRegions
  2.GET GetMovieProviders
  3.GET GetTVProviders
*/
