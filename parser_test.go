package main

import (
	"testing"
)

const file = "tests/input.go"

type F map[string]string

var expectedTags = []Tag{
	tag("TestPackage", "1", "p", F{}),
	tag("go/ast", "4", "i", F{}),
	tag("variable", "7", "v", F{"access": "private", "type": "int"}),
	tag("Constant", "9", "c", F{"access": "public", "type": "string"}),
	tag("OtherConst", "10", "c", F{"access": "public"}),
	tag("Function1", "12", "f", F{"access": "public", "signature": "()", "type": "string"}),
	tag("function2", "15", "f", F{"access": "private", "signature": "(p1, p2 int, p3 *string)"}),
	tag("Field1", "19", "w", F{"access": "public", "ctype": "Struct", "type": "int"}),
	tag("field2", "20", "w", F{"access": "private", "ctype": "Struct", "type": "string"}),
	tag("field3", "21", "w", F{"access": "private", "ctype": "Struct", "type": "*bool"}),
	tag("Struct", "18", "t", F{"access": "public", "type": "struct"}),
	tag("NewStruct", "24", "f", F{"access": "public", "signature": "()", "ctype": "Struct", "type": "*Struct"}),
	tag("myInt", "28", "t", F{"access": "private", "type": "int"}),
	tag("F1", "30", "f", F{"access": "public", "signature": "()", "ctype": "myInt", "type": "[]bool, [2]*string"}),
	tag("Struct", "34", "w", F{"access": "public", "ctype": "TestEmbed", "type": "Struct"}),
	tag("TestEmbed", "33", "t", F{"access": "public", "type": "struct"}),
	tag("InterfaceMethod", "38", "f", F{"access": "public", "signature": "(int)", "ntype": "Interface", "type": "string"}),
	tag("Interface", "37", "n", F{"access": "public", "type": "interface"}),
}

func TestParse(t *testing.T) {
	tags, err := Parse("tests/input.go")
	if err != nil {
		t.Fatalf("Unexpected error: %s", err)
	}

	if len(tags) != len(expectedTags) {
		t.Fatalf("len(tags) == %d, want %d", len(tags), len(expectedTags))
	}

	for i, exp := range expectedTags {
		if tags[i].String() != exp.String() {
			t.Errorf("tag(%d)\n  is:%s\nwant:%s", i, tags[i].String(), exp.String())
		}
	}
}

func tag(n, l, t string, fields F) (tag Tag) {
	tag = Tag{
		Name:    n,
		File:    file,
		Address: l,
		Type:    t,
		Fields:  fields,
	}

	tag.Fields["line"] = l

	return
}
