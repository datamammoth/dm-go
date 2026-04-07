package datamammoth

import "context"

// SupportService handles communication with the support-related endpoints of the
// DataMammoth API v2 (tickets, replies, departments, knowledge base).
type SupportService struct {
	client *Client
}

// ── Tickets ─────────────────────────────────────────────────────────────────

// ListTickets returns a paginated list of the current user's tickets.
//
// API: GET /tickets
func (s *SupportService) ListTickets(ctx context.Context, opts *ListOptions) ([]Ticket, *Pagination, error) {
	var tickets []Ticket
	pagination, err := s.client.doList(ctx, "/tickets", opts, &tickets)
	if err != nil {
		return nil, nil, err
	}
	return tickets, pagination, nil
}

// ListAllTickets returns an iterator over all pages of tickets.
func (s *SupportService) ListAllTickets(opts *ListOptions) *Iterator[Ticket] {
	return newIterator[Ticket](s.client, "/tickets", opts)
}

// GetTicket retrieves a single ticket by ID.
//
// API: GET /tickets/:id
func (s *SupportService) GetTicket(ctx context.Context, id string) (*Ticket, error) {
	var ticket Ticket
	_, err := s.client.get(ctx, "/tickets/"+id, &ticket)
	if err != nil {
		return nil, err
	}
	return &ticket, nil
}

// CreateTicket creates a new support ticket.
//
// API: POST /tickets
func (s *SupportService) CreateTicket(ctx context.Context, params *CreateTicketParams) (*Ticket, error) {
	var ticket Ticket
	_, err := s.client.post(ctx, "/tickets", params, &ticket)
	if err != nil {
		return nil, err
	}
	return &ticket, nil
}

// CloseTicket closes a ticket. Clients can only close tickets via the API.
//
// API: PATCH /tickets/:id
func (s *SupportService) CloseTicket(ctx context.Context, id string) error {
	_, err := s.client.patch(ctx, "/tickets/"+id, &UpdateTicketParams{Status: "CLOSED"}, nil)
	return err
}

// ── Replies ─────────────────────────────────────────────────────────────────

// CreateReply adds a reply to a ticket.
//
// API: POST /tickets/:id/replies
func (s *SupportService) CreateReply(ctx context.Context, ticketID string, params *CreateReplyParams) (*TicketReply, error) {
	var reply TicketReply
	_, err := s.client.post(ctx, "/tickets/"+ticketID+"/replies", params, &reply)
	if err != nil {
		return nil, err
	}
	return &reply, nil
}

// ── Feedback ────────────────────────────────────────────────────────────────

// SubmitFeedback submits CSAT feedback for a ticket.
//
// API: POST /tickets/:id/feedback
func (s *SupportService) SubmitFeedback(ctx context.Context, ticketID string, params *TicketFeedbackParams) (*TicketFeedback, error) {
	var feedback TicketFeedback
	_, err := s.client.post(ctx, "/tickets/"+ticketID+"/feedback", params, &feedback)
	if err != nil {
		return nil, err
	}
	return &feedback, nil
}

// ── Departments ─────────────────────────────────────────────────────────────

// ListDepartments returns the active support departments.
//
// API: GET /tickets/departments
func (s *SupportService) ListDepartments(ctx context.Context) ([]Department, error) {
	var departments []Department
	_, err := s.client.get(ctx, "/tickets/departments", &departments)
	if err != nil {
		return nil, err
	}
	return departments, nil
}

// ── Knowledge Base ──────────────────────────────────────────────────────────

// ListArticles returns KB articles (public). Supports search and category filter.
//
// API: GET /kb/articles
func (s *SupportService) ListArticles(ctx context.Context, opts *ListOptions) ([]KBArticle, *Pagination, error) {
	var articles []KBArticle
	pagination, err := s.client.doList(ctx, "/kb/articles", opts, &articles)
	if err != nil {
		return nil, nil, err
	}
	return articles, pagination, nil
}

// GetArticle retrieves a single KB article by slug.
//
// API: GET /kb/articles/:slug
func (s *SupportService) GetArticle(ctx context.Context, slug string) (*KBArticle, error) {
	var article KBArticle
	_, err := s.client.get(ctx, "/kb/articles/"+slug, &article)
	if err != nil {
		return nil, err
	}
	return &article, nil
}
