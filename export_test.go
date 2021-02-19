package asncidr

import (
	"reflect"
	"testing"
)

func TestParseSingleWithExcepts(t *testing.T) {
	expectResult := []string{"10.0.1.0/24", "10.0.2.0/23", "10.0.4.0/22", "10.0.8.0/21", "10.0.16.0/20", "10.0.32.0/19", "10.0.64.0/18", "10.0.128.0/17", "10.1.0.0/16", "10.2.0.0/15", "10.4.0.0/14", "10.8.0.0/13", "10.16.0.0/12", "10.32.0.0/11", "10.64.0.0/10", "10.128.0.0/9"}
	actualResult := ParseSingleWithExcepts("10.0.0.0/8", []string{"10.0.0.0/24"})
	if !reflect.DeepEqual(expectResult, actualResult) {
		t.Errorf("expected: <%v>, got <%v>", expectResult, actualResult)
	}

}

func TestMergeAllCIDRs(t *testing.T) {
	cidrsA := ParseSingleWithExcepts("10.0.0.0/8", []string{"10.0.0.0/24"})
	cidrsB := ParseSingleWithExcepts("10.0.0.0/16", []string{"10.0.0.0/24"})
	expectResult := []string{"10.0.1.0/24", "10.0.2.0/23", "10.0.4.0/22", "10.0.8.0/21", "10.0.16.0/20", "10.0.32.0/19", "10.0.64.0/18", "10.0.128.0/17", "10.1.0.0/16", "10.2.0.0/15", "10.4.0.0/14", "10.8.0.0/13", "10.16.0.0/12", "10.32.0.0/11", "10.64.0.0/10", "10.128.0.0/9"}
	var flatten []string
	flatten = append(flatten, cidrsA...)
	flatten = append(flatten, cidrsB...)
	actualResult := MergeAllCIDRs(flatten)
	if !reflect.DeepEqual(expectResult, actualResult) {
		t.Errorf("expected:\t<%v>", expectResult)
		t.Errorf("got:\t\t<%v>", actualResult)
	}

	expectResult = []string{"10.0.0.0/16"}
	actualResult = MergeAllCIDRs([]string{"10.0.0.0/16", "10.0.0.0/24", "10.0.0.0/32"})
	if !reflect.DeepEqual(expectResult, actualResult) {
		t.Errorf("expected:\t<%v>", expectResult)
		t.Errorf("got:\t\t<%v>", actualResult)
	}

}
