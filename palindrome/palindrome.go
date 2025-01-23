package palindrome

import (
	"strings"
	"math"
)

func IsPalindrome(s string) bool {
    sLen := len(s)

    if sLen < 1 {
        return true
    }

    var sb strings.Builder
    var sanitizedS strings.Builder

    for i := sLen-1; i >=0; i-- {
        if isAlphaNumeric(s[i]) {
            sb.WriteByte(s[i])
        }
    }

    for i := 0; i < sLen; i++ {
        if isAlphaNumeric(s[i]) {
            sanitizedS.WriteByte(s[i])
        }
    }

    sStraight := strings.ToLower(sanitizedS.String())
    sReversed := strings.ToLower(sb.String())
    
    if len(sStraight) != len(sReversed) {
        return false;
    }

    mLen := math.Min(float64(len(sStraight)), float64(len(sReversed)))

    for i := 0; i < int(mLen); i++ {
        if sStraight[i] != sReversed[i] {
            return false;
        }
    }
    
    return true;
}

func isAlphaNumeric(ch byte) bool {
    return (ch >= 'a' && ch <= 'z') || 
            (ch >= 'A' && ch <= 'Z') || 
            (ch >= '0' && ch <= '9')
}