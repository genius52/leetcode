#include <string>
#include <unordered_map>
using namespace std;

//Input: s = "cbacdcbc"
//Output: "acdb"
class Solution_1081 {
public:
    string smallestSubsequence(string s) {
        std::unordered_map<char,int> record;
        int len = s.length();
        for(int i = 0;i < len;i++){
            if(record.find(s[i]) == record.end()){
                record[s[i]] = 1;
            }else{
                record[s[i]]++;
            }

        }
    }
};