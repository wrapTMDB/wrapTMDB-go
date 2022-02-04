package wraptmdb

/*
 * The MIT License (MIT)
 *
 * Copyright (c) kwangsing3
 *
 * https://github.com/wrapTMDB/wrapTMDB-ts
 *
 */
import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/kwangsing3/http_methods_golang"
)

var Movies movie

type movie struct{}

/********************
 * 1.GET /movie/{movie_id}
 * @description Get the primary information about a movie.
 *  {string} movie_id  Movie ID in TMDB
 *  {string} language  Language to request (optional)
 * @doc https://developers.themoviedb.org/3/movies/get-movie-details
 ********************/
func (r *movie) GetDetail(movie_id string, language string) interface{} {
	//624860
	var token = c_module.GetToken()
	var header = c_module.GetHeader()
	var baseURL = c_module.GetURL()
	var targetURL string = baseURL + c_module.Route.MOVIE + movie_id + ` api_key=` + token

	if language != "" {
		targetURL += `&language=` + language
	}

	if token == "UnitTest_api_key" {
		return targetURL
	}
	var data, err = http_methods_golang.GET(targetURL, header)
	if err != nil {
		log.Print(err)
		return ""
	}

	return data
}

/********************
 * 2.GET /movie/{movie_id}/account_states
 * @description Grab the following account states for a session:
	‧Movie rating
	‧If it belongs to your watchlist
	‧If it belongs to your favourite list
 * @param {int|string} movie_id  Movie ID in TMDB
 * @param {string} session_id
 * @param {string} guest_session_id (optional)
 * @returns JSON
 * @doc https://developers.themoviedb.org/3/movies/get-movie-account-states
 ********************/
func (r *movie) GetAccountStates(
	movie_id string,
	session_id string,
	guest_session_id string,
) interface{} {
	var token = c_module.GetToken()
	var header = c_module.GetHeader()
	var targetURL string = baseURL +
		c_module.Route.MOVIE +
		movie_id +
		"/account_states" +
		` api_key=` + token
	if session_id != "" {
		targetURL += `&session_id=` + session_id
	}

	if guest_session_id != "" {
		targetURL += `&guest_session_id=` + guest_session_id
	}

	if token == "UnitTest_api_key" {
		return targetURL
	}
	var data, err = http_methods_golang.GET(targetURL, header)
	if err != nil {
		log.Print(err)
		return ""
	}
	return data
}

/********************
 * 3.GET /movie/{movie_id}/alternative_titles
 * @description Get all of the alternative titles for a movie.
 * @param {int|string} movie_id  Movie ID in TMDB
 * @param {string} country
 * @returns JSON
 * @doc https://developers.themoviedb.org/3/movies/get-movie-alternative-titles
 ********************/
func (r *movie) GetAlternativetitles(movie_id string, country string) interface{} {
	var token = c_module.GetToken()
	var header = c_module.GetHeader()
	var targetURL string = baseURL +
		c_module.Route.MOVIE +
		movie_id +
		"/alternative_titles" +
		` api_key=` + token
	if country != "" {
		targetURL += `&country=` + country
	}

	if token == "UnitTest_api_key" {
		return targetURL
	}
	var data, err = http_methods_golang.GET(targetURL, header)
	if err != nil {
		log.Print(err)
		return ""
	}
	return data
}

/********************
 * 4.GET /movie/{movie_id}/changes
 * @description Get the changes for a movie. By default only the last 24 hours are returned.
 * You can query up to 14 days in a single query by using the start_date and end_date query parameters.
 * @param {int|string} movie_id  Movie ID in TMDB
 * @param {string} start_date (optional)
 * @param {string} end_date (optional)
 * @param {int} page (optional)
 * @returns JSON
 * @doc https://developers.themoviedb.org/3/movies/get-movie-changes
 ********************/
func (r *movie) GetChanges(movie_id string, start_date string, end_date string, page int) interface{} {
	var token = c_module.GetToken()
	var header = c_module.GetHeader()
	var targetURL string = baseURL +
		c_module.Route.MOVIE +
		movie_id +
		"/changes" +
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
	var data, err = http_methods_golang.GET(targetURL, header)
	if err != nil {
		log.Print(err)
		return ""
	}
	return data
}

/********************
 * 5.GET /movie/{movie_id}/credits
 * @description Get the cast and crew for a movie.
 * @param {int|string} movie_id  Movie ID in TMDB
 * @param {string} language(optional)  Language to request
 * @returns JSON
 * @doc https://developers.themoviedb.org/3/movies/get-movie-credits
 ********************/
