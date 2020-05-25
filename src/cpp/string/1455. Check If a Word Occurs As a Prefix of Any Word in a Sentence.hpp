#include <string>
using namespace std;

class Solution_1455 {
public:
    int isPrefixOfWord(string sentence, string searchWord) {
        int len = sentence.length();
        int begin = 0;
        int end = 0;
        int cnt = 0;
        while(end < len){
            while(end < len && sentence[end] != ' '){
                end++;
            }
            cnt++;
            std::string sub = sentence.substr(begin,end - begin);
            if(sub.find(searchWord) == 0){
                return cnt;
            }
            begin = end + 1;
            end = begin;
        }
        return -1;
    }
};