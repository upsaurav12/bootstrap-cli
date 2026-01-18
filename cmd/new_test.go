package cmd

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReturnUppercase(t *testing.T) {
	tests := []struct {
		in   string
		want string
	}{
		{"user", "User"},
		{"product", "Product"},
		{"", ""},
	}

	for _, tt := range tests {
		assert.Equal(t, tt.want, returnUppercase(tt.in))
	}
}

func TestIsHidden(t *testing.T) {
	hidden, err := IsHidden("/tmp/.env")
	assert.NoError(t, err)
	assert.True(t, hidden)

	notHidden, err := IsHidden("/tmp/main.go")
	assert.NoError(t, err)
	assert.False(t, notHidden)
}

func TestCopyProjectYAML(t *testing.T) {
	tmpDir := t.TempDir()

	src := filepath.Join(tmpDir, "project.yaml")
	err := os.WriteFile(src, []byte("name: test"), 0644)
	assert.NoError(t, err)

	destDir := t.TempDir()

	err = copyProjectYAML(src, destDir)
	assert.NoError(t, err)

	out, err := os.ReadFile(filepath.Join(destDir, "project.yaml"))
	assert.NoError(t, err)
	assert.Equal(t, "name: test", string(out))
}

func TestCopyProjectYAML_EmptySource(t *testing.T) {
	err := copyProjectYAML("", t.TempDir())
	assert.NoError(t, err)
}

func TestWriteSingle_CreatesFile(t *testing.T) {
	tmpDir := t.TempDir()

	data := TemplateData{
		Entity: "User",
	}

	content := []byte("Hello {{.Entity}}")

	err := writeSingle(
		data,
		"example.txt",
		"example.tmpl",
		content,
		tmpDir,
		false,
	)

	assert.NoError(t, err)

	out, err := os.ReadFile(filepath.Join(tmpDir, "user.txt"))
	assert.NoError(t, err)
	assert.Equal(t, "Hello User", string(out))
}

func TestRenderTemplateDir_NoError(t *testing.T) {
	tmpDir := t.TempDir()

	data := TemplateData{
		Entities: []string{},
	}

	err := renderTemplateDir("common", tmpDir, data)

	assert.NoError(t, err)
}
