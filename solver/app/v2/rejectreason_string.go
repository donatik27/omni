// Code generated by "stringer -type=rejectReason -trimprefix=reject"; DO NOT EDIT.

package appv2

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[rejectNone-0]
	_ = x[rejectDestCallReverts-1]
	_ = x[rejectInsufficientFee-2]
	_ = x[rejectInsufficientInventory-3]
	_ = x[rejectUnsupportedToken-4]
	_ = x[rejectUnsupportedDestChain-5]
}

const _rejectReason_name = "NoneDestCallRevertsInsufficientFeeInsufficientInventoryUnsupportedTokenUnsupportedDestChain"

var _rejectReason_index = [...]uint8{0, 4, 19, 34, 55, 71, 91}

func (i rejectReason) String() string {
	if i >= rejectReason(len(_rejectReason_index)-1) {
		return "rejectReason(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _rejectReason_name[_rejectReason_index[i]:_rejectReason_index[i+1]]
}
