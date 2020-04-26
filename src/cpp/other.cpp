//
// Created by 陶诚程 on 2019/4/29.
//

#include <deque>
#include <cmath>
#include <iostream>
#include "other.h"

int reverse_number(int x) {
    bool negative = x < 0;
    std::deque<int> s;
    while(x != 0){
        auto result = x % 10;
        s.push_back(result);
        x = x/10;
    }
    long val = 0;
    int cnt = 0;
    while(!s.empty()){
        long temp = s.back();
        s.pop_back();
        val += temp * pow(10,cnt++);
    }
    std::cout<< val <<" "<<2e31-1<<std::endl;
    int int_max = pow(2,31)-1;
    int int_min = pow(-2,31);
    if(val > int_max)
        return 0;
    if(val < int_min)
        return 0;
    return val;
}

bool isPalindrome(int x) {
    if(x < 0)
        return false;
    if(x == 0)
        return true;
    int new_val = 0;
    int n = 0;
    int origin_val = x;
    while(origin_val != 0)
    {
        int result = origin_val % 10;
        new_val = new_val * 10 +result;

        origin_val = origin_val / 10;
    }
    return new_val == x;
}

std::string longestCommonPrefix(std::vector<std::string>& strs) {
    std::string prefix;
    size_t element_count = strs.size();
    if (element_count == 0)
        return prefix;
    auto len = strs[0].length();

    for(int i = 0;i < len;i++){
        bool same = true;
        auto c = strs[0][i];
        for(int j = 1;j < element_count;j++){
            if(strs[j][i] != c){
                same = false;
                break;
            }
        }
        if(same)
            prefix += c;
        else
            break;
    }
    return prefix;
}

uint32_t reverseBits(uint32_t n) {
    uint32_t result = 0;
    uint32_t mask = 1;
    for(int i = 0;i < 32;i++,n = n >> 1){
        result = result << 1;
        result |= mask & n;
    }
    return result;
}

int countPrimes(int n) {
    if(n <= 1)
        return 0;
    if(n == 2)
        return 1;
    int cnt = 0;
    bool find = false;
    for(int i = 2;i <= n;i++){
        //int val = int(sqrt(i))+1;
        for(int j = 2; j < i;j++){
            if(i%j == 0){
                find = true;
                break;
            }
        }
        if(!find)
        {
            cnt++;
        }
        find = false;
    }
    return cnt+1;
}

