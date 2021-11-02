/*
 * CMS Site Search
 *
 * Use these endpoints for searching content on your HubSpot hosted CMS website(s).
 *
 * API version: v3
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package site_search

// The indexed data in HubSpot
type IndexedData struct {
	// The ID of the document in HubSpot.
	Id string `json:"id"`
	// The type of document. Can be `SITE_PAGE`, `LANDING_PAGE`, `BLOG_POST`, `LISTING_PAGE`, or `KNOWLEDGE_ARTICLE`.
	Type_ string `json:"type"`
	// The indexed fields in HubSpot.
	Fields map[string]SearchHitField `json:"fields"`
}