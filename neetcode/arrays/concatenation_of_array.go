// Problem: Concatenation of Array
// Link: https://neetcode.io/problems/concatenation-of-array/question?list=neetcode250
//
// Given an integer array nums of length n, return an array ans of length 2n
// where ans[i] == nums[i] and ans[i+n] == nums[i]. (Concatenate nums with itself.)
//
// Examples:
// nums = [1, 4, 1, 2]        -> [1, 4, 1, 2, 1, 4, 1, 2]
// nums = [22, 21, 20, 1]     -> [22, 21, 20, 1, 22, 21, 20, 1]
//
// Approach: Preallocate and append nums twice.
// Complexity: Time O(n), Space O(n)

package main

func getConcatenation(nums []int) []int {
    // Preallocate capacity for 2*n; start empty and append.
    res := make([]int, 0, 2*len(nums))
    res = append(res, nums...)
    res = append(res, nums...)
    return res
}
