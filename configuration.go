package wraptmdb

import "github.com/kwangsing3/http_methods_golang"

/*
 * The MIT License (MIT)
 *
 * Copyright (c) kwangsing3
 *
 * https://github.com/wrapTMDB/wrapTMDB-go
 *
 */
type configuration struct{}

var Configuration configuration

/********************
 * 1.GET /configuration
 * @description Get the system wide configuration information.
 * Some elements of the API require some knowledge of this configuration data.
 * The purpose of this is to try and keep the actual API responses as light as possible.
 * It is recommended you cache this data within your application and check for updates every few days.
 * @returns JSON
 * @example To build an image URL, you will need 3 pieces of data. The base_url, size and file_path.
 * Simply combine them all and you will have a fully qualified URL. Here’s an example URL
 * https://image.tmdb.org/t/p/w500/8uO0gUM8aNqYLs1OsTBQiXu0fEv.jpg
 * @doc https://developers.themoviedb.org/3/configuration/get-api-configuration
 ********************/
func (c *configuration) GetAPIConfiguration() interface{} {
	var token = c_module.GetToken()
	var header = c_module.GetHeader()
	var targetURL string = baseURL + c_module.Route.CONFIGURATION + `?api_key=` + token

	if token == "UnitTest_api_key" {
		return targetURL
	}
	var data, _ = http_methods_golang.GET(targetURL, header)
	return data
}

/********************
 * 2.GET /configuration/countries
 * @description Get the list of countries (ISO 3166-1 tags) used throughout TMDB.
 * @returns JSON
 * @doc https://developers.themoviedb.org/3/configuration/get-countries
 ********************/
func (c *configuration) GetCountries() interface{} {
	var token = c_module.GetToken()
	var header = c_module.GetHeader()
	var targetURL string = baseURL +
		c_module.Route.CONFIGURATION +
		"countries" +
		`?api_key=` + token
	if token == "UnitTest_api_key" {
		return targetURL
	}
	var data, _ = http_methods_golang.GET(targetURL, header)
	return data
}

/********************
 * 3.GET /configuration/jobs
 * @description Get a list of the jobs and departments we use on TMDB.
 * @returns JSON
 * @doc https://developers.themoviedb.org/3/configuration/get-jobs
 ********************/
func (c *configuration) GetJobs() interface{} {
	var token = c_module.GetToken()
	var header = c_module.GetHeader()
	var targetURL string = baseURL + c_module.Route.CONFIGURATION + "jobs" + `?api_key=` + token
	if token == "UnitTest_api_key" {
		return targetURL
	}
	var data, _ = http_methods_golang.GET(targetURL, header)
	return data
}

/********************
 * 4.GET /configuration/languages
 * @description Get the list of languages (ISO 639-1 tags) used throughout TMDB.
 * @returns JSON
 * @doc https://developers.themoviedb.org/3/configuration/get-languages
 ********************/
func (c *configuration) GetLanguages() interface{} {
	var token = c_module.GetToken()
	var header = c_module.GetHeader()
	var targetURL string = baseURL +
		c_module.Route.CONFIGURATION +
		"languages" +
		`?api_key=` + token
	if token == "UnitTest_api_key" {
		return targetURL
	}
	var data, _ = http_methods_golang.GET(targetURL, header)
	return data
}

/********************
 * 5.GET /configuration/primary_translations
 * @description Get a list of the officially supported translations on TMDB.
 * @returns JSON
 * @doc https://developers.themoviedb.org/3/configuration/get-primary-translations
 ********************/
func (c *configuration) GetPrimaryTranslations() interface{} {
	var token = c_module.GetToken()
	var header = c_module.GetHeader()
	var targetURL string = baseURL +
		c_module.Route.CONFIGURATION +
		"primary_translations" +
		`?api_key=` + token
	if token == "UnitTest_api_key" {
		return targetURL
	}
	var data, _ = http_methods_golang.GET(targetURL, header)
	return data
}

/********************
 * 6.GET /configuration/timezones
 * @description Get the list of timezones used throughout TMDB.
 * @returns JSON
 * @doc https://developers.themoviedb.org/3/configuration/get-timezones
 ********************/
func (c *configuration) GetTimezones() interface{} {
	var token = c_module.GetToken()
	var header = c_module.GetHeader()
	var targetURL string = baseURL +
		c_module.Route.CONFIGURATION +
		"timezones" +
		`?api_key=` + token
	if token == "UnitTest_api_key" {
		return targetURL
	}
	var data, _ = http_methods_golang.GET(targetURL, header)
	return data
}

/*
1.GET Get API Configuration
2.GET Get Countries
3.GET Get Jobs
4.GET Get Languages
5.GET Get Primary Translations
6.GET Get Timezones
*/
