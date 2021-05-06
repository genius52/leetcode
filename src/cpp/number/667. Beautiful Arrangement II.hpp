#include <algorithm>
#include <vector>
#include <unordered_set>
#include <math.h>
using namespace std;

class Solution_667 {
public:
    vector<int> constructArray(int n, int k){
        std::vector<int> res(n);
        for(int left = 1,right = n,index = 0;index < n;){
            if(k > 1){
                if((k | 1) == k){
                    res[index] = left;
                    res[index + 1] = right;
                }else{
                    res[index] = right;
                    res[index + 1] = left;
                }
                left++;
                right--;
                index += 2;
                k -= 2;
            }else{
                res[index] = left;
                index++;
                left++;
            }
        }
        return res;
    }
    //TLE
//    vector<int> constructArray(int n, int k) {
//        std::vector<int> res(n);
//        for(int i = 0;i < n;i++){
//            res[i] = i + 1;
//        }
//        if(k == 1)
//            return res;
//        while(true){
//            std::unordered_set<int> s;
//            std::next_permutation(res.begin(),res.end());
//            for(int i = 1;i < n;i++){
//                s.insert(abs(res[i] - res[i - 1]));
//            }
//            if(s.size() == k)
//                break;
//        }
//        return res;
//    }
};