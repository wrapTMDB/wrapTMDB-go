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

type tvEpisodes struct{}

var TVEpisodes tvEpisodes

/********************
 * 1.GET /tv/{tv_id}/season/{season_int}/episode/{episode_int}
 * @description Get the TV season details by id.
 * @param {string} tv_id
 * @param {string} season_int
 * @param {string} episode_int
 * @param {string} language (optional)
 * @param {string} append_to_response (optional)
 * @returns JSON
 * @doc https //developers.themoviedb.org/3/tv-episodes/get-tv-episode-details
 ********************/
func (t *tvEpisodes) GetDetails(
	tv_id string,
	season_int string,
	episode_int string,
	language string,
	append_to_response string,
) interface{} {
	var token = c_module.GetToken()
	var header = c_module.GetHeader()
	var targetURL string = baseURL +
		c_module.Route.TV +
		`${tv_id}/` +
		"season/" +
		`${season_int}/` +
		c_module.Route.EPISODE +
		`${episode_int}` +
		` api_key=${token}`
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
 * 2.GET /tv/{tv_id}/season/{season_int}/episode/{episode_int}/account_states
 * @description Get your rating for a episode.
 * @param {string} tv_id
 * @param {string} season_int
 * @param {string} episode_int
 * @param {string} guest_session_id (optional)
 * @param {string} session_id (optional)
 * @returns JSON
 * @doc https //developers.themoviedb.org/3/tv-episodes/get-tv-episode-account-states
 ********************/
func (t *tvEpisodes) GetAccountStates(
	tv_id string,
	season_int string,
	episode_int string,
	guest_session_id string,
	session_id string,
) interface{} {
	var token = c_module.GetToken()
	var header = c_module.GetHeader()
	var targetURL string = baseURL +
		c_module.Route.TV +
		`${tv_id}/` +
		"season/" +
		`${season_int}/` +
		c_module.Route.EPISODE +
		`${episode_int}/` +
		"account_states" +
		` api_key=${token}`
	if guest_session_id != "" {
		targetURL += `&guest_session_id=${guest_session_id}`
	}
	if session_id != "" {
		targetURL += `&session_id=${session_id}`
	}
	if token == "UnitTest_api_key" {
		return targetURL
	}
	var data, _ = http_methods_golang.GET(targetURL, header)
	return data
}

/********************
 * 3.GET /tv/episode/{episode_id}/changes
 * @description Get the changes for a TV season. By default only the last 24 hours are returned.
 * @param {string} episode_id
 * @param {string} start_date (optional)
 * @param {string} end_date (optional)
 * @param {int} page (optional)
 * @returns JSON
 * @doc https //developers.themoviedb.org/3/tv-episodes/get-tv-episode-changes
 ********************/
func (t *tvEpisodes) GetChanges(
	episode_id string,
	start_date string,
	end_date string,
	page int,
) interface{} {
	var token = c_module.GetToken()
	var header = c_module.GetHeader()
	var targetURL string = baseURL +
		c_module.Route.TV +
		c_module.Route.EPISODE +
		`${episode_id}` +
		"/changes" +
		` api_key=${token}`
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
 * 4.GET /tv/{tv_id}/season/{season_int}/episode/{episode_int}/credits
 * @description Get the credits (cast, crew and guest stars) for a TV episode.
 * @param {string} tv_id
 * @param {string} season_int
 * @param {string} episode_int
 * @param {string} language (optional)
 * @returns JSON
 * @doc https //developers.themoviedb.org/3/tv-episodes/get-tv-episode-credits
 ********************/
func (t *tvEpisodes) GetCredits(
	tv_id string,
	season_int string,
	episode_int string,
	language string,
) interface{} {
	var token = c_module.GetToken()
	var header = c_module.GetHeader()
	var targetURL string = baseURL +
		c_module.Route.TV +
		`${tv_id}/` +
		"season/" +
		`${season_int}/` +
		c_module.Route.EPISODE +
		`${episode_int}` +
		"/credits" +
		` api_key=${token}`
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
 * 5.GET /tv/{tv_id}/season/{season_int}/episode/{episode_int}/external_ids
 * @description Get the TV season details by id.
 * @param {string} tv_id
 * @param {string} season_int
 * @param {string} episode_int
 * @returns JSON
 * @doc https //developers.themoviedb.org/3/tv-episodes/get-tv-episode-external-ids
 ********************/
func (t *tvEpisodes) GetExternalIDs(
	tv_id string,
	season_int string,
	episode_int string,
) interface{} {
	var token = c_module.GetToken()
	var header = c_module.GetHeader()
	var targetURL string = baseURL +
		c_module.Route.TV +
		`${tv_id}/` +
		"season/" +
		`${season_int}/` +
		c_module.Route.EPISODE +
		`${episode_int}` +
		"/external_ids" +
		` api_key=${token}`
	if token == "UnitTest_api_key" {
		return targetURL
	}
	var data, _ = http_methods_golang.GET(targetURL, header)
	return data
}

/********************
 * 6.GET /tv/{tv_id}/season/{season_int}/episode/{episode_int}/images
 * @description Get the images that belong to a TV episode.
 * @param {string} tv_id
 * @param {string} season_int
 * @param {string} episode_int
 * @returns JSON
 * @doc https //developers.themoviedb.org/3/tv-episodes/get-tv-episode-images
 ********************/
func (t *tvEpisodes) GetImages(
	tv_id string,
	season_int string,
	episode_int string,
) interface{} {
	var token = c_module.GetToken()
	var header = c_module.GetHeader()
	var targetURL string = baseURL +
		c_module.Route.TV +
		`${tv_id}/` +
		"season/" +
		`${season_int}/` +
		c_module.Route.EPISODE +
		`${episode_int}` +
		"/images" +
		` api_key=${token}`

	if token == "UnitTest_api_key" {
		return targetURL
	}
	var data, _ = http_methods_golang.GET(targetURL, header)
	return data
}

/********************
 * 7.GET /tv/{tv_id}/season/{season_int}/episode/{episode_int}/translations
 * @description Get the translation data for an episode.
 * @param {string} tv_id
 * @param {string} season_int
 * @param {string} episode_int
 * @returns JSON
 * @doc https //developers.themoviedb.org/3/tv-episodes/get-tv-episode-translations
 ********************/
func (t *tvEpisodes) GetTranslations(
	tv_id string,
	season_int string,
	episode_int string,
) interface{} {
	var token = c_module.GetToken()
	var header = c_module.GetHeader()
	var targetURL string = baseURL +
		c_module.Route.TV +
		`${tv_id}/` +
		"season/" +
		`${season_int}/` +
		c_module.Route.EPISODE +
		`${episode_int}` +
		"/translations" +
		` api_key=${token}`

	if token == "UnitTest_api_key" {
		return targetURL
	}
	var data, _ = http_methods_golang.GET(targetURL, header)
	return data
}

/********************
 * 8.POST /tv/{tv_id}/season/{season_int}/episode/{episode_int}/rating
 * @description Rate a TV episode.
 * @param {string} tv_id
 * @param {string} season_int
 * @param {string} episode_int
 * @param {any} query
 * @param {string} session_id (optional)
 * @param {string} guest_session_id (optional)
 * @returns JSON
 * @example query{
 *    "value"  8.5
 * }
 * @doc https //developers.themoviedb.org/3/tv-episodes/rate-tv-episode
 ********************/
func (t *tvEpisodes) PostRateTVEpisode(
	tv_id string,
	season_int string,
	episode_int string,
	query []byte,
	session_id string,
	guest_session_id string,
) interface{} {
	var token = c_module.GetToken()
	var header = c_module.GetHeader()
	var targetURL string = baseURL +
		c_module.Route.TV +
		`${tv_id}/` +
		"season/" +
		`${season_int}/` +
		c_module.Route.EPISODE +
		`${episode_int}` +
		"/rating" +
		` api_key=${token}`
	if guest_session_id != "" {
		targetURL += `&guest_session_id=${guest_session_id}`
	}
	if session_id != "" {
		targetURL += `&session_id=${session_id}`
	}
	if token == "UnitTest_api_key" {
		return targetURL
	}
	var data, _ = http_methods_golang.POST(targetURL, header, query)
	return data
}

/********************
 * 9.DELETE /tv/{tv_id}/season/{season_int}/episode/{episode_int}/rating
 * @description Remove your rating for a TV episode.
 * @param {string} tv_id
 * @param {string} season_int
 * @param {string} episode_int
 * @param {string} session_id (optional)
 * @param {string} guest_session_id (optional)
 * @returns JSON
 * @doc https //developers.themoviedb.org/3/tv-episodes/delete-tv-episode-rating
 ********************/
func (t *tvEpisodes) DeleteDeleteRating(
	tv_id string,
	season_int string,
	episode_int string,
	session_id string,
	guest_session_id string,
) interface{} {
	var token = c_module.GetToken()
	var header = c_module.GetHeader()
	var targetURL string = baseURL +
		c_module.Route.TV +
		`${tv_id}/` +
		"season/" +
		`${season_int}/` +
		c_module.Route.EPISODE +
		`${episode_int}` +
		"/rating" +
		` api_key=${token}`
	if guest_session_id != "" {
		targetURL += `&guest_session_id=${guest_session_id}`
	}

	if session_id != "" {
		targetURL += `&session_id=${session_id}`
	}

	if token == "UnitTest_api_key" {
		return targetURL
	}
	var data, _ = http_methods_golang.DELETE(targetURL, header, nil)
	return data
}

/********************
 * 10.GET /tv/{tv_id}/season/{season_int}/episode/{episode_int}/videos
 * @description Get the videos that have been added to a TV episode.
 * @param {string} tv_id
 * @param {string} season_int
 * @param {string} episode_int
 * @param {string} language (optional)
 * @returns JSON
 * @doc https //developers.themoviedb.org/3/tv-episodes/get-tv-episode-videos
 ********************/
func (t *tvEpisodes) GetVideos(
	tv_id string,
	season_int string,
	episode_int string,
	language string,
) interface{} {
	var token = c_module.GetToken()
	var header = c_module.GetHeader()
	var targetURL string = baseURL +
		c_module.Route.TV +
		`${tv_id}/` +
		"season/" +
		`${season_int}/` +
		c_module.Route.EPISODE +
		`${episode_int}` +
		"/videos" +
		` api_key=${token}`
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
  3.GET Get Changes
  4.GET Get Credits
  5.GET Get External IDs
  6.GET Get Images
  7.GET Get Translations
  8.POST Rate TV Episode
  9.DELETE Delete Rating
  10.GET Get Videos
*/