func (r *movie) GetCredits(movie_id string, language string) interface{} {
	var token = c_module.GetToken()
	var header = c_module.GetHeader()
	var targetURL string = baseURL +
		c_module.Route.MOVIE +
		movie_id +
		"/credits" +
		` api_key=` + token
	if language != "" {
		targetURL += `&language=` + language
	}
	if token == "UnitTest_api_key" {
		return targetURL
	}
	var data, err = http_methods_golang.GET(targetURL, header)
	if err != nil {
		log.Print(err)
		return ""
	}
	return data
}

/********************
 * 6.GET /movie/{movie_id}/external_ids
 * @description Get the external ids for a movie. We currently support the following external sources.
 * @param {int|string} movie_id  Movie ID in TMDB
 * @returns JSON
 * @doc https://developers.themoviedb.org/3/movies/get-movie-external-ids
 ********************/
func (r *movie) GetExternalIDs(movie_id string) interface{} {
	var token = c_module.GetToken()
	var header = c_module.GetHeader()
	var targetURL string = baseURL +
		c_module.Route.MOVIE +
		movie_id +
		"/external_ids" +
		` api_key=` + token
	if token == "UnitTest_api_key" {
		return targetURL
	}
	var data, err = http_methods_golang.GET(targetURL, header)
	if err != nil {
		log.Print(err)
		return ""
	}
	return data
}

/********************
 * 7.GET /movie/{movie_id}/images
 * @description Querying images with a language parameter will filter the results.
 * If you want to include a fallback language (especially useful for backdrops)
 * you can use the include_image_language parameter.
 * This should be a comma seperated value like so include_image_language=en,null.
 * @param {int|string} movie_id  Movie ID in TMDB
 * @param {string} language (optional)
 * @param {string} include_image_language (optional)
 * @doc https://developers.themoviedb.org/3/movies/get-movie-images
 ********************/
func (r *movie) GetImages(movie_id string, language string, include_image_language string) interface{} {
	var token = c_module.GetToken()
	var header = c_module.GetHeader()
	var targetURL string = baseURL +
		c_module.Route.MOVIE +
		movie_id +
		"/images" +
		` api_key=` + token
	if language != "" {
		targetURL += `&language=` + language
	}
	if include_image_language != "" {
		targetURL += `&include_image_language=` + include_image_language
	}
	if token == "UnitTest_api_key" {
		return targetURL
	}
	var data, err = http_methods_golang.GET(targetURL, header)
	if err != nil {
		log.Print(err)
		return ""
	}
	return data
}

/********************
 * 8.GET /movie/{movie_id}/keywords
 * @description Get the keywords that have been added to a movie.
 * @param {int|string} movie_id  Movie ID in TMDB
 * @doc https://developers.themoviedb.org/3/movies/get-movie-keywords
 ********************/
func (r *movie) GetKeywords(movie_id string) interface{} {
	var token = c_module.GetToken()
	var header = c_module.GetHeader()
	var targetURL string = baseURL +
		c_module.Route.MOVIE +
		movie_id +
		"/keywords" +
		` api_key=` + token
	if token == "UnitTest_api_key" {
		return targetURL
	}
	var data, err = http_methods_golang.GET(targetURL, header)
	if err != nil {
		log.Print(err)
		return ""
	}
	return data
}

/********************
 * 9.GET /movie/{movie_id}/lists
 * @description Get a list of lists that this movie belongs to.
 * @param {int|string} movie_id  Movie ID in TMDB
 * @param {string} language(optional)  Language to request
 * @param {int} page (optional)
 * @doc https://developers.themoviedb.org/3/movies/get-movie-lists
 ********************/
func (r *movie) GetLists(movie_id string, language string, page int) interface{} {
	var token = c_module.GetToken()
	var header = c_module.GetHeader()
	var targetURL string = baseURL +
		c_module.Route.MOVIE +
		movie_id +
		"/lists" +
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
	var data, err = http_methods_golang.GET(targetURL, header)
	if err != nil {
		log.Print(err)
		return ""
	}
	return data
}

/********************
 *10.GET /movie/{movie_id}/recommendations
 * @description Get a list of recommended movies for a movie.
 * @param {int|string} movie_id  Movie ID in TMDB
 * @param {string} language(optional)  Language to request
 * @param {int} page (optional)
 * @doc https://developers.themoviedb.org/3/movies/get-movie-recommendations
 ********************/
