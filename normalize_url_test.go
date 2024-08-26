package main

import "testing"

func TestNormalizeURL(t *testing.T) {
	tests := []struct {
		name     string
		inputURL string
		expected string
	}{
		{
			name:     "remove scheme",
			inputURL: "https://blog.boot.dev/path",
			expected: "blog.boot.dev/path",
		},
		// add more test cases here
		{
			name:     "https with /path/",
			inputURL: "https://blog.boot.dev/path/",
			expected: "blog.boot.dev/path",
		},
		{
			name:     "http with /path",
			inputURL: "http://blog.boot.dev/path",
			expected: "blog.boot.dev/path",
		},
		{
			name:     "http with /path/",
			inputURL: "http://blog.boot.dev/path/",
			expected: "blog.boot.dev/path",
		},
		{
			name:     "http with www with /path",
			inputURL: "http://www.blog.boot.dev/path",
			expected: "blog.boot.dev/path",
		},
		{
			name:     "http with www with /path/",
			inputURL: "http://www.blog.boot.dev/path/",
			expected: "blog.boot.dev/path",
		},
		{
			name:     "https with www with /path",
			inputURL: "https://www.blog.boot.dev/path",
			expected: "blog.boot.dev/path",
		},
		{
			name:     "https with www with /path/",
			inputURL: "https://www.blog.boot.dev/path/",
			expected: "blog.boot.dev/path",
		},
	}

	for i, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			actual := normalizeURL(tc.inputURL)
			// if err != nil {
			// 	t.Errorf("Test %v - '%s' FAIL: unexpected error: %v", i, tc.name, err)
			// 	return
			// }
			if actual != tc.expected {
				t.Errorf("Test %v - %s FAIL: expected URL: %v, actual: %v", i, tc.name, tc.expected, actual)
			}
		})
	}
}
