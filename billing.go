package datamammoth

import "context"

// BillingService handles communication with the billing-related endpoints of the
// DataMammoth API v2 (invoices, subscriptions, balance, orders, payments, promo).
type BillingService struct {
	client *Client
}

// ── Invoices ────────────────────────────────────────────────────────────────

// ListInvoices returns a paginated list of the current user's invoices.
//
// API: GET /invoices
func (s *BillingService) ListInvoices(ctx context.Context, opts *ListOptions) ([]Invoice, *Pagination, error) {
	var invoices []Invoice
	pagination, err := s.client.doList(ctx, "/invoices", opts, &invoices)
	if err != nil {
		return nil, nil, err
	}
	return invoices, pagination, nil
}

// ListAllInvoices returns an iterator over all pages of invoices.
func (s *BillingService) ListAllInvoices(opts *ListOptions) *Iterator[Invoice] {
	return newIterator[Invoice](s.client, "/invoices", opts)
}

// GetInvoice retrieves a single invoice by ID.
//
// API: GET /invoices/:id
func (s *BillingService) GetInvoice(ctx context.Context, id string) (*Invoice, error) {
	var invoice Invoice
	_, err := s.client.get(ctx, "/invoices/"+id, &invoice)
	if err != nil {
		return nil, err
	}
	return &invoice, nil
}

