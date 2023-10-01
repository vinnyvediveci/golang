Given an array of integers nums and an integer target, return indices of the two numbers such that they add up to target.

You may assume that each input would have exactly one solution, and you may not use the same element twice.

You can return the answer in any order.

 

Example 1:

Input: nums = [2,7,11,15], target = 9
Output: [0,1]
Explanation: Because nums[0] + nums[1] == 9, we return [0, 1].
Example 2:

Input: nums = [3,2,4], target = 6
Output: [1,2]
Example 3:

Input: nums = [3,3], target = 6
Output: [0,1]
 

Constraints:

2 <= nums.length <= 104
-109 <= nums[i] <= 109
-109 <= target <= 109
Only one valid answer exists.
 

Follow-up: Can you come up with an algorithm that is less than O(n2) time complexity?



Solution 1: Brute Force
Time: O(n^2)
Space: O(1)

// Time: O(n*n) = O(n^2)
// Space: O(1)
func twoSum(nums []int, target int) []int {
    // Time: O(n)
    for i := 0; i < len(nums)-1; i++ {
        // Time: O(n)
        for j := i+1; j < len(nums); j++ {
			// Time: O(1)
            if nums[i] + nums[j] == target {
                return []int{i, j}
            }
        }
    }
    return []int{}
}



Solution 2: Hashmap/Set/Dictionary
Time: O(n)
Space: O(n)

// Time: O(n+n) = O(2n) = O(n)
// Space: O(n)
func twoSum(nums []int, target int) []int {
    // Space: O(n)
    m := make(map[int]int, 0)

    // Time: O(n)
    for idx, num := range nums {
        m[num] = idx
    }

    // Time: O(n)
    for indexX, num := range nums {
        // Time: O(1)
        if indexY, ok := m[target - num]; ok && indexX != indexY {
            return []int{indexX, indexY}
        }
    }

    return []int{}
}