# The isBadVersion API is already defined for you.
# @param version, an integer
# @return a bool
# def isBadVersion(version):

class Solution(object):
    def firstBadVersion(self, n):
        """
        :type n: int
        :rtype: int
        """
        left, right = 0, n
        
        while left < right:
            pivot = (right+left) >> 1
            
            if isBadVersion(pivot):
                right = pivot
            else:
                left = pivot + 1
        
        return left