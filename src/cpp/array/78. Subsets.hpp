#include <vector>
using namespace std;

class Solution_78 {
    vector<vector<int>> v;
public:
    vector<vector<int>> subsets(vector<int>& nums) {
        int len = nums.size();
        if(len == 0)
            return v;
        if(len == 1){
            v.push_back(nums);
            v.push_back(vector<int>());
            return v;
        }
        for(int i = 0;i < len;i++){
            vector<int> select;
            combine(nums,0,select,i);
        }
        v.push_back(nums);
        return v;
    }

    void combine(vector<int>& nums,int step,vector<int>& select_data,int target_num){
        if(select_data.size() == target_num){
            v.push_back(select_data);
            return;
        }
        if(step >= nums.size()){
            return;
        }
        select_data.push_back(nums[step]);
        combine(nums,step + 1,select_data,target_num);
        select_data.pop_back();
        combine(nums,step + 1,select_data,target_num);
    }
};