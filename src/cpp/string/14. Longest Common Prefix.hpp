#include <string>
#include <vector>

class Solution_14 {
public:
    string longestCommonPrefix(vector<string>& strs) {
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
};