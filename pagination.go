package datamammoth

import "context"

// ListPage represents a single page of results from a list endpoint.
type ListPage[T any] struct {
	Data       []T
	Pagination *Pagination
}

// Iterator provides a way to iterate over all pages of a paginated list.
type Iterator[T any] struct {
	client  *Client
	path    string
	opts    ListOptions
	page    int
	done    bool
}

// newIterator creates a new paginated iterator.
func newIterator[T any](client *Client, path string, opts *ListOptions) *Iterator[T] {
	o := ListOptions{Page: 1, PerPage: 20}
	if opts != nil {
		if opts.Page > 0 {
			o.Page = opts.Page
		}
		if opts.PerPage > 0 {
			o.PerPage = opts.PerPage
		}
		o.Sort = opts.Sort
		o.Search = opts.Search
		o.Filter = opts.Filter
	}

	return &Iterator[T]{
		client: client,
		path:   path,
		opts:   o,
		page:   o.Page,
	}
}

// Next fetches the next page of results. Returns false when there are no more pages.
func (it *Iterator[T]) Next(ctx context.Context) (*ListPage[T], bool, error) {
	if it.done {
		return nil, false, nil
	}

	it.opts.Page = it.page
	var data []T
	pagination, err := it.client.doList(ctx, it.path, &it.opts, &data)
	if err != nil {
		return nil, false, err
	}

	result := &ListPage[T]{
		Data:       data,
		Pagination: pagination,
	}

	if pagination == nil || !pagination.HasNext {
		it.done = true
	} else {
		it.page++
	}

	return result, true, nil
}

// All fetches all pages and returns all items concatenated.
func (it *Iterator[T]) All(ctx context.Context) ([]T, error) {
	var all []T
	for {
		page, ok, err := it.Next(ctx)
		if err != nil {
			return nil, err
		}
		if !ok {
			break
		}
		all = append(all, page.Data...)
	}
	return all, nil
}
