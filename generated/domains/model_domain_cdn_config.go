/*
 * Domains
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: v3
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package domains

type DomainCdnConfig struct {
	PortalId                int32  `json:"portalId"`
	Created                 int64  `json:"created"`
	Updated                 int64  `json:"updated"`
	DomainId                int64  `json:"domainId"`
	CertId                  int64  `json:"certId"`
	CertExpirationDate      int64  `json:"certExpirationDate"`
	CdnId                   string `json:"cdnId"`
	CdnGroupId              string `json:"cdnGroupId"`
	SslStatus               string `json:"sslStatus"`
	ValidationMethod        string `json:"validationMethod"`
	HttpBody                string `json:"httpBody"`
	HttpUrlPath             string `json:"httpUrlPath"`
	Cname                   string `json:"cname"`
	CnameTarget             string `json:"cnameTarget"`
	MinimumTlsVersion       string `json:"minimumTlsVersion"`
	AlternateOriginHostname string `json:"alternateOriginHostname"`
	Id                      int64  `json:"id"`
}