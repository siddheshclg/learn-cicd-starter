package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	testCases := []struct{
		input string 
		want string 
		err  bool
		set bool
	}{
		{
			input: "ApiKey 434233342",
			want: "434233342",
			set: true,
		},
		{
			input: "Apikey 24343",
			err: true,
			set: true,
		},
		{
			set: false,
			err: true,
		},
	}
	for _, test := range testCases {
		h := http.Header{} 
		if test.set {
			h.Set("Authorization", test.input)
		}
		got, err := GetAPIKey(h)
		if test.err {
			if err == nil{
				t.Error("Expected error")
			}
		} else {
			if err != nil {
				t.Errorf("Got unexpected error: %v", err)
			} else if test.want != got {
				t.Errorf("Want: %v, Got: %v", test.want, got)
			}
		}
	}
}