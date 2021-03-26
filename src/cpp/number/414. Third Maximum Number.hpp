#include <vector>
using namespace std;

class Solution_414 {
public:
    int thirdMax(vector<int>& nums){
        int len = nums.size();
        int first = nums[0];
        int* second = nullptr;
        int* third = nullptr;
        for(int i = 1;i < len;i++){
            if(nums[i] > first){
                if(second == nullptr){
                    second = new int();
                }else{
                    if(third == nullptr) {
                        third = new int();
                    }
                    *third = *second;
                }
                *second = first;
                first = nums[i];
            }else if(nums[i] == first){
                continue;
            }else{
                if(second == nullptr){
                    second = new int();
                    *second = nums[i];
                }else{
                    if(nums[i] > *second){
                        if(third == nullptr) {
                            third = new int();
                        }
                        *third = *second;
                        *second = nums[i];
                    }else if(nums[i] == *second) {
                        continue;
                    }else{
                        if(third == nullptr){
                            third = new int();
                            *third = nums[i];
                        }else{
                            if(nums[i] > *third){
                                *third =nums[i];
                            }
                        }
                    }
                }
            }
        }
        if(third == nullptr){
            return first;
        }else{
            return *third;
        }
    }

//    int thirdMax(vector<int>& nums) {
//        int len = nums.size();
//        if(len == 1){
//            return nums[0];
//        }
//        else if(len == 2){
//            return nums[0] > nums[1]?nums[0]:nums[1];
//        }
//        std::set<int> s;
//        for(auto e:nums){
//            s.insert(e);
//            if(s.size()>3)
//                s.erase(s.begin());
//        }
//        if(s.size()<3)
//            return *s.rbegin();
//        return *s.begin();
//    }
};