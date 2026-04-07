// Package datamammoth provides a Go client for the DataMammoth API v2.
package datamammoth

import "encoding/json"

// ── Helpers ─────────────────────────────────────────────────────────────────

// String returns a pointer to the given string value.
func String(v string) *string { return &v }

// Int returns a pointer to the given int value.
func Int(v int) *int { return &v }

// Bool returns a pointer to the given bool value.
func Bool(v bool) *bool { return &v }

// Float64 returns a pointer to the given float64 value.
func Float64(v float64) *float64 { return &v }

// ── Pagination ──────────────────────────────────────────────────────────────

// ListOptions specifies common pagination, sorting, search, and filter
// parameters for list endpoints.
type ListOptions struct {
	Page    int               `json:"page,omitempty"`
	PerPage int               `json:"per_page,omitempty"`
	Sort    string            `json:"sort,omitempty"`
	Search  string            `json:"search,omitempty"`
	Filter  map[string]string `json:"filter,omitempty"`
}

// Pagination contains the pagination metadata returned by list endpoints.
type Pagination struct {
	Page        int  `json:"page"`
	PerPage     int  `json:"per_page"`
	Total       int  `json:"total"`
	TotalPages  int  `json:"total_pages"`
	HasNext     bool `json:"has_next"`
	HasPrevious bool `json:"has_previous"`
}

// ── Response Envelope ───────────────────────────────────────────────────────

// Meta contains metadata included in every API response.
type Meta struct {
	RequestID  string      `json:"request_id"`
	Timestamp  string      `json:"timestamp"`
	Version    string      `json:"version"`
	Pagination *Pagination `json:"pagination,omitempty"`
}

// APIError represents a single error object in the API error response.
type APIError struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Field   string `json:"field,omitempty"`
	Docs    string `json:"docs,omitempty"`
}

// Links represents HATEOAS links in the API response.
type Links map[string]string

// ── Server ──────────────────────────────────────────────────────────────────

// ServerSpecs describes the hardware specifications of a server.
type ServerSpecs struct {
	CPU       *int    `json:"cpu,omitempty"`
	RAMGB     *int    `json:"ram_gb,omitempty"`
	DiskGB    *int    `json:"disk_gb,omitempty"`
	DiskType  *string `json:"disk_type,omitempty"`
	Bandwidth *string `json:"bandwidth,omitempty"`
}

// Server represents a provisioned server.
type Server struct {
	ID            string          `json:"id"`
	Hostname      *string         `json:"hostname"`
	Label         *string         `json:"label"`
	Status        string          `json:"status"`
	IPAddress     *string         `json:"ip_address"`
	IPv6Address   *string         `json:"ipv6_address"`
	Region        *string         `json:"region"`
	OSImage       *string         `json:"os_image"`
	Plan          *string         `json:"plan"`
	Specs         *ServerSpecs    `json:"specs,omitempty"`
	ProvisionedAt *string         `json:"provisioned_at"`
	CreatedAt     string          `json:"created_at"`
	UpdatedAt     string          `json:"updated_at"`
	Links         Links           `json:"_links,omitempty"`
	LiveStatus    json.RawMessage `json:"live_status,omitempty"`
	LiveUsage     json.RawMessage `json:"live_usage,omitempty"`
}

// CreateServerParams contains the parameters for creating a server.
type CreateServerParams struct {
	ProductID string   `json:"product_id"`
	ZoneID    string   `json:"zone_id,omitempty"`
	Region    string   `json:"region,omitempty"`
	ImageID   string   `json:"image_id"`
	Hostname  string   `json:"hostname,omitempty"`
	Label     string   `json:"label,omitempty"`
	Password  string   `json:"password,omitempty"`
	SSHKeyIDs []string `json:"ssh_key_ids,omitempty"`
}

// UpdateServerParams contains the parameters for updating a server.
type UpdateServerParams struct {
	Hostname *string `json:"hostname,omitempty"`
	Label    *string `json:"label,omitempty"`
}

// RebuildParams contains the parameters for rebuilding a server.
type RebuildParams struct {
	ImageID      string `json:"image_id"`
	RootPassword string `json:"root_password,omitempty"`
	DefaultUser  string `json:"default_user,omitempty"`
}

// RescueParams contains the parameters for entering rescue mode.
type RescueParams struct {
	RootPassword string `json:"root_password"`
}

// ActionResult represents the result of a server action.
type ActionResult struct {
	Action   string `json:"action"`
	ServerID string `json:"server_id"`
	Status   string `json:"status"`
	Message  string `json:"message"`
}

// Snapshot represents a server snapshot.
type Snapshot struct {
	ID          string  `json:"id"`
	Name        string  `json:"name"`
	Description *string `json:"description"`
	SizeGB      *int    `json:"size_gb"`
	Status      string  `json:"status"`
	CreatedAt   string  `json:"created_at"`
	Links       Links   `json:"_links,omitempty"`
}

// CreateSnapshotParams contains the parameters for creating a snapshot.
type CreateSnapshotParams struct {
	Name        string `json:"name"`
	Description string `json:"description,omitempty"`
}

// SnapshotActionResult represents the result of a snapshot action.
type SnapshotActionResult struct {
	Action     string `json:"action,omitempty"`
	SnapshotID string `json:"snapshot_id"`
	ServerID   string `json:"server_id,omitempty"`
	Status     string `json:"status,omitempty"`
	Message    string `json:"message,omitempty"`
	Deleted    bool   `json:"deleted,omitempty"`
}

// MetricsResult represents server metrics data.
type MetricsResult struct {
	Source  string          `json:"source"`
	Period  string          `json:"period"`
	Count   int             `json:"count,omitempty"`
	Metrics json.RawMessage `json:"metrics,omitempty"`
	Points  json.RawMessage `json:"points,omitempty"`
	Summary json.RawMessage `json:"summary,omitempty"`
	Status  string          `json:"status,omitempty"`
}

// MetricsOptions contains the parameters for fetching server metrics.
type MetricsOptions struct {
	Period string `json:"period,omitempty"` // lastday, lastweek, lastmonth, lastyear
	Source string `json:"source,omitempty"` // provider, agent, db
	Limit  int    `json:"limit,omitempty"`
}

// Event represents a server audit/event entry.
type Event struct {
	Action    string          `json:"action"`
	Timestamp string          `json:"timestamp"`
	ChangedBy string          `json:"changed_by"`
	RequestID string          `json:"request_id"`
	Changes   json.RawMessage `json:"changes"`
	Category  *string         `json:"category"`
}

