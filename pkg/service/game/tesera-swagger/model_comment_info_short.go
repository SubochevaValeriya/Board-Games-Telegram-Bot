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

type CommentInfoShort struct {
	TeseraId        int64          `json:"teseraId,omitempty"`
	Id              int64          `json:"id,omitempty"`
	ParentId        int64          `json:"parentId,omitempty"`
	Title           string         `json:"title,omitempty"`
	Content         string         `json:"content,omitempty"`
	Rating          float64        `json:"rating,omitempty"`
	CreationDateUtc time.Time      `json:"creationDateUtc,omitempty"`
	Author          *UserInfoShort `json:"author,omitempty"`
	Object          *ObjectInfo    `json:"object,omitempty"`
	AnswersTotal    int64          `json:"answersTotal,omitempty"`
}
