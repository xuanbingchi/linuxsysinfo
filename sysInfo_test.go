package linuxsysinfo

import (
	"regexp"
	"testing"
)

func TestCreatCPUInfo(t *testing.T) {
	c, err := CreatCPUInfo()
	if err != nil {
		t.Fatal(err)
	}

	t.Log(c)
}

func TestCreatMemInfo(t *testing.T) {
	c, err := CreatMemInfo()
	if err != nil {
		t.Fatal(err)
	}

	t.Log(c)
}

func TestCreatIfConfigInfos(t *testing.T) {
	i, err := CreatIfConfigInfos()
	if err != nil {
		t.Fatal(err)
	}

	t.Log(i)
}

func TestName(t *testing.T) {

	m := regexp.MustCompile(`^([\w_\(\)]+): +([\d]+)[\D]*$`)
	a := m.FindStringSubmatch("AnonHugePages:   1769472 kB")
	b := m.FindStringSubmatch("HugePages_Total:       0")
	t.Log(a, b)
}
