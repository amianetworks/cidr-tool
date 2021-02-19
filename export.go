package asncidr

// Parse the single CIDR and given excepts to the filtered CIDR
func ParseSingleWithExcepts(allowCIDR string, excepts []string) []string {
	exceptIPNets := getIPRangesFromString(excepts)
	exceptIPNets = getSortedIPRanges(exceptIPNets)
	exceptIPNets = getMergedIPRanges(exceptIPNets)
	var finalAllows []*ipRange
	allow := initIPRangeFromCIDR(allowCIDR)
	from := allow.From
	for _, except := range exceptIPNets {
		if except.containsNet(allow) {
			return []string{}
		}
		if allow.containsNet(except) {
			if !from.Equal(except.From) {
				finalAllows = append(finalAllows, initIPRangeFromRange(from, ipSub(except.From, 1)))
			}
			from = ipAdd(except.To, 1)
		}
	}
	if ipToInt(from) < ipToInt(allow.To) {
		finalAllows = append(finalAllows, initIPRangeFromRange(from, allow.To))
	}
	var result []string
	for _, fAllow := range finalAllows {
		result = append(result, rangeToCIDRs(fAllow.From, fAllow.To)...)
	}
	return result
}

// Giving a list of CIDRS, merge then to a single CIDR list
func MergeAllCIDRs(cidrs []string) []string {
	flattenIPNets := getIPRangesFromString(cidrs)
	flattenIPNets = getSortedIPRanges(flattenIPNets)
	flattenIPNets = getMergedIPRanges(flattenIPNets)
	var result []string
	for _, iprange := range flattenIPNets {
		result = append(result, rangeToCIDRs(iprange.From, iprange.To)...)
	}
	return result
}
