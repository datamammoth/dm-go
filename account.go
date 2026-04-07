package datamammoth

import "context"

// AccountService handles communication with the account-related endpoints of the
// DataMammoth API v2 (profile, api-keys, sessions, 2FA, notifications, activity).
type AccountService struct {
	client *Client
}

// ── Profile ─────────────────────────────────────────────────────────────────

// GetProfile retrieves the current user's profile.
//
// API: GET /me
func (s *AccountService) GetProfile(ctx context.Context) (*UserProfile, error) {
	var profile UserProfile
	_, err := s.client.get(ctx, "/me", &profile)
	if err != nil {
		return nil, err
	}
	return &profile, nil
}

// UpdateProfile updates the current user's profile fields.
//
// API: PATCH /me
func (s *AccountService) UpdateProfile(ctx context.Context, params *UpdateProfileParams) (*UserProfile, error) {
	var profile UserProfile
	_, err := s.client.patch(ctx, "/me", params, &profile)
	if err != nil {
		return nil, err
	}
	return &profile, nil
}

// ChangePassword changes the current user's password.
//
// API: POST /me/change-password
func (s *AccountService) ChangePassword(ctx context.Context, params *ChangePasswordParams) error {
	_, err := s.client.post(ctx, "/me/change-password", params, nil)
	return err
}

// ── API Keys ────────────────────────────────────────────────────────────────

// ListAPIKeys returns the current user's API keys.
//
// API: GET /me/api-keys
func (s *AccountService) ListAPIKeys(ctx context.Context) ([]APIKey, error) {
	var keys []APIKey
	_, err := s.client.get(ctx, "/me/api-keys", &keys)
	if err != nil {
		return nil, err
	}
	return keys, nil
}

// CreateAPIKey creates a new API key. The raw key value is only returned once,
// in the Key field of the response.
//
// API: POST /me/api-keys
func (s *AccountService) CreateAPIKey(ctx context.Context, params *CreateAPIKeyParams) (*APIKey, error) {
	var key APIKey
	_, err := s.client.post(ctx, "/me/api-keys", params, &key)
	if err != nil {
		return nil, err
	}
	return &key, nil
}

// RevokeAPIKey revokes (deactivates) an API key.
//
// API: DELETE /me/api-keys/:id
func (s *AccountService) RevokeAPIKey(ctx context.Context, id string) error {
	_, err := s.client.del(ctx, "/me/api-keys/"+id, nil)
	return err
}

// ── Sessions ────────────────────────────────────────────────────────────────

// ListSessions returns the current user's active sessions.
//
// API: GET /me/sessions
func (s *AccountService) ListSessions(ctx context.Context) ([]Session, error) {
	var sessions []Session
	_, err := s.client.get(ctx, "/me/sessions", &sessions)
	if err != nil {
		return nil, err
	}
	return sessions, nil
}

// RevokeSession revokes (deletes) a session.
//
// API: DELETE /me/sessions/:id
func (s *AccountService) RevokeSession(ctx context.Context, id string) error {
	_, err := s.client.del(ctx, "/me/sessions/"+id, nil)
	return err
}

// ── Two-Factor Authentication ───────────────────────────────────────────────

// TwoFactorStatus retrieves the current 2FA status.
//
// API: POST /me/2fa with action "status"
func (s *AccountService) TwoFactorStatus(ctx context.Context) (*TwoFactorStatus, error) {
	var status TwoFactorStatus
	_, err := s.client.post(ctx, "/me/2fa", &TwoFactorParams{Action: "status"}, &status)
	if err != nil {
		return nil, err
	}
	return &status, nil
}

// EnableTwoFactor enables two-factor authentication.
//
// API: POST /me/2fa with action "enable"
func (s *AccountService) EnableTwoFactor(ctx context.Context) (*TwoFactorStatus, error) {
	var status TwoFactorStatus
	_, err := s.client.post(ctx, "/me/2fa", &TwoFactorParams{Action: "enable"}, &status)
	if err != nil {
		return nil, err
	}
	return &status, nil
}

// DisableTwoFactor disables two-factor authentication.
//
// API: POST /me/2fa with action "disable"
func (s *AccountService) DisableTwoFactor(ctx context.Context) (*TwoFactorStatus, error) {
	var status TwoFactorStatus
	_, err := s.client.post(ctx, "/me/2fa", &TwoFactorParams{Action: "disable"}, &status)
	if err != nil {
		return nil, err
	}
	return &status, nil
}

// ── Notifications ───────────────────────────────────────────────────────────

// ListNotifications returns the current user's notifications (paginated).
//
// API: GET /me/notifications
func (s *AccountService) ListNotifications(ctx context.Context, opts *ListOptions) ([]Notification, *Pagination, error) {
	var notifications []Notification
	pagination, err := s.client.doList(ctx, "/me/notifications", opts, &notifications)
	if err != nil {
		return nil, nil, err
	}
	return notifications, pagination, nil
}

// MarkNotificationsRead marks notifications as read.
//
// API: PATCH /me/notifications
func (s *AccountService) MarkNotificationsRead(ctx context.Context, params *MarkNotificationsReadParams) (*MarkReadResult, error) {
	var result MarkReadResult
	_, err := s.client.patch(ctx, "/me/notifications", params, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// ── Activity ────────────────────────────────────────────────────────────────

// ListActivity returns the current user's activity log (paginated).
//
// API: GET /me/activity
func (s *AccountService) ListActivity(ctx context.Context, opts *ListOptions) ([]Activity, *Pagination, error) {
	var activities []Activity
	pagination, err := s.client.doList(ctx, "/me/activity", opts, &activities)
	if err != nil {
		return nil, nil, err
	}
	return activities, pagination, nil
}

// ListAllActivity returns an iterator over all pages of activity entries.
func (s *AccountService) ListAllActivity(opts *ListOptions) *Iterator[Activity] {
	return newIterator[Activity](s.client, "/me/activity", opts)
}
