package asncidr

import (
	"math"
	"net"
	"reflect"
	"testing"
)

func TestRangeToCIDRs(t *testing.T) {
	expected := []string{"10.0.0.0/24"}
	start := net.ParseIP("10.0.0.0")
	end := net.ParseIP("10.0.0.255")
	actual := rangeToCIDRs(start, end)
	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("expected:\t%v", expected)
		t.Errorf("got:\t%v", actual)
	}
}

func TestCIDRToRange(t *testing.T) {
	_, ipNet, _ := net.ParseCIDR("10.0.0.0/24")
	expectedStart := net.ParseIP("10.0.0.0")
	expectedEnd := net.ParseIP("10.0.0.255")
	actualStart, actualEnd := cidrToRange(ipNet)
	if !expectedStart.Equal(actualStart) {
		t.Errorf("expected:\t%v", expectedStart)
		t.Errorf("got:\t%v", actualStart)
	}
	if !expectedEnd.Equal(actualEnd) {
		t.Errorf("expected:\t%v", expectedEnd)
		t.Errorf("got:\t%v", actualEnd)
	}
}

func TestIPToInt(t *testing.T) {
	expected := math.Pow(2, 32) - 1
	actual := ipToInt(net.ParseIP("255.255.255.255"))
	if uint(expected) != actual {
		t.Errorf("expected:\t%v", expected)
		t.Errorf("got:\t%v", actual)
	}
}

func TestIntToIP(t *testing.T) {
	expected := net.ParseIP("255.255.255.255")
	actual := intToIP(uint(math.Pow(2, 32) - 1))
	if !expected.Equal(actual) {
		t.Errorf("expected:\t%v", expected)
		t.Errorf("got:\t%v", actual)
	}
}

func TestIPAdd(t *testing.T) {
	overflowE := ipAdd(net.ParseIP("255.255.255.255"), 1)
	if overflowE != nil {
		t.Errorf("expected:\t%v", nil)
		t.Errorf("got:\t%v", overflowE)
	}
	expected := net.ParseIP("10.0.0.1")
	actual := ipAdd(net.ParseIP("10.0.0.0"), 1)
	if !expected.Equal(actual) {
		t.Errorf("expected:\t%v", expected)
		t.Errorf("got:\t%v", actual)
	}
}

func TestIPSub(t *testing.T) {
	overflowS := ipSub(net.ParseIP("0.0.0.0"), 1)
	if overflowS != nil {
		t.Errorf("expected:\t%v", nil)
		t.Errorf("got:\t%v", overflowS)
	}
	expected := net.ParseIP("10.0.0.0")
	actual := ipSub(net.ParseIP("10.0.0.1"), 1)
	if !expected.Equal(actual) {
		t.Errorf("expected:\t%v", expected)
		t.Errorf("got:\t%v", actual)
	}
}
