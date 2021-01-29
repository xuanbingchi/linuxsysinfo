package linuxsysinfo

import (
	"regexp"
	"strings"
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

func TestGetInfo3(t *testing.T) {

	getInfo3("")

}

/*
enP1p38s0f1: flags=4099<UP,BROADCAST,MULTICAST>  mtu 1500
        inet 192.168.3.254  netmask 255.255.255.0  broadcast 192.168.3.255
        ether 6c:b3:11:41:65:91  txqueuelen 1000  (Ethernet)
        RX packets 0  bytes 0 (0.0 B)
        RX errors 0  dropped 0  overruns 0  frame 0
        TX packets 0  bytes 0 (0.0 B)
        TX errors 0  dropped 0 overruns 0  carrier 0  collisions 0
        device memory 0xf0e0320000-f0e033ffff

*/

func TestMustCompile(t *testing.T) {

	snowRegexp := regexp.MustCompile(`^ +(inet [\d.]+)  (netmask [\d.]+)  (broadcast [\d.]+)$`)
	params := snowRegexp.FindStringSubmatch("        inet 192.168.3.254  netmask 255.255.255.0")
	//params := snowRegexp.FindStringSubmatch("        inet 192.168.1.254  netmask 255.255.255.0  broadcast 192.168.1.255")
	//flysnowRegexp := regexp.MustCompile(`^http://www.flysnow.org/([\d]{4})/([\d]{2})/([\d]{2})/([\w-]+).html$`)
	//params := flysnowRegexp.FindStringSubmatch("http://www.flysnow.org/2018/01/20/golang-goquery-examples-selector.html")

	t.Log(params)
}

func TestM2(t *testing.T) {
	txt := `enP1p39s0f0: flags=4099<UP,BROADCAST,MULTICAST>  mtu 1500
        inet 192.168.4.254  netmask 255.255.255.0  broadcast 192.168.4.255
        ether 6c:b3:11:41:65:94  txqueuelen 1000  (Ethernet)
        RX packets 0  bytes 0 (0.0 B)
        RX errors 0  dropped 0  overruns 0  frame 0
        TX packets 0  bytes 0 (0.0 B)
        TX errors 0  dropped 0 overruns 0  carrier 0  collisions 0
        device memory 0xf0e0400000-f0e04fffff`
	sp := [...]string{": ",
		"mtu ",
		"inet ",
		"netmask ",
		"broadcast ",
		"ether ",
		"txqueuelen ",
		"RX packets ",
		"RX errors ",
		"TX packets ",
		"TX errors ",
		"device memory ",
	}

	s := strings.Replace(txt, "\n", "", -1)
	r := make([]string, 0)
	for _, spp := range sp {
		index := strings.Index(s, spp)
		if index < 1 {
			continue
		}
		r = append(r, s[:index])
		s = s[index:]
	}

}
