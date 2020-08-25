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
// PayloadDescriptionTestResult struct for PayloadDescriptionTestResult
type PayloadDescriptionTestResult struct {
	DecodingFailureReason string `json:"decodingFailureReason,omitempty"`
	DecodingResult map[string]interface{} `json:"decodingResult,omitempty"`
	DescriptionValid bool `json:"descriptionValid,omitempty"`
	ParsingOk bool `json:"parsingOk,omitempty"`
	TemplatingOK bool `json:"templatingOK,omitempty"`
}