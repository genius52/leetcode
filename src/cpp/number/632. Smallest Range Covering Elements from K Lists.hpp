#include <vector>
#include <queue>
using namespace std;

//You have k lists of sorted integers in non-decreasing order.
// Find the smallest range that includes at least one number from each of the k lists.
//
//We define the range [a, b] is smaller than range [c, d] if b - a < d - c or a < c if b - a == d - c.

//Input: nums = [[4,10,15,24,26],[0,9,12,20],[5,18,22,30]]
//Output: [20,24]
//Explanation:
//List 1: [4, 10, 15, 24,26], 24 is in range [20,24].
//List 2: [0, 9, 12, 20], 20 is in range [20,24].
//List 3: [5, 18, 22, 30], 22 is in range [20,24].
class Solution_632 {
public:
    vector<int> smallestRange(vector<vector<int>>& nums){
        //value:list index:element index
        std::priority_queue<std::pair<int,std::pair<int,int>>,std::deque<std::pair<int,std::pair<int,int>>>, std::greater<std::pair<int,std::pair<int,int>>>> q;
        int len = nums.size();
        int max_num = -2147483648;
        //int min_num = 2147483647;
        int total_length = 0;
        for(int i = 0;i < len;i++){
            total_length += nums[i].size();
        }
        for(int i = 0;i < len;i++){
            if(nums[i][0] > max_num)
                max_num = nums[i][0];
            std::pair<int,std::pair<int,int>> p;
            p.first = nums[i][0];
            p.second.first = i;
            p.second.second = 0;
            q.push(p);
        }
        std::vector<int> res(2);
        int min_diff = 2147483647;//max_num - q.top();
        while (total_length > 0){
            auto p = q.top();//最小的数
            q.pop();
            int cur_diff = max_num - p.first;
            if(cur_diff < min_diff){
                min_diff = cur_diff;
                res[0] = p.first;
                res[1] = max_num;
            }

            int list_num = p.second.first;
            int list_index = p.second.second;
            if(list_index == nums[list_num].size() - 1)
                break;
            std::pair<int,std::pair<int,int>> p2;
            p2.first = nums[list_num][list_index + 1];
            max_num = max(max_num,p2.first);
            p2.second.first = list_num;
            p2.second.second = list_index + 1;
            q.push(p2);
            total_length--;
        }
        return res;
    }

//    vector<int> smallestRange(vector<vector<int>>& nums) {
//        int len = nums.size();
//        int total_length = 0;
//        for(int i = 0;i < len;i++){
//            total_length += nums[i].size();
//        }
//        std::vector<int> res(2);
//        std::vector<std::pair<int,std::pair<int,int>>> data(len);//value : list index: element index
//        for(int i = 0;i < len;i++){
//            std::pair<int,std::pair<int,int>> p;
//            p.first = nums[i][0];
//            p.second.first = i;
//            p.second.second = 0;
//            data[i] = p;
//        }
//        std::sort(data.begin(),data.end(),[](const std::pair<int,std::pair<int,int>> p1,
//                const std::pair<int,std::pair<int,int>>& p2){
//            return p1.first < p2.first;
//        });
//        res[0] = data[0].first;
//        res[1] = data[len - 1].first;
//        int min_diff = res[1] - res[0];
//        while (total_length > 0){
//            auto it = data.begin();
//            int list_num = it->second.first;
//            int list_index = it->second.second;
//            if(list_index == nums[list_num].size() - 1)
//                break;
//            data[0].first = nums[list_num][list_index + 1];
//            data[0].second.first = list_num;
//            data[0].second.second = list_index + 1;
//            std::sort(data.begin(),data.end(),[](const std::pair<int,std::pair<int,int>> p1,
//                                                 const std::pair<int,std::pair<int,int>>& p2){
//                return p1.first < p2.first;
//            });
//            int cur_diff = data[len - 1].first - data[0].first;
//            if(cur_diff < min_diff){
//                min_diff = cur_diff;
//                res[0] = data[0].first;
//                res[1] = data[len - 1].first;
//            }
//            total_length--;
//        }
//        return res;
//    }
};