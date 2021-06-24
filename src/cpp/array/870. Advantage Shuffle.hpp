#include <vector>
#include <map>
#include <set>
#include <algorithm>
using namespace std;

//Input: nums1 = [12,24,8,32], nums2 = [13,25,32,11]
//Output: [24,32,8,12]
class Solution_870 {
public:
    vector<int> advantageCount(vector<int>& A, vector<int>& B){
        int len = A.size();
        if(len <= 1)
            return A;
        std::vector<int> res(len);
        std::multiset<int> s;
        for(auto a : A){
            s.insert(a);
        }
        for(int i = 0;i < len;i++){
            auto it = s.upper_bound(B[i]);
            if(it != s.end()){
                res[i] = *it;
                s.erase(it);
            }else{
                res[i] = *(s.begin());
                s.erase(s.begin());
            }
        }
        return res;
    }

//    vector<int> advantageCount(vector<int>& A, vector<int>& B) {
//        int len = A.size();
//        if (len <= 1)
//            return A;
//        std::map<int,int> map_a;
//        for(auto n : A){
//            map_a[n]++;
//        }
//        std::vector<int> res;
//        for(int i = 0;i < len;i++){
//            auto it = map_a.upper_bound(B[i]);
//            if (it == map_a.end()){
//                auto small = map_a.begin();
//                res.push_back(small->first);
//                small->second--;
//                if (small->second == 0){
//                    map_a.erase(small);
//                }
//            }else{
//                res.push_back(it->first);
//                it->second--;
//                if(it->second == 0)
//                    map_a.erase(it);
//            }
//        }
//        return res;
//    }
};