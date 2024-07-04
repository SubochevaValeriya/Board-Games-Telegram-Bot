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

type GameInfoResponse struct {
	Game               *GameInfo                  `json:"game,omitempty"`
	Collections        *GameInCollectionsResponse `json:"collections,omitempty"`
	OwnersTotal        int64                      `json:"ownersTotal,omitempty"`
	SellTotal          int64                      `json:"sellTotal,omitempty"`
	BuyTotal           int64                      `json:"buyTotal,omitempty"`
	SellTotalAll       int64                      `json:"sellTotalAll,omitempty"`
	BuyTotalAll        int64                      `json:"buyTotalAll,omitempty"`
	CommentsTotal      int64                      `json:"commentsTotal,omitempty"`
	ReportsTotal       int64                      `json:"reportsTotal,omitempty"`
	PhotosTotal        int64                      `json:"photosTotal,omitempty"`
	FilesTotal         int64                      `json:"filesTotal,omitempty"`
	LinksTotal         int64                      `json:"linksTotal,omitempty"`
	VideoExternalTotal int64                      `json:"videoExternalTotal,omitempty"`
	VideoInternalTotal int64                      `json:"videoInternalTotal,omitempty"`
	Photos             []PhotoInfo                `json:"photos,omitempty"`
	Files              []FileInfo                 `json:"files,omitempty"`
	Links              []LinkInfo                 `json:"links,omitempty"`
	Similars           []GameInfoShort            `json:"similars,omitempty"`
	Relateds           []GameInfoShort            `json:"relateds,omitempty"`
	News               []NewsInfoShort            `json:"news,omitempty"`
}