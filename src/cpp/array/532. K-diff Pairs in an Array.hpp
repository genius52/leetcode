#include <vector>
#include <unordered_map>
#include <algorithm>
using namespace std;

//Input: [3, 1, 4, 1, 5], k = 2
//Output: 2
//Explanation: There are two 2-diff pairs in the array, (1, 3) and (3, 5).
//Although we have two 1s in the input, we should only return the number of unique pairs.
class Solution_532 {
public:
//    int findPairs(vector<int>& nums, int k) {
//        std::sort(nums.begin(),nums.end());
//        int len = nums.size();
//        int res = 0;
//        int begin = 0;
//        while(begin < len - 1){
//            while(begin > 0 && nums[begin] == nums[begin - 1])
//                begin++;
//            for(int j = begin + 1;j < len;j++){
//                if(j > begin + 1 && nums[j] == nums[j - 1])
//                    continue;
//                if(nums[j] - nums[begin] == k)
//                    res++;
//                else if(nums[j] - nums[begin] > k)
//                    break;
//            }
//            begin++;
//        }
//        return res;
//    }

    int findPairs(vector<int>& nums, int k) {
        if(k < 0){
            return 0;
        }
        std::unordered_map<int,int> graph;
        int len = nums.size();
        if(len <= 1)
            return 0;
        int res = 0;
        for(int i = 0;i < len;i++){
            graph[nums[i]]++;
        }
        for(auto it = graph.begin();it != graph.end();it++){
            if(k == 0){
                if((*it).second > 1 ){
                    res++;
                }
            }else{
                if(graph.find((*it).first + k) != graph.end()){
                    res++;
                }
            }
        }
        return res;
    }
};