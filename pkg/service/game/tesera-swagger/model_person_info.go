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

type PersonInfo struct {
	Id                  int64        `json:"id,omitempty"`
	TeseraId            int64        `json:"teseraId,omitempty"`
	AuthorId            int64        `json:"authorId,omitempty"`
	FullName            string       `json:"fullName,omitempty"`
	FullNameSecond      string       `json:"fullNameSecond,omitempty"`
	FullNameThird       string       `json:"fullNameThird,omitempty"`
	FirstName           string       `json:"firstName,omitempty"`
	LastName            string       `json:"lastName,omitempty"`
	FirstNameSecond     string       `json:"firstNameSecond,omitempty"`
	LastNameSecond      string       `json:"lastNameSecond,omitempty"`
	Gender              string       `json:"gender,omitempty"`
	Country             *CountryInfo `json:"country,omitempty"`
	City                *CityInfo    `json:"city,omitempty"`
	CountryId           int64        `json:"countryId,omitempty"`
	CityId              int64        `json:"cityId,omitempty"`
	Alias               string       `json:"alias,omitempty"`
	Description         string       `json:"description,omitempty"`
	DescriptionShort    string       `json:"descriptionShort,omitempty"`
	PhotoUrl            string       `json:"photoUrl,omitempty"`
	CreationDateUtc     time.Time    `json:"creationDateUtc,omitempty"`
	ModificationDateUtc time.Time    `json:"modificationDateUtc,omitempty"`
	RatingUser          float64      `json:"ratingUser,omitempty"`
	Rating              float64      `json:"rating,omitempty"`
	RatingTesera        float64      `json:"ratingTesera,omitempty"`
}
