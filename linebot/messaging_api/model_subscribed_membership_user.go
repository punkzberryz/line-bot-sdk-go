/**
 * LINE Messaging API
 * This document describes LINE Messaging API.
 *
 * The version of the OpenAPI document: 0.0.1
 *
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

/**
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

//go:generate python3 ../../generate-code.py
package messaging_api

// SubscribedMembershipUser
// Object containing user membership subscription information.

type SubscribedMembershipUser struct {

	/**
	 * The user&#39;s member number in the membership plan. (Required)
	 */
	MembershipNo int32 `json:"membershipNo"`

	/**
	 * UNIX timestamp at which the user subscribed to the membership. (Required)
	 */
	JoinedTime int32 `json:"joinedTime"`

	/**
	 * Next payment date for membership plan. - Format: yyyy-MM-dd (e.g. 2024-02-08) - Timezone: UTC+9  (Required)
	 */
	NextBillingDate string `json:"nextBillingDate"`

	/**
	 * The period of time in months that the user has been subscribed to a membership plan. If a user previously canceled and then re-subscribed to the same membership plan, only the period after the re-subscription will be counted. (Required)
	 */
	TotalSubscriptionMonths int32 `json:"totalSubscriptionMonths"`
}
