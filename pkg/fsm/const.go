package fsm

// FSM enter status
const (
	Pending        = "Pending"
	Creating       = "Creating"
	Initialization = "Initialization"
	Running        = "Running"
	Terminating    = "Terminating"
)

// FSM event
const (
	CreatingEvent       = "creatingEvent"
	InitializationEvent = "initializationEvent"
	RunningEvent        = "runningEvent"
	TerminatingEvent    = "terminatingEvent"
)
