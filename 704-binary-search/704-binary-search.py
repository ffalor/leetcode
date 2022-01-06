class Solution(object):
    def search(self, nums, target):
        """
        :type nums: List[int]
        :type target: int
        :rtype: int
        """
        
        left, right = 0, len(nums) - 1

        while left <= right:
            
            pivot = left + (right-left)//2
                
            if nums[pivot] == target:
                return pivot
            
            if nums[pivot] < target:
                left = pivot + 1
            else:
                right = pivot - 1
                
                
            
        return -1