// ConsoleAccess represents VNC/SSH console access credentials.
type ConsoleAccess struct {
	Type      string  `json:"type"`
	URL       string  `json:"url"`
	Password  *string `json:"password"`
	ExpiresAt *string `json:"expires_at"`
	ServerID  string  `json:"server_id"`
}

// FirewallRule represents a single firewall rule.
type FirewallRule struct {
	Direction   string `json:"direction"`
	Protocol    string `json:"protocol"`
	Port        string `json:"port,omitempty"`
	SourceIP    string `json:"source_ip,omitempty"`
	DestIP      string `json:"dest_ip,omitempty"`
	Action      string `json:"action"`
	Description string `json:"description,omitempty"`
}

// FirewallConfig represents the firewall configuration for a server.
type FirewallConfig struct {
	ID        *string        `json:"id"`
	ServerID  string         `json:"server_id"`
	Rules     []FirewallRule `json:"rules"`
	Status    string         `json:"status"`
	UpdatedAt *string        `json:"updated_at"`
}

// UpdateFirewallParams contains the parameters for updating firewall rules.
type UpdateFirewallParams struct {
	Rules  []FirewallRule `json:"rules"`
	Status string         `json:"status,omitempty"`
}

// ── Product ─────────────────────────────────────────────────────────────────

// ProductPricing describes the price tiers for a product.
type ProductPricing struct {
	Monthly   *float64 `json:"monthly"`
	Quarterly *float64 `json:"quarterly"`
	Annual    *float64 `json:"annual"`
	OneTime   *float64 `json:"one_time"`
	SetupFee  float64  `json:"setup_fee"`
	Currency  string   `json:"currency"`
}

// Product represents a product in the catalog.
type Product struct {
	ID               string          `json:"id"`
	Name             string          `json:"name"`
	Slug             string          `json:"slug"`
	Description      string          `json:"description"`
	ShortDescription string          `json:"short_description"`
	Type             string          `json:"type"`
	Category         string          `json:"category"`
	Subcategory      *string         `json:"subcategory"`
	Status           string          `json:"status"`
	Pricing          *ProductPricing `json:"pricing,omitempty"`
	Specs            *ServerSpecs    `json:"specs,omitempty"`
	Features         []string        `json:"features,omitempty"`
	IsFeatured       bool            `json:"is_featured"`
	Badge            *string         `json:"badge"`
	BadgeColor       *string         `json:"badge_color,omitempty"`
	PromoLabel       *string         `json:"promo_label,omitempty"`
	PromoPercent     *float64        `json:"promo_percent,omitempty"`
	PromoEndsAt      *string         `json:"promo_ends_at,omitempty"`
	MetaTitle        *string         `json:"meta_title,omitempty"`
	MetaDescription  *string         `json:"meta_description,omitempty"`
	CreatedAt        string          `json:"created_at"`
	UpdatedAt        string          `json:"updated_at"`
	Links            Links           `json:"_links,omitempty"`
}

// ProductOptionChoice represents a choice within a product option.
type ProductOptionChoice struct {
	ID              string  `json:"id"`
	Label           string  `json:"label"`
	Value           string  `json:"value"`
	PriceAdjustment float64 `json:"price_adjustment"`
	IsDefault       bool    `json:"is_default"`
}

// ProductOption represents a configurable option for a product.
type ProductOption struct {
	ID        string                `json:"id"`
	Name      string                `json:"name"`
	Label     string                `json:"label"`
	Type      string                `json:"type"`
	Required  bool                  `json:"required"`
	SortOrder int                   `json:"sort_order"`
	Choices   []ProductOptionChoice `json:"choices"`
}

// ProductAddon represents an addon for a product.
type ProductAddon struct {
	ID           string  `json:"id"`
	Name         string  `json:"name"`
	Description  string  `json:"description"`
	Price        float64 `json:"price"`
	BillingCycle string  `json:"billing_cycle"`
	IsRequired   bool    `json:"is_required"`
}

// ProductPricingDetail contains base and retail pricing for a product.
type ProductPricingDetail struct {
	ProductID     string          `json:"product_id"`
	BasePricing   *ProductPricing `json:"base_pricing"`
	RetailPricing *ProductPricing `json:"retail_pricing"`
	TenantID      *string         `json:"tenant_id"`
}

// Category represents a product category.
type Category struct {
	Name         string `json:"name"`
	Slug         string `json:"slug"`
	ProductCount int    `json:"product_count"`
}

// CreateProductParams contains the parameters for creating a product.
type CreateProductParams struct {
	Name             string                      `json:"name"`
	Slug             string                      `json:"slug"`
	Description      string                      `json:"description,omitempty"`
	ShortDescription string                      `json:"short_description,omitempty"`
	Type             string                      `json:"type,omitempty"`
	Category         string                      `json:"category,omitempty"`
	Subcategory      *string                     `json:"subcategory,omitempty"`
	Status           string                      `json:"status,omitempty"`
	Pricing          *ProductPricing             `json:"pricing,omitempty"`
	Specs            *ServerSpecs                `json:"specs,omitempty"`
	IsFeatured       bool                        `json:"is_featured,omitempty"`
	IsHidden         bool                        `json:"is_hidden,omitempty"`
	Badge            *string                     `json:"badge,omitempty"`
	SortOrder        int                         `json:"sort_order,omitempty"`
	Features         []string                    `json:"features,omitempty"`
	Options          []CreateProductOptionParams `json:"options,omitempty"`
	Addons           []CreateProductAddonParams  `json:"addons,omitempty"`
}

// CreateProductOptionParams contains option params for product creation.
type CreateProductOptionParams struct {
	Name     string                            `json:"name"`
	Label    string                            `json:"label"`
	Type     string                            `json:"type"`
	Required bool                              `json:"required,omitempty"`
	Choices  []CreateProductOptionChoiceParams `json:"choices,omitempty"`
}

// CreateProductOptionChoiceParams contains choice params for product option creation.
type CreateProductOptionChoiceParams struct {
	Label           string  `json:"label"`
	Value           string  `json:"value"`
	PriceAdjustment float64 `json:"price_adjustment,omitempty"`
	IsDefault       bool    `json:"is_default,omitempty"`
}

// CreateProductAddonParams contains addon params for product creation.
type CreateProductAddonParams struct {
	Name         string  `json:"name"`
	Description  string  `json:"description,omitempty"`
	Price        float64 `json:"price"`
	BillingCycle string  `json:"billing_cycle,omitempty"`
	IsRequired   bool    `json:"is_required,omitempty"`
}

