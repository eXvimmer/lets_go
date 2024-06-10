package main

import (
	"testing"
	"time"

	"github.com/eXvimmer/lets_go/internal/assert"
)

func TestHumanDate(t *testing.T) {
	tests := []struct {
		name string
		tm   time.Time
		want string
	}{
		{
			name: "UTC",
			tm:   time.Date(2022, 3, 17, 10, 15, 0, 0, time.UTC),
			want: "17 Mar 2022 at 10:15:00",
		},
		{
			name: "Empty",
			tm:   time.Time{},
			want: "",
		},
		{
			name: "CET",
			tm:   time.Date(2022, 3, 17, 10, 15, 0, 0, time.FixedZone("CET", 1*60*60)),
			want: "17 Mar 2022 at 09:15:00",
		},
	}
	for _, tt := range tests {
		hd := humanDate(tt.tm)
		if hd != tt.want {
			assert.Equal(t, hd, tt.want)
		}
	}
}
