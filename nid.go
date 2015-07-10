/*

Package nid is used to create nids (slugs/tags)

Installation

Just go get the package:

    go get -u github.com/TV4/nid

Usage

A small usage example

    package main

    import (
    	"fmt"

    	"github.com/TV4/nid"
    )

    func main() {
    	fmt.Println(nid.Case("Let's_Dance ")) // lets-dance
    }

*/
package nid

import (
	"regexp"
	"strings"
)

var (
	validPattern  = regexp.MustCompile(`\A[0-9a-zåäö-]*\z`)
	squishPattern = regexp.MustCompile(`\s+`)
	stripPattern  = regexp.MustCompile(`[^0-9a-zåäö-]`)

	dashSpace = strings.NewReplacer("-", " ")
)

// Case returns a nid based on the input text
func Case(text string) string {
	if text == "" {
		return ""
	}

	return strip(transliterate(squish(prepare(text))))
}

// Possible checks if a candidate string is a possible nid
func Possible(candidate string) bool {
	return validPattern.MatchString(candidate)
}

func strip(s string) string {
	return stripPattern.ReplaceAllString(s, "")
}

func transliterate(s string) string {
	return transliterations.Replace(s)
}

func squish(s string) string {
	return squishPattern.ReplaceAllString(s, " ")
}

func prepare(s string) string {
	return strings.TrimSpace(dashSpace.Replace(strings.ToLower(s)))
}

var transliterations = strings.NewReplacer(
	" ", "-",
	"_", "-",
	"×", "x",
	"ß", "ss",
	"à", "a",
	"á", "a",
	"â", "a",
	"ã", "a",
	"ä", "ä",
	"å", "å",
	"æ", "ä",
	"ç", "c",
	"è", "e",
	"é", "e",
	"ê", "e",
	"ë", "e",
	"ì", "i",
	"í", "i",
	"î", "i",
	"ï", "i",
	"ð", "d",
	"ñ", "n",
	"ò", "o",
	"ó", "o",
	"ô", "o",
	"õ", "o",
	"ö", "ö",
	"ø", "ö",
	"ù", "u",
	"ú", "u",
	"û", "u",
	"ü", "u",
	"ý", "y",
	"þ", "th",
	"ÿ", "y",
	"ā", "a",
	"ă", "a",
	"ą", "a",
	"ć", "c",
	"ĉ", "c",
	"ċ", "c",
	"č", "c",
	"ď", "d",
	"đ", "d",
	"ē", "e",
	"ĕ", "e",
	"ė", "e",
	"ę", "e",
	"ě", "e",
	"ĝ", "g",
	"ğ", "g",
	"ġ", "g",
	"ģ", "g",
	"ĥ", "h",
	"ħ", "h",
	"ĩ", "i",
	"ī", "i",
	"ĭ", "i",
	"į", "i",
	"ı", "i",
	"ĳ", "ij",
	"ĵ", "j",
	"ķ", "k",
	"ĸ", "k",
	"ĺ", "l",
	"ļ", "l",
	"ľ", "l",
	"ŀ", "l",
	"ł", "l",
	"ń", "n",
	"ņ", "n",
	"ň", "n",
	"ŉ", "'n",
	"ŋ", "ng",
	"ō", "o",
	"ŏ", "o",
	"ő", "o",
	"œ", "oe",
	"ŕ", "r",
	"ŗ", "r",
	"ř", "r",
	"ś", "s",
	"ŝ", "s",
	"ş", "s",
	"š", "s",
	"ţ", "t",
	"ť", "t",
	"ŧ", "t",
	"ũ", "u",
	"ū", "u",
	"ŭ", "u",
	"ů", "u",
	"ű", "u",
	"ų", "u",
	"ŵ", "w",
	"ŷ", "y",
	"ź", "z",
	"ż", "z",
	"ž", "z",
)
