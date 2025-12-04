package leetcode

func SingleNumber(nums []int) int {
    numSet := make(map[int]struct{})

    for _, num := range nums {
        
        _, found := numSet[num]
        
        if found {
            delete(numSet, num)
        } else {
            numSet[num] = struct{}{}
        }
    }
    
    for key := range numSet {
        return key
    }
    
    return -1
}