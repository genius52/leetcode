#include <vector>
#include <unordered_set>
using namespace std;

class Solution_1207 {
public:
    bool uniqueOccurrences(vector<int>& arr) {
        int count[2001]{};
        for(auto it = arr.begin();it != arr.end(); it++){
            count[*it + 1000] += 1;
        }
        std::unordered_set<int> s;
        for(int i = 0;i < 2001;i++){
            if(count[i] == 0)
                continue;
            auto it = s.find(count[i]);
            if(it != s.end())
                return false;
            s.insert(count[i]);
        }
        return true;
    }
};