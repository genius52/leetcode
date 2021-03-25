#include <vector>
using namespace std;

//Input: people = [[7,0],[4,4],[7,1],[5,0],[6,1],[5,2]]
//Output: [[5,0],[7,0],[5,2],[6,1],[4,4],[7,1]]
class Solution_406 {
public:
    vector<vector<int>> reconstructQueue(vector<vector<int>>& people) {
        std::sort(people.begin(),people.end(),[](const std::vector<int>& p1, const std::vector<int>& p2){
            if(p1[0] > p2[0])
                return true;
            if(p1[0] == p2[0] && p1[1] < p2[1])
                return true;
            return false;
        });
        int len = people.size();
        std::vector<std::vector<int>> res;
        for(int i = 0;i < len;i++){
            res.insert(res.begin() + people[i][1],people[i]);
        }
        return res;
    }
};