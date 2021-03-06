// Code generated by "stringer -type=Op,Attr,Mode -output screen/const_string.go screen/"; DO NOT EDIT

package screen

import "fmt"

const _Op_name = "OpResizeOpClearOpKeyboardOpCursorOpPaletteOpStyleOpPutOpPutRepOpScrollOpFlushOpLogOpEnd"

var _Op_index = [...]uint8{0, 8, 15, 25, 33, 42, 49, 54, 62, 70, 77, 82, 87}

func (i Op) String() string {
	i -= 1
	if i >= Op(len(_Op_index)-1) {
		return fmt.Sprintf("Op(%d)", i+1)
	}
	return _Op_name[_Op_index[i]:_Op_index[i+1]]
}

const (
	_Attr_name_0 = "AttrReverseAttrItalic"
	_Attr_name_1 = "AttrBold"
	_Attr_name_2 = "AttrUnderline"
	_Attr_name_3 = "AttrUndercurl"
	_Attr_name_4 = "AttrEnd"
)

var (
	_Attr_index_0 = [...]uint8{0, 11, 21}
	_Attr_index_1 = [...]uint8{0, 8}
	_Attr_index_2 = [...]uint8{0, 13}
	_Attr_index_3 = [...]uint8{0, 13}
	_Attr_index_4 = [...]uint8{0, 7}
)

func (i Attr) String() string {
	switch {
	case 1 <= i && i <= 2:
		i -= 1
		return _Attr_name_0[_Attr_index_0[i]:_Attr_index_0[i+1]]
	case i == 4:
		return _Attr_name_1
	case i == 8:
		return _Attr_name_2
	case i == 16:
		return _Attr_name_3
	case i == 32:
		return _Attr_name_4
	default:
		return fmt.Sprintf("Attr(%d)", i)
	}
}

const (
	_Mode_name_0 = "ModeBusyModeMouseOn"
	_Mode_name_1 = "ModeNormal"
	_Mode_name_2 = "ModeInsert"
	_Mode_name_3 = "ModeReplace"
	_Mode_name_4 = "ModeEnd"
)

var (
	_Mode_index_0 = [...]uint8{0, 8, 19}
	_Mode_index_1 = [...]uint8{0, 10}
	_Mode_index_2 = [...]uint8{0, 10}
	_Mode_index_3 = [...]uint8{0, 11}
	_Mode_index_4 = [...]uint8{0, 7}
)

func (i Mode) String() string {
	switch {
	case 1 <= i && i <= 2:
		i -= 1
		return _Mode_name_0[_Mode_index_0[i]:_Mode_index_0[i+1]]
	case i == 4:
		return _Mode_name_1
	case i == 8:
		return _Mode_name_2
	case i == 16:
		return _Mode_name_3
	case i == 32:
		return _Mode_name_4
	default:
		return fmt.Sprintf("Mode(%d)", i)
	}
}
