// Link: https://neetcode.io/problems/contains-duplicate
//
// --- Problem Statement ---
// Given an integer array nums, return true if any value
// appears more than once in the array. Otherwise, return false.
//
// --- Examples ---
// Input:  nums = [1, 2, 3, 3]
// Output: true
//
// Input:  nums = [1, 2, 3, 4]
// Output: false
//
// --- Constraints ---
// 1 <= nums.length <= 100000
// -10^9 <= nums[i] <= 10^9
//
// --- Approach ---
// Use a hash map to track numbers that have already been seen.
// If a number is encountered again, return true immediately.
//
// --- Complexity ---
// Time:  O(n)
// Space: O(n)

package main

func hasDuplicate(nums []int) bool {
    seen := make(map[int]bool)

    for _, num := range nums {
        if seen[num] {
            return true
        }
        seen[num] = true
    }

    return false
}
