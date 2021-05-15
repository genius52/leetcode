#include <vector>
#include <unordered_set>
#include <unordered_map>
#include <math.h>
#include <random>
using namespace std;

//Input:
//["Solution","pick","pick","pick"]
//[[3,[1]],[],[],[]]
//Output: [null,0,0,2]
class Solution_710 {
    std::unordered_map<int,int> m_;
    int length = 0;
public:
    Solution_710(int n, vector<int>& blacklist){
        length = n - blacklist.size();
        for(auto num : blacklist){
            m_[num] = 0;
        }
        int right = n - 1;
        for(auto num : blacklist){
            if(num < length){
                while(m_.find(right) != m_.end()) {
                    right--;
                }
                m_[num] = right;
                right--;
            }
        }
    }

    int pick() {
        int num = rand() % length;
        if(m_.find(num) != m_.end())
            return m_[num];
        return num;
    }

    //TLE
//    std::vector<int> data;
//    int length = 0;
//public:
//    Solution(int n, vector<int>& blacklist) {
//        int len = blacklist.size();
//        length = n - len;
//        data.resize(length);
//        std::unordered_set<int> s;
//        for(auto n : blacklist){
//            s.insert(n);
//        }
//        int index = 0;
//        for(int i = 0;i < n;i++){
//            if(s.find(i) == s.end())
//                data[index++] = i;
//        }
//    }
//
//    int pick() {
//        int pos = rand() % length;
//        return data[pos];
//    }
};