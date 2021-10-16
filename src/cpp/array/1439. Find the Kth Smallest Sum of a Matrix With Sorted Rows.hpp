#include <vector>
#include <queue>
#include <set>
using namespace std;

//Input: mat = [[1,3,11],[2,4,6]], k = 5
//Output: 7
//Explanation: Choosing one element from each row, the first k smallest sum are:
//[1,2], [1,4], [3,2], [3,4], [1,6]. Where the 5th sum is 7.
class Solution_1439 {
public:
    int kthSmallest(vector<vector<int>>& mat, int k) {
        int len = mat.size();
        int columns = mat[0].size();
        std::priority_queue<std::pair<int,std::vector<int>>,std::vector<std::pair<int,std::vector<int>>>,
                std::greater<std::pair<int,std::vector<int>>>> q;
        int sum = 0;
        std::vector<int> idx(len);
        std::set<std::vector<int>> visited;
        for(int i = 0;i < len;i++){
            sum += mat[i][0];
            idx[i] = 0;
        }
        q.push({sum,idx});
        visited.insert(idx);
        while (q.size() > 0){
            int cur_len = q.size();
            for(int j = 0;j < cur_len;j++){
                auto cur_sum = q.top().first;
                auto cur_idx = q.top().second;
                q.pop();
                if(k == 1)
                    return cur_sum;
                for(int i = 0;i < len;i++){
                    if(cur_idx[i] + 1 < columns){
                        cur_idx[i]++;
                        if(visited.find(cur_idx) == visited.end()){
                            int next_sum = cur_sum - mat[i][cur_idx[i] - 1] + mat[i][cur_idx[i]];
                            visited.insert(cur_idx);
                            q.push({next_sum,cur_idx});
                        }
                        cur_idx[i]--;
                    }
                }
                k--;
            }
        }
        return 0;
    }
};