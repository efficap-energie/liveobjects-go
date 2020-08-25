/*
 * Live Objects REST API Guide v2.12.2
 *
 * API description for Live Objects service
 *
 * API version: 2.12.2
 * Contact: liveobjects.support@orange.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package liveobjects
// AuditLogMessage a log message that described a functionnal event
type AuditLogMessage struct {
	// category of the log message.
	Category string `json:"category"`
	// log message's specific content
	Content map[string]interface{} `json:"content,omitempty"`
	// date of storage of the log message. ISO 8601 date time.
	Created string `json:"created"`
	// high level description of the log message
	Description string `json:"description"`
	// detailed description of log the message
	DetailedDescription string `json:"detailedDescription,omitempty"`
	// level of the log message.
	Level string `json:"level"`
	Source Source `json:"source,omitempty"`
	// subcategory of the log message.
	Subcategory string `json:"subcategory"`
	// timestamp of the log message. ISO 8601 date time.
	Timestamp string `json:"timestamp"`
	// specific type of the log message.
	Type string `json:"type"`
}
