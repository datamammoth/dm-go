package datamammoth

import "context"

// ProductsService handles communication with the product-related endpoints of the
// DataMammoth API v2.
type ProductsService struct {
	client *Client
}

// List returns a paginated list of products.
//
// API: GET /products
func (s *ProductsService) List(ctx context.Context, opts *ListOptions) ([]Product, *Pagination, error) {
	var products []Product
	pagination, err := s.client.doList(ctx, "/products", opts, &products)
	if err != nil {
		return nil, nil, err
	}
	return products, pagination, nil
}

// ListAll returns an iterator over all pages of products.
func (s *ProductsService) ListAll(opts *ListOptions) *Iterator[Product] {
	return newIterator[Product](s.client, "/products", opts)
}

// Get retrieves a single product by ID.
//
// API: GET /products/:id
func (s *ProductsService) Get(ctx context.Context, id string) (*Product, error) {
	var product Product
	_, err := s.client.get(ctx, "/products/"+id, &product)
	if err != nil {
		return nil, err
	}
	return &product, nil
}

// Create creates a new product (admin only).
//
// API: POST /products
func (s *ProductsService) Create(ctx context.Context, params *CreateProductParams) (*Product, error) {
	var product Product
	_, err := s.client.post(ctx, "/products", params, &product)
	if err != nil {
		return nil, err
	}
	return &product, nil
}

// Update modifies an existing product (admin only).
//
// API: PATCH /products/:id
func (s *ProductsService) Update(ctx context.Context, id string, params *UpdateProductParams) (*Product, error) {
	var product Product
	_, err := s.client.patch(ctx, "/products/"+id, params, &product)
	if err != nil {
		return nil, err
	}
	return &product, nil
}

// Delete removes a product (admin only).
//
// API: DELETE /products/:id
func (s *ProductsService) Delete(ctx context.Context, id string) error {
	_, err := s.client.del(ctx, "/products/"+id, nil)
	return err
}

// ListOptions returns the configurable options for a product.
//
// API: GET /products/:id/options
func (s *ProductsService) ListOptions(ctx context.Context, id string) ([]ProductOption, error) {
	var options []ProductOption
	_, err := s.client.get(ctx, "/products/"+id+"/options", &options)
	if err != nil {
		return nil, err
	}
	return options, nil
}

// ListAddons returns the addons for a product.
//
// API: GET /products/:id/addons
func (s *ProductsService) ListAddons(ctx context.Context, id string) ([]ProductAddon, error) {
	var addons []ProductAddon
	_, err := s.client.get(ctx, "/products/"+id+"/addons", &addons)
	if err != nil {
		return nil, err
	}
	return addons, nil
}

// GetPricing returns the base and retail pricing for a product.
//
// API: GET /products/:id/pricing
func (s *ProductsService) GetPricing(ctx context.Context, id string, tenantID string) (*ProductPricingDetail, error) {
	path := "/products/" + id + "/pricing"
	if tenantID != "" {
		path += "?tenant_id=" + tenantID
	}
	var pricing ProductPricingDetail
	_, err := s.client.get(ctx, path, &pricing)
	if err != nil {
		return nil, err
	}
	return &pricing, nil
}

// ListCategories returns the product categories.
//
// API: GET /categories
func (s *ProductsService) ListCategories(ctx context.Context, tenantID string) ([]Category, error) {
	path := "/categories"
	if tenantID != "" {
		path += "?tenant_id=" + tenantID
	}
	var categories []Category
	_, err := s.client.get(ctx, path, &categories)
	if err != nil {
		return nil, err
	}
	return categories, nil
}
