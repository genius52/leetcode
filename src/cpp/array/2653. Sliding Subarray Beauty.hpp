#include <vector>
#include <queue>
#include <set>
using namespace std;

class Solution_2653 {
public:
    vector<int> getSubarrayBeauty(vector<int>& nums, int k, int x) {
        int len = nums.size();
        std::vector<int> res(len - k + 1);
        int idx = 0;
        int num_cnt[101] = {0};
        int neg_cnt = 0;
        for(int i = 0;i < len;i++){
            num_cnt[nums[i] + 50]++;
            if(nums[i] < 0){
                neg_cnt++;
            }
            if(i < k - 1){
                continue;
            }
            if(neg_cnt >= x){
                int diff = neg_cnt - x;
                for (int j = 49;j >= 0;j--){
                    if(num_cnt[j] == 0)
                        continue;
                    if(num_cnt[j] > diff){
                        res[idx] = j - 50;
                        break;
                    }
                    diff -= num_cnt[j];
                }
            }
            idx++;
            num_cnt[nums[i - k + 1] + 50]--;
            if(nums[i - k + 1] < 0)
                neg_cnt--;
        }
        return res;
    }
};

//class Solution_2653 {
//public:
//    vector<int> getSubarrayBeauty(vector<int>& nums, int k, int x) {
//        std::multiset<int> s;
//        int len = nums.size();
//        std::vector<int> res(len - k + 1);
//        int j = 0;
//        std::multiset<int>::iterator it;
//        for(int i = 0;i < len;i++){
//            s.insert(nums[i]);
//            if(i < k - 1){
//                continue;
//            }
//            if(i == k - 1){
//                it = s.end();
//            }
//            int cnt = std::distance(s.begin(),it);
//            int diff = x - cnt - 1;
//            while(diff > 0){
//                diff--;
//                it++;
//            }
//            while(diff < 0){
//                diff++;
//                it--;
//            }
//            if(*it <= 0)
//                res[j] = *it;
//            s.erase(s.find(nums[i - k + 1]));
//            j++;
//        }
//        return res;
//    }
//};