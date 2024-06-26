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

// SubscribedMembershipPlan
// Object containing information about the membership plan.

type SubscribedMembershipPlan struct {

	/**
	 * Membership plan ID. (Required)
	 */
	MembershipId int32 `json:"membershipId"`

	/**
	 * Membership plan name. (Required)
	 */
	Title string `json:"title"`

	/**
	 * Membership plan description. (Required)
	 */
	Description string `json:"description"`

	/**
	 * List of membership plan perks. (Required)
	 */
	Benefits []string `json:"benefits"`

	/**
	 * Monthly fee for membership plan. (e.g. 1500.00) (Required)
	 */
	Price float64 `json:"price"`

	/**
	 * The currency of membership.price. (Required)
	 */
	Currency SubscribedMembershipPlanCURRENCY `json:"currency"`
}

// SubscribedMembershipPlanCURRENCY type
/* The currency of membership.price. */
type SubscribedMembershipPlanCURRENCY string

// SubscribedMembershipPlanCURRENCY constants
const (
	SubscribedMembershipPlanCURRENCY_JPY SubscribedMembershipPlanCURRENCY = "JPY"

	SubscribedMembershipPlanCURRENCY_TWD SubscribedMembershipPlanCURRENCY = "TWD"

	SubscribedMembershipPlanCURRENCY_THB SubscribedMembershipPlanCURRENCY = "THB"
)
