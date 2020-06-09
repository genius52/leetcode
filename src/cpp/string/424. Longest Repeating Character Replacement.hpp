#include <string>
using namespace std;

//Input:
//s = "AABABBA", k = 1
//
//Output:
//4
//
//Explanation:
//Replace the one 'A' in the middle with 'B' and form "AABBBBA".
//The substring "BBBB" has the longest repeating letters, which is 4.
class Solution_424 {
public:
    int characterReplacement(string s, int k) {
        int len = s.length();
        int count[26] = {0};//the count of char from begin to end
        int begin = 0;
        int end = 0;
        int res = 0;
        int max_count = 0;
        while(end < len){
            count[s[end] - 'A']++;
            max_count = max(max_count,count[s[end] - 'A']);
            if(end - begin + 1 - max_count <= k){
                res = max(res,end - begin + 1);
            }else{
                count[s[begin] - 'A']--;
                max_count = max(max_count,count[s[begin] - 'A']);
                begin++;
            }
            end++;
        }
        return res;
    }
};