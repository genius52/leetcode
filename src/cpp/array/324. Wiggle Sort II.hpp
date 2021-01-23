#include <vector>
using namespace std;

class Solution_324 {
public:
    void swap(int& a,int& b){
        auto tmp = a;
        a = b;
        b = tmp;
    }

    void wiggleSort(std::vector<int>& nums) {
        auto l = nums.size();
        std::sort(nums.begin(),nums.end());
        int mid = nums[l/2];
        std::vector<int> copy(l,mid);
        int odd = 1;//从左往右放，从大放到小，直到中位数
        int i = l - 1;
        while (odd < l && nums[i] > mid){
            copy[odd] = nums[i];
            odd += 2;
            i--;
        }
        int even = l - 1;//从右往左放，从小放到大，直到中位数
        if (l % 2 == 0){
            even--;
        }
        int j = 0;
        while (even >= 0 && nums[j] < mid){
            copy[even] = nums[j];
            even -= 2;
            j++;
        }
        nums = copy;
    }
};