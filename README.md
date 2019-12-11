# go-writer-github

Work in progress. This will replace [go-whosonfirst-readwrite](https://github.com/whosonfirst/go-whosonfirst-readwrite).

## Example

```
package main

import (
	"context"
	"flag"
	"fmt"	
	"github.com/whosonfirst/go-writer"
	_ "github.com/whosonfirst/go-writer-github"		
	"io/ioutil"
	"strings"
	"time"
)

func main() {

	owner := "example"
	repo := "example-repo"
	access_token := "s33kret"
	branch := "example"

	ctx := context.Background()
	
	source := fmt.Sprintf("githubapi://%s/%s?access_token=%s&branch=%s", owner, repo, access_token, branch)
			
	wr, _ := writer.NewWriter(ctx, source)

	now := time.Now()
	msg := fmt.Sprintf("This is a test: %v", now)
	
	br := strings.NewReader(msg)
	fh := ioutil.NopCloser(br)

	wr.Write(ctx, "TEST.md", fh)
}
```

## See also

* https://github.com/whosonfirst/go-writer
* https://github.com/google/go-github