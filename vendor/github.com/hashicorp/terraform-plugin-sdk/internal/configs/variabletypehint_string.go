// Code generated by "stringer -type VariableTypeHint"; DO NOT EDIT.

package configs

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[TypeHintNone-0]
	_ = x[TypeHintString-83]
	_ = x[TypeHintList-76]
	_ = x[TypeHintMap-77]
}

const (
	_VariableTypeHint_name_0 = "TypeHintNone"
	_VariableTypeHint_name_1 = "TypeHintListTypeHintMap"
	_VariableTypeHint_name_2 = "TypeHintString"
)

var (
	_VariableTypeHint_index_1 = [...]uint8{0, 12, 23}
)

func (i VariableTypeHint) String() string {
	switch {
	case i == 0:
		return _VariableTypeHint_name_0
	case 76 <= i && i <= 77:
		i -= 76
		return _VariableTypeHint_name_1[_VariableTypeHint_index_1[i]:_VariableTypeHint_index_1[i+1]]
	case i == 83:
		return _VariableTypeHint_name_2
	default:
		return "VariableTypeHint(" + strconv.FormatInt(int64(i), 10) + ")"
	}
}
