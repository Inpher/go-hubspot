/*
 * OAuthService
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: v1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package oauth

type AccessTokenInfoResponse struct {
	Token                     string   `json:"token"`
	User                      string   `json:"user,omitempty"`
	HubDomain                 string   `json:"hub_domain,omitempty"`
	Scopes                    []string `json:"scopes"`
	ScopeToScopeGroupPks      []int32  `json:"scope_to_scope_group_pks"`
	TrialScopes               []string `json:"trial_scopes"`
	TrialScopeToScopeGroupPks []int32  `json:"trial_scope_to_scope_group_pks"`
	HubId                     int32    `json:"hub_id"`
	AppId                     int32    `json:"app_id"`
	ExpiresIn                 int32    `json:"expires_in"`
	UserId                    int32    `json:"user_id"`
	TokenType                 string   `json:"token_type"`
}