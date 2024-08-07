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

type LoginModel struct {
	// User's login
	Login string `json:"login,omitempty"`
	// User's password
	Password  string `json:"password,omitempty"`
	IpAddress string `json:"ipAddress,omitempty"`
}
