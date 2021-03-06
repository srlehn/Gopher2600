// This file is part of Gopher2600.
//
// Gopher2600 is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// Gopher2600 is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with Gopher2600.  If not, see <https://www.gnu.org/licenses/>.
//
// *** NOTE: all historical versions of this file, as found in any
// git repository, are also covered by the licence, even when this
// notice is not present ***

package future

import (
	"fmt"
	"strings"
)

// Event represents a single occurance of a payload sometime in the future
type Event struct {
	// the future ticker this event belongs to
	tck *Ticker

	// label is a short decription describing the future payload
	label string

	// the number of cycles the event began with
	initialCycles int

	// the number of remaining ticks before the pending action is resolved
	remainingCycles int

	// temporary cessation of ticks
	paused bool

	// completion of the event has been pushed back at least once
	pushed bool

	// the value that is to be the result of the pending action
	payload func()

	payloadWithArg func(interface{})
	payloadArg     interface{}
}

func (ev Event) String() string {
	label := strings.TrimSpace(ev.label)
	if label == "" {
		label = "[unlabelled event]"
	}
	if ev.remainingCycles == -1 {
		return fmt.Sprintf("%s -> imminent", label)
	}
	return fmt.Sprintf("%s -> %d", label, ev.remainingCycles)
}

func (ev *Event) isActive() bool {
	return ev.remainingCycles >= 0
}

func (ev *Event) runPayload() {
	if ev.payloadWithArg != nil {
		ev.payloadWithArg(ev.payloadArg)
	} else {
		ev.payload()
	}
}

// Tick event forward one cycle
func (ev *Event) tick() bool {
	if !ev.isActive() {
		panic("events should not be ticked once they have expired under any circumstances")
	}

	if ev.paused {
		return false
	}

	ev.remainingCycles--

	if ev.remainingCycles == -1 {
		ev.runPayload()

		// tick() should only be called tick the Ticker.Tick() function. moving
		// the event to the inactive list happens there, rather than here.

		return true
	}

	return false
}

// RemainingCycles reports the number of cycles remaining before payload
// function is ran
func (ev Event) RemainingCycles() int {
	return ev.remainingCycles
}

// Force can be used to immediately run the event's payload
//
// It is very important that any references to the event be forgotten once
// Force() has been called
func (ev *Event) Force() {
	if !ev.isActive() {
		panic("cannot do that to a completed event")
	}

	ev.runPayload()
	ev.tck.drop(ev)
	ev.remainingCycles = -1
}

// Drop is be used to remove the event from the ticker queue without executing
// the payload. Because the payload is not run then you should be careful to
// handle any cleanup that might otherwise occur (in the payload).
//
// It is very important that any references to the event be forgotten once
// Drop() has been called
func (ev *Event) Drop() {
	if !ev.isActive() {
		panic("cannot do that to a completed event")
	}
	ev.tck.drop(ev)
	ev.remainingCycles = -1
}

// Push back event completion by restarting the event
func (ev *Event) Push() {
	if !ev.isActive() {
		panic("cannot do that to a completed event")
	}
	ev.remainingCycles = ev.initialCycles
	ev.pushed = true
}

// Pause prevents future calls to Tick() having any effect (until Resume() is
// called)
func (ev *Event) Pause() {
	if !ev.isActive() {
		panic("cannot do that to a completed event")
	}
	ev.paused = true
}

// JustStarted is true if the Tick() function for the event has not yet been
// called
func (ev Event) JustStarted() bool {
	if !ev.isActive() {
		panic("cannot do that to a completed event")
	}
	return ev.remainingCycles == ev.initialCycles && !ev.pushed
}

// AboutToEnd is true if event resolves on next Tick()
func (ev *Event) AboutToEnd() bool {
	// * optimisation: called a lot. pointer to Event to prevent duffcopy

	if !ev.isActive() {
		panic("cannot do that to a completed event")
	}
	return ev.remainingCycles == 0
}
