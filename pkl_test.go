package pkladapter

import (
	"bytes"
	"os"
	"path/filepath"
	"sort"
	"testing"
)

func TestConfigs(t *testing.T) {
	configs, err := filepath.Glob("testdata/config*")
	if err != nil {
		t.Fatal(err)
	}
	l := len(configs)
	if l == 0 {
		t.Fatal("no configs found")
	}
	if l%2 != 0 {
		t.Fatal("number of configs should not be odd")
	}
	sort.Strings(configs)

	a := Adapter{}

	for i := 0; i < l; i += 2 {
		jn, err := os.ReadFile(configs[i])
		if err != nil {
			t.Error(err)
			continue
		}
		pk, err := os.ReadFile(configs[i+1])
		if err != nil {
			t.Error(err)
			continue
		}

		b, _, err := a.Adapt(pk, nil)
		if err != nil {
			t.Error(err)
			continue
		}

		if !bytes.Equal(bytes.TrimSpace(jn), bytes.TrimSpace(b)) {
			t.Errorf("fixture %q failed", configs[i])
			t.Logf("expected: %s\nactual: %s", jn, b)
		}
	}
}
