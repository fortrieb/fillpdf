package fillpdf

import "fmt"

// FormTyp interface for transforming type to FDF
type FormTyp interface {
	// toFdf string representation of this type
	toFdf() string
}

type StringTyp struct {
	Key, Value string
}

func (st StringTypA) toFdf() string {
	return fmt.Sprintf("<< /T (%s) /V (%v)>>\n", st.Key, st.Value)
}

type BoolTyp struct {
	Key   string
	Value bool
}

func (bt BoolTyp) toFdf() string {
	return fmt.Sprintf("<< /T (%s) /V (%v)>>\n", bt.Key, bt.boolToFdf())
}

func (bt BoolTyp) boolToFdf() string {
	var boolStr string
	if bt.Value {
		boolStr = "On"
	} else {
		boolStr = "Off"
	}
	return boolStr
}
