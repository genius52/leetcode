#include <string>
#include <algorithm>
//Input: S = "5F3Z-2e-9-w", K = 4
//
//Output: "5F3Z-2E9W"
//
//Explanation: The string S has been split into two parts, each part has 4 characters.
//Note that the two extra dashes are not needed and can be removed.
class Solution_482 {
public:
    string licenseKeyFormatting(string S, int K) {
        std::string res;
        int len = S.length();
        int end = len - 1;
        while(end >= 0){
            int move = end;
            int j = 0;
            while(move >= 0 && S[move] == '-'){
                move--;
            }
            if(move >= 0 && !res.empty()){
                res += '-';
            }
            while(move >= 0 && j < K){
                if(S[move] == '-'){
                    move--;
                    continue;
                }
                if(S[move] <= 'z' && S[move] >= 'a')
                    res += S[move] - 32;
                else
                    res += S[move];
                move--;
                j++;
            }
            end = move;
        }
        std::reverse(res.rbegin(), res.rend());
        return res;
    }
};