package model



// CreateNotificationRequest represents the request payload for creating a notification.
type CreateNotificationRequest struct {
    UserID     string `json:"user_id"`      // ID of the user to whom the notification is addressed.
    Message    string `json:"message"`     // The notification message.
    RelationID string `json:"relation_id"` // ID of the related entity (e.g., tender_id, bid_id).
    Type       string `json:"type"`        // Type of the notification (e.g., "tender", "bid").
}

// NotificationResponse represents the response after creating a notification.
type NotificationResponse struct {
    Message string `json:"message"` // Response message.
}
