package writer

import (
	"regexp"
	"testing"
)

func TestAssignBranchPrefix(t *testing.T) {

	re_prefix, err := regexp.Compile(`^(\d+)\-[0-9a-f]{8}-[0-9a-f]{4}\-4[0-9a-f]{3}\-[89ab][0-9a-f]{3}\-[0-9a-f]{12}$`)

	if err != nil {
		t.Fatalf("Failed to compile prefix regular expression, %v", err)
	}

	tests_assign := []string{
		"{prefix}-test",
	}

	tests_skip := []string{
		"test",
	}

	for _, b := range tests_assign {

		new_b := AssignBranchPrefix(b)

		if b == new_b {
			t.Fatalf("Failed to assign prefix to '%s'", b)
		}

		if !re_prefix.MatchString(new_b) {
			t.Fatalf("New branch failed prefix regular expression, %s", new_b)
		}
	}

	for _, b := range tests_skip {

		new_b := AssignBranchPrefix(b)

		if b != new_b {
			t.Fatalf("Failed to skip prefix from '%s'", b)
		}
	}

}
