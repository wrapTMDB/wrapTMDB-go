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

type reviews struct{}

var Reviews reviews

/********************
 * 1.GET /review/{review_id}
 * @description Retrieve the details of a movie or TV show review.
 * @param {string} review_id
 * @returns {any} JSON
 * @doc https://developers.themoviedb.org/3/reviews/get-review-details
 ********************/
func (r *reviews) GetDetails(review_id string) interface{} {
	var token = c_module.GetToken()
	var header = c_module.GetHeader()
	var targetURL string = baseURL + c_module.Route.REVIEWS + review_id + `?api_key=` + token
	if token == "UnitTest_api_key" {
		return targetURL
	}
	var data, _ = http_methods_golang.GET(targetURL, header)
	return data
}

/*
1. GET Get Reviews
*/
