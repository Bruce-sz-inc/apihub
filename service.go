package apihub

import "time"

//go:generate counterfeiter . Service

type Service interface {
	// Handle returns the subdomain/host used to access a service.
	Handle() string

	// Start adds a service in the service pool to handle upcoming requests.
	Start() error

	// Stop stops proxying the requests.
	//
	// If kill is false, Apihub stops proxying the requests to one of the backends
	// registered.
	//
	// If kill is true, Apihub stops proxing the requests and remove the service
	// from the service pool.
	Stop(kill bool) error

	// Info returns information about a service.
	Info() (ServiceSpec, error)

	// Addbackend adds a new backend in the list of available be's.
	AddBackend(be BackendInfo) error

	// RemoveBackend removes an existing backend from the list of available be's.
	RemoveBackend(be BackendInfo) error

	// Backends returns all backends in the service.
	Backends() ([]BackendInfo, error)

	// Timeout waits for the duration before returning an error to the client.
	SetTimeout(time.Duration)
}

// ServiceInfo holds information about a service.
type ServiceSpec struct {
	// Handle specifies the subdomain/host used to access the service.
	Handle   string        `json:"handle"`
	Disabled bool          `json:"disabled"`
	Timeout  int           `json:"timeout"`
	Backends []BackendInfo `json:"backends,omitempty"`
}

// Backend holds information about a backend.
type BackendInfo struct {
	Address          string `json:"address"`
	Disabled         bool   `json:"disabled"`
	HeartBeatAddress string `json:"heart_beat_address"`
	HeartBeatTimeout int    `json:"heart_beat_timeout"`
}