// UpdateProductParams contains the parameters for updating a product.
type UpdateProductParams struct {
	Name             *string         `json:"name,omitempty"`
	Slug             *string         `json:"slug,omitempty"`
	Description      *string         `json:"description,omitempty"`
	ShortDescription *string         `json:"short_description,omitempty"`
	Type             *string         `json:"type,omitempty"`
	Category         *string         `json:"category,omitempty"`
	Subcategory      *string         `json:"subcategory,omitempty"`
	Status           *string         `json:"status,omitempty"`
	Pricing          *ProductPricing `json:"pricing,omitempty"`
	Specs            *ServerSpecs    `json:"specs,omitempty"`
	IsFeatured       *bool           `json:"is_featured,omitempty"`
	IsHidden         *bool           `json:"is_hidden,omitempty"`
	Badge            *string         `json:"badge,omitempty"`
	SortOrder        *int            `json:"sort_order,omitempty"`
	Features         []string        `json:"features,omitempty"`
}

// ── Invoice ─────────────────────────────────────────────────────────────────

// InvoiceItem represents a line item on an invoice.
type InvoiceItem struct {
	ID          string  `json:"id"`
	Description string  `json:"description"`
	Quantity    int     `json:"quantity"`
	UnitPrice   float64 `json:"unit_price"`
	Total       float64 `json:"total"`
}

// Invoice represents an invoice.
type Invoice struct {
	ID            string        `json:"id"`
	InvoiceNumber string        `json:"invoice_number"`
	Status        string        `json:"status"`
	Subtotal      float64       `json:"subtotal"`
	Discount      float64       `json:"discount"`
	Tax           float64       `json:"tax"`
	Total         float64       `json:"total"`
	Currency      string        `json:"currency"`
	DueDate       *string       `json:"due_date"`
	PaidDate      *string       `json:"paid_date"`
	CreatedAt     string        `json:"created_at"`
	Items         []InvoiceItem `json:"items"`
	Links         Links         `json:"_links,omitempty"`
}

// PayInvoiceParams contains the parameters for paying an invoice.
type PayInvoiceParams struct {
	PaymentMethodID string `json:"payment_method_id,omitempty"`
	Gateway         string `json:"gateway,omitempty"`
}

// PaymentResult represents the result of a payment action.
type PaymentResult struct {
	InvoiceID   string  `json:"invoice_id"`
	Status      string  `json:"status"`
	Gateway     string  `json:"gateway"`
	Message     string  `json:"message"`
	CheckoutURL *string `json:"checkout_url"`
}

// ── Subscription ────────────────────────────────────────────────────────────

// Subscription represents a billing subscription.
type Subscription struct {
	ID                 string  `json:"id"`
	ProductName        string  `json:"product_name"`
	BillingCycle       string  `json:"billing_cycle"`
	Amount             float64 `json:"amount"`
	Currency           string  `json:"currency"`
	Status             string  `json:"status"`
	AutoCharge         bool    `json:"auto_charge,omitempty"`
	StartDate          *string `json:"start_date"`
	NextDueDate        *string `json:"next_due_date"`
	CancelledAt        *string `json:"cancelled_at"`
	CancellationReason *string `json:"cancellation_reason,omitempty"`
	SuspendedAt        *string `json:"suspended_at,omitempty"`
	CreatedAt          string  `json:"created_at"`
	UpdatedAt          string  `json:"updated_at,omitempty"`
	Links              Links   `json:"_links,omitempty"`
}

// UpdateSubscriptionParams contains the parameters for updating a subscription.
type UpdateSubscriptionParams struct {
	Action            string `json:"action"` // "cancel" or "reactivate"
	Reason            string `json:"reason,omitempty"`
	CancelAtPeriodEnd *bool  `json:"cancel_at_period_end,omitempty"`
}

// ── Balance ─────────────────────────────────────────────────────────────────

// Balance represents the user's account balance.
type Balance struct {
	Balance  float64 `json:"balance"`
	Currency string  `json:"currency"`
	Links    Links   `json:"_links,omitempty"`
}

// TopUpParams contains the parameters for a balance top-up.
type TopUpParams struct {
	Amount  float64 `json:"amount"`
	Gateway string  `json:"gateway,omitempty"`
}

// TopUpResult represents the result of a top-up request.
type TopUpResult struct {
	Status        string  `json:"status"`
	Amount        float64 `json:"amount"`
	Gateway       string  `json:"gateway"`
	TransactionID *string `json:"transaction_id"`
}

// Transaction represents a balance transaction.
type Transaction struct {
	ID          string  `json:"id"`
	Type        string  `json:"type"`
	Amount      float64 `json:"amount"`
	Description string  `json:"description"`
	CreatedAt   string  `json:"created_at"`
}

// ── Order ───────────────────────────────────────────────────────────────────

// OrderItem represents a line item in an order.
type OrderItem struct {
	ID           string          `json:"id"`
	ProductID    string          `json:"product_id"`
	Quantity     int             `json:"quantity"`
	UnitPrice    float64         `json:"unit_price"`
	Total        float64         `json:"total"`
	BillingCycle string          `json:"billing_cycle"`
	Options      json.RawMessage `json:"options,omitempty"`
	Addons       json.RawMessage `json:"addons,omitempty"`
}

// Order represents an order.
type Order struct {
	ID          string      `json:"id"`
	OrderNumber string      `json:"order_number"`
	Status      string      `json:"status"`
	Subtotal    float64     `json:"subtotal"`
	Discount    float64     `json:"discount"`
	Tax         float64     `json:"tax"`
	Total       float64     `json:"total"`
	Currency    string      `json:"currency"`
	Notes       *string     `json:"notes"`
	CreatedAt   string      `json:"created_at"`
	UpdatedAt   string      `json:"updated_at"`
	Items       []OrderItem `json:"items"`
	Links       Links       `json:"_links,omitempty"`
}

// CreateOrderItemParams contains the parameters for an order item.
type CreateOrderItemParams struct {
	ProductID    string                    `json:"product_id"`
	Quantity     int                       `json:"quantity,omitempty"`
	BillingCycle string                    `json:"billing_cycle,omitempty"`
	Options      map[string]interface{}    `json:"options,omitempty"`
	Addons       []CreateOrderAddonParams  `json:"addons,omitempty"`
}

// CreateOrderAddonParams contains addon info within an order item.
type CreateOrderAddonParams struct {
	AddonID  string `json:"addon_id"`
	Quantity int    `json:"quantity,omitempty"`
}