func (r *movie) GetRecommendations(movie_id string, language string, page int) interface{} {
	var token = c_module.GetToken()
	var header = c_module.GetHeader()
	var targetURL string = baseURL +
		c_module.Route.MOVIE +
		movie_id +
		"/recommendations" +
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
	var data, err = http_methods_golang.GET(targetURL, header)
	if err != nil {
		log.Print(err)
		return ""
	}
	return data
}

/********************
 *11.GET /movie/{movie_id}/release_dates
 * @description Get the release date along with the certification for a movie.
 * Release dates support different types:
	 1.Premiere
	 2.Theatrical (limited)
	 3.Theatrical
	 4.Digital
	 5.Physical
	 6.TV
 * @param {int|string} movie_id  Movie ID in TMDB
 * @doc https://developers.themoviedb.org/3/movies/get-movie-release-dates
 ********************/
func (r *movie) GetReleaseDates(movie_id string) interface{} {
	var token = c_module.GetToken()
	var header = c_module.GetHeader()
	var targetURL string = baseURL +
		c_module.Route.MOVIE +
		movie_id +
		"/release_dates" +
		` api_key=` + token
	if token == "UnitTest_api_key" {
		return targetURL
	}
	var data, err = http_methods_golang.GET(targetURL, header)
	if err != nil {
		log.Print(err)
		return ""
	}
	return data
}

/********************
 *12.GET /movie/{movie_id}/reviews
 * @description Get the user reviews for a movie.
 * @param {int|string} movie_id  Movie ID in TMDB
 * @param {string} language(optional)  Language to request
 * @param {int} page (optional)
 * @doc https://developers.themoviedb.org/3/movies/get-movie-reviews
 ********************/
func (r *movie) GetReviews(movie_id string, language string, page int) interface{} {
	var token = c_module.GetToken()
	var header = c_module.GetHeader()
	var targetURL string = baseURL +
		c_module.Route.MOVIE +
		movie_id +
		"/reviews" +
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
	var data, err = http_methods_golang.GET(targetURL, header)
	if err != nil {
		log.Print(err)
		return ""
	}
	return data
}

/********************
 *13.GET /movie/{movie_id}/similar
 * @description Get a list of similar movies. This is "not" the same as the "Recommendation" system you see on the website.
 * These items are assembled by looking at keywords and genres.
 * @param {int|string} movie_id  Movie ID in TMDB
 * @param {string} language(optional)  Language to request
 * @param {int} page (optional)
 * @doc https://developers.themoviedb.org/3/movies/get-similar-movies
 ********************/
func (r *movie) GetSimilar(movie_id string, language string, page int) interface{} {
	var token = c_module.GetToken()
	var header = c_module.GetHeader()
	var targetURL string = baseURL +
		c_module.Route.MOVIE +
		movie_id +
		"/similar" +
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
	var data, err = http_methods_golang.GET(targetURL, header)
	if err != nil {
		log.Print(err)
		return ""
	}
	return data
}

/********************
 *14.GET /movie/{movie_id}/translations
 * @description Get a list of translations that have been created for a movie.
 * @param {int|string} movie_id  Movie ID in TMDB
 * @doc https://developers.themoviedb.org/3/movies/get-movie-translations
 ********************/
func (r *movie) GetTranslations(movie_id string) interface{} {
	var token = c_module.GetToken()
	var header = c_module.GetHeader()
	var targetURL string = baseURL +
		c_module.Route.MOVIE +
		movie_id +
		"/translations" +
		` api_key=` + token
	if token == "UnitTest_api_key" {
		return targetURL
	}
	var data, err = http_methods_golang.GET(targetURL, header)
	if err != nil {
		log.Print(err)
		return ""
	}
	return data
}

/********************
 *15.GET /movie/{movie_id}/videos
 * @description Get the videos that have been added to a movie.
 * @param {int|string} movie_id  Movie ID in TMDB
 * @param {string} language(optional)  Language to request
 * @doc https://developers.themoviedb.org/3/movies/get-movie-videos
 ********************/
func (r *movie) GetVideos(movie_id string, language string) interface{} {
	var token = c_module.GetToken()
	var header = c_module.GetHeader()
	var targetURL string = baseURL +
		c_module.Route.MOVIE +
		movie_id +
		"/videos" +
		` api_key=` + token
	if language != "" {
		targetURL += `&language=` + language
	}
	if token == "UnitTest_api_key" {
		return targetURL
	}
	var data, err = http_methods_golang.GET(targetURL, header)
	if err != nil {
		log.Print(err)
		return ""
	}
	return data
}

