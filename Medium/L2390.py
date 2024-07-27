class Solution:
    def removeStars(self, s: str) -> str:
        st = []

        for c in s:
            if c != '*':
                st.append(c)
            elif st:
                st.pop()

        return ''.join(st)
