package prensentationcheck

import (
	"path/filepath"
	"testing"

	"golang.org/x/tools/go/analysis/analysistest"
)

func Test(t *testing.T) {
	testdata := filepath.Join(analysistest.TestData())
	analysistest.Run(t, testdata, Analyzer, "a/presentation")
}
