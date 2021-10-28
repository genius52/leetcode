#include <queue>
#include <vector>
using namespace std;

class Solution_1494 {
public:
    int minNumberOfSemesters(int n, vector<vector<int>>& relations, int k){
        return 0;
    }
    //priority_queue + depth is wrong solution
//    int dfs_minNumberOfSemesters(int n,std::vector<std::vector<bool>>& graph,int cur,std::vector<std::pair<int,int>>& node_depth){
//        if(node_depth[cur - 1].first != -1)
//            return node_depth[cur - 1].first;
//        int depth = 0;
//        for(int i = 1;i <= n;i++){
//            if(!graph[cur][i])
//                continue;
//            depth = max(depth,1 + dfs_minNumberOfSemesters(n,graph,i,node_depth));
//        }
//        node_depth[cur - 1].first = depth;
//        return depth;
//    }
//    int minNumberOfSemesters(int n, vector<vector<int>>& relations, int k) {
//        std::vector<int> in(n + 1);
//        std::vector<std::vector<bool>> graph(n + 1,std::vector<bool>(n + 1));
//        std::vector<std::pair<int,int>> node_depth(n);//depth,index
//        for(int i = 0;i < n;i++){
//            node_depth[i] = {-1,i + 1};
//        }
//        for(auto edge : relations){
//            graph[edge[0]][edge[1]] = true;
//            in[edge[1]]++;
//        }
//        for(int i = 1;i <= n;i++){
//            dfs_minNumberOfSemesters(n,graph,i,node_depth);
//        }
//        std::vector<bool> visited(n + 1);
//        std::priority_queue<std::pair<int,int>,std::deque<std::pair<int,int>>, std::less<std::pair<int,int>>> q;
//        //将入度为0，加入队列,深度最大的节点优先访问
//        for(int i = 0;i < n;i++){
//            if(in[node_depth[i].second] == 0){
//                q.push(node_depth[i]);
//                visited[node_depth[i].second] = true;
//            }
//        }
//        int steps = 0;
//        while (!q.empty()){
//            std::vector<std::pair<int,int>> next;
//            for(int times = 0;times < k && !q.empty();times++){
//                std::pair<int,int> cur = q.top();
//                q.pop();
//                for(int j = 1;j <= n;j++){
//                    if(graph[cur.second][j]){
//                        in[j]--;
//                        if(in[j] == 0 && !visited[j]){
//                            visited[j] = true;
//                            next.push_back(node_depth[j - 1]);
//                        }
//                    }
//                }
//            }
//            for (int i = 0; i < next.size(); ++i) {
//                q.push(next[i]);
//            }
//            steps++;
//        }
//        return steps;
//    }
};