/********************
*16.GET /movie/{movie_id}/watch/providers
* @description Powered by JustWatch, you can query this method to get a list of the availabilities per country by provider.
 This is not going to return full deep links, but rather, it"s just enough information to display what"s available where.
 You can link to the provided TMDB URL to help support TMDB and provide the actual deep links to the content.
* @param {int|string} movie_id  Movie ID in TMDB
* @doc https://developers.themoviedb.org/3/movies/get-movie-watch-providers
********************/
func (r *movie) GetWatchProviders(movie_id string) interface{} {
	var token = c_module.GetToken()
	var header = c_module.GetHeader()
	var targetURL string = baseURL +
		c_module.Route.MOVIE +
		movie_id +
		"/watch/providers" +
		` api_key=` + token
	if token == "UnitTest_api_key" {
		return targetURL
	}
	var data, err = http_methods_golang.GET(targetURL, header)
	if err != nil {
		log.Print(err)
		return ""
	}
	return data
}

/********************
 *17.POST /movie/{movie_id}/rating
 * @description Rate a movie.
 * A valid session or guest session ID is required.
 * @param {int|string} movie_id  Movie ID in TMDB
 * @param {any} query POST query
 * @param {string} guest_session_id (optional)
 * @param {string} session_id (optional)
 * @doc https://developers.themoviedb.org/3/movies/rate-movie
 * @example query {
 *  "value":8.5
 * }
 ********************/
func (r *movie) PostRateMovie(movie_id string, query interface{}, session_id string, guest_session_id string) interface{} {
	var token = c_module.GetToken()
	var targetURL string = baseURL +
		c_module.Route.MOVIE +
		movie_id +
		"/rating" +
		` api_key=` + token
	if session_id != "" {
		targetURL += `&session_id=` + session_id
	}
	if guest_session_id != "" {
		targetURL += `&guest_session_id=` + guest_session_id
	}
	var header = map[string]string{
		"api_key": token,
	}

	if token == "UnitTest_api_key" {
		return targetURL
	}
	conent, _ := json.Marshal(query)
	var data, _ = http_methods_golang.POST(targetURL, header, conent)
	return data
}

/********************
 * 18.DELETE /movie/{movie_id}/rating
 * @description Remove your rating for a movie.
 * A valid session or guest session ID is required
 * @param {int|string} movie_id  Movie ID in TMDB
 * @param {string} guest_session_id (optional)
 * @param {string} session_id (optional)
 * @doc https://developers.themoviedb.org/3/movies/delete-movie-rating
 ********************/
