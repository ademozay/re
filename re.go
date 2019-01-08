package re

import (
	"errors"
	"time"
)

// Try, at first, executes given task with no timer started.
// If an error occurs, it starts a Ticker and a Timer and executes
// the task every time Ticker ticks unless Timer is finish.
// If Timer finishes, Try returns the latest error returned from the task.
func Try(task func() error, period, timeout time.Duration) error {

	if task == nil {
		return errors.New("task is nil")
	}

	if err := task(); err == nil {
		return nil
	}

	ticker := time.NewTicker(period)
	timer := time.NewTimer(timeout)

	defer func() {
		ticker.Stop()
		timer.Stop()
	}()

	var err error

	for {
		select {
		case <-ticker.C:
			if err = task(); err != nil {
				continue
			}
			return nil
		case <-timer.C:
			return err
		}
	}

}