// CreateOrderParams contains the parameters for creating an order.
type CreateOrderParams struct {
	Items           []CreateOrderItemParams `json:"items"`
	PromoCode       string                  `json:"promo_code,omitempty"`
	Notes           string                  `json:"notes,omitempty"`
	PaymentMethodID string                  `json:"payment_method_id,omitempty"`
}

// ── Payment Method ──────────────────────────────────────────────────────────

// PaymentMethod represents a stored payment method.
type PaymentMethod struct {
	ID          string  `json:"id"`
	Gateway     string  `json:"gateway"`
	Type        string  `json:"type"`
	Label       string  `json:"label"`
	Last4       *string `json:"last4"`
	Brand       *string `json:"brand"`
	ExpiryMonth *int    `json:"expiry_month"`
	ExpiryYear  *int    `json:"expiry_year"`
	IsDefault   bool    `json:"is_default"`
	IsActive    bool    `json:"is_active"`
	CreatedAt   string  `json:"created_at"`
}

// AddPaymentMethodParams contains the parameters for adding a payment method.
type AddPaymentMethodParams struct {
	Gateway      string `json:"gateway"`
	GatewayToken string `json:"gateway_token"`
	SetDefault   bool   `json:"set_default,omitempty"`
}

// ── Promo ───────────────────────────────────────────────────────────────────

// PromoValidation represents the result of a promo code validation.
type PromoValidation struct {
	Valid                bool     `json:"valid"`
	Code                 string   `json:"code"`
	Reason               string   `json:"reason,omitempty"`
	Type                 string   `json:"type,omitempty"`
	Value                float64  `json:"value,omitempty"`
	Currency             string   `json:"currency,omitempty"`
	Description          *string  `json:"description,omitempty"`
	Scope                string   `json:"scope,omitempty"`
	MinOrderAmount       *float64 `json:"min_order_amount,omitempty"`
	MaxDiscount          *float64 `json:"max_discount,omitempty"`
	ApplicableProducts   []string `json:"applicable_products,omitempty"`
	ApplicableCategories []string `json:"applicable_categories,omitempty"`
	ExpiresAt            *string  `json:"expires_at,omitempty"`
}

// ── Ticket ──────────────────────────────────────────────────────────────────

// TicketReply represents a reply to a ticket.
type TicketReply struct {
	ID             string `json:"id"`
	Message        string `json:"message"`
	IsStaff        bool   `json:"is_staff"`
	SenderType     string `json:"sender_type"`
	IsInternalNote bool   `json:"is_internal_note"`
	CreatedAt      string `json:"created_at"`
}

// Ticket represents a support ticket.
type Ticket struct {
	ID           string        `json:"id"`
	TicketNumber string        `json:"ticket_number"`
	Subject      string        `json:"subject"`
	Department   string        `json:"department"`
	Priority     string        `json:"priority"`
	Status       string        `json:"status"`
	Channel      string        `json:"channel,omitempty"`
	Category     string        `json:"category,omitempty"`
	CreatedAt    string        `json:"created_at"`
	UpdatedAt    string        `json:"updated_at"`
	ClosedAt     *string       `json:"closed_at"`
	Replies      []TicketReply `json:"replies,omitempty"`
}

// CreateTicketParams contains the parameters for creating a ticket.
type CreateTicketParams struct {
	Subject    string `json:"subject"`
	Department string `json:"department"`
	Priority   string `json:"priority,omitempty"`
	Message    string `json:"message"`
	Category   string `json:"category,omitempty"`
}

// UpdateTicketParams contains the parameters for updating a ticket (client).
type UpdateTicketParams struct {
	Status string `json:"status"`
}

// CreateReplyParams contains the parameters for creating a ticket reply.
type CreateReplyParams struct {
	Message string `json:"message"`
}

// TicketFeedbackParams contains the parameters for submitting ticket feedback.
type TicketFeedbackParams struct {
	Rating  int    `json:"rating"`
	Comment string `json:"comment,omitempty"`
}

// TicketFeedback represents the feedback response.
type TicketFeedback struct {
	ID         string  `json:"id"`
	TicketID   string  `json:"ticket_id"`
	Rating     int     `json:"rating"`
	Comment    *string `json:"comment"`
	AnsweredAt string  `json:"answered_at"`
}

// Department represents a support department.
type Department struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Slug        string `json:"slug"`
	Description string `json:"description"`
	EmailPrefix string `json:"email_prefix"`
	SortOrder   int    `json:"sort_order"`
}

// ── Knowledge Base ──────────────────────────────────────────────────────────

// KBArticle represents a knowledge base article.
type KBArticle struct {
	ID              string   `json:"id"`
	Title           string   `json:"title"`
	Slug            string   `json:"slug"`
	Type            string   `json:"type"`
	Category        string   `json:"category"`
	Tags            []string `json:"tags"`
	Excerpt         string   `json:"excerpt,omitempty"`
	ContentMD       string   `json:"content_md,omitempty"`
	Status          string   `json:"status"`
	Author          string   `json:"author"`
	PublishedAt     *string  `json:"published_at"`
	ViewCount       int      `json:"view_count"`
	HelpfulCount    int      `json:"helpful_count"`
	NotHelpfulCount int      `json:"not_helpful_count"`
	CreatedAt       string   `json:"created_at"`
	UpdatedAt       string   `json:"updated_at"`
}

// ── Account ─────────────────────────────────────────────────────────────────

// UserProfile represents the current user's profile.
type UserProfile struct {
	ID               string  `json:"id"`
	Email            string  `json:"email"`
	Name             *string `json:"name"`
	FirstName        *string `json:"first_name"`
	LastName         *string `json:"last_name"`
	Phone            *string `json:"phone"`
	AvatarURL        *string `json:"avatar_url"`
	Locale           *string `json:"locale"`
	TwoFactorEnabled bool    `json:"two_factor_enabled"`
	EmailVerified    *string `json:"email_verified"`
	LastLoginAt      *string `json:"last_login_at"`
	CreatedAt        string  `json:"created_at"`
}

// UpdateProfileParams contains the parameters for updating the user profile.
type UpdateProfileParams struct {
	FirstName *string `json:"first_name,omitempty"`
	LastName  *string `json:"last_name,omitempty"`
	Phone     *string `json:"phone,omitempty"`
	Locale    *string `json:"locale,omitempty"`
	AvatarURL *string `json:"avatar_url,omitempty"`
}

