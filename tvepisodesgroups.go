package wraptmdb_go

import "github.com/kwangsing3/http_methods_golang"

/*
 * The MIT License (MIT)
 *
 * Copyright (c) kwangsing3
 *
 * https://github.com/wrapTMDB/wrapTMDB-go
 *
 */
type tvepisodesgroups struct{}

var TVepisodesgroups tvepisodesgroups

/********************
 * 1.GET /tv/episode_group/{id}
 * @description Get the details of a TV episode group. Groups support 7 different types which are enumerated as the following:
 * 1.Original air date
 * 2.Absolute
 * 3.DVD
 * 4.Digital
 * 5.Story arc
 * 6.Production
 * 7.TV
 * @param {number|string} id
 * @returns {any} JSON
 * @doc https://developers.themoviedb.org/3/tv-episode-groups/get-tv-episode-group-details
 ********************/
func (t *tvepisodesgroups) GetDetails(id string, language string) interface{} {
	var token = c_module.GetToken()
	var header = c_module.GetHeader()
	var targetURL string = baseURL +
		c_module.Route.TV +
		c_module.Route.EPISODEGROUPS +
		`` + id +
		`?api_key=` + token
	if language != "" {
		targetURL += `&language=` + language
	}
	if token == "UnitTest_api_key" {
		return targetURL
	}
	var data, _ = http_methods_golang.GET(targetURL, header)
	return data
}

/*
  1.GET GetDetails
*/
