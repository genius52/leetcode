#include <vector>
#include <string>
#include <map>
using namespace std;

class Solution_2213 {
public:
    vector<int> longestRepeating(string s, string queryCharacters, vector<int>& queryIndices) {
        std::map<int,int> char_position[26];
        int max_len[26] = {0};
        int str_len = s.size();
        int left = 0;
        int right = 0;
        while (left < str_len){
            while (right < str_len && s[right] == s[left]){
                right++;
            }
            char_position[s[left] - 'a'][left] = right - left;
            if(right - left > max_len[s[left] - 'a']){
                max_len[s[left] - 'a'] = right - left;
            }
            left = right;
        }
        int l = queryIndices.size();
        std::vector<int> res(l);
        for(int i = 0;i < l;i++){
            // queryIndices[i]
            //查看char_position[queryCharacters[i] - 'a']前面和后面的位置，判断最大长度是否增长
        }
        return res;
    }
};