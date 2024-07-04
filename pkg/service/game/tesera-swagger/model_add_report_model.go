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

type AddReportModel struct {
	// Game Id
	GameId int64 `json:"gameId,omitempty"`
	// Report title
	Title string `json:"title,omitempty"`
	// Report descr
	Description string `json:"description,omitempty"`
	// Place Id, ref
	PlaceId    int64  `json:"placeId,omitempty"`
	NumPlayers int64  `json:"numPlayers,omitempty"`
	GameDate   string `json:"gameDate,omitempty"`
	GameTime   int64  `json:"gameTime,omitempty"`
	GameRating int64  `json:"gameRating,omitempty"`
	IsGamePage bool   `json:"isGamePage,omitempty"`
	AccessType int64  `json:"accessType,omitempty"`
	ResultType int64  `json:"resultType,omitempty"`
	IsDraft    bool   `json:"isDraft,omitempty"`
}