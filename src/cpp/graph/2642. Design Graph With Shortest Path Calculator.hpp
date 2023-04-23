#include <vector>
#include <unordered_map>
using namespace std;

class Graph {
    //std::unordered_map<int,std::unordered_map<int,int>> from_to_weight_;
    std::vector<std::vector<int>> from_to_;
    int n_ = 0;
    int invalid_ = 1 << 31 - 1;
public:
    Graph(int n, vector<vector<int>>& edges) {
        n_ = n;
        from_to_.resize(n);
        for(int i = 0;i < n;i++){
            from_to_[i].resize(n);
            for(int j = 0;j < n;j++){
                if(i != j)
                    from_to_[i][j] = invalid_;
            }
        }
        for(auto edge : edges){
            int from = edge[0];
            int to = edge[1];
            int weight = edge[2];
            from_to_[from][to] = min(from_to_[from][to],weight);
        }
        for(int mid = 0;mid < n;mid++) {
            for (int start = 0; start < n; start++) {
                for (int end = 0; end < n; end++) {
                    if (start == end)
                        continue;
                    if(from_to_[start][mid] != invalid_ && from_to_[mid][end] != invalid_){
                        from_to_[start][end] = min(from_to_[start][end],from_to_[start][mid] + from_to_[mid][end]);
                    }
                }
            }
        }
    }

    void addEdge(vector<int> edge) {
        int from = edge[0];
        int to = edge[1];
        int weight = edge[2];
        from_to_[from][to] = min(from_to_[from][to],weight);
        for (int start = 0; start < n_; start++) {
            for (int end = 0; end < n_; end++) {
                if (start == end)
                    continue;
                if(from_to_[start][from] != invalid_ && from_to_[from][end] != invalid_){
                    from_to_[start][end] = min(from_to_[start][end],from_to_[start][from] + from_to_[from][end]);
                }
            }
        }
        for (int start = 0; start < n_; start++) {
            for (int end = 0; end < n_; end++) {
                if (start == end)
                    continue;
                if(from_to_[start][to] != invalid_ && from_to_[to][end] != invalid_){
                    from_to_[start][end] = min(from_to_[start][end],from_to_[start][to] + from_to_[to][end]);
                }
            }
        }
    }

    int shortestPath(int node1, int node2) {
        if(from_to_[node1][node2] != invalid_){
            return from_to_[node1][node2];
        }else{
            return -1;
        }
    }
};