// ChangePasswordParams contains the parameters for changing the password.
type ChangePasswordParams struct {
	CurrentPassword string `json:"current_password"`
	NewPassword     string `json:"new_password"`
}

// APIKey represents an API key.
type APIKey struct {
	ID         string   `json:"id"`
	Name       string   `json:"name"`
	KeyPrefix  string   `json:"key_prefix"`
	Key        string   `json:"key,omitempty"` // Only on creation
	Scopes     []string `json:"scopes"`
	AllowedIPs []string `json:"allowed_ips"`
	RateLimit  int      `json:"rate_limit"`
	ExpiresAt  *string  `json:"expires_at"`
	LastUsedAt *string  `json:"last_used_at"`
	IsActive   bool     `json:"is_active"`
	CreatedAt  string   `json:"created_at"`
}

// CreateAPIKeyParams contains the parameters for creating an API key.
type CreateAPIKeyParams struct {
	Name          string   `json:"name"`
	Scopes        []string `json:"scopes,omitempty"`
	AllowedIPs    []string `json:"allowed_ips,omitempty"`
	ExpiresInDays *int     `json:"expires_in_days,omitempty"`
}

// Session represents an active session.
type Session struct {
	ID        string `json:"id"`
	IPAddress string `json:"ip_address"`
	UserAgent string `json:"user_agent"`
	Expires   string `json:"expires"`
}

// TwoFactorParams contains the parameters for 2FA operations.
type TwoFactorParams struct {
	Action string `json:"action"` // "enable", "disable", "status"
}

// TwoFactorStatus represents the current 2FA status.
type TwoFactorStatus struct {
	TwoFactorEnabled bool    `json:"two_factor_enabled"`
	Method           *string `json:"method,omitempty"`
	VerifiedAt       *string `json:"verified_at,omitempty"`
	Message          string  `json:"message,omitempty"`
}

// Notification represents a notification entry.
type Notification struct {
	ID        string          `json:"id"`
	Action    string          `json:"action"`
	Entity    string          `json:"entity"`
	EntityID  string          `json:"entity_id"`
	Details   json.RawMessage `json:"details"`
	Status    string          `json:"status"`
	CreatedAt string          `json:"created_at"`
}

// MarkNotificationsReadParams contains the parameters for marking notifications as read.
type MarkNotificationsReadParams struct {
	NotificationIDs []string `json:"notification_ids"`
}

// MarkReadResult represents the result of marking notifications as read.
type MarkReadResult struct {
	Acknowledged   int `json:"acknowledged"`
	TotalRequested int `json:"total_requested"`
}

// Activity represents an activity log entry.
type Activity struct {
	ID        string          `json:"id"`
	Action    string          `json:"action"`
	Entity    string          `json:"entity"`
	EntityID  string          `json:"entity_id"`
	Details   json.RawMessage `json:"details"`
	IPAddress string          `json:"ip_address"`
	UserAgent string          `json:"user_agent"`
	Status    string          `json:"status"`
	CreatedAt string          `json:"created_at"`
}

// ── Admin Types ─────────────────────────────────────────────────────────────

// UserSummary represents a user in admin list view.
type UserSummary struct {
	ID          string  `json:"id"`
	Email       string  `json:"email"`
	Name        *string `json:"name"`
	FirstName   *string `json:"first_name"`
	LastName    *string `json:"last_name"`
	IsActive    bool    `json:"is_active"`
	Status      string  `json:"status"`
	LastLoginAt *string `json:"last_login_at"`
	CreatedAt   string  `json:"created_at"`
}

// UserRole represents a role assignment.
type UserRole struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Slug string `json:"slug"`
}

// UserDetail represents a user in admin detail view.
type UserDetail struct {
	ID               string     `json:"id"`
	Email            string     `json:"email"`
	Name             *string    `json:"name"`
	FirstName        *string    `json:"first_name"`
	LastName         *string    `json:"last_name"`
	Phone            *string    `json:"phone"`
	AvatarURL        *string    `json:"avatar_url"`
	Locale           *string    `json:"locale"`
	IsActive         bool       `json:"is_active"`
	Status           string     `json:"status"`
	TwoFactorEnabled bool       `json:"two_factor_enabled"`
	LastLoginAt      *string    `json:"last_login_at"`
	CreatedAt        string     `json:"created_at"`
	UpdatedAt        string     `json:"updated_at"`
	Roles            []UserRole `json:"roles"`
}

// CreateUserParams contains the parameters for creating a user.
type CreateUserParams struct {
	Email     string   `json:"email"`
	Name      string   `json:"name,omitempty"`
	FirstName string   `json:"first_name,omitempty"`
	LastName  string   `json:"last_name,omitempty"`
	Password  string   `json:"password,omitempty"`
	RoleIDs   []string `json:"role_ids,omitempty"`
}

// UpdateUserParams contains the parameters for updating a user.
type UpdateUserParams struct {
	Email     *string  `json:"email,omitempty"`
	Name      *string  `json:"name,omitempty"`
	FirstName *string  `json:"first_name,omitempty"`
	LastName  *string  `json:"last_name,omitempty"`
	Phone     *string  `json:"phone,omitempty"`
	IsActive  *bool    `json:"is_active,omitempty"`
	RoleIDs   []string `json:"role_ids,omitempty"`
}

// RolePermission represents a permission assigned to a role.
type RolePermission struct {
	Module string `json:"module"`
	Action string `json:"action"`
}

// Role represents a role.
type Role struct {
	ID          string           `json:"id"`
	Name        string           `json:"name"`
	Slug        string           `json:"slug"`
	Description *string          `json:"description"`
	IsSystem    bool             `json:"is_system"`
	Priority    int              `json:"priority"`
	UserCount   int              `json:"user_count"`
	Permissions []RolePermission `json:"permissions"`
	CreatedAt   string           `json:"created_at"`
	UpdatedAt   string           `json:"updated_at"`
}

// CreateRoleParams contains the parameters for creating a role.
type CreateRoleParams struct {
	Name          string   `json:"name"`
	Slug          string   `json:"slug"`
	Description   string   `json:"description,omitempty"`
	Priority      int      `json:"priority,omitempty"`
	PermissionIDs []string `json:"permission_ids,omitempty"`
}

// UpdateRoleParams contains the parameters for updating a role.
type UpdateRoleParams struct {
	Name          *string  `json:"name,omitempty"`
	Description   *string  `json:"description,omitempty"`
	Priority      *int     `json:"priority,omitempty"`
	PermissionIDs []string `json:"permission_ids,omitempty"`
}

