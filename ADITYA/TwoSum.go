func twoSum(nums []int, target int) []int {
    
    for i:= 0; i < len(nums) ; i++{
        find := target - nums [i]
        for j := i+1; j < len(nums) ; j++{
            if (find == nums[j]){
             return []int{i,j}
            }
        }
    }
    return nil
}
