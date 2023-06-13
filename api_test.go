package writer

import (
	"context"
	"flag"
	"fmt"
	"strings"
	"testing"
	"time"

	"github.com/whosonfirst/go-ioutil"
	"github.com/whosonfirst/go-writer/v3"
)

var source = flag.String("source", "", "A valid githubapi:// (go-writer-github) URI.")
var uri = flag.String("uri", "", "The URI to write your file to.")

func TestAPIWriter(t *testing.T) {

	ctx := context.Background()

	if *source == "" {
		t.Skip()
	}

	if *uri == "" {
		t.Skip()
	}

	wr, err := writer.NewWriter(ctx, *source)

	if err != nil {
		t.Fatal(err)
	}

	now := time.Now()
	msg := fmt.Sprintf("This is a test: %v", now)

	br := strings.NewReader(msg)
	fh, err := ioutil.NewReadSeekCloser(br)

	if err != nil {
		t.Fatalf("Failed to create new io.ReadSeekCloser, %v", err)
	}

	_, err = wr.Write(ctx, *uri, fh)

	if err != nil {
		t.Fatalf("Failed to write %s, %v", *uri, err)
	}
}
