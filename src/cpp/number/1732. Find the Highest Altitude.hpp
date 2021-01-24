#include <vector>
using namespace std;

class Solution_1732 {
public:
    int largestAltitude(vector<int>& gain) {
        int cur_height = 0;
        int len = gain.size();
        int max_height = 0;
        for(int i = 0;i < len;i++){
            cur_height += gain[i];
            if (cur_height > max_height){
                max_height = cur_height;
            }
        }
        return max_height;
    }
};