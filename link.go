package link

import (
	"fmt"
	"net"
	"sync"
)

const (
	metaDataNodeId      = "link.node_id"
	metaDataNodeVersion = "link.node_version"
	capabilityLink      = "link"
)

type link struct {
	// Config contains all dependencies as well as information about the node Link is
	// running on.
	*Config

	// listener is the listener of the Link SCADA capability.
	listener net.Listener

	// running is set true if Link is running
	running     bool
	runningLock sync.Mutex
}

// New creates a new instance of a Link interface, that allows access to
// functionality of linked HCP services.
func New(config *Config) (Link, error) {
	if config == nil {
		return nil, fmt.Errorf("failed to initialize link library: config must be provided")
	}

	return &link{
		Config: config,
	}, nil
}

// Start implements Link interface.
//
// It will set the Link specific meta-data values and expose the Link specific capability.
func (l *link) Start() error {
	l.runningLock.Lock()
	defer l.runningLock.Unlock()

	// Check if Link is already running
	if l.running {
		return nil
	}

	// Configure Link specific meta-data
	l.ScadaProvider.SetMetaValue(metaDataNodeId, l.NodeID)
	l.ScadaProvider.SetMetaValue(metaDataNodeVersion, l.NodeVersion)

	// Start listening on Link capability
	listener, err := l.ScadaProvider.Listen(capabilityLink)
	if err != nil {
		return fmt.Errorf("failed to start listening on the %q capability: %w", capabilityLink, err)
	}

	// TODO: Handle requests
	l.listener = listener

	// Mark Link as running
	l.running = true

	return nil
}

// Stop implements Link interface.
//
// It will unset the Link specific meta-data value and stop to expose the Link capability.
func (l *link) Stop() error {
	l.runningLock.Lock()
	defer l.runningLock.Unlock()

	// Check if Link is already stopped
	if !l.running {
		return nil
	}

	// Clear Link specific meta-data
	l.ScadaProvider.SetMetaValue(metaDataNodeId, "")
	l.ScadaProvider.SetMetaValue(metaDataNodeVersion, "")

	// Stop listening on the Link capability
	err := l.listener.Close()
	if err != nil {
		return fmt.Errorf("failed to close listener for %q capability: %w", capabilityLink, err)
	}

	// Reset listener
	l.listener = nil

	// Mark Link as stopped
	l.running = false

	return nil
}
