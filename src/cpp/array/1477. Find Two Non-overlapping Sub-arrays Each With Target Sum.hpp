#include <vector>
#include <unordered_map>
#include <set>
using namespace std;

//Input: arr = [3,1,1,1,5,1,2,1], target = 3
//Output: 3
class Solution_1477 {
public:
    bool is_overlap(int begin1,int end1,int begin2,int end2){
        if((begin1 <= begin2 && begin2 <= end1) || (begin1 <= end2 && end2 <= end1)
        || (begin2 <= begin1 && begin1 <= end2) || (begin2 <= end1 && end1 <= end2))
            return true;
        return false;
    }

    int minSumOfLengths(vector<int>& arr, int target) {
        int len = arr.size();
        if(len < 1)
            return -1;
        std::vector<int> leftdp(len);
        std::vector<int> rightdp(len);
        std::unordered_map<int,int> record;//key: sum from 0 to i, value: start to end
        leftdp[0] = arr[0];
        for(int i = 1;i <= len;i++){
            leftdp[i] = leftdp[i - 1] + arr[i];
        }
        for(int i = len - 2;i >= 0;i--){
            rightdp[i] = rightdp[i + 1] + arr[i];
        }
        for(int i = 0;i < len;i++){

        }
        return -1;
    }
};