#include <vector>
#include <unordered_map>
#include <set>
using namespace std;

class Solution_575 {
public:
    int distributeCandies(vector<int>& candyType) {
        int target = candyType.size() / 2;
        std::unordered_map<int,int> type_count;
        for(auto c : candyType){
            type_count[c]++;
        }
        int types = type_count.size();
        if(types >= target){
            return target;
        }else{
            return types;
        }
    }
};