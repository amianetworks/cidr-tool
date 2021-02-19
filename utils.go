package asncidr

import (
	"fmt"
	"math"
	"net"
	"strconv"
	"strings"
)

// Giving the start and end of ip range, convert it to the list of CIDR
func rangeToCIDRs(ipStart net.IP, ipEnd net.IP) []string {
	start := ipToInt(ipStart)
	end := ipToInt(ipEnd)
	var result []string

	for end >= start {
		var maxSize byte = 32
		for maxSize > 0 {
			mask := iMask(int(maxSize - 1))
			maskBase := start & mask
			if maskBase != start {
				break
			}
			maxSize--
		}

		x := math.Log(float64(end-start+1)) / math.Log(2)
		maxDiff := byte(32 - math.Floor(x))
		if maxSize < maxDiff {
			maxSize = maxDiff
		}
		ip := intToIP(start)
		result = append(result, fmt.Sprintf("%s/%d", ip, maxSize))
		start += uint(math.Pow(2, float64(32-maxSize)))
	}
	return result
}

// Giving the CIDR, convert to the, get the start and end of the range
func cidrToRange(network *net.IPNet) (net.IP, net.IP) {
	// the first IP
	firstIP := network.IP

	// the last IP is the network address OR NOT the mask address
	prefixLen, bits := network.Mask.Size()
	if prefixLen == bits {
		// make sure that our two slices are distinct, since they would be in all other cases.
		lastIP := make([]byte, len(firstIP))
		copy(lastIP, firstIP)
		return firstIP, lastIP
	}

	firstIPInt := ipToInt(firstIP)
	hostLen := uint(32) - uint(prefixLen)
	var lastIPInt uint = 1
	lastIPInt = lastIPInt << hostLen
	lastIPInt = lastIPInt - 1
	lastIPInt = lastIPInt | firstIPInt

	firstIP.To4()
	return firstIP, intToIP(lastIPInt)
}

func iMask(s int) uint {
	return uint(math.Round(math.Pow(2, 32) - math.Pow(2, float64(32-s))))
}

// Convert int to ip
func intToIP(intIP uint) net.IP {
	var sbIP string
	sbIP += strconv.Itoa(int(intIP >> 24))
	sbIP += "."
	sbIP += strconv.Itoa(int((intIP & 0x00FFFFFF) >> 16))
	sbIP += "."
	sbIP += strconv.Itoa(int((intIP & 0x0000FFFF) >> 8))
	sbIP += "."
	sbIP += strconv.Itoa(int(intIP & 0x000000FF))

	return net.ParseIP(sbIP)
}

// Convert ip to int
func ipToInt(ipStr net.IP) uint {
	ipAddressInArray := strings.Split(ipStr.String(), ".")
	var num int = 0
	for x := 3; x >= 0; x-- {
		ip, err := strconv.Atoi(ipAddressInArray[3-x])
		if err != nil {
			fmt.Println(err)
			return 0
		}
		num |= ip << (x << 3)
	}
	return uint(num)
}

// Add numbers to ip
func ipAdd(ip net.IP, num uint) net.IP {
	return intToIP(ipToInt(ip) + num)
}

// Subtract numbers to ip
func ipSub(ip net.IP, num uint) net.IP {
	return intToIP(ipToInt(ip) - num)
}
