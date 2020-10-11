#include <vector>
using namespace std;
#include <unordered_map>
#include <unordered_set>
using namespace std;

class Solution_5536 {
public:
    int maximalNetworkRank(int n, vector<vector<int>>& roads) {
        std::unordered_map<int, std::unordered_set<int>> graph;
        int len = roads.size();
        for (size_t i = 0; i < len; i++)
        {
            graph[roads[i][0]].insert(roads[i][1]);
            graph[roads[i][1]].insert(roads[i][0]);
        }
        int res = 0;
        for (int i = 0; i < n; i++) {
            int l1 = 0;
            if (graph.find(i) != graph.end()) {
                l1 = graph[i].size();
            }
            else {
                continue;
            }
            for (int j = i + 1; j < n; j++) {
                int l2 = 0;
                if (graph.find(j) != graph.end()) {
                    l2 = graph[j].size();
                    if (graph[j].find(i) != graph[j].end()) {
                        l2--;
                    }
                }
                if (l1 + l2 > res)
                    res = l1 + l2;
            }
        }
        return res;
    }
};