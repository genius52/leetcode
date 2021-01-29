#include <string>
using namespace std;

//Input: croakOfFrogs = "crcoakroak"
//Output: 2
//Explanation: The minimum number of frogs is two.
//The first frog could yell "crcoakroak".
//The second frog could yell later "crcoakroak".
//Input: croakOfFrogs = "croakcroa"
//Output: -1
class Solution_1419 {
public:
    int minNumberOfFrogs(string croakOfFrogs) {
        int cnt_c = 0;
        int cnt_r = 0;
        int cnt_o = 0;
        int cnt_a = 0;
        int cnt_k = 0;
        int len = croakOfFrogs.size();
        int res = 1;
        for(int i = 0;i < len;i++){
            if(croakOfFrogs[i] == 'c'){
                cnt_c++;
            }else if(croakOfFrogs[i] == 'r'){
                cnt_r++;
            }else if(croakOfFrogs[i] == 'o'){
                cnt_o++;
            }else if(croakOfFrogs[i] == 'a'){
                cnt_a++;
            }else if(croakOfFrogs[i] == 'k'){
                cnt_k++;
                if(cnt_c - cnt_k >= res - 1){
                    res = cnt_c - cnt_k + 1;
                }
                if(cnt_r - cnt_k >= res - 1){
                    res = cnt_r - cnt_k + 1;
                }
                if(cnt_o - cnt_k >= res - 1){
                    res = cnt_o - cnt_k + 1;
                }
                if(cnt_a - cnt_k >= res - 1){
                    res = cnt_a - cnt_k + 1;
                }
            }
            if(cnt_k > cnt_a || cnt_a > cnt_o || cnt_o > cnt_r || cnt_r > cnt_c){
                return -1;
            }
        }
        if(cnt_c != cnt_r || cnt_c != cnt_o || cnt_c != cnt_a || cnt_c != cnt_k){
            return -1;
        }
        return res;
    }
};