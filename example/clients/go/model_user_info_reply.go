/*
 *
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version:
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type UserInfoReply struct {
	Name        string   `json:"name,omitempty"`
	Age         int32    `json:"age,omitempty"`
	Birthday    string   `json:"birthday,omitempty"`
	Description string   `json:"description,omitempty"`
	Tag         []string `json:"tag,omitempty"`
}
