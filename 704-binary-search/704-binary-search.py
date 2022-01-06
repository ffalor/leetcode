class Solution(object):
    def search(self, nums, target):
        """
        :type nums: List[int]
        :type target: int
        :rtype: int
        """
        
        left = 0
        right = len(nums) - 1
        
        while left <= right:
            
            pivot = left+((right-left)>>1)
                
            if nums[pivot] == target:
                return pivot
            
            if nums[pivot] < target:
                left = pivot + 1
            else:
                right = pivot - 1
                
                
            
        return -1