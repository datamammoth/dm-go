package datamammoth

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
)

// ServersService handles communication with the server-related endpoints of the
// DataMammoth API v2.
type ServersService struct {
	client *Client
}

// List returns a paginated list of servers.
//
// API: GET /servers
func (s *ServersService) List(ctx context.Context, opts *ListOptions) ([]Server, *Pagination, error) {
	var servers []Server
	pagination, err := s.client.doList(ctx, "/servers", opts, &servers)
	if err != nil {
		return nil, nil, err
	}
	return servers, pagination, nil
}

// ListAll returns an iterator over all pages of servers.
func (s *ServersService) ListAll(opts *ListOptions) *Iterator[Server] {
	return newIterator[Server](s.client, "/servers", opts)
}

// Get retrieves a single server by ID.
//
// API: GET /servers/:id
func (s *ServersService) Get(ctx context.Context, id string) (*Server, error) {
	var server Server
	_, err := s.client.get(ctx, "/servers/"+id, &server)
	if err != nil {
		return nil, err
	}
	return &server, nil
}

// Create provisions a new server. This is an async operation that returns a task.
//
// API: POST /servers
func (s *ServersService) Create(ctx context.Context, params *CreateServerParams) (*TaskAccepted, error) {
	var task TaskAccepted
	_, err := s.client.post(ctx, "/servers", params, &task)
	if err != nil {
		return nil, err
	}
	return &task, nil
}

// Update modifies a server's hostname or label.
//
// API: PATCH /servers/:id
func (s *ServersService) Update(ctx context.Context, id string, params *UpdateServerParams) (*Server, error) {
	var server Server
	_, err := s.client.patch(ctx, "/servers/"+id, params, &server)
	if err != nil {
		return nil, err
	}
	return &server, nil
}

// Delete terminates a server. This is an async operation that returns a task.
//
// API: DELETE /servers/:id
func (s *ServersService) Delete(ctx context.Context, id string) (*TaskAccepted, error) {
	var task TaskAccepted
	_, err := s.client.del(ctx, "/servers/"+id, &task)
	if err != nil {
		return nil, err
	}
	return &task, nil
}

// PowerOn starts a server.
//
// API: POST /servers/:id/actions/power-on
func (s *ServersService) PowerOn(ctx context.Context, id string) (*ActionResult, error) {
	return s.doAction(ctx, id, "power-on", nil)
}

// PowerOff stops a server (hard power off).
//
// API: POST /servers/:id/actions/power-off
func (s *ServersService) PowerOff(ctx context.Context, id string) (*ActionResult, error) {
	return s.doAction(ctx, id, "power-off", nil)
}

// Reboot reboots a server.
//
// API: POST /servers/:id/actions/reboot
func (s *ServersService) Reboot(ctx context.Context, id string) (*ActionResult, error) {
	return s.doAction(ctx, id, "reboot", nil)
}

// Shutdown gracefully shuts down a server.
//
// API: POST /servers/:id/actions/shutdown
func (s *ServersService) Shutdown(ctx context.Context, id string) (*ActionResult, error) {
	return s.doAction(ctx, id, "shutdown", nil)
}

// Rebuild rebuilds a server with a new OS image.
//
// API: POST /servers/:id/actions/rebuild
func (s *ServersService) Rebuild(ctx context.Context, id string, params *RebuildParams) (*ActionResult, error) {
	return s.doAction(ctx, id, "rebuild", params)
}

// Rescue boots a server into rescue mode.
//
// API: POST /servers/:id/actions/rescue
func (s *ServersService) Rescue(ctx context.Context, id string, params *RescueParams) (*ActionResult, error) {
	return s.doAction(ctx, id, "rescue", params)
}

