package nid

import (
	"fmt"
	"regexp"
	"testing"
)

func TestCase(t *testing.T) {
	for i, tt := range []struct {
		in   string
		want string
		nid  *Nid
	}{
		// should nid case "" to ""
		{"", "", Default},

		// should change space to dash
		{"foo bar bee", "foo-bar-bee", Default},

		// should downcase all chars in the Swedish alphabet if allowing åäö
		{"ABCDEFGHIJKLMNOPQRSTUVWXYZÅÄÖ", "abcdefghijklmnopqrstuvwxyzåäö", WithÅÄÖ},

		// should downcase all chars and transliterate
		{"ABCDEFGHIJKLMNOPQRSTUVWXYZÅÄÖ", "abcdefghijklmnopqrstuvwxyzaao", Default},

		// should downcase all chars matched by the strip pattern
		{"ABCDEFGHIJKLMNOPQRSTUVWXYZÅÄÖ", "abcdefghijklmnopqrstuvwxyz", New(AllowÅÄÖ, SetStripPattern(regexp.MustCompile(`[^a-z-]`)))},

		// should transliterate
		{"ÀÁÂÃÄÅÆÇ_ÈÉÊËÌÍÎÏÐÑÒÓÔÕÖ×ØÙÚÛÜ", "aaaaäåäc-eeeeiiiidnooooöxöuuuu", WithÅÄÖ},
		{"ÀÁÂÃÄÅÆÇ_ÈÉÊËÌÍÎÏÐÑÒÓÔÕÖ×ØÙÚÛÜ", "aaaaaaac-eeeeiiiidnoooooxouuuu", Default},
		{"ÝÞßàáâãäåæçèéêëìíîïðñòóôõöøùúû", "ythssaaaaäåäceeeeiiiidnooooööuuu", WithÅÄÖ},
		{"ÝÞßàáâãäåæçèéêëìíîïðñòóôõöøùúû", "ythssaaaaaaaceeeeiiiidnoooooouuu", Default},
		{"üýþÿĀāĂăĄąĆćĈĉĊċČčĎďĐđĒēĔĕĖėĘę", "uythyaaaaaaccccccccddddeeeeeeee", Default},
		{"ĚěĜĝĞğĠġĢģĤĥĦħĨĩĪīĬĭĮįİıĲĳĴĵĶķ", "eegggggggghhhhiiiiiiiiiiijijjjkk", Default},
		{"ĸĹĺĻļĽľĿŀŁłŃńŅņŇňŉŊŋŌōŎŏŐőŒœŔŕ", "kllllllllllnnnnnnnngngoooooooeoerr", Default},
		{"ŖŗŘřŚśŜŝŞşŠšŢţŤťŦŧŨũŪūŬŭŮůŰűŲų", "rrrrssssssssttttttuuuuuuuuuuuu", Default},
		{"ŴŵŶŷŸŹźŻżŽž", "wwyyyzzzzzz", Default},

		// should ignore all chars not included in a-z, åäö, -, 0-9
		{"kale8^79'0-", "kale8790", Default},
		{"kale8^79'0-", "kale", New(SetStripPattern(regexp.MustCompile(`[^a-z]`)))},

		// should convert diacritical characters
		{"Dürén Ibrahimović", "duren-ibrahimovic", Default},

		// preserves åäö ÅÄÖ if allowing åäö
		{"ÅÄÖåäö", "åäöåäö", WithÅÄÖ},

		// converts `¨´^
		{"ÈÉËÊèéëê", "eeeeeeee", Default},
		{"ÀÁÂàáâ", "aaaaaa", Default},
		{"Üü", "uu", Default},
		{"ČĆÇčćç", "cccccc", Default},
		{"Ññ", "nn", Default},
		{"Ïï", "ii", Default},
		{"ÆØæø", "aoao", Default},
		{"ÆØæø", "äöäö", WithÅÄÖ},

		// converts _ to -
		{"Let's_Dance", "lets-dance", Default},
		{"N___F5__9hf3m2iDyO4F__rjyD", "n-f5-9hf3m2idyo4f-rjyd", Default},

		// removes -- from name
		{"Let's -- da-da-dance", "lets-da-da-dance", Default},

		// removes surrounding and double space from name and tag
		{" Fångarna     på  fortet   ", "fangarna-pa-fortet", Default},
		{" Fångarna     på  fortet   ", "fångarna-på-fortet", WithÅÄÖ},

		// removes –, —
		{"Arn – Tempelriddaren", "arn-tempelriddaren", Default},
		{"Arn — Tempelriddaren", "arn-tempelriddaren", Default},

		// replaces repeating dashes with single dash
		{"Alvinnn!! & the Chipmunks", "alvinnn-the-chipmunks", Default},
		{"alvinnn--the---chipmunks", "alvinnn-the-chipmunks", Default},
		{"alvinnn--the-chipmunks", "alvinnn-the-chipmunks", Default},
	} {
		if got := tt.nid.Case(tt.in); got != tt.want {
			t.Fatalf(`[%d] tt.nid.Case(%q) = %q, want %q`, i, tt.in, got, tt.want)
		}
	}
}

func ExampleCase() {
	fmt.Println(Case("Let's_Dance "))
	// Output: lets-dance
}

func BenchmarkPartialPrepare(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Default.prepare("Dürén Ibrahimović")
	}
}

func BenchmarkPartialPrepareSquish(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Default.squish(Default.prepare("Dürén Ibrahimović"))
	}
}

func BenchmarkPartialPrepareSquishTransliterate(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Default.transliterate(Default.squish(Default.prepare("Dürén Ibrahimović")))
	}
}

func BenchmarkPartialPrepareSquishTransliterateStrip(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Default.strip(Default.transliterate(Default.squish(Default.prepare("Dürén Ibrahimović"))))
	}
}

func BenchmarkCaseEmpty(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Case("")
	}
}

func BenchmarkCaseIgnore(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Case("kale8^79'0-")
	}
}

func BenchmarkCaseSquish(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Case(" Fångarna     på  fortet   ")
	}
}

func BenchmarkCaseDiacritical(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Case("Dürén Ibrahimović")
	}
}

func TestPossible(t *testing.T) {
	for _, tt := range []struct {
		in   string
		want bool
		nid  *Nid
	}{
		// does not accept nothingness by default
		{"", false, Default},

		// accepts nothingness when allowing åäö
		{"", true, WithÅÄÖ},

		// accepts åäö
		{"räksmörgås", true, WithÅÄÖ},

		// does not accept åäö if we change the valid pattern
		{"räksmörgås", false, Default},

		// rejects double dashes
		{"foo--bar", false, Default},

		// rejects diacritical marks
		{"dürén-ibrahimović", false, Default},

		// rejects non-letters
		{"foo bar", false, Default},
		{"foo/bar", false, Default},
		{"foo\n", false, Default},

		// rejects upper-case letters
		{"FOO", false, Default},
	} {
		if got := tt.nid.Possible(tt.in); got != tt.want {
			t.Errorf(`tt.nid.Possible(%q) = %v, want %v`, tt.in, got, tt.want)
		}
	}
}

func ExamplePossible() {
	fmt.Println(Possible("Zlatan Ibrahimović"))
	// Output: false
}

func TestSquish(t *testing.T) {
	for _, tt := range []struct {
		in   string
		want string
	}{
		{"foo  bar   baz", "foo bar baz"},
	} {
		if got := Default.squish(tt.in); got != tt.want {
			t.Errorf(`squish(%q) = %q, want %q`, tt.in, got, tt.want)
		}
	}
}
