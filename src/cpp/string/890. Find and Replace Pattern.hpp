#include <vector>
#include <map>
using namespace std;

class Solution {
public:
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
};