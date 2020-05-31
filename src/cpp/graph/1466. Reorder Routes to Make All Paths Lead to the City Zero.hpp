#include <vector>
#include <deque>
#include <unordered_map>
#include <unordered_set>
using namespace std;

class Solution_1466 {
public:
    int minReorder(int n, vector<vector<int>>& connections) {
        std::deque<int> q;
        q.push_back(0);
        std::deque<int> copy_q;
        std::vector<bool> visited;
        visited.resize(n);
        std::unordered_map<int, std::unordered_set<int>> distance;
        std::unordered_map<int, vector<int>> graph;
        for (auto it = connections.begin(); it != connections.end(); it++) {
            graph[(*it)[0]].push_back((*it)[1]);
            distance[(*it)[0]].insert((*it)[1]);
            distance[(*it)[1]].insert((*it)[0]);
        }
        int cnt = 0;
        while (!q.empty()) {
            while (!q.empty()) {
                int point = q.front();
                vector<int> target = graph[point];
                for (auto it = target.begin(); it != target.end(); it++) {
                    if (!visited[*it])
                        cnt++;
                }
                graph.erase(q.front());
                q.pop_front();
                visited[point] = true;
                for (auto it = distance[point].begin(); it != distance[point].end(); it++) {
                    if(!visited[*it])
                        copy_q.push_back(*it);
                }
            }
            q = copy_q;
            copy_q.clear();
        }
        return cnt;
    }
};