package link

// Link offers functionality for linked HCP resources.
type Link interface {
	// Start will expose Link functionality to the control-plane. The SCADAProvider
	// used by Link will need to be started separately.
	Start() error

	// Stop will stop exposing Link functionality to the control-plane.
	Stop() error
}
