class Solution:
    def getOrder(self, tasks: List[List[int]]) -> List[int]:
        dic=defaultdict(list)
        
        for i in range(len(tasks)):
            dic[tasks[i][0]].append((tasks[i][1],i))
        
        
        ans=[]
        keys=sorted(dic.keys())
        
        while keys:
            k=keys.pop(0)
            pq=dic[k]
            heapq.heapify(pq)
            time=k
            
            while pq:
                p_time,ind=heapq.heappop(pq)
                ans.append(ind)
                time+=p_time
                while keys:
                    if keys[0]>time:
                        break
                    for item in dic[keys.pop(0)]:
                        heapq.heappush(pq,item)
        return ans
