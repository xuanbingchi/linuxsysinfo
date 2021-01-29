package linuxsysinfo

import (
	"testing"
)

func TestGetTwoColumns(t *testing.T) {
	c, err := CreatCPUInfo()
	if err != nil {
		t.Fatal(err)
	}

	t.Log(c)
}
