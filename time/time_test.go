package util

import (
	"reflect"
	"testing"
	"time"
)

func parseTimeString(str string) time.Time {
	t, _ := time.Parse("2006-01-02 15:04:05", str)
	return t
}

func TestEndTimeOfDay(t *testing.T) {
	type args struct {
		t time.Time
	}
	tests := []struct {
		name string
		args args
		want time.Time
	}{
		// TODO: Add test cases.
		{
			args: args{
				t: parseTimeString("2020-12-01 12:00:12"),
			},
			want: parseTimeString("2020-12-02 00:00:00"),
		},
		{
			args: args{
				t: parseTimeString("2020-12-01 00:00:00"),
			},
			want: parseTimeString("2020-12-02 00:00:00"),
		},
		{
			args: args{
				t: parseTimeString("2020-02-29 05:00:00"),
			},
			want: parseTimeString("2020-03-01 00:00:00"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := EndTimeOfDay(tt.args.t); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("EndTimeOfDay() = %v, want %v", got, tt.want)
			}
		})
	}
}
