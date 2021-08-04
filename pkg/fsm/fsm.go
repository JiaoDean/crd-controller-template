package fsm

import (
	"github.com/looplab/fsm"
)

// Events defines the events of fsm
var Events = fsm.Events{
	{Name: CreatingEvent, Src: []string{Pending}, Dst: Creating},
	{Name: InitializationEvent, Src: []string{Creating}, Dst: Initialization},
	{Name: RunningEvent, Src: []string{Initialization}, Dst: Running},
	{Name: TerminatingEvent, Src: []string{Pending, Creating, Initialization, Running}, Dst: Terminating},
}
