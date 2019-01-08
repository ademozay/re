package re

import (
	"errors"
	"testing"
	"time"
)

func TestTry(t *testing.T) {

	// means try task 5 times in a second
	const (
		period  = time.Millisecond * 200
		timeout = time.Second
	)

	task := func(currHit, targetHit int) error {
		if currHit <= targetHit {
			return errors.New("not yet")
		}
		return nil
	}

	var hit int

	tests := []struct {
		name    string
		task    func() error
		wantErr bool
	}{
		{
			name:    "nil task",
			task:    nil,
			wantErr: true,
		},
		{
			name:    "no err func",
			task:    func() error { return nil },
			wantErr: false,
		},
		{
			name: "create err 4 times",
			task: func() error {
				hit++
				return task(hit, 4)
			},
			wantErr: false,
		},
		{
			name: "create err 5 times",
			task: func() error {
				hit++
				return task(hit, 5)
			},
			wantErr: false,
		},
		{
			name: "create err 6 times",
			task: func() error {
				hit++
				return task(hit, 6)
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			hit = 0
			if err := Try(tt.task, period, timeout); (err != nil) != tt.wantErr {
				t.Errorf("Try() err: %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
