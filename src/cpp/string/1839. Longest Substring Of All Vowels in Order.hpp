#include <string>
#include <set>
#include <unordered_map>
using namespace std;

//Input: word = "aeiaaioaaaaeiiiiouuuooaauuaeiu"
//Output: 13
//Explanation: The longest beautiful substring in word is "aaaaeiiiiouuu" of length 13.
class Solution_1839 {
public:
    int longestBeautifulSubstring(string word) {
        //std::unordered_map<char,std::set<int>> char_pos;
        int len = word.size();
        if(len < 5)
            return 0;
        int max_len = 0;
        int scores = 0;
        int left = 0;
        for(int i = 0;i < len;i++){
            if(word[i] == 'a'){
                if(scores != 1){
                    scores = 1;
                    left = i;
                }
            }else if(word[i] == 'e'){
                if(scores == 1){
                    scores++;
                }else if(scores != 2){
                    scores = 0;
                }
            }else if(word[i] == 'i'){
                if(scores == 2){
                    scores++;
                }else if(scores != 3){
                    scores = 0;
                }
            }else if(word[i] == 'o'){
                if(scores == 3){
                    scores++;
                }else if(scores != 4){
                    scores = 0;
                }
            }else if(word[i] == 'u'){
                if(scores == 4){
                    scores++;
                }
                if(scores == 5){
                    int cur_len = i - left + 1;
                    if(cur_len > max_len)
                        max_len = cur_len;
                }else{
                    scores = 0;
                }
            }
        }
        return max_len;
    }
};