// PayInvoice initiates payment for an invoice.
//
// API: POST /invoices/:id/pay
func (s *BillingService) PayInvoice(ctx context.Context, id string, params *PayInvoiceParams) (*PaymentResult, error) {
	var result PaymentResult
	_, err := s.client.post(ctx, "/invoices/"+id+"/pay", params, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// ── Subscriptions ───────────────────────────────────────────────────────────

// ListSubscriptions returns a paginated list of the current user's subscriptions.
//
// API: GET /subscriptions
func (s *BillingService) ListSubscriptions(ctx context.Context, opts *ListOptions) ([]Subscription, *Pagination, error) {
	var subs []Subscription
	pagination, err := s.client.doList(ctx, "/subscriptions", opts, &subs)
	if err != nil {
		return nil, nil, err
	}
	return subs, pagination, nil
}

// ListAllSubscriptions returns an iterator over all pages of subscriptions.
func (s *BillingService) ListAllSubscriptions(opts *ListOptions) *Iterator[Subscription] {
	return newIterator[Subscription](s.client, "/subscriptions", opts)
}

// GetSubscription retrieves a single subscription by ID.
//
// API: GET /subscriptions/:id
func (s *BillingService) GetSubscription(ctx context.Context, id string) (*Subscription, error) {
	var sub Subscription
	_, err := s.client.get(ctx, "/subscriptions/"+id, &sub)
	if err != nil {
		return nil, err
	}
	return &sub, nil
}

// UpdateSubscription cancels or reactivates a subscription.
//
// API: PATCH /subscriptions/:id
func (s *BillingService) UpdateSubscription(ctx context.Context, id string, params *UpdateSubscriptionParams) (*Subscription, error) {
	var sub Subscription
	_, err := s.client.patch(ctx, "/subscriptions/"+id, params, &sub)
	if err != nil {
		return nil, err
	}
	return &sub, nil
}

// CancelSubscription is a convenience method to cancel a subscription.
func (s *BillingService) CancelSubscription(ctx context.Context, id string, reason string) (*Subscription, error) {
	return s.UpdateSubscription(ctx, id, &UpdateSubscriptionParams{
		Action:            "cancel",
		Reason:            reason,
		CancelAtPeriodEnd: Bool(true),
	})
}

// ReactivateSubscription is a convenience method to reactivate a subscription.
func (s *BillingService) ReactivateSubscription(ctx context.Context, id string) (*Subscription, error) {
	return s.UpdateSubscription(ctx, id, &UpdateSubscriptionParams{
		Action: "reactivate",
	})
}

// ── Balance ─────────────────────────────────────────────────────────────────

// GetBalance retrieves the current user's prepaid balance.
//
// API: GET /balance
func (s *BillingService) GetBalance(ctx context.Context) (*Balance, error) {
	var balance Balance
	_, err := s.client.get(ctx, "/balance", &balance)
	if err != nil {
		return nil, err
	}
	return &balance, nil
}

// TopUp initiates a balance top-up.
//
// API: POST /balance
func (s *BillingService) TopUp(ctx context.Context, params *TopUpParams) (*TopUpResult, error) {
	var result TopUpResult
	_, err := s.client.post(ctx, "/balance", params, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// ListTransactions returns a paginated list of balance transactions.
//
// API: GET /balance/transactions
func (s *BillingService) ListTransactions(ctx context.Context, opts *ListOptions) ([]Transaction, *Pagination, error) {
	var txns []Transaction
	pagination, err := s.client.doList(ctx, "/balance/transactions", opts, &txns)
	if err != nil {
		return nil, nil, err
	}
	return txns, pagination, nil
}

// ListAllTransactions returns an iterator over all pages of balance transactions.
func (s *BillingService) ListAllTransactions(opts *ListOptions) *Iterator[Transaction] {
	return newIterator[Transaction](s.client, "/balance/transactions", opts)
}

// ── Orders ──────────────────────────────────────────────────────────────────

// ListOrders returns a paginated list of the current user's orders.
//
// API: GET /orders
func (s *BillingService) ListOrders(ctx context.Context, opts *ListOptions) ([]Order, *Pagination, error) {
	var orders []Order
	pagination, err := s.client.doList(ctx, "/orders", opts, &orders)
	if err != nil {
		return nil, nil, err
	}
	return orders, pagination, nil
}

// ListAllOrders returns an iterator over all pages of orders.
func (s *BillingService) ListAllOrders(opts *ListOptions) *Iterator[Order] {
	return newIterator[Order](s.client, "/orders", opts)
}

// GetOrder retrieves a single order by ID.
//
// API: GET /orders/:id
func (s *BillingService) GetOrder(ctx context.Context, id string) (*Order, error) {
	var order Order
	_, err := s.client.get(ctx, "/orders/"+id, &order)
	if err != nil {
		return nil, err
	}
	return &order, nil
}

// CreateOrder places a new order.
//
// API: POST /orders
func (s *BillingService) CreateOrder(ctx context.Context, params *CreateOrderParams) (*Order, error) {
	var order Order
	_, err := s.client.post(ctx, "/orders", params, &order)
	if err != nil {
		return nil, err
	}
	return &order, nil
}

// ── Payment Methods ─────────────────────────────────────────────────────────

// ListPaymentMethods returns the current user's stored payment methods.
//
// API: GET /payment-methods
func (s *BillingService) ListPaymentMethods(ctx context.Context) ([]PaymentMethod, error) {
	var methods []PaymentMethod
	_, err := s.client.get(ctx, "/payment-methods", &methods)
	if err != nil {
		return nil, err
	}
	return methods, nil
}

// AddPaymentMethod adds a new payment method.
//
// API: POST /payment-methods
func (s *BillingService) AddPaymentMethod(ctx context.Context, params *AddPaymentMethodParams) error {
	_, err := s.client.post(ctx, "/payment-methods", params, nil)
	return err
}

// DeletePaymentMethod removes a payment method by ID.
//
// API: DELETE /payment-methods?id=:id
func (s *BillingService) DeletePaymentMethod(ctx context.Context, id string) error {
	_, err := s.client.del(ctx, "/payment-methods?id="+id, nil)
	return err
}

// ── Promo ───────────────────────────────────────────────────────────────────

// ValidatePromo validates a promotional code.
//
// API: GET /promo/validate?code=:code
func (s *BillingService) ValidatePromo(ctx context.Context, code string) (*PromoValidation, error) {
	var result PromoValidation
	_, err := s.client.get(ctx, "/promo/validate?code="+code, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
