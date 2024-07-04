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

type ReportPlayersGroupsResponse struct {
	PlayerGroups *ReportPlayerGroups       `json:"playerGroups,omitempty"`
	PlayerOrder  []ReportPlayersGroupOrder `json:"playerOrder,omitempty"`
}