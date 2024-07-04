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

type AddPlayerToReportModel struct {
	// Player Id
	CustomPlayerId int64 `json:"customPlayerId,omitempty"`
	// User Id
	UserPlayerId int64 `json:"userPlayerId,omitempty"`
	IsWinner     bool  `json:"isWinner,omitempty"`
	Points       int64 `json:"points,omitempty"`
	TeamId       int64 `json:"teamId,omitempty"`
	SideId       int64 `json:"sideId,omitempty"`
}