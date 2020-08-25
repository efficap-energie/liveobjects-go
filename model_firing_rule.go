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
// FiringRule defines if and how the matched data from MatchingRules will fire an event. You can restrict it for instance to trigger maximum 1 event per hour, or only once until you use the FiringGuard API to re-allow it. Associated to aggregationKeys, it will apply the firing type (eg 1 per hour) based on the value of these keys (eg 'streamId').
type FiringRule struct {
	// the list of jsonPath in the Data that will define on which group of data this FiringRule should be set. For instance 'streamId', 'metadata.source', 'value.type'.
	AggregationKeys []string `json:"aggregationKeys,omitempty"`
	// activate or not the FiringRule. Default is false.
	Enabled bool `json:"enabled,omitempty"`
	// define the type of firing mechanism : ONCE, SLEEP, or ALWAYS
	FiringType string `json:"firingType"`
	// id of the FiringRule. Should be null when used for POST.
	Id string `json:"id,omitempty"`
	// the list of MatchingRule ids that will be handeld by this FiringRule.
	MatchingRuleIds []string `json:"matchingRuleIds,omitempty"`
	// name of the FiringRule. Must be unique.
	Name string `json:"name"`
	// sleep duration of the FiringRule. Is defined as a ISO-8601 Period string, restricted to days, hours, minutes and seconds. 1 day will always be equivalent to 24H, regardless of daylight saving time. eg. : 'P1D', 'PT30M'. Must be set only for 'SLEEP' FiringType.
	SleepDuration string `json:"sleepDuration,omitempty"`
}