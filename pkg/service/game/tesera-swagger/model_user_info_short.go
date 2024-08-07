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

type UserInfoShort struct {
	TeseraId  int64  `json:"teseraId,omitempty"`
	Id        int64  `json:"id,omitempty"`
	Login     string `json:"login,omitempty"`
	Name      string `json:"name,omitempty"`
	AvatarUrl string `json:"avatarUrl,omitempty"`
	Rating    int64  `json:"rating,omitempty"`
	TeseraUrl string `json:"teseraUrl,omitempty"`
}
