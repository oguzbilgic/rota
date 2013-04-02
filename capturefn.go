package rota

import (
	"unicode"
)

// CaptureFns is a slice of CaptureFns
type CaptureFns []CaptureFn

// ParseRotaPattern parses the rota's pattern and returns a CaptureFns.
func ParseRotaPattern(pattern string) CaptureFns {
	// [TODO] This function will parse the given pattern but for not
	// it is static
	return CaptureFns{
		ConstCaptureFn("/articles/"),
		IntVarCaptureFn(),
	}
}

// CaptureFn is a type of function which tries to capture information
// from the begining of the given string.
type CaptureFn func(path string) (match bool, rest string)

// ConstCaptureFn generates a CaptureFn which tries to capture the given string from
// the given path.
func ConstCaptureFn(str string) CaptureFn {
	return func(path string) (bool, string) {
		if len(path) < len(str) {
			return false, path
		}

		if path[:len(str)] == str {
			return true, path[len(str):]
		}

		return false, path
	}
}

// VarCaptureFn generates a CaptureFn which tries to capture the int variable from
// the given path.
func VarCaptureFn(rangeTab *unicode.RangeTable) CaptureFn {
	return func(path string) (bool, string) {
		pos := -1

		for pos+1 < len(path) {
			if unicode.Is(rangeTab, rune(path[pos+1])) {
				pos++
			} else {
				break
			}
		}

		if pos > -1 {
			return true, path[pos+1:]
		}

		return false, path
	}
}

// IntVarCaptureFn generates a CaptureFn which tries to capture the int variable
// from the given path.
func IntVarCaptureFn() CaptureFn {
	return VarCaptureFn(unicode.Digit)
}

// StrVarCaptureFn generates a CaptureFn which tries to capture the string variable
// from the given path.
func StrVarCaptureFn() CaptureFn {
	return VarCaptureFn(unicode.Letter)
}
