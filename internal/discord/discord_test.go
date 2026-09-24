package discord

import "testing"

func TestVersionGreater(t *testing.T) {
	cases := []struct {
		a, b string
		want bool
	}{
		{"1.0.9190", "1.0.9189", true},
		{"1.0.9200", "1.0.9199", true},
		{"1.0.9190", "1.0.9190", false},
		{"1.0.100", "1.1.0", false},
	}
	for _, tc := range cases {
		if got := versionGreater(tc.a, tc.b); got != tc.want {
			t.Fatalf("versionGreater(%q, %q) = %v, want %v", tc.a, tc.b, got, tc.want)
		}
	}
}
