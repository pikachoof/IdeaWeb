package models

import "time"

type Session struct {
	ID        string    `json:"id"`
	UserID    uint      `json:"user_id"`
	ExpiresAt time.Time `json:"expires_at"`
	// IPAddress - to check if a session is requested from a different IP, to prevent session stealing
	// UserAgent - same as IPAddress, but prevents using the session from different device
}
