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

type MapLocationFullInfo struct {
	Type_            string           `json:"type,omitempty"`
	User             *UserFullInfo    `json:"user,omitempty"`
	Photos           []PhotoInfo      `json:"photos,omitempty"`
	Collections      []CollectionInfo `json:"collections,omitempty"`
	TradeCollections []CollectionInfo `json:"tradeCollections,omitempty"`
	LocationMessage  string           `json:"locationMessage,omitempty"`
	ViewEmail        bool             `json:"viewEmail,omitempty"`
	Email            string           `json:"email,omitempty"`
	InGuest          bool             `json:"inGuest,omitempty"`
	InHome           bool             `json:"inHome,omitempty"`
	InClub           bool             `json:"inClub,omitempty"`
	IamPlayer        bool             `json:"iamPlayer,omitempty"`
	IamClub          bool             `json:"iamClub,omitempty"`
	IamShop          bool             `json:"iamShop,omitempty"`
	IamEvent         bool             `json:"iamEvent,omitempty"`
	Id               int64            `json:"id,omitempty"`
	Address          string           `json:"address,omitempty"`
	CoordFirst       float64          `json:"coordFirst,omitempty"`
	CoordSecond      float64          `json:"coordSecond,omitempty"`
	TailFirst        int64            `json:"tailFirst,omitempty"`
	TailSecond       int64            `json:"tailSecond,omitempty"`
	ViewType         int64            `json:"viewType,omitempty"`
}