#include <vector>
#include <queue>
using namespace std;

class Solution_1489 {
public:
    int get_parent(std::vector<int>& groups,int index){
        if(groups[index] != index){
            groups[index] = get_parent(groups,groups[index]);
        }
        return groups[index];
    }
    //kruskal
    //1 将边按权重排序
    //2 每个节点设置组号
    //3 按权重从小到大遍历每条边，如果目标和来源属于不同组则合并，否则忽略(不然会形成环路)
    vector<vector<int>> findCriticalAndPseudoCriticalEdges(int n, vector<vector<int>>& edges){
        int len = edges.size();
        std::vector<std::vector<int>> graph(len);
        for (int i = 0; i < len; ++i) {
            graph[i] = {edges[i][0],edges[i][1],edges[i][2],i};
            //graph[i].push_back(i);
        }
        std::sort(graph.begin(),graph.end(),[](const std::vector<int>& v1,const std::vector<int>& v2){
            return v1[2] < v2[2];
        });
        std::vector<int> trace;
        std::vector<bool> visited(n);
        std::vector<int> groups(n);
        for (int i = 0; i < n; ++i) {
            groups[i] = i;
        }
        int min_distance = 0;

        for (int i = 0; i < len; ++i) {
            int from = graph[i][0];
            int to = graph[i][1];
            int dis = graph[i][2];
            int group1 = get_parent(groups,from);
            int group2 = get_parent(groups,to);
            if(group1 != group2){
                groups[to] = group1;//当前选中的节点，归属到来源节点的组
                min_distance += dis;
                visited[to] = true;
            }
        }
        std::vector<std::vector<int>> res;
        return res;
    }

    //prime
//把一个初始顶点的所有边加入优先队列
//取出最短的边，把这条边的另一个顶点相关的边加入队列
//再取出最小的边，重复下去，直到所有顶点加入过了
//最小生成树无法获取所有边的索引
//    bool is_critical(int n,vector<vector<int>>& edges,std::vector<std::vector<std::pair<int,int>>>& graph,int del_idx,int min_distance){
//        std::priority_queue<std::pair<int,std::pair<int,int>>,std::deque<std::pair<int,std::pair<int,int>>>,
//                std::greater<std::pair<int,std::pair<int,int>>>> q;
//        std::vector<bool> visited_node(n);
//        int cur_distance = 0;
//        int edge_cnt = 0;
//        visited_node[0] = true;
//        for(int j = 0;j < graph[0].size();j++){ //将i节点所有的边放入队列
//            if(graph[0][j].second == del_idx)
//                continue;
//            q.push({edges[graph[0][j].second][2],{graph[0][j].first,graph[0][j].second}});
//        }
//        while (!q.empty()){
//            auto cur = q.top();//找距离最小的
//            q.pop();
//            int dst_node = cur.second.first; //目标节点
//            int distance = cur.first;//距离
//            if(!visited_node[dst_node]){
//                visited_node[dst_node] = true;
//                cur_distance += distance;
//                if(cur_distance > min_distance)
//                    return true;
//                //result[i].push_back(cur.second.second);
//                edge_cnt++;
//                if(edge_cnt == n - 1){
//                    break;
//                }
//                for(int k = 0;k < graph[dst_node].size();k++){
//                    if(visited_node[graph[dst_node][k].first])//目标节点
//                        continue;
//                    if(graph[dst_node][k].second == del_idx)
//                        continue;
//                    q.push({edges[graph[dst_node][k].second][2],{graph[dst_node][k].first,graph[dst_node][k].second}});
//                }
//            }
//        }
//        if(edge_cnt != n - 1)
//            return true;
//        return cur_distance != min_distance;
//    }
//
//    vector<vector<int>> findCriticalAndPseudoCriticalEdges(int n, vector<vector<int>>& edges) {
//        std::vector<std::vector<std::pair<int,int>>> graph(n);//graph[i][j]: 点i指向
//        int len = edges.size();
//        for(int i = 0;i < len;i++){
//            graph[edges[i][0]].push_back({edges[i][1],i});//dst,index
//            graph[edges[i][1]].push_back({edges[i][0],i});
//        }
//        std::set<int> result;
//        int min_distance = 0;
//        //weight,idx
//        for (int i = 0; i < n; ++i) {
//            std::priority_queue<std::pair<int,std::pair<int,int>>,std::deque<std::pair<int,std::pair<int,int>>>,
//                    std::greater<std::pair<int,std::pair<int,int>>>> q;
//            std::vector<bool> visited_node(n);
//            int edge_cnt = 0;
//            int cur_distance = 0;
//            visited_node[i] = true;
//            for(int j = 0;j < graph[i].size();j++){ //将i节点所有的边放入队列
//                q.push({edges[graph[i][j].second][2],{graph[i][j].first,graph[i][j].second}});
//                //节点i的边{distance,from,to,index in edges}
//            }
//            while (!q.empty()){
//                auto cur = q.top();//找距离最小的
//                q.pop();
//                int dst_node = cur.second.first; //目标节点
//                int distance = cur.first;//距离
//                if(!visited_node[dst_node]){
//                    visited_node[dst_node] = true;
//                    cur_distance += distance;
//                    result.insert(cur.second.second);
//                    edge_cnt++;
//                    if(edge_cnt == n - 1){
//                        break;
//                    }
//                    for(int k = 0;k < graph[dst_node].size();k++){
//                        if(visited_node[graph[dst_node][k].first])//目标节点
//                            continue;
//                        q.push({edges[graph[dst_node][k].second][2],{graph[dst_node][k].first,graph[dst_node][k].second}});
//                    }
//                }
//            }
//            min_distance = cur_distance;
//        }
//        std::vector<std::vector<int>> res(2);
//        for (auto it = result.begin(); it != result.end(); ++it) {
//            bool ret = is_critical(n,edges,graph,*it,min_distance);
//            if(ret){
//                res[0].push_back(*it);
//            }else{
//                res[1].push_back(*it);
//            }
//        }
//        return res;
//    }
};