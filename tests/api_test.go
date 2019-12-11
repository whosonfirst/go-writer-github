package writer

import (
	"context"
	"flag"
	"github.com/whosonfirst/go-writer"
	_ "github.com/whosonfirst/go-writer-github"		
	"io/ioutil"
	"testing"
	"strings"
	"time"
	"fmt"
)

var source = flag.String("source", "", "...")
var uri = flag.String("uri", "", "...")

func TestAPIWriter(t *testing.T) {

	ctx := context.Background()

	if *source == "" {
		t.Fatal("Missing -source parameter")
	}

	if *uri == "" {
		t.Fatal("Missing -uri parameter")
	}
		
	wr, err := writer.NewWriter(ctx, *source)

	if err != nil {
		t.Fatal(err)
	}

	now := time.Now()
	msg := fmt.Sprintf("This is a test: %v", now)
	
	br := strings.NewReader(msg)
	fh := ioutil.NopCloser(br)

	err = wr.Write(ctx, *uri, fh)

	if err != nil {
		t.Fatal(err)
	}
}
