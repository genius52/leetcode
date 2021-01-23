#include <string>
using namespace std;

class Solution_942 {
public:
    vector<int> diStringMatch(string S) {
        vector<int> res;
        int low = 0;
        int high = S.length();
        for(int i = 0;i <= S.length();i++){
            if(i == S.length()){
                res.push_back(low);
                break;
            }
            res.push_back(S[i] == 'I'? low++:high--);
        }
        return res;
    }
};