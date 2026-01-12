// Problem: Range Sum Query 2D - Immutable
// Link: https://neetcode.io/problems/range-sum-query-2d-immutable/question
//
// --- Problem Statement ---
// You are given a 2D matrix. Design a data structure that supports
// sumRegion(row1, col1, row2, col2), which returns the sum of elements
// inside the rectangle defined by its top-left (row1, col1) and
// bottom-right (row2, col2).
//
// The solution must support sumRegion queries in O(1) time.
//
// --- Intuition ---
// A naive approach would iterate over every cell inside the rectangle
// for each query, costing O(rows * cols) time per query. Since there can
// be many queries, this is too slow.
//
// Instead, we preprocess the matrix using a 2D prefix sum array.
// Each cell in the prefix matrix stores the sum of all elements
// from (0,0) up to that cell.
//
// Using inclusion–exclusion, we can compute the sum of any rectangle
// by combining four prefix sums, giving O(1) query time after O(m*n)
// preprocessing.
//
// --- Prefix Sum Definition ---
// Let sumMat[r][c] be the sum of all elements in the rectangle
// from (0,0) to (r-1,c-1) in the original matrix.
//
// --- Inclusion–Exclusion Formula ---
// sumRegion = bottomRight
//           - above
//           - left
//           + topLeft
//
// --- Complexity ---
// Preprocessing: O(m * n)
// Query:         O(1)
// Space:         O(m * n)

package main

type NumMatrix struct {
    sumMat [][]int
}

// Constructor builds the 2D prefix sum matrix.
// An extra row and column of zeros are added to simplify bounds handling.
func Constructor(matrix [][]int) NumMatrix {
    rows, cols := len(matrix), len(matrix[0])

    // sumMat has size (rows+1) x (cols+1)
    sumMat := make([][]int, rows+1)
    for i := range sumMat {
        sumMat[i] = make([]int, cols+1)
    }

    // Build prefix sums row by row
    for r := 0; r < rows; r++ {
        rowPrefix := 0
        for c := 0; c < cols; c++ {
            rowPrefix += matrix[r][c]
            // Current sum = left row sum + sum from rows above
            sumMat[r+1][c+1] = rowPrefix + sumMat[r][c+1]
        }
    }

    return NumMatrix{sumMat: sumMat}
}

// SumRegion returns the sum of elements within the rectangle
// defined by (row1, col1) and (row2, col2), inclusive.
func (this *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
    // Shift indices by +1 to match prefix sum coordinates
    row1++
    col1++
    row2++
    col2++

    bottomRight := this.sumMat[row2][col2]
    above := this.sumMat[row1-1][col2]
    left := this.sumMat[row2][col1-1]
    topLeft := this.sumMat[row1-1][col1-1]

    return bottomRight - above - left + topLeft
}

/*
Example Usage:

numMatrix := Constructor([][]int{
    {3, 0, 1, 4, 2},
    {5, 6, 3, 2, 1},
    {1, 2, 0, 1, 5},
    {4, 1, 0, 1, 7},
    {1, 0, 3, 0, 5},
})

numMatrix.SumRegion(2, 1, 4, 3) // 8
numMatrix.SumRegion(1, 1, 2, 2) // 11
numMatrix.SumRegion(1, 2, 2, 4) // 12
*/