// Tenant represents a tenant.
type Tenant struct {
	ID               string  `json:"id"`
	Name             string  `json:"name"`
	Slug             string  `json:"slug"`
	Domain           *string `json:"domain"`
	LogoURL          *string `json:"logo_url"`
	TenantType       string  `json:"tenant_type"`
	CompanyName      *string `json:"company_name"`
	CompanyEmail     *string `json:"company_email"`
	Timezone         string  `json:"timezone"`
	DefaultLocale    string  `json:"default_locale"`
	RegistrationMode string  `json:"registration_mode"`
	IsActive         bool    `json:"is_active"`
	UserCount        int     `json:"user_count"`
	CreatedAt        string  `json:"created_at"`
	UpdatedAt        string  `json:"updated_at"`
}

// TenantDetail represents a tenant in detail view.
type TenantDetail struct {
	ID                 string   `json:"id"`
	Name               string   `json:"name"`
	Slug               string   `json:"slug"`
	Domain             *string  `json:"domain"`
	LogoURL            *string  `json:"logo_url"`
	LogoIconURL        *string  `json:"logo_icon_url"`
	FaviconURL         *string  `json:"favicon_url"`
	PrimaryColor       *string  `json:"primary_color"`
	SecondaryColor     *string  `json:"secondary_color"`
	AccentColor        *string  `json:"accent_color"`
	TextColor          *string  `json:"text_color"`
	TenantType         string   `json:"tenant_type"`
	OwnerID            *string  `json:"owner_id"`
	CompanyName        *string  `json:"company_name"`
	CompanyEmail       *string  `json:"company_email"`
	CompanyPhone       *string  `json:"company_phone"`
	Timezone           string   `json:"timezone"`
	DefaultLocale      string   `json:"default_locale"`
	EnabledLocales     []string `json:"enabled_locales"`
	EnabledModules     []string `json:"enabled_modules"`
	RegistrationMode   string   `json:"registration_mode"`
	DefaultMarkupType  string   `json:"default_markup_type"`
	DefaultMarkupValue float64  `json:"default_markup_value"`
	IsActive           bool     `json:"is_active"`
	UserCount          int      `json:"user_count"`
	CreatedAt          string   `json:"created_at"`
	UpdatedAt          string   `json:"updated_at"`
}

// UpdateTenantParams contains the parameters for updating a tenant.
type UpdateTenantParams struct {
	Name           *string `json:"name,omitempty"`
	Domain         *string `json:"domain,omitempty"`
	LogoURL        *string `json:"logo_url,omitempty"`
	CompanyName    *string `json:"company_name,omitempty"`
	CompanyEmail   *string `json:"company_email,omitempty"`
	CompanyPhone   *string `json:"company_phone,omitempty"`
	Timezone       *string `json:"timezone,omitempty"`
	DefaultLocale  *string `json:"default_locale,omitempty"`
	PrimaryColor   *string `json:"primary_color,omitempty"`
	SecondaryColor *string `json:"secondary_color,omitempty"`
	AccentColor    *string `json:"accent_color,omitempty"`
	IsActive       *bool   `json:"is_active,omitempty"`
}

// AdminInvoice represents an invoice in admin list view (includes user info).
type AdminInvoice struct {
	Invoice
	IsAutoGenerated bool            `json:"is_auto_generated"`
	Notes           *string         `json:"notes,omitempty"`
	User            *UserInline     `json:"user,omitempty"`
	UpdatedAt       string          `json:"updated_at,omitempty"`
}

// UserInline represents an inline user reference.
type UserInline struct {
	ID    string  `json:"id"`
	Email string  `json:"email"`
	Name  *string `json:"name"`
}

// AdminServer represents a server in admin list view (includes owner info).
type AdminServer struct {
	ID            string      `json:"id"`
	Provider      string      `json:"provider"`
	ExternalID    string      `json:"external_id"`
	Hostname      *string     `json:"hostname"`
	Label         *string     `json:"label"`
	IPAddress     *string     `json:"ip_address"`
	IPv6Address   *string     `json:"ipv6_address"`
	OSImage       *string     `json:"os_image"`
	Plan          *string     `json:"plan"`
	Region        *string     `json:"region"`
	Status        string      `json:"status"`
	Specs         interface{} `json:"specs"`
	ProvisionedAt *string     `json:"provisioned_at"`
	LastCheckedAt *string     `json:"last_checked_at"`
	Owner         *UserInline `json:"owner,omitempty"`
	CreatedAt     string      `json:"created_at"`
	UpdatedAt     string      `json:"updated_at"`
}

// AdminTicket represents a ticket in admin list view.
type AdminTicket struct {
	ID                       string      `json:"id"`
	TicketNumber             string      `json:"ticket_number"`
	Subject                  string      `json:"subject"`
	Department               string      `json:"department"`
	Priority                 string      `json:"priority"`
	Status                   string      `json:"status"`
	UserID                   string      `json:"user_id"`
	Channel                  string      `json:"channel,omitempty"`
	Category                 string      `json:"category,omitempty"`
	LifecycleState           string      `json:"lifecycle_state,omitempty"`
	AssignedAgentPersonaName *string     `json:"assigned_agent_persona_name,omitempty"`
	Brand                    interface{} `json:"brand,omitempty"`
	CreatedAt                string      `json:"created_at"`
	UpdatedAt                string      `json:"updated_at"`
	ClosedAt                 *string     `json:"closed_at"`
	LastReply                *string     `json:"last_reply,omitempty"`
	ReplyCount               int         `json:"reply_count"`
}

// AdminUpdateTicketParams contains the parameters for admin ticket update.
type AdminUpdateTicketParams struct {
	Status                 *string `json:"status,omitempty"`
	Priority               *string `json:"priority,omitempty"`
	AssignedAgentPersonaID *string `json:"assigned_agent_persona_id,omitempty"`
	AssignedTier           *int    `json:"assigned_tier,omitempty"`
	Department             *string `json:"department,omitempty"`
	Category               *string `json:"category,omitempty"`
}

// Lead represents a sales lead.
type Lead struct {
	ID              string   `json:"id"`
	Email           *string  `json:"email"`
	Phone           *string  `json:"phone"`
	FirstName       *string  `json:"first_name"`
	LastName        *string  `json:"last_name"`
	Company         *string  `json:"company"`
	JobTitle        *string  `json:"job_title"`
	Source          string   `json:"source"`
	SourceDetail    *string  `json:"source_detail"`
	Status          string   `json:"status"`
	Score           *int     `json:"score"`
	Industry        *string  `json:"industry"`
	Country         *string  `json:"country"`
	Region          *string  `json:"region"`
	Tags            []string `json:"tags"`
	LastContactedAt *string  `json:"last_contacted_at"`
	ContactCount    int      `json:"contact_count"`
	ConvertedAt     *string  `json:"converted_at"`
	CreatedAt       string   `json:"created_at"`
	UpdatedAt       string   `json:"updated_at"`
}

