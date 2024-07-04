package game

type Game struct {
	Id                  int     `json:"id"`
	TeseraId            int     `json:"teseraId"`
	BggId               int     `json:"bggId"`
	Title               string  `json:"title"`
	Title2              string  `json:"title2"`
	Title3              string  `json:"title3"`
	Alias               string  `json:"alias"`
	DescriptionShort    string  `json:"descriptionShort"`
	Description         string  `json:"description"`
	PhotoUrl            string  `json:"photoUrl"`
	PhotoUrl2           string  `json:"photoUrl2"`
	Year                int     `json:"year"`
	RatingUser          float64 `json:"ratingUser"`
	N10Rating           float64 `json:"n10Rating"`
	N20Rating           float64 `json:"n20Rating"`
	BggRating           float64 `json:"bggRating"`
	BggGeekRating       float64 `json:"bggGeekRating"`
	BggNumVotes         int     `json:"bggNumVotes"`
	NumVotes            int     `json:"numVotes"`
	PlayersMin          int     `json:"playersMin"`
	PlayersMax          int     `json:"playersMax"`
	PlayersMinRecommend int     `json:"playersMinRecommend"`
	PlayersMaxRecommend int     `json:"playersMaxRecommend"`
	PlayersAgeMin       int     `json:"playersAgeMin"`
	TimeToLearn         int     `json:"timeToLearn"`
	PlaytimeMin         int     `json:"playtimeMin"`
	PlaytimeMax         int     `json:"playtimeMax"`
	CommentsTotal       int     `json:"commentsTotal"`
	CommentsTotalNew    int     `json:"commentsTotalNew"`
	TeseraUrl           string  `json:"teseraUrl"`
	IsAddition          bool    `json:"isAddition"`
}
