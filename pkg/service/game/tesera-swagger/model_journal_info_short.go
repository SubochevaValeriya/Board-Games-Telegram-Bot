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

type JournalInfoShort struct {
	TeseraId            int64     `json:"teseraId,omitempty"`
	Alias               string    `json:"alias,omitempty"`
	Title               string    `json:"title,omitempty"`
	Content             string    `json:"content,omitempty"`
	CreationDateUtc     time.Time `json:"creationDateUtc,omitempty"`
	ModificationDateUtc time.Time `json:"modificationDateUtc,omitempty"`
	PublicationDateUtc  time.Time `json:"publicationDateUtc,omitempty"`
	DatePost            time.Time `json:"datePost,omitempty"`
	PhotoUrl            string    `json:"photoUrl,omitempty"`
	Rating              float64   `json:"rating,omitempty"`
	DateGame            time.Time `json:"dateGame,omitempty"`
	Place               string    `json:"place,omitempty"`
	CommentsTotal       int64     `json:"commentsTotal,omitempty"`
	NumVotes            int64     `json:"numVotes,omitempty"`
}
