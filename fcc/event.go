//
// Copyright(c) 2017 Uli Fuchs <ufuchs@gmx.com>
// MIT Licensed
//

// [ The true delight is in the finding out rather than in the knowing.         ]
// [                                                             -Isaac Asimov- ]

package fcc

import "time"

type (
	Event struct {
		Num      int
		Interval time.Duration
		Min      int
		Max      int
		OnChange []int
	}
)

//
//
//
func NewEventList(device []Device) []Event {

	events := []Event{}

	// for _, d := range deviceAttr {

	//  e := &Event{
	//      Num:      d.Num,
	//      Interval: 240,
	//  }

	//  events[e.Num] = e

	// }

	return events

}
