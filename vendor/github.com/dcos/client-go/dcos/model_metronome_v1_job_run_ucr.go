/*
 * DC/OS
 *
 * DC/OS API
 *
 * API version: 1.0.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package dcos

type MetronomeV1JobRunUcr struct {
	Image MetronomeV1JobRunUcrImage `json:"image"`
	// Run this docker image in privileged mode
	Privileged bool `json:"privileged,omitempty"`
}
