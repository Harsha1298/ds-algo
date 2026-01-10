
// topKFrequent returns the k most frequent elements in nums.
// It uses a frequency map and a "bucket sort" idea where
// freq[i] holds all numbers that appear exactly i times.
func topKFrequent(nums []int, k int) []int {
    // Count occurrences of each number: value -> frequency
    count := make(map[int]int)

    // Buckets: index is frequency, value is list of numbers with that frequency.
    // Max possible frequency for any number is len(nums),
    // so we create len(nums)+1 buckets (including index 0 which stays unused).
    freq := make([][]int, len(nums)+1)

    // Build the frequency map.
    for _, num := range nums {
        count[num]++
    }

    // Place each number into the bucket corresponding to its frequency.
    for num, cnt := range count {
        freq[cnt] = append(freq[cnt], num)
    }

    // Collect results by scanning buckets from high frequency to low.
    res := []int{}
    for i := len(freq) - 1; i > 0; i-- { // start from the highest possible frequency
        for _, num := range freq[i] {    // iterate all numbers that appear i times
            res = append(res, num)
            if len(res) == k {           // stop once we have k elements
                return res
            }
        }
    }

    // If k is larger than the number of unique elements,
    // we return whatever we collected (could be < k).
    return res
}
