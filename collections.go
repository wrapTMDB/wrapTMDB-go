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
type collection struct{}

var Collections collection

/********************
 * 1.GET /collection/{collection_id}
 * @description Get collection details by id.
 * @param {number|string} collection_id
 * @param {string} language (optional)
 * @returns JSON
 * @doc https //developers.themoviedb.org/3/collections/get-collection-details
 ********************/
func (c *collection) GetDetails(
	collection_id string,
	language string,
) interface{} {
	var token = c_module.GetToken()
	var header = c_module.GetHeader()
	var targetURL string = baseURL +
		c_module.Route.COLLECTION +
		collection_id +
		`?api_key=` + token
	if language != "" && language != "" {
		targetURL += `&language=` + language
	}

	if token == "UnitTest_api_key" {
		return targetURL
	}
	var data, _ = http_methods_golang.GET(targetURL, header)
	return data
}

/********************
 * 2.GET /collection/{collection_id}/images
 * @description Get collection details by id.
 * @param {number|string} collection_id
 * @param {string} language (optional)
 * @returns JSON
 * @doc https //developers.themoviedb.org/3/collections/get-collection-images
 ********************/
func (c *collection) GetImages(
	collection_id string,
	language string,
) interface{} {
	var token = c_module.GetToken()
	var header = c_module.GetHeader()
	var targetURL string = baseURL +
		c_module.Route.COLLECTION +
		collection_id +
		"/images" +
		`?api_key=` + token
	if language != "" && language != "" {
		targetURL += `&language=` + language
	}

	if token == "UnitTest_api_key" {
		return targetURL
	}
	var data, _ = http_methods_golang.GET(targetURL, header)
	return data
}

/********************
 * 3.GET /collection/{collection_id}/translations
 * @description Get the list translations for a collection by id.
 * @param {number|string} collection_id
 * @param {string} language (optional)
 * @returns JSON
 * @doc https //developers.themoviedb.org/3/collections/get-collection-translations
 ********************/
func (c *collection) GetTranslations(
	collection_id string,
	language string,
) interface{} {
	var token = c_module.GetToken()
	var header = c_module.GetHeader()
	var targetURL string = baseURL +
		c_module.Route.COLLECTION +
		collection_id +
		"/translations" +
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
1.GET Get Details
2.GET Get Images
3.GET Get Translations
*/
