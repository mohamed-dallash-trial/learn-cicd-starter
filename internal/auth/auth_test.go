package auth

import (
	"errors"
	"net/http"
	"reflect"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	tests := map[string]struct {
		input     http.Header
		wantValue string
		wantErr   error
	}{
		"normal": {
			input:     http.Header{"Authorization": {"ApiKey abcdef"}},
			wantValue: "abcdef",
			wantErr:   nil,
		},
		"no auth header": {
			input:     http.Header{},
			wantValue: "",
			wantErr:   ErrNoAuthHeaderIncluded,
		},
		"malformed auth header": {
			input:     http.Header{"Authorization": {"ApiKEY abcdef"}},
			wantValue: "",
			wantErr:   errors.New("malformed authorization header"),
		},
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			gotValue, gotErr := GetAPIKey(tc.input)
			if gotValue != tc.wantValue || !reflect.DeepEqual(gotErr, tc.wantErr) {
				//note that using DeepEqual for errors is bad
				t.Fatalf("\nexpected: %v, %v\ngot: %v, %v", tc.wantValue, tc.wantErr, gotValue, gotErr)
			}
		})
	}

}
