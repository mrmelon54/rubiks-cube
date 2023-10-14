// Code generated by "stringer -type Move"; DO NOT EDIT.

package rubiks_cube

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[Up-0]
	_ = x[Down-1]
	_ = x[Front-2]
	_ = x[Back-3]
	_ = x[Right-4]
	_ = x[Left-5]
	_ = x[UpPrime-6]
	_ = x[DownPrime-7]
	_ = x[FrontPrime-8]
	_ = x[BackPrime-9]
	_ = x[RightPrime-10]
	_ = x[LeftPrime-11]
}

const _Move_name = "UpDownFrontBackRightLeftUpPrimeDownPrimeFrontPrimeBackPrimeRightPrimeLeftPrime"

var _Move_index = [...]uint8{0, 2, 6, 11, 15, 20, 24, 31, 40, 50, 59, 69, 78}

func (i Move) String() string {
	if i >= Move(len(_Move_index)-1) {
		return "Move(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _Move_name[_Move_index[i]:_Move_index[i+1]]
}
