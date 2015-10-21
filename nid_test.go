package nid

import (
	"fmt"
	"testing"
)

func TestCase(t *testing.T) {
	for _, tt := range []struct {
		in   string
		want string
	}{
		// should nid case "" to ""
		{"", ""},

		// should change space to dash
		{"foo bar bee", "foo-bar-bee"},

		// should downcase all chars in the Swedish alphabet
		{"ABCDEFGHIJKLMNOPQRSTUVWXYZÅÄÖ", "abcdefghijklmnopqrstuvwxyzaao"},

		// should transliterate
		{"ÀÁÂÃÄÅÆÇ_ÈÉÊËÌÍÎÏÐÑÒÓÔÕÖ×ØÙÚÛÜ", "aaaaaaac-eeeeiiiidnoooooxouuuu"},
		{"ÝÞßàáâãäåæçèéêëìíîïðñòóôõöøùúû", "ythssaaaaaaaceeeeiiiidnoooooouuu"},
		{"üýþÿĀāĂăĄąĆćĈĉĊċČčĎďĐđĒēĔĕĖėĘę", "uythyaaaaaaccccccccddddeeeeeeee"},
		{"ĚěĜĝĞğĠġĢģĤĥĦħĨĩĪīĬĭĮįİıĲĳĴĵĶķ", "eegggggggghhhhiiiiiiiiiiijijjjkk"},
		{"ĸĹĺĻļĽľĿŀŁłŃńŅņŇňŉŊŋŌōŎŏŐőŒœŔŕ", "kllllllllllnnnnnnnngngoooooooeoerr"},
		{"ŖŗŘřŚśŜŝŞşŠšŢţŤťŦŧŨũŪūŬŭŮůŰűŲų", "rrrrssssssssttttttuuuuuuuuuuuu"},
		{"ŴŵŶŷŸŹźŻżŽž", "wwyyyzzzzzz"},

		// should ignore all chars not included in a-z, åäö, -, 0-9
		{"kale8^79'0-", "kale8790"},

		// should convert diacritical characters
		{"Dürén Ibrahimović", "duren-ibrahimovic"},

		// does not preserve åäö
		{"ÅÄÖåäö", "aaoaao"},

		// converts `¨´^
		{"ÈÉËÊèéëê", "eeeeeeee"},
		{"ÀÁÂàáâ", "aaaaaa"},
		{"Üü", "uu"},
		{"ČĆÇčćç", "cccccc"},
		{"Ññ", "nn"},
		{"Ïï", "ii"},
		{"ÆØæø", "aoao"},

		// converts _ to -
		{"Let's_Dance", "lets-dance"},
		{"N___F5__9hf3m2iDyO4F__rjyD", "n-f5-9hf3m2idyo4f-rjyd"},

		// removes -- from name
		{"Let's -- da-da-dance", "lets-da-da-dance"},

		// removes surrounding and double space from name and tag
		{" Fångarna     på  fortet   ", "fangarna-pa-fortet"},

		// removes –, —
		{"Arn – Tempelriddaren", "arn-tempelriddaren"},
		{"Arn — Tempelriddaren", "arn-tempelriddaren"},
	} {
		if got := Case(tt.in); got != tt.want {
			t.Errorf(`nid.Case(%q) = %q, want %q`, tt.in, got, tt.want)
		}
	}
}

func ExampleCase() {
	fmt.Println(Case("Let's_Dance "))
	// Output: lets-dance
}

func BenchmarkPartialPrepare(b *testing.B) {
	for i := 0; i < b.N; i++ {
		prepare("Dürén Ibrahimović")
	}
}

func BenchmarkPartialPrepareSquish(b *testing.B) {
	for i := 0; i < b.N; i++ {
		squish(prepare("Dürén Ibrahimović"))
	}
}

func BenchmarkPartialPrepareSquishTransliterate(b *testing.B) {
	for i := 0; i < b.N; i++ {
		transliterate(squish(prepare("Dürén Ibrahimović")))
	}
}

func BenchmarkPartialPrepareSquishTransliterateStrip(b *testing.B) {
	for i := 0; i < b.N; i++ {
		strip(transliterate(squish(prepare("Dürén Ibrahimović"))))
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
	}{
		// accepts nothingness
		{"", true},

		// does not accept åäö
		{"räksmörgås", false},

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
		if got := squish(tt.in); got != tt.want {
			t.Errorf(`squish(%q) = %q, want %q`, tt.in, got, tt.want)
		}
	}
}
