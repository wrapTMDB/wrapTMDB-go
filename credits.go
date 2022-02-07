package wraptmdb

import "github.com/kwangsing3/http_methods_golang"

/*
 * The MIT License (MIT)
 *
 * Copyright (c) kwangsing3
 *
 * https://github.com/wrapTMDB/wraptmdb_go
 *
 */
type credits struct{}

var Credits credits

/********************
 * 1.GET /credit/{credit_id}
 * @description Get a movie or TV credit details by id.
 * @param {number|string} credit_id
 * @returns JSON
 * @doc https://developers.themoviedb.org/3/credits/get-credit-details
 ********************/
func (c *credits) GetDetails(credit_id string) interface{} {
	var token = c_module.GetToken()
	var header = c_module.GetHeader()
	var targetURL string = baseURL + c_module.Route.CREDIT + credit_id + `?api_key=` + token
	if token == "UnitTest_api_key" {
		return targetURL
	}
	var data, _ = http_methods_golang.GET(targetURL, header)
	return data
}

/*
1. GET Get Details
*/
