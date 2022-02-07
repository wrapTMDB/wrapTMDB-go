package wraptmdb_go

/*
 * The MIT License (MIT)
 *
 * Copyright (c) kwangsing3
 *
 * https://github.com/wrapTMDB/wrapTMDB-go
 *
 */
import (
	"fmt"

	"github.com/kwangsing3/http_methods_golang"
)

type tv struct{}

var TV tv

/********************
* 1.GET /tv/{tv_id}
* @description Get the primary TV show details by id.
* @param { int string} tv_id  Series/TV ID in TMDB
* @param {string} language(optional)  Language to request
* @returns { int} JSON
* @doc https://developers.themoviedb.org/3/tv/get-tv-details
********************/
func (t *tv) GetDetails(tv_id string, language string) interface{} {
	var token = c_module.GetToken()
	var header = c_module.GetHeader()
	var targetURL string = baseURL + c_module.Route.TV + tv_id + `  api_key=` + token

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
 * 2.GET /tv/{tv_id}/account_states
 * @description Grab the following account states for a session:
	  ‧TV show rating
	  ‧If it belongs to your watchlist
	  ‧If it belongs to your favourite list
 * @param { int string} tv_id  TV ID in TMDB
 * @param {string} session_id
 * @param {string} guest_session_id (optional)
 * @returns JSON
 * @doc https://developers.themoviedb.org/3/tv/get-tv-account-states
 ********************/
func (t *tv) GetAccountStates(
	tv_id string,
	session_id string,
	guest_session_id string,
) interface{} {
	var token = c_module.GetToken()
	var header = c_module.GetHeader()
	var targetURL string = baseURL +
		c_module.Route.TV +
		tv_id +
		"/account_states" +
		`  api_key=` + token
	if session_id != "" {
		targetURL += `&session_id=` + session_id
	}

	if guest_session_id != "" {
		targetURL += `&guest_session_id=` + guest_session_id
	}

	if token == "UnitTest_api_key" {
		return targetURL
	}
	var data, _ = http_methods_golang.GET(targetURL, header)
	return data
}

/********************
 * 3.GET /tv/{tv_id}/aggregate_credits
 * @description Get the aggregate credits (cast and crew) that have been added to a TV show.
 * This call differs from the main credits call in that it does not return the newest season but rather,
 * is a view of all the entire cast & crew for all episodes belonging to a TV show.
 * @param { int string} tv_id  TV ID in TMDB
 * @param {string} language
 * @returns JSON
 * @doc https://developers.themoviedb.org/3/tv/get-tv-aggregate-credits
 ********************/
func (t *tv) GetAggregateCredits(tv_id string, language string) interface{} {
	var token = c_module.GetToken()
	var header = c_module.GetHeader()
	var targetURL string = baseURL +
		c_module.Route.TV +
		tv_id +
		"/aggregate_credits" +
		`  api_key=` + token
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
 * 4.GET /tv/{tv_id}/alternative_titles
 * @description Returns all of the alternative titles for a TV show.
 * @param { int string} tv_id  TV ID in TMDB
 * @param {string} language
 * @returns JSON
 * @doc https://developers.themoviedb.org/3/tv/get-tv-alternative-titles
 ********************/
func (t *tv) GetAlternativetitles(tv_id string, language string) interface{} {
	var token = c_module.GetToken()
	var header = c_module.GetHeader()
	var targetURL string = baseURL +
		c_module.Route.TV +
		tv_id +
		"/alternative_titles" +
		`  api_key=` + token
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
 * 5.GET /tv/{tv_id}/changes
 * @description Get the changes for a movie. By default only the last 24 hours are returned.
 * You can query up to 14 days in a single query by using the start_date and end_date query parameters.
 * @param { int string} tv_id  TV ID in TMDB
 * @param { int string} start_date (optional)
 * @param { int string} end_date (optional)
 * @param { int} page (optional)
 * @returns JSON
 * @doc https://developers.themoviedb.org/3/tv/get-tv-changes
 ********************/
func (t *tv) GetChanges(tv_id string, start_date string, end_date string, page int) interface{} {
	var token = c_module.GetToken()
	var header = c_module.GetHeader()
	var targetURL string = baseURL + c_module.Route.TV + tv_id + "/changes" + `  api_key=` + token
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
 * 6.GET /tv/{tv_id}/content_ratings
 * @description Get the list of content ratings (certifications) that have been added to a TV show.
 * @param { int string} tv_id  TV ID in TMDB
 * @param {string} language(optional)  Language to request
 * @returns JSON
 * @doc https://developers.themoviedb.org/3/tv/get-tv-content-ratings
 ********************/
func (t *tv) GetContentRatings(tv_id string, language string) interface{} {
	var token = c_module.GetToken()
	var header = c_module.GetHeader()
	var targetURL string = baseURL +
		c_module.Route.TV +
		tv_id +
		"/content_ratings" +
		`  api_key=` + token
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
 * 7.GET /tv/{tv_id}/credits
 * @description Get the cast and crew for a movie.
 * @param { int string} tv_id  TV ID in TMDB
 * @param {string} language(optional)  Language to request
 * @returns JSON
 * @doc https://developers.themoviedb.org/3/movies/get-movie-credits
 ********************/
func (t *tv) GetCredits(tv_id string, language string) interface{} {
	var token = c_module.GetToken()
	var header = c_module.GetHeader()
	var targetURL string = baseURL + c_module.Route.TV + tv_id + "/credits" + `  api_key=` + token
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
 * 8.GET /tv/{tv_id}/episode_groups
 * @description Get all of the episode groups that have been created for a TV show.
 * With a group ID you can call the GetTVEpisodeGroup method.
 * @param { int string} tv_id  TV ID in TMDB
 * @param {string} language(optional)  Language to request
 * @returns JSON
 * @doc https://developers.themoviedb.org/3/tv/get-tv-episode-groups
 ********************/
func (t *tv) GetEpisodeGroup(
	tv_id string,
	language string,
) interface{} {
	var token = c_module.GetToken()
	var header = c_module.GetHeader()
	var targetURL string = baseURL +
		c_module.Route.TV +
		tv_id +
		"/episode_groups" +
		`  api_key=` + token
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
 * 9.GET /tv/{tv_id}/external_ids
 * @description Get the external ids for a TV show. We currently support the following external sources.
 * @param { int string} tv_id  tv ID in TMDB
 * @param {string} language(optional)  Language to request
 * @returns JSON
 * @doc https://developers.themoviedb.org/3/tv/get-tv-external-ids
 ********************/
func (t *tv) GetExternalIDs(
	tv_id string,
	language string,
) interface{} {
	var token = c_module.GetToken()
	var header = c_module.GetHeader()
	var targetURL string = baseURL +
		c_module.Route.TV +
		tv_id +
		"/external_ids" +
		`  api_key=` + token
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
 * 10.GET /tv/{tv_id}/images
 * @description Get the images that belong to a TV show.
 * @param { int string} tv_id  tv ID in TMDB
 * @param {string} language (optional)
 * @doc https://developers.themoviedb.org/3/tv/get-tv-images
 ********************/
func (t *tv) GetImages(tv_id string, language string) interface{} {
	var token = c_module.GetToken()
	var header = c_module.GetHeader()
	var targetURL string = baseURL + c_module.Route.TV + tv_id + "/images" + `  api_key=` + token
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
 * 11.GET /tv/{tv_id}/keywords
 * @description Get the keywords that have been added to a TV show.
 * @param { int string} tv_id  tv ID in TMDB
 * @doc https://developers.themoviedb.org/3/tv/get-tv-keywords
 ********************/
func (t *tv) GetKeywords(tv_id string) interface{} {
	var token = c_module.GetToken()
	var header = c_module.GetHeader()
	var targetURL string = baseURL + c_module.Route.TV + tv_id + "/keywords" + `  api_key=` + token
	if token == "UnitTest_api_key" {
		return targetURL
	}
	var data, _ = http_methods_golang.GET(targetURL, header)
	return data
}

/********************
 * 12.GET /tv/{tv_id}/recommendations
 * @description Get the list of TV show recommendations for this item.
 * @param { int string} tv_id  tv ID in TMDB
 * @param {string} language (optional)
 * @param { int} page (optional)
 * @doc https://developers.themoviedb.org/3/tv/get-tv-recommendations
 ********************/
func (t *tv) GetRecommendations(
	tv_id string,
	language string,
	page int,
) interface{} {
	var token = c_module.GetToken()
	var header = c_module.GetHeader()
	var targetURL string = baseURL +
		c_module.Route.TV +
		tv_id +
		"/recommendations" +
		`  api_key=` + token
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
 *13.GET /tv/{tv_id}/reviews
 * @description Get the reviews for a TV show.
 * @param { int string} tv_id  tv ID in TMDB
 * @param {string} language(optional)  Language to request
 * @param { int} page (optional)
 * @doc https://developers.themoviedb.org/3/tv/get-tv-reviews
 ********************/
func (t *tv) GetReviews(
	tv_id string,
	language string,
	page int,
) interface{} {
	var token = c_module.GetToken()
	var header = c_module.GetHeader()
	var targetURL string = baseURL + c_module.Route.TV + tv_id + "/reviews" + `  api_key=` + token
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
 *14.GET /tv/{tv_id}/screened_theatrically
 * @description Get a list of seasons or episodes that have been screened in a film festival or theatre.
 * @param { int string} tv_id  tv ID in TMDB
 * @doc https://developers.themoviedb.org/3/tv/get-screened-theatrically
 ********************/
func (t *tv) GetScreenedTheatrically(tv_id string) interface{} {
	var token = c_module.GetToken()
	var header = c_module.GetHeader()
	var targetURL string = baseURL +
		c_module.Route.TV +
		tv_id +
		"/screened_theatrically" +
		`  api_key=` + token

	if token == "UnitTest_api_key" {
		return targetURL
	}
	var data, _ = http_methods_golang.GET(targetURL, header)
	return data
}

/********************
 *15.GET /tv/{tv_id}/similar
 * @description Get a list of similar TV shows. These items are assembled by looking at keywords and genres.
 * @param { int string} tv_id  tv ID in TMDB
 * @param {string} language(optional)  Language to request
 * @param { int} page (optional)
 * @doc https://developers.themoviedb.org/3/tv/get-similar-tv-shows
 ********************/
func (t *tv) GetSimilarTVShows(
	tv_id string,
	language string,
	page int,
) interface{} {
	var token = c_module.GetToken()
	var header = c_module.GetHeader()
	var targetURL string = baseURL + c_module.Route.TV + tv_id + "/similar" + `  api_key=` + token
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
 *16.GET /tv/{tv_id}/translations
 * @description Get a list of the translations that exist for a TV show.
 * @param { int string} tv_id  tv ID in TMDB
 * @doc https://developers.themoviedb.org/3/movies/get-movie-translations
 ********************/
func (t *tv) GetTranslations(tv_id string) interface{} {
	var token = c_module.GetToken()
	var header = c_module.GetHeader()
	var targetURL string = baseURL +
		c_module.Route.TV +
		tv_id +
		"/translations" +
		`  api_key=` + token
	if token == "UnitTest_api_key" {
		return targetURL
	}
	var data, _ = http_methods_golang.GET(targetURL, header)
	return data
}

/********************
 *17.GET /tv/{tv_id}/videos
 * @description Get the videos that have been added to a TV show.
 * @param { int string} tv_id  tv ID in TMDB
 * @param {string} language(optional)  Language to request
 * @doc https://developers.themoviedb.org/3/tv/get-tv-videos
 ********************/
func (t *tv) GetVideos(tv_id string, language string) interface{} {
	var token = c_module.GetToken()
	var header = c_module.GetHeader()
	var targetURL string = baseURL + c_module.Route.TV + tv_id + "/videos" + `  api_key=` + token
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
*18.GET /tv/{tv_id}/watch/providers
* @description Powered by JustWatch, you can query this method to get a list of the availabilities per country by provider.
 This is not going to return full deep links, but rather, it"s just enough information to display what"s available where.
 You can link to the provided TMDB URL to help support TMDB and provide the actual deep links to the content.
* @param { int string} tv_id  tv ID in TMDB
* @doc https://developers.themoviedb.org/3/movies/get-movie-watch-providers
********************/
func (t *tv) GetWatchProviders(tv_id string) interface{} {
	var token = c_module.GetToken()
	var header = c_module.GetHeader()

	var targetURL string = baseURL +
		c_module.Route.TV +
		tv_id +
		"/watch/providers" +
		`  api_key=` + token
	if token == "UnitTest_api_key" {
		return targetURL
	}
	var data, _ = http_methods_golang.GET(targetURL, header)
	return data
}

/********************
 *19.POST /tv/{tv_id}/rating
 * @description Rate a TV show.
 * A valid session or guest session ID is required.
 * @param { int string} tv_id  tv ID in TMDB
 * @param {[]byte} query POST query
 * @param {string} guest_session_id (optional)
 * @param {string} session_id (optional)
 * @doc https://developers.themoviedb.org/3/tv/rate-tv-show
 * @example query {
 *  "value":8.5
 * }
 ********************/
func (t *tv) PostRateTVShow(
	tv_id string,
	query []byte,
	session_id string,
	guest_session_id string,
) interface{} {
	var token = c_module.GetToken()
	var targetURL string = baseURL + c_module.Route.TV + tv_id + "/rating" + `  api_key=` + token
	var header = c_module.GetHeader()

	if session_id != "" {
		targetURL += `&session_id=` + session_id
	}

	if guest_session_id != "" {
		targetURL += `&guest_session_id=` + guest_session_id
	}

	if token == "UnitTest_api_key" {
		return targetURL
	}
	var data, _ = http_methods_golang.POST(targetURL, header, query)
	return data
}

/********************
 * 20.DELETE /tv/{tv_id}/rating
 * @description Remove your rating for a TV show.
 * A valid session or guest session ID is required
 * @param { int string} tv_id  tv ID in TMDB
 * @param {string} guest_session_id (optional)
 * @param {string} session_id
 * @doc https://developers.themoviedb.org/3/tv/delete-tv-show-rating
 ********************/
func (t *tv) DeleteRating(
	tv_id string,
	session_id string,
	guest_session_id string,
) interface{} {
	var token = c_module.GetToken()
	var header = c_module.GetHeader()

	var targetURL string = baseURL + c_module.Route.TV + tv_id + "/rating" + `  api_key=` + token
	if session_id != "" {
		targetURL += `&session_id=` + session_id
	}
	if guest_session_id != "" {
		targetURL += `&guest_session_id=` + guest_session_id
	}

	if token == "UnitTest_api_key" {
		return targetURL
	}
	var data, _ = http_methods_golang.DELETE(targetURL, header, nil)
	return data
	/*no respone*/
}

/********************
 * 21.GET /tv/latest
 * @description Get the most newly created TV show. This is a live response and will continuously change.
 * @param {string} language(optional)  Language to request
 * @doc https://developers.themoviedb.org/3/tv/get-latest-tv
 ********************/
func (t *tv) GetLatest(language string) interface{} {
	var token = c_module.GetToken()
	var header = c_module.GetHeader()
	var targetURL string = baseURL + c_module.Route.TV + "latest" + `  api_key=` + token

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
 * 22.GET /tv/airing_today
 * @description Get a list of TV shows that are airing today. This query is purely day based as we do not currently support airing times.
 * You can specify a timezone to offset the day calculation.
 * Without a specified timezone, this query defaults to EST (Eastern Time UTC-05:00).
 * @param {string} language(optional)  Language to request
 * @param { int} page(optional)
 * @doc https://developers.themoviedb.org/3/tv/get-tv-airing-today
 ********************/
func (t *tv) GetTVAiringToday(language string, page int) interface{} {
	var token = c_module.GetToken()
	var header = c_module.GetHeader()
	var targetURL string = baseURL + c_module.Route.TV + "airing_today" + `  api_key=` + token
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
 * 23.GET /tv/on_the_air
 * @description Get a list of shows that are currently on the air.
 * This query looks for []byte TV show that has an episode with an air date in the next 7 days.
 * @param {string} language(optional)  Language to request
 * @param { int} page(optional)
 * @doc https://developers.themoviedb.org/3/tv/get-tv-on-the-air
 ********************/
func (t *tv) GetTVOnTheAir(language string, page int) interface{} {
	var token = c_module.GetToken()
	var header = c_module.GetHeader()
	var targetURL string = baseURL + c_module.Route.TV + "on_the_air" + `  api_key=` + token
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
 * 24.GET /tv/popular
 * @description Get a list of the current popular TV shows on TMDB. This list updates daily.
 * @param {string} language(optional)  Language to request
 * @param { int} page(optional)
 * @doc https://developers.themoviedb.org/3/tv/get-popular-tv-shows
 ********************/
func (t *tv) GetPopular(language string, page int) interface{} {
	var token = c_module.GetToken()
	var header = c_module.GetHeader()
	var targetURL string = baseURL + c_module.Route.TV + "popular" + `  api_key=` + token
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
 * 25.GET /tv/top_rated
 * @description Get a list of the top rated TV shows on TMDB.
 * @param {string} language(optional)  Language to request
 * @param { int} page(optional)
 * @doc https://developers.themoviedb.org/3/tv/get-top-rated-tv
 ********************/
func (t *tv) GetTopRated(language string, page int) interface{} {
	var token = c_module.GetToken()
	var header = c_module.GetHeader()
	var targetURL string = baseURL + c_module.Route.TV + "top_rated" + `  api_key=` + token
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
 TV
 1.GET Get Details
 2.GET Get Account States
 3.GET Get Aggregate Credits
 4.GET Get Alternative Titles
 5.GET Get Changes
 6.GET Get Content Ratings
 7.GET Get Credits
 8.GET Get Episode Groups
 9.GET Get External IDs
 10.GET Get Images
 11.GET Get Keywords
 12.GET Get Recommendations
 13.GET Get Reviews
 14.GET Get Screened Theatrically
 15.GET Get Similar TV Shows
 16.GET Get Translations
 17.GET Get Videos
 18.GET Get Watch Providers
 19.POST Rate TV Show
 20.DELETE Delete Rating
 21.GET Get Latest
 22.GET Get TV Airing Today
 23.GET Get TV On The Air
 24.GET Get Popular
 25.GET Get Top Rated
*/
