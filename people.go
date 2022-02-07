package wraptmdb

import (
	"fmt"

	"github.com/kwangsing3/http_methods_golang"
)

/*
 * The MIT License (MIT)
 *
 * Copyright (c) kwangsing3
 *
 * https://github.com/wrapTMDB/wraptmdb-go
 *
 */

type people struct{}

var People people

/********************
 * 1.GET /person/{person_id}
 * @description Get the primary person details by id.
 * @param {int|string} person_id
 * @param {string} language
 * @returns {any} JSON
 * @doc https://developers.themoviedb.org/3/people/get-person-details
 ********************/
func (p *people) GetDetails(
	person_id string,
	language string,
	append_to_response string,
) interface{} {
	var token = c_module.GetToken()
	var header = c_module.GetHeader()
	var targetURL string = baseURL + c_module.Route.PEOPLE + person_id + ` api_key=` + token
	if language != "" {
		targetURL += `&language=` + language
	}
	if append_to_response != "" {
		targetURL += `&append_to_response=` + append_to_response
	}
	if token == "UnitTest_api_key" {
		return targetURL
	}
	var data, _ = http_methods_golang.GET(targetURL, header)
	return data
}

/********************
 * 2.GET /person/{person_id}/changes
 * @description Get the changes for a person. By default only the last 24 hours are returned.
 * @param {int|string} person_id
 * @param {string} end_date (optional)
 * @param {int} page (optional)
 * @param {string} start_date (optional)
 * @returns {any} JSON
 * @doc https://developers.themoviedb.org/3/people/get-person-changes
 ********************/
func (p *people) GetChanges(
	person_id string,
	start_date string,
	end_date string,
	page int,
) interface{} {
	var token = c_module.GetToken()
	var header = c_module.GetHeader()
	var targetURL string = baseURL +
		c_module.Route.PEOPLE +
		person_id + `changes` +
		` api_key=` + token
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
 * 3.GET /person/{person_id}/movie_credits
 * @description Get the movie credits for a person.
 * @param {int|string} person_id
 * @param {string} language(optional)
 * @returns {any} JSON
 * @doc https://developers.themoviedb.org/3/people/get-person-movie-credits
 ********************/
func (p *people) GetMovieCredits(
	person_id string,
	language string,
) interface{} {
	var token = c_module.GetToken()
	var header = c_module.GetHeader()
	var targetURL string = baseURL +
		c_module.Route.PEOPLE +
		person_id + `movie_credits` +
		` api_key=` + token
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
 * 4.GET /person/{person_id}/tv_credits
 * @description Get the TV show credits for a person.
 * @param {int|string} person_id
 * @param {string} language(optional)
 * @returns {any} JSON
 * @doc https://developers.themoviedb.org/3/people/get-person-tv-credits
 ********************/
func (p *people) GetTVCredits(
	person_id string,
	language string,
) interface{} {
	var token = c_module.GetToken()
	var header = c_module.GetHeader()
	var targetURL string = baseURL +
		c_module.Route.PEOPLE +
		person_id + `tv_credits` +
		` api_key=` + token
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
 * 5.GET /person/{person_id}/combined_credits
 * @description Get the movie and TV credits together in a single response.
 * @param {int|string} person_id
 * @param {string} language(optional)
 * @returns {any} JSON
 * @doc https://developers.themoviedb.org/3/people/get-person-combined-credits
 ********************/
func (p *people) GetCombinedCredits(
	person_id string,
	language string,
) interface{} {
	var token = c_module.GetToken()
	var header = c_module.GetHeader()
	var targetURL string = baseURL +
		c_module.Route.PEOPLE +
		person_id + `combined_credits` +
		` api_key=` + token
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
 * 6.GET /person/{person_id}/external_ids
 * @description Get the external ids for a person. We currently support the following external sources.
 * @param {int|string} person_id
 * @param {string} language(optional)
 * @returns {any} JSON
 * @doc https://developers.themoviedb.org/3/people/get-person-external-ids
 ********************/
func (p *people) GetExternalIDs(
	person_id string,
	language string,
) interface{} {
	var token = c_module.GetToken()
	var header = c_module.GetHeader()
	var targetURL string = baseURL +
		c_module.Route.PEOPLE +
		person_id + `external_ids` +
		` api_key=` + token
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
 * 7.GET /person/{person_id}/images
 * @description Get the images for a person.
 * @param {int|string} person_id
 * @returns {any} JSON
 * @doc https://developers.themoviedb.org/3/people/get-person-images
 ********************/
func (p *people) GetImages(person_id string) interface{} {
	var token = c_module.GetToken()
	var header = c_module.GetHeader()
	var targetURL string = baseURL +
		c_module.Route.PEOPLE +
		person_id + `images` +
		` api_key=` + token
	if token == "UnitTest_api_key" {
		return targetURL
	}
	var data, _ = http_methods_golang.GET(targetURL, header)
	return data
}

/********************
 * 8.GET /person/{person_id}/tagged_images
 * @description Get the primary person details by id.
 * @param {int|string} person_id
 * @param {string} language(optional)
 * @param {int} page(optional)
 * @returns {any} JSON
 * @doc https://developers.themoviedb.org/3/people/get-tagged-images
 ********************/
func (p *people) GetTaggedImages(
	person_id string,
	language string,
	page int,
) interface{} {
	var token = c_module.GetToken()
	var header = c_module.GetHeader()
	var targetURL string = baseURL +
		c_module.Route.PEOPLE +
		person_id + `tagged_images` +
		` api_key=` + token
	if language != "" {
		targetURL += `&language=` + language
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
 * 9.GET /person/{person_id}/translations
 * @description Get a list of translations that have been created for a person.
 * @param {int|string} person_id
 * @param {string} language(optional)
 * @returns {any} JSON
 * @doc https://developers.themoviedb.org/3/people/get-person-translations
 ********************/
func (p *people) GetTranslations(
	person_id string,
	language string,
) interface{} {
	var token = c_module.GetToken()
	var header = c_module.GetHeader()
	var targetURL string = baseURL +
		c_module.Route.PEOPLE +
		person_id + `translations` +
		` api_key=` + token
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
 * 10.GET /person/latest
 * @description Get the most newly created person. This is a live response and will continuously change.
 * @param {string} language(optional)
 * @returns {any} JSON
 * @doc https://developers.themoviedb.org/3/people/get-latest-person
 ********************/
func (p *people) GetLatest(language string) interface{} {
	var token = c_module.GetToken()
	var header = c_module.GetHeader()
	var targetURL string = baseURL + c_module.Route.PEOPLE + "latest" + ` api_key=` + token
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
 * 11.GET /person/{person_id}
 * @description Get the primary person details by id.
 * @param {string} language(optional)
 * @param {int} page(optional)
 * @returns {any} JSON
 * @doc https://developers.themoviedb.org/3/people/get-person-details
 ********************/
func (p *people) GetPopular(language string, page int) interface{} {
	var token = c_module.GetToken()
	var header = c_module.GetHeader()
	var targetURL string = baseURL + c_module.Route.PEOPLE + "popular" + ` api_key=` + token
	if language != "" {
		targetURL += `&language=` + language
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
  1.GET Get Details
  2.GET Get Changes
  3.GET Get Movie Credits
  4.GET Get TV Credits
  5.GET Get Combined Credits
  6.GET Get External IDs
  7.GET Get Images
  8.GET Get Tagged Images
  9.GET Get Translations
  10.GET Get Latest
  11.GET Get Popular
*/
