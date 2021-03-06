// Code generated by "stringer -type=topologyOption"; DO NOT EDIT.

package protocol

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[toHostName-1]
	_ = x[toHostPortnumber-2]
	_ = x[toTenantName-3]
	_ = x[toLoadfactor-4]
	_ = x[toVolumeID-5]
	_ = x[toIsPrimary-6]
	_ = x[toIsCurrentSession-7]
	_ = x[toServiceType-8]
	_ = x[toNetworkDomain-9]
	_ = x[toIsStandby-10]
	_ = x[toAllIPAddresses-11]
	_ = x[toAllHostNames-12]
}

const _topologyOption_name = "toHostNametoHostPortnumbertoTenantNametoLoadfactortoVolumeIDtoIsPrimarytoIsCurrentSessiontoServiceTypetoNetworkDomaintoIsStandbytoAllIPAddressestoAllHostNames"

var _topologyOption_index = [...]uint8{0, 10, 26, 38, 50, 60, 71, 89, 102, 117, 128, 144, 158}

func (i topologyOption) String() string {
	i -= 1
	if i < 0 || i >= topologyOption(len(_topologyOption_index)-1) {
		return "topologyOption(" + strconv.FormatInt(int64(i+1), 10) + ")"
	}
	return _topologyOption_name[_topologyOption_index[i]:_topologyOption_index[i+1]]
}
