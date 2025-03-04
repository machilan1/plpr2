package testhelper

import (
	"os"
	"path/filepath"
	"runtime"
	"testing"

	"github.com/google/go-cmp/cmp"
)

// TestDataPath returns a path corresponding to a path relative to the calling
// test file. For convenience, rel is assumed to be "/"-delimited.
//
// It panics on failure.
func TestDataPath(rel string) string {
	_, filename, _, ok := runtime.Caller(1)
	if !ok {
		panic("unable to determine relative path")
	}
	return filepath.Clean(filepath.Join(filepath.Dir(filename), filepath.FromSlash(rel)))
}

func CompareWithGolden(t *testing.T, got, filename string, update bool) {
	t.Helper()
	if update {
		writeGolden(t, filename, got)
	} else {
		want := readGolden(t, filename)
		if diff := cmp.Diff(want, got); diff != "" {
			t.Errorf("%s: mismatch (-want, +got):\n%s", filename, diff)
		}
	}
}

func writeGolden(t *testing.T, name string, data string) {
	filename := filepath.Join("testdata", name)
	if err := os.WriteFile(filename, []byte(data), 0o644); err != nil { //nolint:gosec
		t.Fatal(err)
	}
	t.Logf("wrote %s", filename)
}

func readGolden(t *testing.T, name string) string {
	data, err := os.ReadFile(filepath.Join("testdata", name))
	if err != nil {
		t.Fatal(err)
	}
	return string(data)
}
