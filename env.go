package packr

import (
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

// ABR-286814 - importing github.com/gobuffalo/envy causes searching for go binary

// GoPath returns the current GOPATH env var
// or if it's missing, the default.
func GoPath() string {
	out, err := exec.Command("go", "env", "GOPATH").Output()
	if err == nil {
		return strings.TrimSpace(string(out))
	}
	return filepath.Join(os.Getenv("HOME"), "go")
}

// GoBin returns the current GO_BIN env var
// or if it's missing, a default of "go"
func GoBin() string {
	go_bin := os.Getenv("GO_BIN")
	if go_bin == "" {
		return "go"
	}
	return go_bin
}
