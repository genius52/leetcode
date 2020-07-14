#include <vector>
#include <unordered_map>
using namespace std;

class TreeAncestor {
    std::vector<std::vector<int>> record;
public:
    TreeAncestor(int n, vector<int>& parent) {
        record.resize(n);
        for (int i = 0; i < n; i++) {
            auto par = parent[i];
            while (par != -1) {
                record[i].push_back(par);
                par = parent[par];
            }
        }
    }

    int getKthAncestor(int node, int k) {
        int len = record[node].size();
        if (len < k)
            return -1;
        else
            return record[node][k - 1];
    }
};