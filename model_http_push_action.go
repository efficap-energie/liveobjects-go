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
// HttpPushAction struct for HttpPushAction
type HttpPushAction struct {
	// A raw string or a mustache (https://mustache.github.io/mustache.5.html) template describing the webhook body. If empty, the raw event will be used.
	Content string `json:"content,omitempty"`
	// A map of header to be passed in the http POST headers
	Headers map[string][]string `json:"headers,omitempty"`
	// The json path to extract from the considered message (or event), it will be taken as the root datacontext object when combined with a mustache template in content
	JsonPath string `json:"jsonPath,omitempty"`
	// Indicate if a retry policy should be set up in case of a http push delivery failure
	RetryOnFailure bool `json:"retryOnFailure,omitempty"`
	// The url of the target service webhook (only the ports 80, 443, 8080, 8443 and 9243 are allowed).
	WebhookUrl string `json:"webhookUrl"`
}
