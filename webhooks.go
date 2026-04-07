package datamammoth

import "context"

// WebhooksService handles communication with the webhook-related endpoints of the
// DataMammoth API v2.
type WebhooksService struct {
	client *Client
}

// List returns all webhook subscriptions.
//
// API: GET /webhooks
func (s *WebhooksService) List(ctx context.Context) ([]Webhook, error) {
	var webhooks []Webhook
	_, err := s.client.get(ctx, "/webhooks", &webhooks)
	if err != nil {
		return nil, err
	}
	return webhooks, nil
}

// Get retrieves a webhook subscription by ID.
//
// API: GET /webhooks/:id
func (s *WebhooksService) Get(ctx context.Context, id string) (*Webhook, error) {
	var webhook Webhook
	_, err := s.client.get(ctx, "/webhooks/"+id, &webhook)
	if err != nil {
		return nil, err
	}
	return &webhook, nil
}

// Create creates a new webhook subscription.
//
// API: POST /webhooks
func (s *WebhooksService) Create(ctx context.Context, params *CreateWebhookParams) (*Webhook, error) {
	var webhook Webhook
	_, err := s.client.post(ctx, "/webhooks", params, &webhook)
	if err != nil {
		return nil, err
	}
	return &webhook, nil
}

// Update modifies a webhook subscription.
//
// API: PATCH /webhooks/:id
func (s *WebhooksService) Update(ctx context.Context, id string, params *UpdateWebhookParams) (*Webhook, error) {
	var webhook Webhook
	_, err := s.client.patch(ctx, "/webhooks/"+id, params, &webhook)
	if err != nil {
		return nil, err
	}
	return &webhook, nil
}

// Delete removes a webhook subscription.
//
// API: DELETE /webhooks/:id
func (s *WebhooksService) Delete(ctx context.Context, id string) error {
	_, err := s.client.del(ctx, "/webhooks/"+id, nil)
	return err
}

// ListDeliveries returns a paginated list of webhook deliveries.
//
// API: GET /webhooks/:id/deliveries
func (s *WebhooksService) ListDeliveries(ctx context.Context, id string, opts *ListOptions) ([]WebhookDelivery, *Pagination, error) {
	var deliveries []WebhookDelivery
	pagination, err := s.client.doList(ctx, "/webhooks/"+id+"/deliveries", opts, &deliveries)
	if err != nil {
		return nil, nil, err
	}
	return deliveries, pagination, nil
}

// SendTest sends a test webhook delivery.
//
// API: POST /webhooks/:id/test
func (s *WebhooksService) SendTest(ctx context.Context, id string) (*WebhookTestResult, error) {
	var result WebhookTestResult
	_, err := s.client.post(ctx, "/webhooks/"+id+"/test", nil, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// ListEventTypes returns the available webhook event types.
//
// API: GET /webhooks/events
func (s *WebhooksService) ListEventTypes(ctx context.Context) ([]WebhookEventType, error) {
	var events []WebhookEventType
	_, err := s.client.get(ctx, "/webhooks/events", &events)
	if err != nil {
		return nil, err
	}
	return events, nil
}
