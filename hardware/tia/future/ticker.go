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
	"container/list"
	"strings"
)

// the number of elements in the list of events. through observation it has
// been determined that there is never more than 6 elements per ticker queue
// required at any one time. an additional element is needed as the sentry
// element
const poolSize = 7

// Ticker is used to group payloads for future triggering.
type Ticker struct {
	Label          string
	pool           list.List
	activeSentinal *list.Element
}

// NewTicker is the only method of initialisation for the Ticker type
func NewTicker(label string) *Ticker {
	tck := &Ticker{Label: label}

	// push empty elements into the pool
	for i := 0; i < poolSize; i++ {
		tck.pool.PushBack(&Event{tck: tck, remainingCycles: -1})
	}

	// the pool begins with no active elements. the active sentinal is
	// therefore at the very front of the pool
	tck.activeSentinal = tck.pool.Front()

	return tck
}

func (tck Ticker) String() string {
	s := strings.Builder{}
	for e := tck.pool.Front(); e != nil; e = e.Next() {
		if e == tck.activeSentinal {
			break
		}
		s.WriteString(tck.Label)
		s.WriteString(": ")
		s.WriteString(e.Value.(*Event).String())
		s.WriteString("\n")
	}
	return strings.TrimRight(s.String(), "\n")
}

// Tick moves the pending action counter on one step
func (tck *Ticker) Tick() bool {
	// return false unless at least one event has been succesfully ticked
	r := false

	// walk through the pool until we reach the active sentinal
	e := tck.pool.Front()
	for e != nil && e != tck.activeSentinal {
		if e.Value.(*Event).tick() {
			// an event has ticked. update return value
			r = true

			// take note of next element
			n := e.Next()

			// move to front of inactive queue (after the active sentinal)
			tck.pool.MoveToBack(e)

			// continue for loop with next element noted above. we cannot just
			// use e.Next() because we've moved it in the list so e.Next() will
			// not return the node we expect
			e = n
		} else {
			e = e.Next()
		}
	}

	return r
}

func (tck *Ticker) drop(ev *Event) {
	// walk through the pool until we reach the active sentinal
	e := tck.pool.Front()
	for e != nil && e != tck.activeSentinal {
		if ev == e.Value.(*Event) {
			// move to front of inactive queue (after the active sentinal)
			tck.pool.MoveToBack(e)

			// event has been found so we can return. job done.
			return
		}

		e = e.Next()
	}

	// we should never reach this point. if we do something has gone very, very
	// wrong. it is okay to panic.
	panic("cannot drop an event that is not in the list of active events")
}
