package arraysandstrings

import (
	"testing"
)

func TestIsUnique(t *testing.T) {
	if IsUnique("abcdefg") != true {
		t.Fail()
	}
	if IsUnique("aabbccddeeffgg") != false {
		t.Fail()
	}
}

func TestCheckPermutation(t *testing.T) {
	if CheckPermutation("bca", "abc") != true {
		t.Fail()
	}
	if CheckPermutation("def", "abc") != false {
		t.Fail()
	}
}

func TestURLify(t *testing.T) {
	type args struct {
		s       string
		trueLen int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"Spaces", args{"Mr John Smith    ", 13}, "Mr%20John%20Smith"},
		{"No spaces", args{"IAintGotNoSpaces", 16}, "IAintGotNoSpaces"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := URLify(tt.args.s, tt.args.trueLen); got != tt.want {
				t.Errorf("URLify() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPalindromePermutation(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"Is palindrome", args{"Tact Coa"}, true},
		{"Is not palindrome", args{"REPLCEME"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := PalindromePermutation(tt.args.s); got != tt.want {
				t.Errorf("PalindromePermutation() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOneAway(t *testing.T) {
	type args struct {
		before string
		after  string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"Zero", args{"pale", "pale"}, true},
		{"Remove in", args{"pale", "ple"}, true},
		{"Remove last", args{"pales", "pale"}, true},
		{"Replace", args{"pale", "bale"}, true},
		{"More than one away", args{"pale", "bake"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := OneAway(tt.args.before, tt.args.after); got != tt.want {
				t.Errorf("OneAway() = %v, want %v", got, tt.want)
			}
		})
	}
}
