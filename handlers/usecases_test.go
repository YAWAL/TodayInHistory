package handlers

import (
	"reflect"
	"testing"

	"github.com/YAWAL/TodayInHistory/model"
)

func Test_validDate(t *testing.T) {

	tests := []struct {
		name  string
		month string
		day   string
		want  bool
	}{
		{"valid date",
			"1",
			"2",
			true,
		},
		{"invalid date",
			"14",
			"2",
			false,
		},
		{"invalid zero nums in date",
			"0",
			"0",
			false,
		},
		{"invalid month in date",
			"14",
			"1",
			false,
		},
		{"invalid day in date",
			"1",
			"32",
			false,
		},
		{"invalid day in date - negative number",
			"1",
			"-2",
			false,
		},
		{"invalid month in date - negative number",
			"-1",
			"1",
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := validDate(tt.month, tt.day); got != tt.want {
				t.Errorf("validDate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_request(t *testing.T) {

	tests := []struct {
		name            string
		url             string
		wantHistoryData model.HistoryData
		wantErr         bool
	}{
		{
			"valid url",
			"https://history.muffinlabs.com/date/9/2",
			model.HistoryData{
				URL:  "https://wikipedia.org/wiki/September_2",
				Date: "September 2",
			},
			false,
		},
		{
			"invalid url",
			"https://invalidApi.com",
			model.HistoryData{},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotHistoryData, err := request(tt.url)
			if (err != nil) != tt.wantErr {
				t.Errorf("request() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotHistoryData.URL, tt.wantHistoryData.URL) {
				t.Errorf("request() = %v, want %v", gotHistoryData.URL, tt.wantHistoryData.URL)
			}
			if !reflect.DeepEqual(gotHistoryData.Date, tt.wantHistoryData.Date) {
				t.Errorf("request() = %v, want %v", gotHistoryData.Date, tt.wantHistoryData.Date)
			}
		})
	}
}