func (s *ServersService) doAction(ctx context.Context, id, action string, body interface{}) (*ActionResult, error) {
	path := fmt.Sprintf("/servers/%s/actions/%s", id, action)
	var result ActionResult
	_, err := s.client.do(ctx, http.MethodPost, path, body, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// ListSnapshots returns all snapshots for a server.
//
// API: GET /servers/:id/snapshots
func (s *ServersService) ListSnapshots(ctx context.Context, id string) ([]Snapshot, error) {
	var snapshots []Snapshot
	_, err := s.client.get(ctx, "/servers/"+id+"/snapshots", &snapshots)
	if err != nil {
		return nil, err
	}
	return snapshots, nil
}

// CreateSnapshot creates a new snapshot of a server.
//
// API: POST /servers/:id/snapshots
func (s *ServersService) CreateSnapshot(ctx context.Context, id string, params *CreateSnapshotParams) (*Snapshot, error) {
	var snapshot Snapshot
	_, err := s.client.post(ctx, "/servers/"+id+"/snapshots", params, &snapshot)
	if err != nil {
		return nil, err
	}
	return &snapshot, nil
}

// DeleteSnapshot deletes a snapshot.
//
// API: DELETE /servers/:id/snapshots/:snapId
func (s *ServersService) DeleteSnapshot(ctx context.Context, serverID, snapshotID string) (*SnapshotActionResult, error) {
	var result SnapshotActionResult
	_, err := s.client.del(ctx, fmt.Sprintf("/servers/%s/snapshots/%s", serverID, snapshotID), &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// RestoreSnapshot restores (rolls back) a snapshot on a server.
//
// API: POST /servers/:id/snapshots/:snapId
func (s *ServersService) RestoreSnapshot(ctx context.Context, serverID, snapshotID string) (*SnapshotActionResult, error) {
	var result SnapshotActionResult
	_, err := s.client.post(ctx, fmt.Sprintf("/servers/%s/snapshots/%s", serverID, snapshotID), nil, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// GetMetrics retrieves server metrics.
//
// API: GET /servers/:id/metrics
func (s *ServersService) GetMetrics(ctx context.Context, id string, opts *MetricsOptions) (*MetricsResult, error) {
	params := url.Values{}
	if opts != nil {
		if opts.Period != "" {
			params.Set("period", opts.Period)
		}
		if opts.Source != "" {
			params.Set("source", opts.Source)
		}
		if opts.Limit > 0 {
			params.Set("limit", strconv.Itoa(opts.Limit))
		}
	}

	path := "/servers/" + id + "/metrics"
	if q := params.Encode(); q != "" {
		path += "?" + q
	}

	var result MetricsResult
	_, err := s.client.get(ctx, path, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// ListEvents returns the audit log / event history for a server.
//
// API: GET /servers/:id/events
func (s *ServersService) ListEvents(ctx context.Context, id string, opts *ListOptions) ([]Event, *Pagination, error) {
	var events []Event
	pagination, err := s.client.doList(ctx, "/servers/"+id+"/events", opts, &events)
	if err != nil {
		return nil, nil, err
	}
	return events, pagination, nil
}

// GetConsole retrieves console access (VNC/SSH) for a server.
//
// API: GET /servers/:id/console
func (s *ServersService) GetConsole(ctx context.Context, id string) (*ConsoleAccess, error) {
	var console ConsoleAccess
	_, err := s.client.get(ctx, "/servers/"+id+"/console", &console)
	if err != nil {
		return nil, err
	}
	return &console, nil
}

// GetFirewall retrieves the firewall configuration for a server.
//
// API: GET /servers/:id/firewall
func (s *ServersService) GetFirewall(ctx context.Context, id string) (*FirewallConfig, error) {
	var fw FirewallConfig
	_, err := s.client.get(ctx, "/servers/"+id+"/firewall", &fw)
	if err != nil {
		return nil, err
	}
	return &fw, nil
}

// UpdateFirewall replaces the firewall rules for a server.
//
// API: PUT /servers/:id/firewall
func (s *ServersService) UpdateFirewall(ctx context.Context, id string, params *UpdateFirewallParams) (*FirewallConfig, error) {
	var fw FirewallConfig
	_, err := s.client.put(ctx, "/servers/"+id+"/firewall", params, &fw)
	if err != nil {
		return nil, err
	}
	return &fw, nil
}
