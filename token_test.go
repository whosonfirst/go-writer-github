package writer

import (
	"context"
	_ "gocloud.dev/runtimevar/constantvar"
	"testing"
)

func TestEnsureGitHubAccessToken(t *testing.T) {

	ctx := context.Background()

	expected_uri := "githubapi://sfomuseum-data/sfomuseum-data-collection?access_token=s33kret"

	writer_uri := "githubapi://sfomuseum-data/sfomuseum-data-collection?access_token={access_token}"
	token_uri := "constant://?val=s33kret"

	final_uri, err := EnsureGitHubAccessToken(ctx, writer_uri, token_uri)

	if err != nil {
		t.Fatalf("Failed to ensure github access token in '%s', %v", writer_uri, err)
	}

	if final_uri != expected_uri {
		t.Fatalf("Unexpected final URI: %s", final_uri)
	}

	stdout_uri := "stdout://"

	final_uri, err = EnsureGitHubAccessToken(ctx, stdout_uri, "")

	if err != nil {
		t.Fatalf("Failed to ensure github access token in '%s', %v", stdout_uri, err)
	}

	if final_uri != stdout_uri {
		t.Fatalf("Unexpected final URI for %s: %s", stdout_uri, final_uri)
	}
}
