func quickSort(nums []int) {
    if  len(nums) > 1 {

        pivot := fetchPivot(nums)
        quickSort(nums[0:pivot])
        quickSort(nums[pivot+1:])
    }
    
    return
    
}

func fetchPivot(nums []int) int {
    pivot := 0
    
    i := 0
    j := len(nums) - 1
    for i < j {
        for ; nums[pivot] >= nums[i]; i++ {
            if i == len(nums) - 1 {
               break;   
            }
        }
     

        for ; nums[pivot] < nums[j]; j-- {
            if j == 0 {
                break;
            }
        }
        
	    if i < j {
        	nums[i], nums[j] = nums[j], nums[i]
	    }
    }
  
    nums[pivot], nums[j] = nums[j], nums[pivot]
    return j;
}
