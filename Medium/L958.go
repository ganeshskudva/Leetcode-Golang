package Medium

# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right
class Solution:
    def isCompleteTree(self, root: Optional[TreeNode]) -> bool:
        end, q = False, deque()
        q.append(root)
        
        while q:
            curr = q.popleft()
            
            if not curr:
                end = True
            else:
                if end:
                    return False
                q.append(curr.left)
                q.append(curr.right)
        
        return True
