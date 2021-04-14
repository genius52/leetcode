#include <vector>
#include <map>
#include <string>
using namespace std;

class Solution_506 {
public:
    vector<string> findRelativeRanks(vector<int>& nums) {
        std::map<int,int,std::greater<int>> map;
        int len = nums.size();
        std::vector<std::string> v(len);
        for (int i = 0;i<len;i++)
        {
            map[nums[i]] = i;
        }
        int rank = 1;
        std::stringstream stream;
        for (auto it = map.begin();it != map.end();it++,rank++)
        {
            if (rank == 1)
                v[(*it).second] = "Gold Medal";
            else if (rank == 2)
                v[(*it).second] = "Silver Medal";
            else if (rank == 3)
                v[(*it).second] = "Bronze Medal";
            else
            {
                v[(*it).second] = std::to_string(rank);
            }
        }
        return v;
    }
};