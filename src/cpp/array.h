//
// Created by 陶诚程 on 2019/2/21.
//

#ifndef LEET_ARRAY_H
#define LEET_ARRAY_H

#endif //LEET_ARRAY_H
#include <vector>
#include <map>
#include <iostream>
#include <random>
#include <iostream>     // std::cout
#include <algorithm>    // std::next_permutation, std::sort, std::reverse
#include <ctime>

int remove_duplicate_number(std::vector<int>& nums);

int allow_duplicate_number(std::vector<int>& nums);

int fibex(int n);
int fibex2(int n);

int lastStoneWeight(std::vector<int>& stones);
class Solution_array {
    int len = 0;
public:
    int rob(std::vector<int>& nums) {
        len = nums.size();

        if(len == 0)
            return 0;
        if(len == 1)
            return nums[0];
        int max_choose_cur = recursive(nums,2) + nums[0];
        int max_skip_cur = recursive(nums,1);
        return max_choose_cur > max_skip_cur?max_choose_cur:max_skip_cur;
    }

    int recursive(std::vector<int>& nums,int begin){
        if(begin >= len)
            return 0;
        int max_choose_cur = recursive(nums,begin + 2) + nums[begin];
        int max_skip_cur = recursive(nums,begin + 1);
        auto max = max_choose_cur > max_skip_cur?max_choose_cur:max_skip_cur;
        return max;
    }
};


//Input: votes = ["ABC","ACB","ABC","ACB","ACB"]
//Output: "ACB"
namespace s1366{
    class Solution {
    public:
        std::string rankTeams(std::vector<std::string>& votes);
    };
}

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