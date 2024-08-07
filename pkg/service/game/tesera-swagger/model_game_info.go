/*
 * Tesera API
 *
 * API for Tesera
 *
 * API version: v1
 * Contact:
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type GameInfo struct {
	Id                  int64   `json:"id,omitempty"`
	TeseraId            int64   `json:"teseraId,omitempty"`
	BggId               int64   `json:"bggId,omitempty"`
	Title               string  `json:"title,omitempty"`
	Title2              string  `json:"title2,omitempty"`
	Title3              string  `json:"title3,omitempty"`
	Alias               string  `json:"alias,omitempty"`
	DescriptionShort    string  `json:"descriptionShort,omitempty"`
	Description         string  `json:"description,omitempty"`
	PhotoUrl            string  `json:"photoUrl,omitempty"`
	PhotoUrl2           string  `json:"photoUrl2,omitempty"`
	Year                int64   `json:"year,omitempty"`
	RatingUser          float64 `json:"ratingUser,omitempty"`
	N10Rating           float64 `json:"n10Rating,omitempty"`
	N20Rating           float64 `json:"n20Rating,omitempty"`
	RatingMy            float64 `json:"ratingMy,omitempty"`
	BggRating           float64 `json:"bggRating,omitempty"`
	BggGeekRating       float64 `json:"bggGeekRating,omitempty"`
	BggNumVotes         int64   `json:"bggNumVotes,omitempty"`
	NumVotes            int64   `json:"numVotes,omitempty"`
	PlayersMin          int64   `json:"playersMin,omitempty"`
	PlayersMax          int64   `json:"playersMax,omitempty"`
	PlayersMinRecommend int64   `json:"playersMinRecommend,omitempty"`
	PlayersMaxRecommend int64   `json:"playersMaxRecommend,omitempty"`
	PlayersAgeMin       int64   `json:"playersAgeMin,omitempty"`
	TimeToLearn         int64   `json:"timeToLearn,omitempty"`
	PlaytimeMin         int64   `json:"playtimeMin,omitempty"`
	PlaytimeMax         int64   `json:"playtimeMax,omitempty"`
	CommentsTotal       int64   `json:"commentsTotal,omitempty"`
	CommentsTotalNew    int64   `json:"commentsTotalNew,omitempty"`
	TeseraUrl           string  `json:"teseraUrl,omitempty"`
	IsAddition          bool    `json:"isAddition,omitempty"`
}
