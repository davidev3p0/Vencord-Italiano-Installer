package asar

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestWriteContainsBootstrap(t *testing.T) {
	out := filepath.Join(t.TempDir(), "app.asar")
	patcher := `C:\\Users\\Test\\AppData\\Roaming\\Vencord\\dist\\patcher.js`

	if err := Write(out, patcher); err != nil {
		t.Fatal(err)
	}
	data, err := os.ReadFile(out)
	if err != nil {
		t.Fatal(err)
	}
	text := string(data)
	if !strings.Contains(text, "index.js") || !strings.Contains(text, "package.json") {
		t.Fatalf("missing ASAR entries")
	}
	if !strings.Contains(text, "patcher.js") {
		t.Fatalf("missing patcher bootstrap")
	}
}
