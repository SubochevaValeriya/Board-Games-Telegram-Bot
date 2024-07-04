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

type ReportPlayerInfo struct {
	Id               int64          `json:"id,omitempty"`
	Name             string         `json:"name,omitempty"`
	Type_            string         `json:"type,omitempty"`
	ReferencedUserId int64          `json:"referencedUserId,omitempty"`
	ReferencedUser   *UserInfoShort `json:"referencedUser,omitempty"`
}