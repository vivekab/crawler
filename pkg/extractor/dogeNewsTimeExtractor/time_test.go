package dogenewstimeextractor

import (
	"reflect"
	"testing"
	"time"
)

func Test_dogeNewsTime_ToTime(t *testing.T) {

	type args struct {
		input string
	}
	tests := []struct {
		name    string
		d       *dogeNewsTime
		args    args
		wantT   time.Time
		wantErr bool
	}{
		{
			name: "SuccessCase",
			d: &dogeNewsTime{},
			args: args{
				input: "by David Cox on May 28, 2021 at 11:42 am",
			},
			wantT: time.Date(2021,time.May,28,11,42,0,0,time.UTC),
			wantErr: false,
		},
		{
			name: "SuccessCaseWithMothInOtherPlace",
			d: &dogeNewsTime{},
			args: args{
				input: "by May Cox on May 28, 2021 at 11:42 am",
			},
			wantT: time.Date(2021,time.May,28,11,42,0,0,time.UTC),
			wantErr: false,
		},
		{
			name: "SuccessCaseWithExtraSpaceAtEnd",
			d: &dogeNewsTime{},
			args: args{
				input: "by May Cox on May 28, 2021 at 11:42 am   ",
			},
			wantT: time.Date(2021,time.May,28,11,42,0,0,time.UTC),
			wantErr: false,
		},
		{
			name: "SuccessCaseWithExtraSpaceAtEnd",
			d: &dogeNewsTime{},
			args: args{
				input: "by June Cox on June 28, 2021 at 11:42 am   ",
			},
			wantT: time.Date(2021,time.June,28,11,42,0,0,time.UTC),
			wantErr: false,
		},
		{
			name: "SuccessCaseWithExtraSpaceAtEnd",
			d: &dogeNewsTime{},
			args: args{
				input: "by January Cox on January 28, 2021 at 11:42 am   ",
			},
			wantT: time.Date(2021,time.January,28,11,42,0,0,time.UTC),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &dogeNewsTime{}
			gotT, err := d.Extract(tt.args.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("dogeNewsTime.ToTime() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotT, tt.wantT) {
				t.Errorf("dogeNewsTime.ToTime() = %v, want %v", gotT, tt.wantT)
			}
		})
	}
}
