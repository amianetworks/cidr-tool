package asncidr

import (
	"net"
	"sort"
)

type ipRange struct {
	Original *net.IPNet
	From     net.IP
	To       net.IP
}

func initIPRangeFromCIDR(cidr string) *ipRange {
	_, ipNet, err := net.ParseCIDR(cidr)
	if err != nil {
		return nil
	}
	from, to := cidrToRange(ipNet)
	return &ipRange{
		Original: ipNet,
		From:     from,
		To:       to,
	}
}

func initIPRangeFromRange(start net.IP, end net.IP) *ipRange {
	return &ipRange{
		From: start,
		To:   end,
	}
}

// The ip range contains the target IP. In other word, the target IP is in the ip range
func (iprange ipRange) containsIP(targetIP net.IP) bool {
	return ipToInt(iprange.From) <= ipToInt(targetIP) && ipToInt(targetIP) <= ipToInt(iprange.To)
}

// The ip range's start IP (From) is in the target ip range but the end IP (To) is not in the target ip range
func (iprange ipRange) startIPInNet(target *ipRange) bool {
	return target.containsIP(iprange.From) && !target.containsIP(iprange.To)
}

// The ip range's end IP (To) is in the target ip range but the start IP (From) is not in the target ip range
func (iprange ipRange) endIPInNet(target *ipRange) bool {
	return !target.containsIP(iprange.From) && target.containsIP(iprange.To)
}

// The ip range contains the target ip range (both start and end IP of target is in the iprange)
func (iprange ipRange) containsNet(target *ipRange) bool {
	return iprange.containsIP(target.From) && iprange.containsIP(target.To)
}

// Convert the list of CIDRs to the list of ipRange
func getIPRangesFromString(list []string) []*ipRange {
	var result []*ipRange
	for _, cidr := range list {
		ipRange := initIPRangeFromCIDR(cidr)

		result = append(result, ipRange)
	}
	return result
}

// Sort the ipRange to the ascending order
func getSortedIPRanges(list []*ipRange) []*ipRange {
	sort.Slice(list, func(i, j int) bool {
		if ipToInt(list[i].From) == ipToInt(list[j].From) {
			return ipToInt(list[i].To) < ipToInt(list[j].To)
		} else {
			return ipToInt(list[i].From) < ipToInt(list[j].From)
		}
	})
	return list
}

//	Merge the ipRange. If two ipRange intersected, take the min start (From) and max end (To).
func getMergedIPRanges(list []*ipRange) []*ipRange {
	var merged []*ipRange
	if len(list) == 0 {
		return merged
	}
	j := 1
	currentAllow := list[0]
	for j < len(list) {
		finder := list[j]
		if currentAllow.containsNet(finder) {
			// do nothing, looking for the net
		} else if finder.containsNet(currentAllow) {
			// finder contains current , replace current by the larger range
			currentAllow = finder
		} else if currentAllow.startIPInNet(finder) {
			// current start ip in finder, merge it to a larger current
			currentAllow = initIPRangeFromRange(currentAllow.From, finder.To)
		} else if currentAllow.endIPInNet(finder) {
			// current end ip in finder, merge it to a larger current
			currentAllow = initIPRangeFromRange(finder.From, currentAllow.To)
		} else {
			// not intersected, add current allow to the result
			merged = append(merged, currentAllow)
			currentAllow = finder
		}
		j++
	}
	merged = append(merged, currentAllow)
	return merged
}
