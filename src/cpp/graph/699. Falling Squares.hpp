#include <vector>
#include <map>
using namespace std;

//Input: positions = [[1,2],[2,3],[6,1]]
//Output: [2,5,5]
//Explanation:
//After the first drop, the tallest stack is square 1 with a height of 2.
//After the second drop, the tallest stack is squares 1 and 2 with a height of 5.
//After the third drop, the tallest stack is still squares 1 and 2 with a height of 5.
//Thus, we return an answer of [2, 5, 5].
class Solution_699 {
public:
    vector<int> fallingSquares(vector<vector<int>>& positions) {
        int len = positions.size();
        std::vector<int> res(len);
        for(int i = 0;i < len;i++){
            int start = positions[i][0];
            int edge_len = positions[i][1];
            res[i] += edge_len;
            for(int j = i + 1;j < len;j++){
                int next_start = positions[j][0];
                int next_edge_len = positions[j][1];
                if((next_start + next_edge_len) > start && next_start < (start + edge_len)){
                    res[j] = max(res[j],res[i]);
                }
            }

        }
        int max_height = 0;
        for(int i = 0;i < len;i++){
            res[i] = max(res[i],max_height);
            max_height = res[i];
        }
        return res;
    }
};