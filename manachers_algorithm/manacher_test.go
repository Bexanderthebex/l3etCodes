package manachers_algorithm

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMain(m *testing.M) {
	m.Run()
}

func TestLongestPalindrome1(t *testing.T) {
	testString := "xabaxbbbbb"

	expectedAnswer := "xabax"
	ans := longestPalindrome(testString)

	assert.Equal(t, expectedAnswer, ans)
}

func TestLongestPalindrome2(t *testing.T) {
	testString := "cbbd"

	expectedAnswer := "bb"
	ans := longestPalindrome(testString)

	assert.Equal(t, expectedAnswer, ans)
}

func TestLongestPalindrome3(t *testing.T) {
	testString := "ccc"

	expectedAnswer := "ccc"
	ans := longestPalindrome(testString)

	assert.Equal(t, expectedAnswer, ans)
}

func TestLongestPalindrome4(t *testing.T) {
	testString := "kyyrjtdplseovzwjkykrjwhxquwxsfsorjiumvxjhjmgeueafubtonhlerrgsgohfosqssmizcuqryqomsipovhhodpfyudtusjhonlqabhxfahfcjqxyckycstcqwxvicwkjeuboerkmjshfgiglceycmycadpnvoeaurqatesivajoqdilynbcihnidbizwkuaoegmytopzdmvvoewvhebqzskseeubnretjgnmyjwwgcooytfojeuzcuyhsznbcaiqpwcyusyyywqmmvqzvvceylnuwcbxybhqpvjumzomnabrjgcfaabqmiotlfojnyuolostmtacbwmwlqdfkbfikusuqtupdwdrjwqmuudbcvtpieiwteqbeyfyqejglmxofdjksqmzeugwvuniaxdrunyunnqpbnfbgqemvamaxuhjbyzqmhalrprhnindrkbopwbwsjeqrmyqipnqvjqzpjalqyfvaavyhytetllzupxjwozdfpmjhjlrnitnjgapzrakcqahaqetwllaaiadalmxgvpawqpgecojxfvcgxsbrldktufdrogkogbltcezflyctklpqrjymqzyzmtlssnavzcquytcskcnjzzrytsvawkavzboncxlhqfiofuohehaygxidxsofhmhzygklliovnwqbwwiiyarxtoihvjkdrzqsnmhdtdlpckuayhtfyirnhkrhbrwkdymjrjklonyggqnxhfvtkqxoicakzsxmgczpwhpkzcntkcwhkdkxvfnjbvjjoumczjyvdgkfukfuldolqnauvoyhoheoqvpwoisniv"

	expectedAnswer := "qahaq"
	ans := longestPalindrome(testString)

	assert.Equal(t, expectedAnswer, ans)
}
