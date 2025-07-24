package writer

import (
	"fmt"
	"strings"
	"time"

	"github.com/google/uuid"
)

const BRANCH_UUID_PREFIX string = "{prefix}-"

func AssignBranchPrefix(branch string) string {

	if !strings.HasPrefix(branch, BRANCH_UUID_PREFIX) {
		return branch
	}

	id := uuid.New()
	now := time.Now()

	prefix := fmt.Sprintf("%d-%s-", now.Unix(), id.String())
	return strings.Replace(branch, BRANCH_UUID_PREFIX, prefix, 1)
}
