// Code generated by "stringer -type=FrameMode"; DO NOT EDIT.

package mpg123

import "fmt"

const _FrameMode_name = "ModeStereoModeJointModeDualModeMono"

var _FrameMode_index = [...]uint8{0, 10, 19, 27, 35}

func (i FrameMode) String() string {
	if i < 0 || i >= FrameMode(len(_FrameMode_index)-1) {
		return fmt.Sprintf("FrameMode(%d)", i)
	}
	return _FrameMode_name[_FrameMode_index[i]:_FrameMode_index[i+1]]
}
