package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	tests := map[string]struct {
		headers http.Header
		want    string
		wantErr bool
	}{
		"Simple": {
			headers: makeHeaders(map[string]string{"Authorization": "ApiKey my-key"}),
			want:    "my-key",
			wantErr: false,
		},
		"No Authorization header": {
			headers: makeHeaders(map[string]string{"Content-Type": "text/json"}),
			want:    "",
			wantErr: true,
		},
		"Malformed Authorization header": {
			headers: makeHeaders(map[string]string{"Authorization": "Api Key my-key"}),
			want:    "",
			wantErr: true,
		},
		"Forced failure": {
			headers: makeHeaders(map[string]string{"Authorization": "Nope"}),
			want:    "Value that can never exist",
			wantErr: false,
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got, err := GetAPIKey(tc.headers)
			if err != nil && tc.wantErr == false {
				t.Fatalf("%s: unexpected error: %v", name, err)
			}
			if got != tc.want {
				t.Fatalf("%s: expected: %v, got: %v", name, tc.want, got)
			}
			if err == nil && tc.wantErr == true {
				t.Fatalf("%s: expected error not found", name)
			}
		})
	}
}

func makeHeaders(m map[string]string) http.Header {
	h := http.Header{}
	for k, v := range m {
		h.Set(k, v)
	}
	return h
}
