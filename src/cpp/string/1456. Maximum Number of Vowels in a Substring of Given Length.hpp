#include <string>
using namespace std;

//Input: s = "leetcode", k = 3
//Output: 2
//Explanation: "lee", "eet" and "ode" contain 2 vowels.
class Solution_1456 {
public:
    int maxVowels(string s, int k) {
        int len = s.length();
        int cnt = 0;
        for(int i = 0;i < k;i++){
            if(s[i] == 'a'||s[i] == 'e'||s[i] == 'i'||s[i] == 'o'||s[i] == 'u'){
                cnt++;
            }
        }
        int max = cnt;
        int begin = 1;
        int end = k;
        while(end < len){
            if(s[begin - 1] == 'a'||s[begin - 1] == 'e'||s[begin - 1] == 'i'||s[begin - 1] == 'o'||s[begin - 1] == 'u'){
                cnt--;
            }
            if(s[end] == 'a'||s[end] == 'e'||s[end] == 'i'||s[end] == 'o'||s[end] == 'u'){
                cnt++;
            }
            if (cnt > max){
                max = cnt;
            }
            begin++;
            end++;
        }
        return max;
    }
};