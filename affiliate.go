package datamammoth

import "context"

// AffiliateService handles communication with the affiliate-related endpoints of the
// DataMammoth API v2.
type AffiliateService struct {
	client *Client
}

// GetDashboard retrieves the current user's affiliate dashboard.
//
// API: GET /affiliate/me
func (s *AffiliateService) GetDashboard(ctx context.Context) (*Affiliate, error) {
	var aff Affiliate
	_, err := s.client.get(ctx, "/affiliate/me", &aff)
	if err != nil {
		return nil, err
	}
	return &aff, nil
}

// ListCommissions returns a paginated list of the current affiliate's commissions.
//
// API: GET /affiliate/commissions
func (s *AffiliateService) ListCommissions(ctx context.Context, opts *ListOptions) ([]Commission, *Pagination, error) {
	var commissions []Commission
	pagination, err := s.client.doList(ctx, "/affiliate/commissions", opts, &commissions)
	if err != nil {
		return nil, nil, err
	}
	return commissions, pagination, nil
}

// ListAllCommissions returns an iterator over all pages of commissions.
func (s *AffiliateService) ListAllCommissions(opts *ListOptions) *Iterator[Commission] {
	return newIterator[Commission](s.client, "/affiliate/commissions", opts)
}

// ListReferrals returns a paginated list of the current affiliate's referrals.
//
// API: GET /affiliate/referrals
func (s *AffiliateService) ListReferrals(ctx context.Context, opts *ListOptions) ([]Referral, *Pagination, error) {
	var referrals []Referral
	pagination, err := s.client.doList(ctx, "/affiliate/referrals", opts, &referrals)
	if err != nil {
		return nil, nil, err
	}
	return referrals, pagination, nil
}

// ListAllReferrals returns an iterator over all pages of referrals.
func (s *AffiliateService) ListAllReferrals(opts *ListOptions) *Iterator[Referral] {
	return newIterator[Referral](s.client, "/affiliate/referrals", opts)
}

// RequestPayout requests a payout of available affiliate earnings.
//
// API: POST /affiliate/payout-request
func (s *AffiliateService) RequestPayout(ctx context.Context) (*PayoutRequest, error) {
	var payout PayoutRequest
	_, err := s.client.post(ctx, "/affiliate/payout-request", nil, &payout)
	if err != nil {
		return nil, err
	}
	return &payout, nil
}

// ListMaterials returns the available marketing materials.
//
// API: GET /affiliate/materials
func (s *AffiliateService) ListMaterials(ctx context.Context) ([]MarketingMaterial, error) {
	var materials []MarketingMaterial
	_, err := s.client.get(ctx, "/affiliate/materials", &materials)
	if err != nil {
		return nil, err
	}
	return materials, nil
}
