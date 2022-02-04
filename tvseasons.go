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

type tvseason struct{}

var TVseason tvseason

/********************
 * 1.GET /tv/{tv_id}/season/{season_number}
 * @description Get the TV season details by id.
 * @param {string} tv_id
 * @param {string} season_number
 * @param {string} language (optional)
 * @param {string} append_to_response (optional)
 * @returns JSON
 * @doc https://developers.themoviedb.org/3/tv-seasons/get-tv-season-details
 ********************/
func (t *tvseason) GetDetails(
	tv_id string,
	season_number string,
	language string,
	append_to_response string,
) interface{} {
	var token = c_module.GetToken()
	var header = c_module.GetHeader()
	var targetURL string = baseURL +
		c_module.Route.TV +
		`${tv_id}/` +
		"season/" +
		`${season_number}` +
		`?api_key=${token}`
	if language != "" {
		targetURL += `&language=${language}`
	}
	if append_to_response != "" {
		targetURL += `&append_to_response=${append_to_response}`
	}
	if token == "UnitTest_api_key" {
		return targetURL
	}
	var data, _ = http_methods_golang.GET(targetURL, header)
	return data
}

/********************
 * 2.GET /tv/{tv_id}/season/{season_number}/account_states
 * @description Returns all of the user ratings for the season"s episodes.
 * @param {string} tv_id
 * @param {string} season_number
 * @param {string} session_id (optional)
 * @param {string} guest_session_id (optional)
 * @param {string} language (optional)
 * @returns JSON
 * @doc https://developers.themoviedb.org/3/tv-seasons/get-tv-season-account-states
 ********************/
func (t *tvseason) GetAccountStates(
	tv_id string,
	season_number string,
	session_id string,
	language string,
	guest_session_id string,
) interface{} {
	var token = c_module.GetToken()
	var header = c_module.GetHeader()
	var targetURL string = baseURL +
		c_module.Route.TV +
		`${tv_id}/` +
		"season/" +
		`${season_number}` +
		"/account_states" +
		`?api_key=${token}`
	if session_id != "" {
		targetURL += `&session_id=${session_id}`
	}

	if language != "" {
		targetURL += `&language=${language}`
	}

	if guest_session_id != "" {
		targetURL += `&guest_session_id=${guest_session_id}`
	}
	if token == "UnitTest_api_key" {
		return targetURL
	}
	var data, _ = http_methods_golang.GET(targetURL, header)
	return data
}

/********************
 * 3.GET /tv/{tv_id}/season/{season_number}/aggregate_credits
 * @description Get the aggregate credits for TV season.
 * @param {string} tv_id
 * @param {string} season_number
 * @param {string} language (optional)
 * @returns JSON
 * @doc https://developers.themoviedb.org/3/tv-seasons/get-tv-season-aggregate-credits
 ********************/
func (t *tvseason) GetAggregateCredits(
	tv_id string,
	season_number string,
	language string,
) interface{} {
	var token = c_module.GetToken()
	var header = c_module.GetHeader()
	var targetURL string = baseURL +
		c_module.Route.TV +
		`${tv_id}/` +
		"season/" +
		`${season_number}` +
		"/aggregate_credits" +
		`?api_key=${token}`
	if language != "" {
		targetURL += `&language=${language}`
	}
	if token == "UnitTest_api_key" {
		return targetURL
	}
	var data, _ = http_methods_golang.GET(targetURL, header)
	return data
}

/********************
 * 4.GET /tv/season/{season_id}/changes
 * @description Get the changes for a TV season. By default only the last 24 hours are returned.
 * @param {string} season_id
 * @param {string} start_date (optional)
 * @param {string} end_date (optional)
 * @param {int} page (optional)
 * @returns JSON
 * @doc https://developers.themoviedb.org/3/tv-seasons/get-tv-season-changes
 ********************/
