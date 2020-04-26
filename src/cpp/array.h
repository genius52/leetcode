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

//384
/**
 * Your Solution object will be instantiated and called as such:
 * Solution* obj = new Solution(nums);
 * vector<int> param_1 = obj->reset();
 * vector<int> param_2 = obj->shuffle();
 */
namespace s384{
    class Solution {
    public:
        Solution(std::vector<int>& nums) {
            default_nums = nums;
            current_nums = nums;
            random1.seed(time(NULL));
            random2.seed(time(NULL));
        }

        /** Resets the array to its original configuration and return it. */
        std::vector<int> reset() {
            current_nums = default_nums;
            return current_nums;
        }

        /** Returns a random shuffling of the array. */
        std::vector<int> shuffle() {
            std::uniform_int_distribution<int> dis1(0,current_nums.size());
            std::uniform_int_distribution<int> dis2(0,current_nums.size());
            int ran1 = dis1(random1);
            int ran2 = dis2(random2);
            if(current_nums.size() > 0){
                int i = ran1;
                int j = ran2;
                int temp = current_nums[i];
                current_nums[i] = current_nums[j];
                current_nums[j] = temp;
            }
            std::next_permutation(current_nums.begin(),current_nums.end());
            return current_nums;
        }
    private:
        std::vector<int> default_nums;
        std::vector<int> current_nums;
        std::default_random_engine random1;
        std::default_random_engine random2;
    };
//    class Solution {
//    public:
//        Solution(std::vector<int>& nums) {
//            default_nums = nums;
//            cur_nums = nums;
//            possibility = fac(default_nums.size());
//            random.seed(time(NULL));
//        }
//
//        /** Resets the array to its original configuration and return it. */
//        std::vector<int> reset() {
//            cur_nums = default_nums;
//            return cur_nums;
//        }
//
//        /** Returns a random shuffling of the array. */
//        std::vector<int> shuffle() {
//            std::uniform_int_distribution<int> dis(1,possibility);
//            int target = dis(random);
//            int step = 0;
//            perm(cur_nums,0,step,target);
//            return cur_nums;
//        }
//    private:
//        void swap(int& a,int& b){
//            int tmp = a;
//            a = b;
//            b = tmp;
//        }
//        bool perm(std::vector<int>& nums,int cur_pos,int& cur_steps,int target){
//            if (cur_pos >= nums.size()){
//                cur_steps++;
//                if (cur_steps == target){
//                    return true;
//                }
//                return false;
//            }
//            for (int i = cur_pos;i < nums.size();i++){
//                swap(nums[i],nums[cur_pos]);
//                if(perm(cur_nums,cur_pos+1,cur_steps,target)){
//                    return true;
//                }
//                swap(nums[i],nums[cur_pos]);
//            }
//            return false;
//        }
//
//        template<typename T>
//        inline T fac(T num){
//            T res = 0;
//            if(num == 1)
//                return 1;
//            if (num > 1){
//                res = num * fac(num - 1);
//            }
//            return res;
//        }
//        std::vector<int> default_nums;
//        std::vector<int> cur_nums;
//        long possibility = 0;
//        std::default_random_engine random;;
//    };
}

//Input: nums = [8,1,2,2,3]
//Output: [4,0,1,1,3]
namespace s1365{
    class Solution {
    public:
        std::vector<int> smallerNumbersThanCurrent(std::vector<int>& nums) {
            std::vector<int> record(101);
            int max_num = -1;
            for (auto i : nums){
                record[i]++;
                if (i > max_num){
                    max_num = i;
                }
            }
            int sum = 0;
            for (int i = 0;i <= max_num;i++){
                if (record[i] == 0)
                    continue;
                int smaller = sum;
                sum += record[i];
                record[i] = smaller;
            }
            std::vector<int> res(nums.size());
            for (int i = 0;i < nums.size();i++){
                res[i] = record[nums[i]];
            }
            return res;
        }
    };
}

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