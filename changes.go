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
 * https  //github.com/wrapTMDB/wrapTMDB-ts
 *
 */
type change struct{}

var Changes change

/********************
 * 1.GET /movie/changes
 * @description Get a list of all of the movie ids that have been changed in the past 24 hours.
 * You can query it for up to 14 days worth of changed IDs at a time with
 * the "start_date" and "end_date" query parameters. 100 items are returned per page.
 * @param {int|string} start_date (optional)
 * @param {int|string} end_date (optional)
 * @param {int} page (optional)
 * @returns JSON
 * @doc https  //developers.themoviedb.org/3/changes/get-movie-change-list
 ********************/
func (c *change) GetMovieChangeList(
	start_date string,
	end_date string,
	page int,
) interface{} {
	var token = c_module.GetToken()
	var header = c_module.GetHeader()
	var targetURL string = baseURL + c_module.Route.MOVIE + "changes" + ` api_key=` + token
	if start_date != "" {
		targetURL += `&start_date=` + start_date
	}
	if end_date != "" {
		targetURL += `&end_date=` + end_date
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
 * 2.GET /tv/changes
 * @description Get a list of all of the TV show ids that have been changed in the past 24 hours.
 * You can query it for up to 14 days worth of changed IDs at a time
 * with the "start_date" and "end_date" query parameters. 100 items are returned per page.
 * @param {int|string} start_date (optional)
 * @param {int|string} end_date (optional)
 * @param {int} page (optional)
 * @returns JSON
 * @doc https  //developers.themoviedb.org/3/changes/get-tv-change-list
 ********************/
func (c *change) GetTVChangeList(
	start_date string,
	end_date string,
	page int,
) interface{} {
	var token = c_module.GetToken()
	var header = c_module.GetHeader()
	var targetURL string = baseURL + c_module.Route.TV + "changes" + ` api_key=` + token
	if start_date != "" {
		targetURL += `&start_date=` + start_date
	}

	if end_date != "" {
		targetURL += `&end_date=` + end_date
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
 * 3.GET /person/changes
 * @description Get a list of all of the person ids that have been changed in the past 24 hours.
 * You can query it for up to 14 days worth of changed IDs at a time
 * with the "start_date" and "end_date" query parameters. 100 items are returned per page.
 * @param {int|string} start_date (optional)
 * @param {int|string} end_date (optional)
 * @param {int} page (optional)
 * @returns JSON
 * @doc https  //developers.themoviedb.org/3/changes/get-person-change-list
 ********************/
func (c *change) GetPersonChangesList(
	start_date string,
	end_date string,
	page int,
) interface{} {
	var token = c_module.GetToken()
	var header = c_module.GetHeader()
	var targetURL string = baseURL + "person" + "/changes" + ` api_key=` + token
	if start_date != "" {
		targetURL += `&start_date=` + start_date
	}

	if end_date != "" {
		targetURL += `&end_date=` + end_date
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

/*
//1.GET Get Movie Change List
//2.GET Get TV Change List
3.GET Get Person Change List
*/
