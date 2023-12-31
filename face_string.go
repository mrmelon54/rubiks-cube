// Code generated by "stringer -type Face"; DO NOT EDIT.

package rubiks_cube

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[FaceUp-0]
	_ = x[FaceDown-1]
	_ = x[FaceFront-2]
	_ = x[FaceBack-3]
	_ = x[FaceRight-4]
	_ = x[FaceLeft-5]
}

const _Face_name = "FaceUpFaceDownFaceFrontFaceBackFaceRightFaceLeft"

var _Face_index = [...]uint8{0, 6, 14, 23, 31, 40, 48}

func (i Face) String() string {
	if i >= Face(len(_Face_index)-1) {
		return "Face(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _Face_name[_Face_index[i]:_Face_index[i+1]]
}
