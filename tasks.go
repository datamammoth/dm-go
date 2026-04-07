package datamammoth

import "context"

// TasksService handles communication with the task-related endpoints of the
// DataMammoth API v2.
type TasksService struct {
	client *Client
}

// List returns a paginated list of async tasks.
//
// API: GET /tasks
func (s *TasksService) List(ctx context.Context, opts *ListOptions) ([]Task, *Pagination, error) {
	var tasks []Task
	pagination, err := s.client.doList(ctx, "/tasks", opts, &tasks)
	if err != nil {
		return nil, nil, err
	}
	return tasks, pagination, nil
}

// ListAll returns an iterator over all pages of tasks.
func (s *TasksService) ListAll(opts *ListOptions) *Iterator[Task] {
	return newIterator[Task](s.client, "/tasks", opts)
}

// Get retrieves a task by ID.
//
// API: GET /tasks/:id
func (s *TasksService) Get(ctx context.Context, id string) (*Task, error) {
	var task Task
	_, err := s.client.get(ctx, "/tasks/"+id, &task)
	if err != nil {
		return nil, err
	}
	return &task, nil
}

// Cancel cancels a pending or running task.
//
// API: DELETE /tasks/:id
func (s *TasksService) Cancel(ctx context.Context, id string) (*CancelledTask, error) {
	var result CancelledTask
	_, err := s.client.del(ctx, "/tasks/"+id, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
