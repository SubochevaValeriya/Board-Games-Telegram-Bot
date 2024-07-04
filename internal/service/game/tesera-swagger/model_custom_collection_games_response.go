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

type CustomCollectionGamesResponse struct {
	Collection *CustomCollectionInfo                `json:"collection,omitempty"`
	List       *PagedListOfCustomCollectionGameInfo `json:"list,omitempty"`
}