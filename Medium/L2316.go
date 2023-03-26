package Medium

class Solution {
public:
    void dfs(vector<int>adj[],int src,vector<bool>&vis,int &counter){
        if(vis[src]) return;
        vis[src]=true;
        counter++;
        for(auto ele:adj[src]){
            if(!vis[ele]){
                dfs(adj,ele,vis,counter);
            }
        }

    }
    long long countPairs(int n, vector<vector<int>>& edges) {
        vector<int>adj[n];
        for(auto ele:edges){
            adj[ele[0]].push_back(ele[1]);
            adj[ele[1]].push_back(ele[0]);
        }
        long long res = 0;
        vector<bool>vis(n,false);
        vector<int>temp;
        for(int i = 0;i<n;i++){
            if(!vis[i]){
                int counter = 0;
                dfs(adj,i,vis,counter);
               temp.push_back(counter);
            }
        }
        int total = 0;
        for(int i = 0;i<temp.size();i++){
            res+=(long)((long)temp[i]*(long)(n-total-temp[i]));
            total+=temp[i];
        }
        return res;
    }
};
