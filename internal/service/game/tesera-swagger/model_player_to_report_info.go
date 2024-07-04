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

type PlayerToReportInfo struct {
	Id             int64             `json:"id,omitempty"`
	ReportId       int64             `json:"reportId,omitempty"`
	CustomPlayerId int64             `json:"customPlayerId,omitempty"`
	UserPlayerId   int64             `json:"userPlayerId,omitempty"`
	TeamId         int64             `json:"teamId,omitempty"`
	IsWinner       bool              `json:"isWinner,omitempty"`
	Points         int64             `json:"points,omitempty"`
	SideId         int64             `json:"sideId,omitempty"`
	UsageCount     int64             `json:"usageCount,omitempty"`
	UserPlayer     *UserInfoShort    `json:"userPlayer,omitempty"`
	CustomPlayer   *ReportPlayerInfo `json:"customPlayer,omitempty"`
}