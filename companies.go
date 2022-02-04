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
type companies struct{}

var Companies companies

/********************
 * 1.GET /company/{company_id}
 * @description Get a companies details by id.
 * @param { int|string} company_id
 * @returns JSON
 * @doc https //developers.themoviedb.org/3/companies/get-company-details
 ********************/
func (c *companies) GetDetails(company_id string) interface{} {
	var token = c_module.GetToken()
	var header = c_module.GetHeader()
	var targetURL string = baseURL + c_module.Route.COMPANY + `` + company_id + `?api_key=` + token
	if token == "UnitTest_api_key" {
		return targetURL
	}
	var data, _ = http_methods_golang.GET(targetURL, header)
	return data
}

/********************
 * 2.GET /company/{company_id}/alternative_names
 * @description Get the alternative names of a company.
 * @param { int|string} company_id
 * @returns JSON
 * @doc https //developers.themoviedb.org/3/companies/get-company-details
 ********************/
func (c *companies) GetAlternativeNames(company_id string) interface{} {
	var token = c_module.GetToken()
	var header = c_module.GetHeader()
	var targetURL string = baseURL +
		c_module.Route.COMPANY +
		company_id +
		"/alternative_names" +
		`?api_key=` + token
	if token == "UnitTest_api_key" {
		return targetURL
	}
	var data, _ = http_methods_golang.GET(targetURL, header)
	return data
}

/********************
 * 3.GET /company/{company_id}/images
 * @description Get a companies logos by id.
 * @param { int|string} company_id
 * @returns JSON
 * @doc https //developers.themoviedb.org/3/companies/get-company-details
 ********************/
func (c *companies) GetImages(company_id string) interface{} {
	var token = c_module.GetToken()
	var header = c_module.GetHeader()
	var targetURL string = baseURL +
		c_module.Route.COMPANY +
		company_id +
		"/images" +
		`?api_key=` + token
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
