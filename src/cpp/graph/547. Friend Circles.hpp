#include <vector>
#include <queue>
#include <unordered_set>
#include <unordered_map>
using namespace std;

class Solution_547 {
public:
    int findCircleNum(vector<vector<int>>& M) {
        int len = M.size();
        std::unordered_map<int, std::unordered_set<int>> graph;
        for (int i = 0; i < len; i++) {
            for (int j = 0; j < len; j++) {
                if (i == j)
                    continue;
                if (M[i][j] == 1)
                    graph[i].insert(j);
            }
        }
        std::unordered_set<int> visited;
        int cycles = 0;
        for (int i = 0; i < len; i++) {
            if (graph.find(i) == graph.end()) {//'i' has no friends
                if (visited.find(i) == visited.end()) {
                    cycles++;
                    visited.insert(i);
                }
                continue;
            }
            if (visited.find(i) != visited.end()) {
                continue;
            }
            cycles++;
            std::queue<int> q;
            q.push(i);
            visited.insert(i);
            while (!q.empty()) {
                int l = q.size();
                for (int i = 0; i < l; i++) {
                    int top = q.front();
                    q.pop();
                    for (auto r : graph[top]) {
                        if (visited.find(r) == visited.end()) {
                            q.push(r);
                            visited.insert(r);
                        }
                    }
                }
            }
        }
        return cycles;
    }
};