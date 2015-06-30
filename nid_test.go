package nid

import (
	"testing"
)

func TestNidCase(t *testing.T) {
	for _, tt := range []struct {
		in   string
		want string
	}{
		// should nid case "" to ""
		{"", ""},

		// should change space to dash
		{"foo bar bee", "foo-bar-bee"},

		// should downcase all chars in the Swedish alphabet
		{"ABCDEFGHIJKLMNOPQRSTUVWXYZÅÄÖ", "abcdefghijklmnopqrstuvwxyzåäö"},

		// should ignore all chars not included in a-z, åäö, -, 0-9
		{"kale8^79'0-", "kale8790"},

		// should convert diacritical characters
		{"Dürén Ibrahimović", "duren-ibrahimovic"},

		// preserves åäö ÅÄÖ
		{"ÅÄÖåäö", "åäöåäö"},

		// converts `¨´^
		{"ÈÉËÊèéëê", "eeeeeeee"},
		{"ÀÁÂàáâ", "aaaaaa"},
		{"Üü", "uu"},
		{"ČĆÇčćç", "cccccc"},
		{"Ññ", "nn"},
		{"Ïï", "ii"},
		{"ÆØæø", "äöäö"},

		// converts _ to -
		{"Let's_Dance", "lets-dance"},

		// removes -- from name
		{"Let's -- da-da-dance", "lets-da-da-dance"},

		// removes surrounding and double space from name and tag
		{" Fångarna     på  fortet   ", "fångarna-på-fortet"},
	} {
		if got := Case(tt.in); got != tt.want {
			t.Errorf(`nid.Case(%q) = %q, want %q`, tt.in, got, tt.want)
		}
	}
}

func TestNidPossible(t *testing.T) {
	for _, tt := range []struct {
		in   string
		want bool
	}{
		// accepts nothingness
		{"", true},

		// accepts åäö
		{"räksmörgås", true},

		// rejects diacritical marks
		{"dürén-ibrahimović", false},

		// rejects non-letters
		{"foo bar", false},
		{"foo/bar", false},
		{"foo\n", false},

		// rejects upper-case letters
		{"FOO", false},
	} {
		if got := Possible(tt.in); got != tt.want {
			t.Errorf(`nid.Possible(%q) = %v, want %v`, tt.in, got, tt.want)
		}
	}
}
