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

import (
	"time"
)

type PostInfo struct {
	Id                  int64          `json:"id,omitempty"`
	AuthorId            int64          `json:"authorId,omitempty"`
	Title               string         `json:"title,omitempty"`
	Alias               string         `json:"alias,omitempty"`
	Content             string         `json:"content,omitempty"`
	ContentShort        string         `json:"contentShort,omitempty"`
	Rating              float64        `json:"rating,omitempty"`
	CreationDateUtc     time.Time      `json:"creationDateUtc,omitempty"`
	ModificationDateUtc time.Time      `json:"modificationDateUtc,omitempty"`
	PublicationDateUtc  time.Time      `json:"publicationDateUtc,omitempty"`
	IsDraft             bool           `json:"isDraft,omitempty"`
	CommentsTotal       int64          `json:"commentsTotal,omitempty"`
	Author              *UserInfoShort `json:"author,omitempty"`
}