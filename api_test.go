package writer

import (
	"bytes"
	"context"
	"github.com/whosonfirst/go-writer"
	"io/ioutil"
	"testing"
)

func TestAPIWriter(t *testing.T) {

	ctx := context.Background()

	wr_uri := ""
	uri := ""

	wr, err := writer.NewWriter(wr_uri)

	if err != nil {
		t.Fatal(err)
	}

	br := bytes.NewReader("This is a test")
	fh := ioutil.NopCloser(br)

	err = wr.Write(ctx, uri, fh)

	if err != nil {
		t.Fatal(err)
	}
}
