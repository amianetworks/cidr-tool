# ASN CIDR Tool

## Description
ASN CIDR Tool is a Golang package main for handling parse/merge the CIDRs. Check the following use case for the usage.

## Functions

### func ParseSingleWithExcepts(allowCIDR string, excepts []string) []string

Description: the function is for parsing given allow CIDR and except CIDRs, returning all allowed CIDRs by filter out the except.
Use case: In k8s network policy, they have:
```yaml
ipBlock:
  cidr: 172.17.0.0/16
  except:
    - 172.17.1.0/24
```
For some cases, you want to get all allowed CIDRs. So you can use ParseSingleWithExcepts, for example:
Call `ParseSingleWithExcepts("172.17.0.0/16", []string{"172.17.1.0/24"})` iill get
`[172.17.0.0/24 172.17.2.0/23 172.17.4.0/22 172.17.8.0/21 172.17.16.0/20 172.17.32.0/19 172.17.64.0/18 172.17.128.0/17]`

### func MergeAllCIDRs(cidrs []string) []string

Description: the function is for merge all CIDRs in the list.

For example:
Call `MergeAllCIDRs([]string{"10.0.0.0/16", "10.0.0.0/24", "10.0.0.0/32"})`will get
`["10.0.0.16]`

## How to run the test

run `go test -v`.