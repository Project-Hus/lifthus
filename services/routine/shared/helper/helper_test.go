package helper

import "testing"

func TestSlugify(t *testing.T) {
	// test slugify
	t.Run("test slugify", func(t *testing.T) {
		// test cases
		testCases := []struct {
			name     string
			input    string
			expected string
		}{
			{
				"test slugify with space",
				"test slugify",
				"test-slugify",
			},
			{
				"test slugify with special characters",
				"test@slugify",
				"test%40slugify",
			},
			{
				"test slugify with special characters",
				"test#slugify",
				"test%23slugify",
			},
		}

		// run test cases
		for _, tc := range testCases {
			t.Run(tc.name, func(t *testing.T) {
				actual := Slugify(tc.input)
				if actual != tc.expected {
					t.Errorf("expected: %s, actual: %s", tc.expected, actual)
				}
			})
		}
	})
}
