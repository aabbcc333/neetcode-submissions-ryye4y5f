func findDuplicate(nums []int) int {
	result := 0
	sort.Ints(nums)
    for i := range nums{
		if nums[i+1] == nums[i]{
         result = nums[i+1]
		 break
		}
	}
	return result
}
