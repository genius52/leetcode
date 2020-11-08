#include <vector>
#include <map>
#include <algorithm>
using namespace std;

class Solution_870 {
public:
    vector<int> advantageCount(vector<int>& A, vector<int>& B) {
        int len = A.size();
        if (len <= 1)
            return A;
        std::map<int,int> map_a;
        for(auto n : A){
            map_a[n]++;
        }
        std::vector<int> res;
        for(int i = 0;i < len;i++){
            auto it = map_a.upper_bound(B[i]);
            if (it == map_a.end()){
                auto small = map_a.begin();
                res.push_back(small->first);
                small->second--;
                if (small->second == 0){
                    map_a.erase(small);
                }
            }else{
                res.push_back(it->first);
                it->second--;
                if(it->second == 0)
                    map_a.erase(it);
            }
        }
        return res;
    }
};