// CreateLeadParams contains the parameters for creating a lead.
type CreateLeadParams struct {
	Email        string   `json:"email,omitempty"`
	Phone        string   `json:"phone,omitempty"`
	FirstName    string   `json:"first_name,omitempty"`
	LastName     string   `json:"last_name,omitempty"`
	Company      string   `json:"company,omitempty"`
	JobTitle     string   `json:"job_title,omitempty"`
	Source       string   `json:"source"`
	SourceDetail string   `json:"source_detail,omitempty"`
	Industry     string   `json:"industry,omitempty"`
	Country      string   `json:"country,omitempty"`
	Region       string   `json:"region,omitempty"`
	Tags         []string `json:"tags,omitempty"`
	Notes        string   `json:"notes,omitempty"`
}

// AuditEntry represents an audit log entry.
type AuditEntry struct {
	ID        string          `json:"id"`
	UserID    string          `json:"user_id"`
	Action    string          `json:"action"`
	Entity    string          `json:"entity"`
	EntityID  string          `json:"entity_id"`
	IPAddress string          `json:"ip_address"`
	UserAgent string          `json:"user_agent"`
	Status    string          `json:"status"`
	Details   json.RawMessage `json:"details"`
	CreatedAt string          `json:"created_at"`
}

// DashboardStats represents admin dashboard KPI stats.
type DashboardStats struct {
	Users    DashboardUsersStats    `json:"users"`
	Servers  DashboardServersStats  `json:"servers"`
	Revenue  DashboardRevenueStats  `json:"revenue"`
	Tickets  DashboardTicketsStats  `json:"tickets"`
	Invoices DashboardInvoicesStats `json:"invoices"`
	Period   DashboardPeriod        `json:"period"`
}

// DashboardUsersStats contains user stats.
type DashboardUsersStats struct {
	Total  int `json:"total"`
	Active int `json:"active"`
}

// DashboardServersStats contains server stats.
type DashboardServersStats struct {
	Total  int `json:"total"`
	Active int `json:"active"`
}

// DashboardRevenueStats contains revenue stats.
type DashboardRevenueStats struct {
	ThisMonth float64 `json:"this_month"`
	Currency  string  `json:"currency"`
}

// DashboardTicketsStats contains ticket stats.
type DashboardTicketsStats struct {
	Open int `json:"open"`
}

// DashboardInvoicesStats contains invoice stats.
type DashboardInvoicesStats struct {
	Total  int `json:"total"`
	Unpaid int `json:"unpaid"`
}

// DashboardPeriod contains the time period for stats.
type DashboardPeriod struct {
	MonthStart  string `json:"month_start"`
	GeneratedAt string `json:"generated_at"`
}

// MasqueradeResult represents the result of a masquerade action.
type MasqueradeResult struct {
	TargetUserID string  `json:"target_user_id,omitempty"`
	TargetName   *string `json:"target_name,omitempty"`
	TargetEmail  string  `json:"target_email,omitempty"`
	Masquerading bool    `json:"masquerading"`
}

// V1UsageStats represents v1 API usage statistics.
type V1UsageStats struct {
	Endpoints      json.RawMessage `json:"endpoints"`
	TotalV1Hits    int             `json:"total_v1_hits"`
	MigratedCount  int             `json:"migrated_count"`
	UnmigratedCount int            `json:"unmigrated_count"`
	TotalTracked   int             `json:"total_tracked"`
}

// ── Affiliate ───────────────────────────────────────────────────────────────

// Affiliate represents an affiliate account.
type Affiliate struct {
	ID                  string  `json:"id"`
	Status              string  `json:"status"`
	ReferralCode        string  `json:"referral_code"`
	CustomSlug          *string `json:"custom_slug"`
	PayoutMethod        string  `json:"payout_method"`
	TotalEarnings       float64 `json:"total_earnings"`
	TotalPaid           float64 `json:"total_paid"`
	PendingBalance      float64 `json:"pending_balance"`
	AvailableBalance    float64 `json:"available_balance"`
	LifetimeReferrals   int     `json:"lifetime_referrals"`
	LifetimeConversions int     `json:"lifetime_conversions"`
	EnrolledAt          string  `json:"enrolled_at"`
	CreatedAt           string  `json:"created_at"`
}

// Commission represents an affiliate commission.
type Commission struct {
	ID               string  `json:"id"`
	Status           string  `json:"status"`
	CommissionType   string  `json:"commission_type"`
	CommissionRate   float64 `json:"commission_rate"`
	OrderAmount      float64 `json:"order_amount"`
	CommissionAmount float64 `json:"commission_amount"`
	BonusAmount      float64 `json:"bonus_amount"`
	Currency         string  `json:"currency"`
	Description      string  `json:"description"`
	IsRecurring      bool    `json:"is_recurring"`
	RecurringMonth   int     `json:"recurring_month"`
	HoldUntil        *string `json:"hold_until"`
	ApprovedAt       *string `json:"approved_at"`
	PaidAt           *string `json:"paid_at"`
	CreatedAt        string  `json:"created_at"`
}

// Referral represents an affiliate referral.
type Referral struct {
	ID               string   `json:"id"`
	LandingPage      string   `json:"landing_page"`
	RefSource        string   `json:"ref_source"`
	SubID            *string  `json:"sub_id"`
	SignedUpAt       *string  `json:"signed_up_at"`
	ConvertedAt      *string  `json:"converted_at"`
	IsConverted      bool     `json:"is_converted"`
	ConversionAmount *float64 `json:"conversion_amount"`
	CreatedAt        string   `json:"created_at"`
}

// PayoutRequest represents a payout request.
type PayoutRequest struct {
	ID          string  `json:"id"`
	Status      string  `json:"status"`
	Method      string  `json:"method"`
	Amount      float64 `json:"amount"`
	Currency    string  `json:"currency"`
	Fees        float64 `json:"fees"`
	NetAmount   float64 `json:"net_amount"`
	RequestedAt string  `json:"requested_at"`
	CreatedAt   string  `json:"created_at"`
}

