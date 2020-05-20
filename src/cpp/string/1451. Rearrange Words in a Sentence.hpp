
//Input: text = "Keep calm and code on"
//Output: "On and keep calm code"
//Explanation: Output is ordered as follows:
//"On" 2 letters.
//"and" 3 letters.
//"keep" 4 letters in case of tie order by position in original text.
//"calm" 4 letters.
//"code" 4 letters.
#include <string>
#include <map>
using namespace std;

class Solution_1451 {
public:
    string arrangeWords(string text) {
        std::map<int,std::vector<std::string>> record;
        int len = text.length();
        int begin = 0;
        int end = 1;
        while(end < len){
            if(text[end] == ' '){
                int l = end - begin;
                record[l].push_back(text.substr(begin,l));
                begin = end + 1;
                end = begin + 1;
            }else{
                end++;
            }
        }
        int l = end - begin;
        record[l].push_back(text.substr(begin,l));
        std::string res;
        for(auto it = record.begin();it != record.end();it++){
            std::vector<std::string> word_list = (*it).second;
            for(auto it2 = word_list.begin();it2 != word_list.end();it2++){
                std::string word = *it2;
                if(!res.empty())
                    res += " ";
                for(int i = 0;i < word.length();i++){
                    if(res.empty()){
                        if('a' <= word[i] && word[i] <= 'z'){
                            res += word[i] - 32;
                        }else{
                            res += word[i];
                        }
                    }else{
                        if('a' <= word[i] && word[i] <= 'z'){
                            res += word[i];
                        }else{
                            res += word[i] + 32;
                        }
                    }
                }
            }
        }
        return res;
    }
};