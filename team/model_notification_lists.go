/*
 * Teams API
 *
 *  ## Overview An API for creating, retrieving, updating, and deleting teams ## Authentication To authenticate with SignalFx, the following operations require a session  token associated with a SignalFx user that has administrative privileges:<br>   * Create a team - **POST** `/team`   * Update a team - **PUT** `/team/{id}`   * Delete a team - **DELETE** `/team/{id}`   * Update team members - **PUT** `/team/{id}/members`  You can authenticate the following operations with either an org token or a session token. The session token  doesn't need to be associated with a SignalFx user that has administrative privileges:<br>   * Retrieve teams using a query - **GET** `/team`   * Retrieve a team using its ID - **GET** `/team/{id}`
 *
 * API version: 3.1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package team

import (
	"github.com/signalfx/signalfx-go/notification"
)

// Specifies the notification targets to which SignalFx sends alerts when you specify the team as an alert recipient
type NotificationLists struct {
	// List of notification destinations to use for **undefined** alerts category.
	Default []*notification.Notification `json:"default,omitempty"`
	// List of notification destinations to use for **critical** alerts
	Critical []*notification.Notification `json:"critical,omitempty"`
	// List of notification destinations to use for **warning** alerts
	Warning []*notification.Notification `json:"warning,omitempty"`
	// List of notification destinations to use for **major** alerts
	Major []*notification.Notification `json:"major,omitempty"`
	// List of notification destinations to use for **minor** alerts
	Minor []*notification.Notification `json:"minor,omitempty"`
	// List of notification destinations to use for **information** alerts.
	Info []*notification.Notification `json:"info,omitempty"`
}
