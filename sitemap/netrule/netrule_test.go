// netrule_test.go
package netrule

import (
	"errors"
	"fmt"
	"sync"
	"testing"
	"time"

	"github.com/stretchrcom/testify/assert"
)

/*
func Test_(t *testing.T) {

}

func Benchmark_(b *testing.B) {

}
*/

func Test_TimeAccessRule_SingleNotifyCase(t *testing.T) {
	time_access := NewTimeAccessRule(time.Millisecond * 100)

	err := test_AccessChannel(t, "expected false from ok Case 1", time_access, time.Second*10)
	assert.Nil(t, err, fmt.Sprintf("expected nil error from access_channel test Case 1"))
}

func Test_TimeAccessRule_MultipleNotifyCase(t *testing.T) {
	time_access := NewTimeAccessRule(time.Millisecond * 100)
	var done sync.WaitGroup
	done.Add(3)
	test1 := func() {
		err := test_AccessChannel(t, "expected false from ok Case 1", time_access, time.Second*10)
		assert.Nil(t, err, fmt.Sprintf("expected nil error from access_channel test Case 1"))
		done.Done()
	}
	test2 := func() {
		err := test_AccessChannel(t, "expected false from ok Case 2", time_access, time.Second*10)
		assert.Nil(t, err, fmt.Sprintf("expected nil error from access_channel test Case 2"))
		done.Done()
	}
	test3 := func() {
		err := test_AccessChannel(t, "expected false from ok Case 3", time_access, time.Second*10)
		assert.Nil(t, err, fmt.Sprintf("expected nil error from access_channel test Case 3"))
		done.Done()
	}
	go test1()
	go test2()
	go test3()
	done.Wait()
}

func test_AccessChannel(t *testing.T, err_msg string, timer NetAccessRule, timeout time.Duration) (err error) {
	timeout_chan := make(chan struct{}, 1)
	timeout_func := func() {
		time.Sleep(timeout)
		close(timeout_chan)

	}
	go timeout_func()
	i := 0
	for i < 5 {
		access_chan := timer.AccessChannel()
		break_loop := false
		for !break_loop {
			select {
			case <-timeout_chan:
				return errors.New(fmt.Sprintf("Access channel timed out at iteration %i", i))
			case _, ok := <-access_chan:
				assert.False(t, ok, err_msg)
				i++
				break_loop = true
			}
		}
	}
	return nil
}
