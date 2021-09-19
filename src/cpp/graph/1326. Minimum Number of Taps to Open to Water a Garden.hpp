#include <vector>
#include <map>
#include <set>
using namespace std;

//Input: n = 5, ranges = [3,4,1,1,0,0]
//Output: 1
//Explanation: The tap at point 0 can cover the interval [-3,3]
//The tap at point 1 can cover the interval [-3,5]
//The tap at point 2 can cover the interval [1,3]
//The tap at point 3 can cover the interval [2,4]
//The tap at point 4 can cover the interval [4,4]
//The tap at point 5 can cover the interval [5,5]
//Opening Only the second tap will water the whole garden [0,5]
class Solution_1326 {
public:
    int minTaps(int n, vector<int>& ranges) {
        std::map<int,std::set<int>,std::greater<int>> start_end;
        int l = ranges.size();
        for(int i = 0;i < l;i++){
            if(ranges[i] == 0)
                continue;
            int start = i - ranges[i];
            if(start < 0)
                start = 0;
            int end = i + ranges[i];
            start_end[start].insert(end);
        }
        int res = 0;
        int pre_start = 0;
        int pre_end = 0;
        while(true){
            auto find = start_end.lower_bound(pre_end);//小于等于pre_end
            if(find == start_end.end())
                return -1;
            int furthest = 0;
            int furthest_start;
            for(auto it = find;it != start_end.end();it++){
                if((*it).first < pre_start)
                    break;
                if(*it->second.rbegin() > furthest){
                    furthest_start = (*it).first;
                    furthest = *it->second.rbegin();
                }
            }
            res++;
            if (pre_end == furthest){
                return -1;
            }
            pre_start = furthest_start;
            pre_end = furthest;
            if(pre_end >= n)
                return res;
        }
        return -1;
    }
};