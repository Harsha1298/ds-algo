// Problem: Valid Anagram
// Link: https://neetcode.io/problems/valid-anagram
//
// Given two strings s and t, return true if t is an anagram of s.
// An anagram contains the same characters with the same frequencies,
// but possibly in a different order.
//
// Examples:
// s = "racecar", t = "carrace" -> true
// s = "jar", t = "jam"         -> false
//
// Constraints:
// s and t consist of lowercase English letters.
//
// Approach:
// Use a fixed-size frequency array for letters 'a' to 'z'.
// Increment counts for s and decrement for t, then verify all counts are zero.
//
// Complexity:
// Time:  O(n)
// Space: O(1)

package main

func isAnagram(s string, t string) bool {
    if len(s) != len(t) {
        return false
    }

    var count [26]int
    for i := 0; i < len(s); i++ {
        count[s[i]-'a']++
        count[t[i]-'a']--
    }

    for i := 0; i < 26; i++ {
        if count[i] != 0 {
            return false
        }
    }
    return true
}
