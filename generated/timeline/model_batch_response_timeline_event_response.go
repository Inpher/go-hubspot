/*
 * Timeline events
 *
 * This feature allows an app to create and configure custom events that can show up in the timelines of certain CRM objects like contacts, companies, tickets, or deals. You'll find multiple use cases for this API in the sections below.
 *
 * API version: v3
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package timeline

import (
	"time"
)

// The state of the batch event request.
type BatchResponseTimelineEventResponse struct {
	// The status of the batch response. Should always be COMPLETED if processed.
	Status string `json:"status"`
	// Successfully created events.
	Results []TimelineEventResponse `json:"results"`
	// The time the request occurred.
	RequestedAt time.Time `json:"requestedAt,omitempty"`
	// The time the request began processing.
	StartedAt time.Time `json:"startedAt"`
	// The time the request was completed.
	CompletedAt time.Time         `json:"completedAt"`
	Links       map[string]string `json:"links,omitempty"`
}