#include <string>
#include <queue>
using namespace std;

//Input: "I speak Goat Latin"
//Output: "Imaa peaksmaaa oatGmaaaa atinLmaaaaa"

class Solution_824 {
public:
    string toGoatLatin(string S) {
        int len = S.size();
        int begin = 0;
        int end = 0;
        std::queue<std::string> q;
        while(end < len){
            if(S[end] == ' '){
                q.push(S.substr(begin,end - begin));
                begin = end + 1;
            }
            end++;
        }
        q.push(S.substr(begin,end - begin));
        std::string res;
        int a_cnt = 1;
        while(!q.empty()){
            std::string sub = q.front();
            q.pop();
            if(sub[0] != 'a' && sub[0] != 'e' && sub[0] != 'i' && sub[0] != 'o' && sub[0] != 'u' &&
                    sub[0] != 'A' && sub[0] != 'E' && sub[0] != 'I' && sub[0] != 'O' && sub[0] != 'U'){
                sub = sub.substr(1,sub.size() - 1) + sub[0];
            }
            sub += "ma";
            if(!res.empty()){
                res += " ";
            }
            std::string postfix = string(a_cnt,'a');
            res = res + sub + postfix;
            a_cnt++;
        }
        return res;
    }
};