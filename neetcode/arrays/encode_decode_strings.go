// Problem: Encode and Decode Strings
// Link: https://neetcode.io/problems/string-encode-and-decode/question?list=neetcode250
//
// --- Intuition ---
// We need a way to join multiple strings into one string such that we can
// unambiguously split them back into the originals, even if they contain
// any ASCII characters (including delimiters like '#', digits, spaces, etc.).
//
// A robust pattern for this is *length-prefix encoding*:
// For each string s, write: "<length>#<content>". For example:
//   "Hello" -> "5#Hello"
//   ""      -> "0#"
// Concatenating records gives a single stream like:
//   ["Hello","World"] -> "5#Hello5#World"
// During decoding, we read the number before '#', then slice exactly that many
// bytes after the '#'. This avoids ambiguity because we rely on the explicit
// length, not on the content (which may contain any characters).
//
// --- Approach ---
// Encode: iterate over strs, and for each string:
//   - write its byte length, then a '#', then the string itself.
//   - Use strings.Builder for efficient concatenation.
// Decode: scan the encoded string left-to-right:
//   - parse digits until '#', convert to length L,
//   - slice the next L bytes as the string,
//   - repeat until the end.
// Add simple bounds checks to guard against malformed input.
//
// --- Complexity ---
// Let N be the number of strings and M be the total length of all strings.
// Encode:  O(M) time, O(M) output size.
// Decode:  O(M) time, produces O(M) total string data.
// Space overhead is linear in the total data processed.
//

package main

import (
    "strconv"
    "strings"
)

type Solution struct{}

// Encode concatenates strings using length-prefix records: "<len>#<string>".
// Example: ["Hello","World"] -> "5#Hello5#World".
// This scheme is safe for any ASCII content because we don't rely on a
// special delimiter inside the content; we trust the explicit length.
func (s *Solution) Encode(strs []string) string {
    var b strings.Builder
    for _, str := range strs {
        // Write the length in decimal, then a separator '#', then the content.
        b.WriteString(strconv.Itoa(len(str)))
        b.WriteByte('#')
        b.WriteString(str)
    }
    return b.String()
}

// Decode parses the encoded stream by reading the decimal length up to '#',
// then extracting exactly that many bytes. If the input is malformed (e.g.,
// missing '#', invalid length, or length exceeding remaining data), it returns nil.
func (s *Solution) Decode(encoded string) []string {
    var res []string
    for i := 0; i < len(encoded); {
        // Find the position of the next '#'
        j := i
        for j < len(encoded) && encoded[j] != '#' {
            j++
        }
        // Malformed: no '#' found
        if j == len(encoded) {
            return nil
        }
        // Parse the length [i:j]
        length, err := strconv.Atoi(encoded[i:j])
        if err != nil || length < 0 {
            return nil
        }
        // Slice out the next 'length' bytes after '#'
        start := j + 1
        end := start + length
        if end > len(encoded) {
            return nil
        }
        res = append(res, encoded[start:end])
        // Advance to the next record
        i = end
    }
    return res
}
