#include <string>
#include <math.h>
using namespace std;

//Input: s = "pwwkew"
//Output: 3
//Explanation: The answer is "wke", with the length of 3.
//Notice that the answer must be a substring, "pwke" is a subsequence and not a substring.
class Solution_3 {
public:
    int lengthOfLongestSubstring(string s) {
        int len = s.length();
        int begin = 0;
        int end = 0;
        int record[255] = {};
        int max_len = 0;
        while(end < len){
            int pos = s[end];
            record[pos]++;
            if(record[pos] > 1){
                while(record[pos] > 1){
                    record[s[begin]]--;
                    begin++;
                }
                end++;
            }else{
                max_len = max(max_len,end - begin + 1);
                end++;
            }
        }
        return max_len;
    }
};