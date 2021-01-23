#include <vector>
using namespace std;

class Solution_46 {
    vector<vector<int>> v;
public:
    vector<vector<int>> permute(vector<int>& nums) {
        int len = nums.size();
        if(len == 0)
            return v;
        if(len == 1){
            v.push_back(nums);
            return v;
        }
        perm(nums,0,len - 1);
        return v;
    }

    void perm(vector<int>& nums,int begin,int end){
        if(begin == end){
            v.push_back(nums);
        }
        else{
            for(int i = begin;i <= end;i++){
                int tmp = nums[begin];
                nums[begin] = nums[i];
                nums[i] = tmp;
                perm(nums,begin + 1,end);
                tmp = nums[begin];
                nums[begin] = nums[i];
                nums[i] = tmp;
            }

        }
    }
};