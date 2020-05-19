//
// Created by 陶诚程 on 2020-05-17.
//
#include <string>
using namespace std;

//Input: s = "abbcccddddeeeeedcba"
//Output: 5
//Explanation: The substring "eeeee" is of length 5 with the character 'e' only.
class Solution_1446 {
public:
    int maxPower(string s) {
        int len = s.size();
        if(len <= 1)
            return len;
        int slow = 0;
        int fast = 1;
        int max_len = 1;
        while(fast < len){
            while (s[slow] == s[fast]){
                fast++;
            }
            if((fast - slow) > max_len){
                max_len = fast - slow;
            }
            slow = fast;
            fast++;
        }
        return max_len;
    }
};