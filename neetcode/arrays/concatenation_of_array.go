// Problem: Concatenation of Array
// Link: https://neetcode.io/problems/concatenation-of-array/question?list=neetcode250
//
// --- Problem Statement ---
// You are given an integer array nums of length n.
// Create an array ans of length 2n where:
//   - ans[i]     == nums[i]
//   - ans[i + n] == nums[i]
// for all 0 <= i < n (0-indexed).
//
// Specifically, ans is the concatenation of two nums arrays.
// Return the array ans.
//
// --- Examples ---
// Example 1:
//   Input:  nums = [1, 4, 1, 2]
//   Output: [1, 4, 1, 2, 1, 4, 1, 2]
//
// Example 2:
//   Input:  nums = [22, 21, 20, 1]
//   Output: [22, 21, 20, 1, 22, 21, 20, 1]
//
// --- Constraints ---
//   1 <= nums.length <= 1000
//   1 <= nums[i]      <= 1000
//
// --- Approach ---
// We need to return nums concatenated with itself. In Go, this can be done
// efficiently by preallocating a slice with capacity 2*n and appending nums twice.
// Preallocating capacity avoids intermediate allocations during append.
//
// --- Complexity ---
// Time:  O(n)   — each element is appended twice overall.
// Space: O(n)   — the result slice has length 2*n.
//
// --- Notes ---
// - This is the idiomatic Go pattern for building slices dynamically:
//     make([]T, 0, expectedCapacity) + append
// - Alternatively, you can use copy() twice into a pre-sized slice (see alt version below).

package main

// getConcatenation returns a new slice that is nums concatenated with itself.
// Example: [1,2,3] -> [1,2,3,1,2,3]
func getConcatenation(nums []int) []int {
    // Preallocate capacity for 2*n elements to avoid reallocations.
    // Start with length 0 because we plan to grow the slice via append.
    res := make([]int, 0, 2*len(nums))

    // Append all elements of nums once.
    res = append(res, nums...)

    // Append nums again to complete the concatenation.
    res = append(res, nums...)

    return res
}

/*
--- Alternative Implementation (using copy) ---

func getConcatenation(nums []int) []int {
    // Allocate the final length upfront (2*n).
    res := make([]int, 2*len(nums))
    // Copy nums into the first half.
    copy(res, nums)
    // Copy nums into the second half.
    copy(res[len(nums):], nums)
    return res
}
*/

/*
--- Optional: Simple sanity check (uncomment to run) ---

import "fmt"

func main() {
    fmt.Println(getConcatenation([]int{1, 4, 1, 2}))        // [1 4 1 2 1 4 1 2]
    fmt.Println(getConcatenation([]int{22, 21, 20, 1}))     // [22 21 20 1 22 21 20 1]
}
*/
