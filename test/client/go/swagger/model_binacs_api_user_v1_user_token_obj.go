/*
 * user.proto
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: version not set
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type BinacsApiUserV1UserTokenObj struct {
	AccessToken string `json:"access_token,omitempty"`
	ExpireTime string `json:"expire_time,omitempty"`
	RefreshToken string `json:"refresh_token,omitempty"`
}