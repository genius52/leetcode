#include <vector>
#include <string>
#include <unordered_set>
using namespace std;

class Solution_500 {
public:
    vector<string> findWords(vector<string>& words) {
        std::unordered_set<char> first{'q','w','e','r','t','y','u','i','o','p'};
        std::unordered_set<char> second{'a','s','d','f','g','h','j','k','l'};
        std::unordered_set<char> third{'z','x','c','v','b','n','m'};
        std::vector<std::string> res;
        for(auto w : words){
            int first_cnt = 0;
            int second_cnt = 0;
            int third_cnt = 0;
            bool match = true;
            for(auto c : w){
                if(c < 'a'){
                    c += 32;
                }
                if(first.find(c) != first.end()){
                    first_cnt++;
                }else if(second.find(c) != second.end()){
                    second_cnt++;
                }else if(third.find(c) != third.end()){
                    third_cnt++;
                }
                if((first_cnt > 0 && second_cnt > 0) || (first_cnt > 0 && third_cnt > 0) ||
                    (second_cnt > 0 && third_cnt > 0)){
                    match = false;
                    break;
                }
            }
            if(match){
                res.push_back(w);
            }
        }
        return res;
    }
};