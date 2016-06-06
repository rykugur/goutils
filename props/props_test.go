package props

import "testing"

func TestProps(t *testing.T) {
	propsFile := "props_test.properties"

	props, err := Load(propsFile)

	if err != nil {
		t.Error("Load returned nil")
	}

	// props should be length 5
	propsLen := len(props)
	if propsLen != 5 {
		t.Error("props map expected length 5, actual=", propsLen)
	}

	// pick a few random props
	if props["a.fun.prop"] != "porp.nuf.a" {
		t.Error("expected \"porp.nuf.a\", actual=", props["a.fun.prop"])
	}
	if props["numbers"] != "123" {
		t.Error("expected \"123\", actual=", props["numbers"])
	}
}
