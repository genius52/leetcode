//
// Created by 陶诚程 on 2019/4/29.
//

#ifndef LEET_OTHER_H
#define LEET_OTHER_H
#include <string>
#include <sstream>
#include <vector>
#include <queue>
#include <algorithm>
#include <functional>
#include <condition_variable>
#include <mutex>
using namespace std;

int reverse_number(int x);
bool isPalindrome(int x);

std::string longestCommonPrefix(std::vector<std::string>& strs);

uint32_t reverseBits(uint32_t n);

int countPrimes(int n);

std::vector<std::vector<int>> permute(std::vector<int>& nums);
std::vector<std::vector<int>> subsets(std::vector<int>& nums);

bool hasAlternatingBits(int n);

void quick_sort(std::vector<int>& A,int begin,int end);

//295
/**
 * Your MedianFinder object will be instantiated and called as such:
 * MedianFinder* obj = new MedianFinder();
 * obj->addNum(num);
 * double param_2 = obj->findMedian();
 */
class MedianFinder {
public:
    /** initialize your data structure here. */
    MedianFinder() {
    }

    void addNum(int num) {
        auto small_size = smalltop_queue.size();
        auto big_size = bigtop_queue.size();
        if (small_size == 0){
            smalltop_queue.push(num);
            return;
        }
        if (small_size > big_size){
            if (num > smalltop_queue.top()){
                bigtop_queue.push(smalltop_queue.top());
                smalltop_queue.pop();
                smalltop_queue.push(num);
            }else{
                bigtop_queue.push(num);
            }
        }else if (small_size < big_size){
            if (num < bigtop_queue.top()){
                smalltop_queue.push(bigtop_queue.top());
                bigtop_queue.pop();
                bigtop_queue.push(num);
            }else{
                smalltop_queue.push(num);
            }
        }else{
            if (num <= bigtop_queue.top()){
                bigtop_queue.push(num);
            }else{
                smalltop_queue.push(num);
            }
        }
    }

    double findMedian() {
        if (smalltop_queue.size() == bigtop_queue.size()){
            return double(smalltop_queue.top() + bigtop_queue.top()) / 2;
        }else if(smalltop_queue.size() > bigtop_queue.size()){
            return smalltop_queue.top();
        }else{
            return bigtop_queue.top();
        }
    }
    std::priority_queue<int,std::vector<int>,std::greater<int> > smalltop_queue;
    std::priority_queue<int,std::vector<int>,std::less<int> > bigtop_queue;
};


//Input: "2*3-4*5"
// Output: [-34, -14, -10, -10, 10]
// Explanation: 
// (2*(3-(4*5))) = -34 
// ((2*3)-(4*5)) = -14 
// ((2*(3-4))*5) = -10 
// (2*((3-4)*5)) = -10 
// (((2*3)-4)*5) = 10
//class Solution_241 {
//public:
//    vector<int> diffWaysToCompute(string input) {
//        std::vector<int> res;
//        int l = input.length();
//        std::deque<char> nums;
//        std::deque<char> operators;
//        int num_begin = 0;
//        for (int i = 0;i < l;){
//            if (input[i] == '+' || input[i] == '-' || input[i] == '*' ){
//                operators.push_back(input[i]);
//                i++;
//            }else{
//                int begin = i;
//                while(i < (l - 1) && input[i + 1] != '+' && input[i + 1] != '-' && input[i + 1] != '*'){
//                    i++;
//                }
//                std::stringstream ss;
//                ss<<input.substr(begin,i - begin + 1);
//                int n;
//                ss>>n;
//                nums.push_back(n);
//                i++;
//            }
//        }
//        return res;
//    }
//};
#endif //LEET_OTHER_H
