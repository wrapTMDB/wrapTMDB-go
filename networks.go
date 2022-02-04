package wraptmdb

import "github.com/kwangsing3/http_methods_golang"

/*
 * The MIT License (MIT)
 *
 * Copyright (c) kwangsing3
 *
 * https //github.com/wrapTMDB/wrapTMDB-go
 *
 */

type network struct{}

var Network network

/********************
 * 1.GET /network/{network_id}
 * @description Get the details of a network.
 * @param {number|string} network_id
 * @returns {any} JSON
 * @doc https //developers.themoviedb.org/3/networks/get-network-details
 ********************/
func (n *network) GetDetails(network_id string) interface{} {
	var token = c_module.GetToken()
	var header = c_module.GetHeader()
	var targetURL string = baseURL + c_module.Route.NETWORK + `${network_id}` + `?api_key=${token}`
	if token == "UnitTest_api_key" {
		return targetURL
	}
	var data, _ = http_methods_golang.GET(targetURL, header)
	return data
}

/********************
 * 2.GET /network/{network_id}/alternative_names
 * @description Get the alternative names of a network.
 * @param {number|string} network_id
 * @returns JSON
 * @doc https //developers.themoviedb.org/3/networks/get-network-alternative-names
 ********************/
func (n *network) GetAlternativeNames(network_id string) interface{} {
	var token = c_module.GetToken()
	var header = c_module.GetHeader()
	var targetURL string = baseURL +
		c_module.Route.NETWORK +
		`${network_id}` +
		"/alternative_names" +
		`?api_key=${token}`
	if token == "UnitTest_api_key" {
		return targetURL
	}
	var data, _ = http_methods_golang.GET(targetURL, header)
	return data
}

/********************
 * 3.GET /network/{network_id}/images
 * @description Get the TV network logos by id.
 * @param {number|string} network_id
 * @returns JSON
 * @doc https //developers.themoviedb.org/3/networks/get-network-images
 ********************/
func (n *network) GetImages(network_id string) interface{} {
	var token = c_module.GetToken()
	var header = c_module.GetHeader()
	var targetURL string = baseURL +
		c_module.Route.NETWORK +
		`${network_id}` +
		"/images" +
		`?api_key=${token}`
	if token == "UnitTest_api_key" {
		return targetURL
	}
	var data, _ = http_methods_golang.GET(targetURL, header)
	return data
}

/*
  1.GET Get Details
  2.GET Get Alternative Names
  3.GET Get Images
*/
