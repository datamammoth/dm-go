package datamammoth

import (
	"context"
	"fmt"
)

// AdminService handles communication with the admin-only endpoints of the
// DataMammoth API v2 (users, roles, tenants, invoices, servers, tickets, leads,
// audit, dashboard, masquerade).
type AdminService struct {
	client *Client
}

// ── Users ───────────────────────────────────────────────────────────────────

// ListUsers returns a paginated list of users (admin).
//
// API: GET /admin/users
func (s *AdminService) ListUsers(ctx context.Context, opts *ListOptions) ([]UserSummary, *Pagination, error) {
	var users []UserSummary
	pagination, err := s.client.doList(ctx, "/admin/users", opts, &users)
	if err != nil {
		return nil, nil, err
	}
	return users, pagination, nil
}

// ListAllUsers returns an iterator over all pages of users.
func (s *AdminService) ListAllUsers(opts *ListOptions) *Iterator[UserSummary] {
	return newIterator[UserSummary](s.client, "/admin/users", opts)
}

// GetUser retrieves a user by ID (admin).
//
// API: GET /admin/users/:id
func (s *AdminService) GetUser(ctx context.Context, id string) (*UserDetail, error) {
	var user UserDetail
	_, err := s.client.get(ctx, "/admin/users/"+id, &user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// CreateUser creates a new user (admin).
//
// API: POST /admin/users
func (s *AdminService) CreateUser(ctx context.Context, params *CreateUserParams) (*UserSummary, error) {
	var user UserSummary
	_, err := s.client.post(ctx, "/admin/users", params, &user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// UpdateUser updates a user (admin).
//
// API: PATCH /admin/users/:id
func (s *AdminService) UpdateUser(ctx context.Context, id string, params *UpdateUserParams) (*UserDetail, error) {
	var user UserDetail
	_, err := s.client.patch(ctx, "/admin/users/"+id, params, &user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// DeleteUser soft-deletes (deactivates) a user (admin).
//
// API: DELETE /admin/users/:id
func (s *AdminService) DeleteUser(ctx context.Context, id string) error {
	_, err := s.client.del(ctx, "/admin/users/"+id, nil)
	return err
}

// ── Roles ───────────────────────────────────────────────────────────────────

// ListRoles returns all roles for the tenant (admin).
//
// API: GET /admin/roles
func (s *AdminService) ListRoles(ctx context.Context) ([]Role, error) {
	var roles []Role
	_, err := s.client.get(ctx, "/admin/roles", &roles)
	if err != nil {
		return nil, err
	}
	return roles, nil
}

// GetRole retrieves a role by ID (admin).
//
// API: GET /admin/roles/:id
func (s *AdminService) GetRole(ctx context.Context, id string) (*Role, error) {
	var role Role
	_, err := s.client.get(ctx, "/admin/roles/"+id, &role)
	if err != nil {
		return nil, err
	}
	return &role, nil
}

// CreateRole creates a new role (admin).
//
// API: POST /admin/roles
func (s *AdminService) CreateRole(ctx context.Context, params *CreateRoleParams) (*Role, error) {
	var role Role
	_, err := s.client.post(ctx, "/admin/roles", params, &role)
	if err != nil {
		return nil, err
	}
	return &role, nil
}

// UpdateRole updates a role (admin).
//
// API: PATCH /admin/roles/:id
func (s *AdminService) UpdateRole(ctx context.Context, id string, params *UpdateRoleParams) (*Role, error) {
	var role Role
	_, err := s.client.patch(ctx, "/admin/roles/"+id, params, &role)
	if err != nil {
		return nil, err
	}
	return &role, nil
}

// DeleteRole deletes a role (admin). System roles and roles with assigned users
// cannot be deleted.
//
// API: DELETE /admin/roles/:id
func (s *AdminService) DeleteRole(ctx context.Context, id string) error {
	_, err := s.client.del(ctx, "/admin/roles/"+id, nil)
	return err
}

// ── Tenants ─────────────────────────────────────────────────────────────────

// ListTenants returns a paginated list of tenants (admin).
//
// API: GET /admin/tenants
func (s *AdminService) ListTenants(ctx context.Context, opts *ListOptions) ([]Tenant, *Pagination, error) {
	var tenants []Tenant
	pagination, err := s.client.doList(ctx, "/admin/tenants", opts, &tenants)
	if err != nil {
		return nil, nil, err
	}
	return tenants, pagination, nil
}

// GetTenant retrieves a tenant by ID (admin).
//
// API: GET /admin/tenants/:id
func (s *AdminService) GetTenant(ctx context.Context, id string) (*TenantDetail, error) {
	var tenant TenantDetail
	_, err := s.client.get(ctx, "/admin/tenants/"+id, &tenant)
	if err != nil {
		return nil, err
	}
	return &tenant, nil
}

// UpdateTenant updates a tenant's settings (admin).
//
// API: PATCH /admin/tenants/:id
func (s *AdminService) UpdateTenant(ctx context.Context, id string, params *UpdateTenantParams) (*TenantDetail, error) {
	var tenant TenantDetail
	_, err := s.client.patch(ctx, "/admin/tenants/"+id, params, &tenant)
	if err != nil {
		return nil, err
	}
	return &tenant, nil
}

// ── Admin Invoices ──────────────────────────────────────────────────────────

// ListAllInvoices returns a paginated list of all invoices for the tenant (admin).
//
// API: GET /admin/invoices
func (s *AdminService) ListAllInvoices(ctx context.Context, opts *ListOptions) ([]AdminInvoice, *Pagination, error) {
	var invoices []AdminInvoice
	pagination, err := s.client.doList(ctx, "/admin/invoices", opts, &invoices)
	if err != nil {
		return nil, nil, err
	}
	return invoices, pagination, nil
}

// ── Admin Servers ───────────────────────────────────────────────────────────

// ListAllServers returns a paginated list of all servers for the tenant (admin).
//
// API: GET /admin/servers
func (s *AdminService) ListAllServers(ctx context.Context, opts *ListOptions) ([]AdminServer, *Pagination, error) {
	var servers []AdminServer
	pagination, err := s.client.doList(ctx, "/admin/servers", opts, &servers)
	if err != nil {
		return nil, nil, err
	}
	return servers, pagination, nil
}

// ── Admin Tickets ───────────────────────────────────────────────────────────

// ListAllTickets returns a paginated list of all tickets for the tenant (admin).
//
// API: GET /admin/tickets
func (s *AdminService) ListAllTickets(ctx context.Context, opts *ListOptions) ([]AdminTicket, *Pagination, error) {
	var tickets []AdminTicket
	pagination, err := s.client.doList(ctx, "/admin/tickets", opts, &tickets)
	if err != nil {
		return nil, nil, err
	}
	return tickets, pagination, nil
}

// UpdateTicket updates a ticket (admin: status, priority, assignment).
//
// API: PATCH /admin/tickets/:id
func (s *AdminService) UpdateTicket(ctx context.Context, id string, params *AdminUpdateTicketParams) (*AdminTicket, error) {
	var ticket AdminTicket
	_, err := s.client.patch(ctx, "/admin/tickets/"+id, params, &ticket)
	if err != nil {
		return nil, err
	}
	return &ticket, nil
}

// ── Leads ───────────────────────────────────────────────────────────────────

// ListLeads returns a paginated list of leads (admin).
//
// API: GET /admin/leads
func (s *AdminService) ListLeads(ctx context.Context, opts *ListOptions) ([]Lead, *Pagination, error) {
	var leads []Lead
	pagination, err := s.client.doList(ctx, "/admin/leads", opts, &leads)
	if err != nil {
		return nil, nil, err
	}
	return leads, pagination, nil
}

// ListAllLeads returns an iterator over all pages of leads.
func (s *AdminService) ListAllLeads(opts *ListOptions) *Iterator[Lead] {
	return newIterator[Lead](s.client, "/admin/leads", opts)
}

// CreateLead creates a new lead (admin).
//
// API: POST /admin/leads
func (s *AdminService) CreateLead(ctx context.Context, params *CreateLeadParams) (*Lead, error) {
	var lead Lead
	_, err := s.client.post(ctx, "/admin/leads", params, &lead)
	if err != nil {
		return nil, err
	}
	return &lead, nil
}

// ── Audit Log ───────────────────────────────────────────────────────────────

// ListAuditLog returns a paginated list of audit log entries (admin).
//
// API: GET /admin/audit-log
func (s *AdminService) ListAuditLog(ctx context.Context, opts *ListOptions) ([]AuditEntry, *Pagination, error) {
	var entries []AuditEntry
	pagination, err := s.client.doList(ctx, "/admin/audit-log", opts, &entries)
	if err != nil {
		return nil, nil, err
	}
	return entries, pagination, nil
}

// ListAllAuditLog returns an iterator over all pages of audit log entries.
func (s *AdminService) ListAllAuditLog(opts *ListOptions) *Iterator[AuditEntry] {
	return newIterator[AuditEntry](s.client, "/admin/audit-log", opts)
}

// ── Dashboard ───────────────────────────────────────────────────────────────

// GetDashboardStats returns admin dashboard KPI stats.
//
// API: GET /admin/dashboard/stats
func (s *AdminService) GetDashboardStats(ctx context.Context) (*DashboardStats, error) {
	var stats DashboardStats
	_, err := s.client.get(ctx, "/admin/dashboard/stats", &stats)
	if err != nil {
		return nil, err
	}
	return &stats, nil
}

// ── Masquerade ──────────────────────────────────────────────────────────────

// StartMasquerade starts masquerading as another user (admin).
//
// API: POST /admin/masquerade/:userId
func (s *AdminService) StartMasquerade(ctx context.Context, userID string) (*MasqueradeResult, error) {
	var result MasqueradeResult
	_, err := s.client.post(ctx, fmt.Sprintf("/admin/masquerade/%s", userID), nil, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// StopMasquerade stops the current masquerade session.
//
// API: DELETE /admin/masquerade/:userId
func (s *AdminService) StopMasquerade(ctx context.Context, userID string) (*MasqueradeResult, error) {
	var result MasqueradeResult
	_, err := s.client.del(ctx, fmt.Sprintf("/admin/masquerade/%s", userID), &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// ── V1 Usage ────────────────────────────────────────────────────────────────

// GetV1Usage returns usage statistics for v1 API endpoints (admin).
//
// API: GET /admin/v1-usage
func (s *AdminService) GetV1Usage(ctx context.Context) (*V1UsageStats, error) {
	var stats V1UsageStats
	_, err := s.client.get(ctx, "/admin/v1-usage", &stats)
	if err != nil {
		return nil, err
	}
	return &stats, nil
}
