package config

import (
	"gotest.tools/assert"
	"testing"
)

func TestNewRunConfig(t *testing.T) {
	var data = `
log_config:
    Log_level: 1,
    log_path: "/tmp/test.log"
    log_type: "stdout"
finder_config:
    ignore_dir:
        - .git
        -  .idea
    ignore_file:
        -  readme.md
    timeout: 15 
`

	c, err := NewRunConfig([]byte(data))
	if err != nil {
		t.Error(err)
		return
	}

	if c.LogLevel != 1 && c.LogPath != "/tmp/test.log" && c.LogType != "stdout" {
		t.Error("parser log config error")
	}

	assert.DeepEqual(t, c.IgnoreDir, []string{".git", ".idea"})
	assert.DeepEqual(t, c.IgnoreFile, []string{"readme.md"})
	assert.Assert(t, c.Timeout == 15)

}
