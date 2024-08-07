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

type PagedListOfCustomCollectionGameInfo struct {
	Offset             int32                      `json:"offset,omitempty"`
	Limit              int32                      `json:"limit,omitempty"`
	Total              int32                      `json:"total,omitempty"`
	TotalOther         int32                      `json:"totalOther,omitempty"`
	TotalPages         int32                      `json:"totalPages,omitempty"`
	Data               []CustomCollectionGameInfo `json:"data,omitempty"`
	HasPreviousPage    bool                       `json:"hasPreviousPage,omitempty"`
	HasNextPage        bool                       `json:"hasNextPage,omitempty"`
	NextPageNumber     int32                      `json:"nextPageNumber,omitempty"`
	PreviousPageNumber int32                      `json:"previousPageNumber,omitempty"`
}