std::vector<std::vector<int>> v;
void perm(std::vector<int>& nums,int begin,int end){
    if(begin == end){
        v.push_back(nums);
    }
    else{
        for(int i = (begin + 1);i <= end;i++){
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

std::vector<std::vector<int>> permute(std::vector<int>& nums) {
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


void combine(std::vector<int>& nums,int step,std::vector<int>& select_data,int target_num){
    if(select_data.size() == target_num){
        v.push_back(select_data);
        return;
    }
    if(step >= nums.size()){
        return;
    }
    select_data.push_back(nums[step]);
    //std::cout<<nums[step];
    combine(nums,step + 1,select_data,target_num);
    select_data.pop_back();
    combine(nums,step + 1,select_data,target_num);
}

std::vector<std::vector<int>> subsets(std::vector<int>& nums) {
    int len = nums.size();
    if(len == 0)
        return v;
    if(len == 1){
        v.push_back(nums);
        return v;
    }
    for(int i = 1;i < len;i++){
        std::vector<int> select;
        combine(nums,0,select,i);
    }
    return v;
}

bool hasAlternatingBits(int n) {
    if(n == 1 || n == 0)
        return true;
    long tag = 1;
    if((n | 1) == n)
        tag = 1;
    else
        tag = 1 << 1;
    int sum = tag;
    while(sum < n){
        tag = tag<<2;
        sum = sum + tag;

    }
    return sum == n;
}

void quick_sort(std::vector<int>& A,int begin,int end){
    if(begin > end)
        return;
    int i = begin;
    int j = end;
    int temp = A[i];
    while(i < j){
        while(i < j && A[j] < temp)
            j--;
        if(i < j){
            A[i] = A[j];
            i++;
        }
        while(i < j && A[i] > temp)
            i++;
        if(i < j){
            A[j] = A[i];
            j--;
        }
    }
    A[i] = temp;
    quick_sort(A,begin,i-1);
    quick_sort(A,i+1,end);
}

bool isLongPressedName(std::string name, std::string typed) {
    int len = name.length();
    int len_t = typed.length();
    if(len > len_t)
        return false;
    int i,j;
    while(i < len && j < len_t){
        char c = name[i];
        if(name[i] != typed[j])
            return false;
        i++;
        j++;
        while(c == typed[j]){
            j++;
        }
    }
    return true;
}

//Input: words = ["abc","deq","mee","aqq","dkd","ccc"], pattern = "abb"
//Output: ["mee","aqq"]
#include <map>
std::vector<std::string> findAndReplacePattern(std::vector<std::string>& words, std::string pattern) {
    std::vector<std::string> res;
    int len = words.size();
    if(len == 0)
        return res;

    for(int i = 0;i < len;i++){
        bool bmatch = true;
        std::map<char,char> m1,m2;
        for(int j = 0;j < words[i].size();j++){
            if(m1.find(words[i][j]) == m1.end()){
                m1[words[i][j]] = pattern[j];
                if(m2.find(pattern[j]) != m2.end()){
                    bmatch = false;
                    break;
                }
                m2[pattern[j]] = words[i][j];
            }
            else{
                if(m1[words[i][j]] != pattern[j] || m2[pattern[j]] != words[i][j]){
                    bmatch = false;
                    break;
                }
            }
        }
         if(bmatch){
             res.push_back(words[i]);
         }
    }
    return res;
}


bool judge(int i){
    int origin = i;
    while(i > 0){
        int mod = i % 10;
        if(mod == 0)
            return false;
        if(origin % mod != 0)
            return false;
        i = i/10;
    }
    return true;
}

std::vector<int> selfDividingNumbers(int left, int right) {
    std::vector<int> v;
    if(left > right)
        return v;
    for(int i = left;i <= right;i++){
        if(judge(i))
            v.push_back(i);
    }
    return v;
}

//1195
//If the number is divisible by 3, output "fizz".
// If the number is divisible by 5, output "buzz".
// If the number is divisible by both 3 and 5, output "fizzbuzz".
// For example, for n = 15, we output: 1, 2, fizz, 4, buzz, fizz, 7, 8, fizz, buzz, 11, fizz, 13, 14, fizzbuzz.
//Implement a multithreaded version of FizzBuzz with four threads. The same instance of FizzBuzz will be passed to four different threads:

// Thread A will call fizz() to check for divisibility of 3 and outputs fizz.
// Thread B will call buzz() to check for divisibility of 5 and outputs buzz.
// Thread C will call fizzbuzz() to check for divisibility of 3 and 5 and outputs fizzbuzz.
// Thread D will call number() which should only output the numbers.
class FizzBuzz {
private:
    int n;
    int cur_num = 1;
    std::mutex mtx;
    //std::mutex mtx_change_num;
    std::condition_variable cv;
public:
    FizzBuzz(int n) {
        this->n = n;
        cv.notify_all();
    }

    // printFizz() outputs "fizz".
    void fizz(function<void()> printFizz) {
        while (cur_num <= n){
            std::unique_lock<mutex> lock(mtx);
            cv.wait(lock, [this](){
                return this->cur_num > n ||(this->cur_num % 3 == 0 && this->cur_num % 5 != 0);
            });
            if(this->cur_num > n)
                break; 
            printFizz();
            this->cur_num++;
            lock.unlock();
            cv.notify_all();
        }
    }

    // printBuzz() outputs "buzz".
    void buzz(function<void()> printBuzz) {
        while(cur_num <= n){
            std::unique_lock<mutex> lock(mtx);
            cv.wait(lock, [this](){
                return this->cur_num > n || (this->cur_num % 3 != 0 && this->cur_num % 5 == 0);
            });
            if(this->cur_num > n)
                break; 
            printBuzz();
            this->cur_num++;
            //lock.unlock();
            cv.notify_all();
        }
    }

    // printFizzBuzz() outputs "fizzbuzz".
	void fizzbuzz(function<void()> printFizzBuzz) {
        while(cur_num <= n){
            std::unique_lock<mutex> lock(mtx);
            cv.wait(lock,[this](){
                return this->cur_num > n || (this->cur_num % 3 == 0 && this->cur_num % 5 == 0);
            });
            if(this->cur_num > n)
                break;
            printFizzBuzz();
            this->cur_num++;
            //lock.unlock();
            cv.notify_all();
        }
    }

    // printNumber(x) outputs "x", where x is an integer.
    void number(function<void(int)> printNumber) {
        while(cur_num <= n){
            std::unique_lock<mutex> lock(mtx);
            cv.wait(lock,[this](){
                return this->cur_num > n ||(this->cur_num % 3 != 0 && this->cur_num % 5 != 0);
            });
            if(this->cur_num > n)
                break;
            printNumber(cur_num);
            this->cur_num++;
            //lock.unlock();
            cv.notify_all();
        }
    }
};
