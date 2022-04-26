#include <vector>
#include <unordered_map>
using namespace std;

class TreeAncestor {
    std::vector<std::vector<int>> record;//record[i][j]: node i's jth ancestor
public:
    TreeAncestor(int n, vector<int>& parent) {
        record.resize(n);
        for(int i = 0;i < n;i++){
            record[i].push_back(parent[i]);
        }

    }

    int getKthAncestor(int node, int k) {

    }
};