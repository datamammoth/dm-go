package datamammoth

import "context"

// ZonesService handles communication with the zone-related endpoints of the
// DataMammoth API v2.
type ZonesService struct {
	client *Client
}

// List returns the available hosting zones.
//
// API: GET /zones
func (s *ZonesService) List(ctx context.Context, opts *ListOptions) ([]Zone, error) {
	path := s.client.buildListURL("/zones", opts)
	var zones []Zone
	_, err := s.client.get(ctx, path, &zones)
	if err != nil {
		return nil, err
	}
	return zones, nil
}

// ListImages returns the OS images available in a zone.
//
// API: GET /zones/:id/images
func (s *ZonesService) ListImages(ctx context.Context, zoneID string, opts *ListOptions) ([]OSImage, error) {
	path := s.client.buildListURL("/zones/"+zoneID+"/images", opts)
	var images []OSImage
	_, err := s.client.get(ctx, path, &images)
	if err != nil {
		return nil, err
	}
	return images, nil
}

// GetResellerOverview retrieves the reseller dashboard overview.
//
// API: GET /reseller
func (s *ZonesService) GetResellerOverview(ctx context.Context) (*ResellerOverview, error) {
	var overview ResellerOverview
	_, err := s.client.get(ctx, "/reseller", &overview)
	if err != nil {
		return nil, err
	}
	return &overview, nil
}
