package stack_test

import (
	"flag"
	"io"
	"os"
	"testing"

	"github.com/creachadair/hypercard/stack"
)

var testStack = flag.String("stack", "", "Path to .stk file to test")

func TestParseOne(t *testing.T) {
	if *testStack == "" {
		t.Skip("Skipping test; no --stack path is given")
	}
	f, err := os.Open(*testStack)
	if err != nil {
		t.Fatalf("Opening input: %v", err)
	}
	defer f.Close()

	for {
		blk, err := stack.ParseOne(f)
		if err == io.EOF {
			t.Log("Reached EOF")
			break
		} else if err != nil {
			t.Fatalf("Error parsing: %v", err)
		}

		t.Logf(`Size: %d bytes; Type: %q; ID: %d; Data: [%d bytes]`,
			blk.Size, blk.Type, blk.ID, len(blk.Data))
	}
}
