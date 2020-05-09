package commander

import (
	cm "github.com/cloudrain21/backup-blog/module/commander"
	"github.com/magiconair/properties/assert"
	"testing"
)

func ExampleNewSSH() {
	var commander cm.Commander
	commander = cm.NewSSH("www.cloudrain21.com", "rain", "rain")

	commander.Connect()
	commander.RunCommand("ls -al")
	commander.Close()
}

func TestNewSSH(t *testing.T) {
	t.Run("connection test", func(t *testing.T) {
		var commander cm.Commander
		commander = cm.NewSSH("www.cloudrain21.com", "rain", "rain")

		var err error = nil
		want := err
		got := commander.Connect()

		assert.Equal(t, want, got)
	})
}
