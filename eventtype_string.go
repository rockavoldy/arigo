// Code generated by "stringer -type=EventType"; DO NOT EDIT.

package arigo

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[StartEvent-0]
	_ = x[PauseEvent-1]
	_ = x[StopEvent-2]
	_ = x[CompleteEvent-3]
	_ = x[BTCompleteEvent-4]
	_ = x[ErrorEvent-5]
}

const _EventType_name = "StartEventPauseEventStopEventCompleteEventBTCompleteEventErrorEvent"

var _EventType_index = [...]uint8{0, 10, 20, 29, 42, 57, 67}

func (i EventType) String() string {
	if i >= EventType(len(_EventType_index)-1) {
		return "EventType(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _EventType_name[_EventType_index[i]:_EventType_index[i+1]]
}