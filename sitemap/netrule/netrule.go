package netrule

import (
	//"net/http"
	"sync"
	"time"
)

type NetAccessRule interface {
	// Asynchronously checks and sends a signal on AccessChannel when it is free to access
	// Returns a channel that is closed when it is free to access
	AccessChannel() <-chan struct{}
	// Accessed is called after every net access
	Accessed()
}

// One rule cannot be used across multiple instances that read from the channel
type TimeAccessRule struct {
	access_chan   chan struct{}
	timer         *time.Timer
	duration      time.Duration
	run_lock      sync.Mutex
	running_check bool
}

func NewTimeAccessRule(d time.Duration) *TimeAccessRule {
	rule := &TimeAccessRule{}
	rule.access_chan = make(chan struct{}, 1)
	rule.duration = d
	rule.timer = time.NewTimer(d)
	return rule
}

func (rule *TimeAccessRule) AccessChannel() <-chan struct{} {
	go rule.sendOnTimeout()
	return rule.access_chan
}

func (rule *TimeAccessRule) sendOnTimeout() {
	rule.run_lock.Lock()
	if !rule.running_check {
		rule.running_check = true
		<-rule.timer.C
		rule.timer.Reset(rule.duration)
		close(rule.access_chan)
		rule.access_chan = make(chan struct{}, 1)
		rule.running_check = false
	}
	rule.run_lock.Unlock()
}

// Accessed is called after every net access
func (rule *TimeAccessRule) Accessed() {
	rule.timer.Reset(rule.duration)
}

func (rule *TimeAccessRule) SetDelay(d time.Duration) {
	rule.duration = d
	rule.timer.Reset(d)
}