func (t *tvseason) GetChanges(
	season_id string,
	start_date string,
	end_date string,
	page int,
) interface{} {
	var token = c_module.GetToken()
	var header = c_module.GetHeader()
	var targetURL string = baseURL +
		c_module.Route.TV +
		"season/" +
		`${season_id}` +
		"/changes" +
		`?api_key=${token}`
	if start_date != "" {
		targetURL += `&start_date=${start_date}`
	}

	if end_date != "" {
		targetURL += `&end_date=${end_date}`
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
 * 5.GET /tv/{tv_id}/season/{season_number}/credits
 * @description Get the credits for TV season.
 * @param {string} tv_id
 * @param {string} season_number
 * @param {string} language (optional)
 * @returns JSON
 * @doc https://developers.themoviedb.org/3/tv-seasons/get-tv-season-credits
 ********************/
func (t *tvseason) GetCredits(
	tv_id string,
	season_number string,
	language string,
) interface{} {
	var token = c_module.GetToken()
	var header = c_module.GetHeader()
	var targetURL string = baseURL +
		c_module.Route.TV +
		`${tv_id}/` +
		"season/" +
		`${season_number}` +
		"/credits" +
		`?api_key=${token}`
	if language != "" {
		targetURL += `&language=${language}`
	}
	if token == "UnitTest_api_key" {
		return targetURL
	}
	var data, _ = http_methods_golang.GET(targetURL, header)
	return data
}

/********************
 * 6.GET /tv/{tv_id}/season/{season_number}/external_ids
 * @description Get the TV season details by id.
 * @param {string} tv_id
 * @param {string} season_number
 * @param {string} language (optional)
 * @returns JSON
 * @doc https://developers.themoviedb.org/3/tv-seasons/get-tv-season-details
 ********************/
func (t *tvseason) GetExternalIDs(
	tv_id string,
	season_number string,
	language string,
) interface{} {
	var token = c_module.GetToken()
	var header = c_module.GetHeader()
	var targetURL string = baseURL +
		c_module.Route.TV +
		`${tv_id}/` +
		"season/" +
		`${season_number}` +
		"/external_ids" +
		`?api_key=${token}`
	if language != "" {
		targetURL += `&language=${language}`
	}

	if token == "UnitTest_api_key" {
		return targetURL
	}
	var data, _ = http_methods_golang.GET(targetURL, header)
	return data
}

/********************
 * 7.GET /tv/{tv_id}/season/{season_number}/images
 * @description Get the images that belong to a TV season.
 * @param {string} tv_id
 * @param {string} season_number
 * @param {string} language (optional)
 * @returns JSON
 * @doc https://developers.themoviedb.org/3/tv-seasons/get-tv-season-images
 ********************/
func (t *tvseason) GetImages(
	tv_id string,
	season_number string,
	language string,
) interface{} {
	var token = c_module.GetToken()
	var header = c_module.GetHeader()
	var targetURL string = baseURL +
		c_module.Route.TV +
		`${tv_id}/` +
		"season/" +
		`${season_number}` +
		"/images" +
		`?api_key=${token}`
	if language != "" {
		targetURL += `&language=${language}`
	}

	if token == "UnitTest_api_key" {
		return targetURL
	}
	var data, _ = http_methods_golang.GET(targetURL, header)
	return data
}

/********************
 * 8.GET /tv/{tv_id}/season/{season_number}/translations
 * @description Get the credits for TV season.
 * @param {string} tv_id
 * @param {string} season_number
 * @param {string} language (optional)
 * @returns JSON
 * @doc https://developers.themoviedb.org/3/tv-seasons/get-tv-season-translations
 ********************/
func (t *tvseason) GetTranslations(
	tv_id string,
	season_number string,
	language string,
) interface{} {
	var token = c_module.GetToken()
	var header = c_module.GetHeader()
	var targetURL string = baseURL +
		c_module.Route.TV +
		`${tv_id}/` +
		"season/" +
		`${season_number}` +
		"/translations" +
		`?api_key=${token}`
	if language != "" {
		targetURL += `&language=${language}`
	}

	if token == "UnitTest_api_key" {
		return targetURL
	}
	var data, _ = http_methods_golang.GET(targetURL, header)
	return data
}

/********************
 * 9.GET /tv/{tv_id}/season/{season_number}/videos
 * @description Get the videos that have been added to a TV season.
 * @param {string} tv_id
 * @param {string} season_number
 * @param {string} language (optional)
 * @returns JSON
 * @doc https://developers.themoviedb.org/3/tv-seasons/get-tv-season-videos
 ********************/
func (t *tvseason) GetVideos(
	tv_id string,
	season_number string,
	language string,
) interface{} {
	var token = c_module.GetToken()
	var header = c_module.GetHeader()
	var targetURL string = baseURL +
		c_module.Route.TV +
		`${tv_id}/` +
		"season/" +
		`${season_number}` +
		"/videos" +
		`?api_key=${token}`
	if language != "" {
		targetURL += `&language=${language}`
	}

	if token == "UnitTest_api_key" {
		return targetURL
	}
	var data, _ = http_methods_golang.GET(targetURL, header)
	return data
}

/*
  1.GET Get Details
  2.GET Get Account States
  3.GET Get Aggregate Credits
  4.GET Get Changes
  5.GET Get Credits
  6.GET Get External IDs
  7.GET Get Images
  8.GET Get Translations
  9.GET Get Videos
*/
