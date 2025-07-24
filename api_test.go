package writer

/*

go test -v -run TestAPIWriter -writer-uri 'githubapi-branch://sfomuseum-data/sfomuseum-data-test-2021?to-branch={prefix}-test&access_token={ACCESS_TOKEN}&author=thisisaaronland&email=thisisaaronland@localhost' -uri debug/test2.txt

*/

import (
	"context"
	"flag"
	"fmt"
	"log/slog"
	"strings"
	"testing"
	"time"

	"github.com/whosonfirst/go-ioutil"
	"github.com/whosonfirst/go-writer/v3"
)

var writer_uri = flag.String("writer-uri", "", "A valid githubapi:// (go-writer-github) URI.")
var uri = flag.String("uri", "", "The URI to write your file to.")

func TestAPIWriter(t *testing.T) {

	ctx := context.Background()

	if *writer_uri == "" {
		slog.Info("-writer-uri flag is empty, skipping test.")
		t.Skip()
	}

	if *uri == "" {
		slog.Info("-uri flag is empty, skipping test.")
		t.Skip()
	}

	wr, err := writer.NewWriter(ctx, *writer_uri)

	if err != nil {
		t.Fatalf("Failed to create writer, %v", err)
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

	err = wr.Close(ctx)

	if err != nil {
		t.Fatalf("Failed to close writer, %v", err)
	}
}
