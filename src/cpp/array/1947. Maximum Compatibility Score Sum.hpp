#include <vector>
using namespace std;

class Solution_1947 {
public:
    int max = 0;
    void perm(vector<vector<int>>& graph, vector<vector<int>>& students, vector<vector<int>>& mentors, int l1, int pos) {
        if (pos == l1) {
            int cur = 0;
            for (int i = 0; i < l1; i++) {
                cur += graph[i][i];
            }
            if (cur > max)
                max = cur;
            return;
        }
        for (int i = pos; i < l1; i++) {
            std::swap(students[i], students[pos]);
            std::swap(graph[i], graph[pos]);
            perm(graph, students, mentors, l1, pos + 1);
            std::swap(graph[i], graph[pos]);
            std::swap(students[i], students[pos]);
        }
        return;
    }
    int maxCompatibilitySum(vector<vector<int>>& students, vector<vector<int>>& mentors) {
        int l1 = students.size();
        int l2 = students[0].size();
        vector<vector<int>> graph(l1, std::vector<int>(l1));
        for (int i = 0; i < l1; i++) {
            for (int j = 0; j < l1; j++) {
                int same = 0;
                for (int k = 0; k < l2; k++) {
                    if (students[i][k] == mentors[j][k]) {
                        same++;
                    }
                }
                graph[i][j] = same;
            }
        }
        perm(graph, students, mentors, l1, 0);
        return max;
    }
};