// MarketingMaterial represents affiliate marketing material.
type MarketingMaterial struct {
	ID        string  `json:"id"`
	Name      string  `json:"name"`
	Type      string  `json:"type"`
	Content   string  `json:"content"`
	ImageURL  *string `json:"image_url"`
	Width     *int    `json:"width"`
	Height    *int    `json:"height"`
	SortOrder int     `json:"sort_order"`
	CreatedAt string  `json:"created_at"`
	UpdatedAt string  `json:"updated_at"`
}

// ── Webhook ─────────────────────────────────────────────────────────────────

// Webhook represents a webhook subscription.
type Webhook struct {
	ID        string   `json:"id"`
	URL       string   `json:"url"`
	Events    []string `json:"events"`
	Secret    string   `json:"secret"`
	IsActive  bool     `json:"is_active"`
	CreatedAt string   `json:"created_at"`
	UpdatedAt string   `json:"updated_at"`
}

// CreateWebhookParams contains the parameters for creating a webhook.
type CreateWebhookParams struct {
	URL    string   `json:"url"`
	Events []string `json:"events"`
}

// UpdateWebhookParams contains the parameters for updating a webhook.
type UpdateWebhookParams struct {
	URL      *string  `json:"url,omitempty"`
	Events   []string `json:"events,omitempty"`
	IsActive *bool    `json:"is_active,omitempty"`
}

// WebhookDelivery represents a webhook delivery log entry.
type WebhookDelivery struct {
	ID             string          `json:"id"`
	WebhookID      string          `json:"webhook_id"`
	Event          string          `json:"event"`
	URL            string          `json:"url"`
	RequestBody    json.RawMessage `json:"request_body"`
	ResponseStatus int             `json:"response_status"`
	ResponseBody   string          `json:"response_body"`
	Duration       int             `json:"duration_ms"`
	Success        bool            `json:"success"`
	Attempt        int             `json:"attempt"`
	CreatedAt      string          `json:"created_at"`
}

// WebhookTestResult represents the result of a test delivery.
type WebhookTestResult struct {
	ID             string `json:"id"`
	Event          string `json:"event"`
	ResponseStatus int    `json:"response_status"`
	Success        bool   `json:"success"`
	Duration       int    `json:"duration_ms"`
}

// WebhookEventType represents an available webhook event type.
type WebhookEventType struct {
	Event       string `json:"event"`
	Category    string `json:"category"`
	Description string `json:"description"`
}

// ── Task ────────────────────────────────────────────────────────────────────

// Task represents an async task.
type Task struct {
	ID           string          `json:"id"`
	Type         string          `json:"type"`
	Status       string          `json:"status"`
	ResourceID   *string         `json:"resource_id"`
	ResourceType *string         `json:"resource_type"`
	Progress     *int            `json:"progress"`
	Result       json.RawMessage `json:"result"`
	Error        *string         `json:"error"`
	CreatedAt    string          `json:"created_at"`
	UpdatedAt    string          `json:"updated_at"`
	StartedAt    *string         `json:"started_at"`
	CompletedAt  *string         `json:"completed_at"`
	Links        Links           `json:"_links,omitempty"`
}

// TaskAccepted represents a 202 Accepted response with a task ID.
type TaskAccepted struct {
	TaskID string `json:"task_id"`
	Status string `json:"status"`
}

// CancelledTask represents the result of cancelling a task.
type CancelledTask struct {
	ID          string `json:"id"`
	Status      string `json:"status"`
	CancelledAt string `json:"cancelled_at"`
}

// ── Zone ────────────────────────────────────────────────────────────────────

// Zone represents a hosting zone (region/datacenter).
type Zone struct {
	ID              string  `json:"id"`
	Slug            string  `json:"slug"`
	Name            string  `json:"name"`
	Description     *string `json:"description"`
	Country         string  `json:"country"`
	City            *string `json:"city"`
	Continent       string  `json:"continent"`
	FlagEmoji       *string `json:"flag_emoji"`
	SortOrder       int     `json:"sort_order"`
	LatencyEndpoint *string `json:"latency_endpoint"`
	ProviderSlug    string  `json:"provider_slug,omitempty"`
	ProviderRegion  string  `json:"provider_region,omitempty"`
	Links           Links   `json:"_links,omitempty"`
}

// OSImage represents an OS image available in a zone.
type OSImage struct {
	ID           string   `json:"id"`
	Name         string   `json:"name"`
	Description  *string  `json:"description"`
	Family       string   `json:"family"`
	Version      string   `json:"version"`
	Architecture *string  `json:"architecture"`
	MinDiskGB    *int     `json:"min_disk_gb"`
	IsCustom     bool     `json:"is_custom"`
	Price        float64  `json:"price"`
	Currency     string   `json:"currency"`
	Status       *string  `json:"status"`
	SizeMB       *int     `json:"size_mb"`
}

// ── Reseller ────────────────────────────────────────────────────────────────

// ResellerOverview represents the reseller dashboard data.
type ResellerOverview struct {
	Tenant   ResellerTenant   `json:"tenant"`
	Earnings ResellerEarnings `json:"earnings"`
	Clients  ResellerClients  `json:"clients"`
	Pricing  ResellerPricing  `json:"pricing"`
}

// ResellerTenant represents basic tenant info for reseller view.
type ResellerTenant struct {
	ID       string  `json:"id"`
	Name     string  `json:"name"`
	Slug     string  `json:"slug"`
	Domain   *string `json:"domain"`
	IsActive bool    `json:"is_active"`
}

// ResellerEarnings represents reseller earnings summary.
type ResellerEarnings struct {
	TotalGross        float64 `json:"total_gross"`
	TotalCost         float64 `json:"total_cost"`
	TotalNet          float64 `json:"total_net"`
	TotalTransactions int     `json:"total_transactions"`
}

// ResellerClients represents reseller client counts.
type ResellerClients struct {
	Total  int `json:"total"`
	Active int `json:"active"`
}

// ResellerPricing represents reseller pricing config.
type ResellerPricing struct {
	DefaultMarkupType  string `json:"default_markup_type"`
	DefaultMarkupValue float64 `json:"default_markup_value"`
	ProductOverrides   int     `json:"product_overrides"`
}

// ── Generic Deletion Response ───────────────────────────────────────────────

// DeleteResult represents a generic deletion result.
type DeleteResult struct {
	Deleted bool   `json:"deleted,omitempty"`
	ID      string `json:"id,omitempty"`
	Revoked bool   `json:"revoked,omitempty"`
}

// MessageResult represents a generic message result.
type MessageResult struct {
	Message string `json:"message"`
}
