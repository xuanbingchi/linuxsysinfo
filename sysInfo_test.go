package linuxsysinfo

import (
	"encoding/json"
	"testing"
)

func TestCreatCPUInfo(t *testing.T) {
	i, err := CreatCPUInfo()
	if err != nil {
		t.Fatal(err)
	}
	ii, err := json.MarshalIndent(i, "", "\t")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(string(ii))
}

func TestCreatMemInfo(t *testing.T) {
	i, err := CreatMemInfo()
	if err != nil {
		t.Fatal(err)
	}

	ii, err := json.MarshalIndent(i, "", "\t")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(string(ii))
}

func TestCreatIfConfigInfos(t *testing.T) {
	i, err := CreatIfConfigInfos()
	if err != nil {
		t.Fatal(err)
	}

	ii, err := json.MarshalIndent(i, "", "\t")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(string(ii))
}

func TestCreatVersionInfo(t *testing.T) {
	i, err := CreatVersionInfo()
	if err != nil {
		t.Fatal(err)
	}

	ii, err := json.MarshalIndent(i, "", "\t")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(string(ii))
}