func (r *movie) DeleteRating(movie_id string, session_id string, guest_session_id string) interface{} {
	var token = c_module.GetToken()
	var header = c_module.GetHeader()
	var baseURL = c_module.GetURL()
	var targetURL string = baseURL +
		c_module.Route.MOVIE +
		movie_id +
		"/rating" +
		` api_key=` + token
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
 * 19.GET /movie/latest
 * @description Get the most newly created movie. This is a live response and will continuously change.
 * @param {string} language(optional)  Language to request
 * @doc https://developers.themoviedb.org/3/movies/get-latest-movie
 ********************/
func (r *movie) GetLatest(language string) interface{} {
	var token = c_module.GetToken()
	var header = c_module.GetHeader()
	var baseURL = c_module.GetURL()
	var targetURL string = baseURL + c_module.Route.MOVIE + "latest" + ` api_key=` + token
	if language != "" {
		targetURL += `&language=` + language
	}
	if token == "UnitTest_api_key" {
		return targetURL
	}
	var data, err = http_methods_golang.GET(targetURL, header)
	if err != nil {
		log.Print(err)
		return ""
	}
	return data
}

/********************
 * 20.GET /movie/now_playing
 * @description Get a list of movies in theatres.
 * This is a release type query that looks for all movies that have a release type of 2 or 3 within the specified date range.
 * You can optionally specify a region prameter which will narrow the search to only look for theatrical release dates within the specified country.
 * @param {string} language(optional)  Language to request
 * @param {int} page(optional)
 * @param {string} region(optional)
 * @doc https://developers.themoviedb.org/3/movies/get-latest-movie
 ********************/
func (r *movie) GetNowPlaying(language string, page int, region string) interface{} {
	var token = c_module.GetToken()
	var header = c_module.GetHeader()
	var baseURL = c_module.GetURL()
	var targetURL string = baseURL + c_module.Route.MOVIE + "now_playing" + ` api_key=` + token
	if language != "" {
		targetURL += `&language=` + language
	}
	if page > 0 {
		targetURL += `&page=` + fmt.Sprint(page)
	}
	if region != "" {
		targetURL += `&region=` + region
	}
	if token == "UnitTest_api_key" {
		return targetURL
	}
	var data, err = http_methods_golang.GET(targetURL, header)
	if err != nil {
		log.Print(err)
		return ""
	}
	return data
}

/********************
 * 21.GET /movie/popular
 * @description Get a list of the current popular movies on TMDB. This list updates daily.
 * @param {string} language(optional)  Language to request
 * @param {int} page(optional)
 * @param {string} region(optional)
 * @doc https://developers.themoviedb.org/3/movies/get-popular-movies
 ********************/
func (r *movie) GetPopular(language string, page int, region string) interface{} {
	var token = c_module.GetToken()
	var header = c_module.GetHeader()
	var baseURL = c_module.GetURL()
	var targetURL string = baseURL + c_module.Route.MOVIE + "popular" + ` api_key=` + token
	if language != "" {
		targetURL += `&language=` + language
	}
	if page > 0 {
		targetURL += `&page=` + fmt.Sprint(page)
	}
	if region != "" {
		targetURL += `&region=` + region
	}
	if token == "UnitTest_api_key" {
		return targetURL
	}
	var data, err = http_methods_golang.GET(targetURL, header)
	if err != nil {
		log.Print(err)
		return ""
	}
	return data
}

/********************
 * 22.GET /movie/top_rated
 * @description Get the top rated movies on TMDB.
 * @param {string} language(optional)  Language to request
 * @param {int} page(optional)
 * @param {string} region(optional)
 * @doc https://developers.themoviedb.org/3/movies/get-top-rated-movies
 ********************/
func (r *movie) GetTopRated(language string, page int, region string) interface{} {
	var token = c_module.GetToken()
	var header = c_module.GetHeader()
	var baseURL = c_module.GetURL()
	var targetURL string = baseURL + c_module.Route.MOVIE + "top_rated" + ` api_key=` + token
	if language != "" {
		targetURL += `&language=` + language
	}
	if page > 0 {
		targetURL += `&page=` + fmt.Sprint(page)
	}
	if region != "" {
		targetURL += `&region=` + region
	}
	if token == "UnitTest_api_key" {
		return targetURL
	}
	var data, err = http_methods_golang.GET(targetURL, header)
	if err != nil {
		log.Print(err)
		return ""
	}
	return data
}

/********************
 * 23.GET /movie/upcoming
 * @description Get a list of upcoming movies in theatres.
 * This is a release type query that looks for all movies that have a release type of 2 or 3 within the specified date range.
 * You can optionally specify a region prameter which will narrow the search to only look for theatrical release dates within the specified country.
 * @param {string} language(optional)  Language to request
 * @param {int} page(optional)
 * @param {string} region(optional)
 * @doc https://developers.themoviedb.org/3/movies/get-upcoming
 ********************/
func (r *movie) GetUpcoming(language string, page int, region string) interface{} {
	var token = c_module.GetToken()
	var header = c_module.GetHeader()
	var baseURL = c_module.GetURL()
	var targetURL string = baseURL + c_module.Route.MOVIE + "upcoming" + ` api_key=` + token
	if language != "" {
		targetURL += `&language=` + language
	}
	if page > 0 {
		targetURL += `&page=` + fmt.Sprint(page)
	}
	if region != "" {
		targetURL += `&region=` + region
	}
	if token == "UnitTest_api_key" {
		return targetURL
	}
	var data, err = http_methods_golang.GET(targetURL, header)
	if err != nil {
		log.Print(err)
		return ""
	}
	return data
}

/*
MOVIES
1.GET Get Details
2.GET Get Account States
3.GET Get Alternative Titles
4.GET Get Changes
5.GET Get Credits
6.GET Get External IDs
7.GET Get Images
8.GET Get Keywords
9.GET Get Lists
10.GET Get Recommendations
11.GET Get Release Dates
12.GET Get Reviews
13.GET Get Similar Movies
14.GET Get Translations
15.GET Get Videos
16.GET Get Watch Providers
17.POST Rate Movie
18.DE var E De var e Rating
19.GET Get Latest
20.GET Get Now Playing
21.GET Get Popular
22.GET Get Top Rated
23.GET Get Upcoming
*/
