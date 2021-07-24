#include <vector>
#include <map>
using namespace std;

//输入：[10,13,12,14,15]
//输出：2
//解释：
//从起始索引 i = 0 出发，我们可以跳到 i = 2，（因为 A[2] 是 A[1]，A[2]，A[3]，A[4] 中大于或等于 A[0] 的最小值），然后我们就无法继续跳下去了。
//从起始索引 i = 1 和 i = 2 出发，我们可以跳到 i = 3，然后我们就无法继续跳下去了。
//从起始索引 i = 3 出发，我们可以跳到 i = 4，到达数组末尾。
//从起始索引 i = 4 出发，我们已经到达数组末尾。
//总之，我们可以从 2 个不同的起始索引（i = 3, i = 4）出发，通过一定数量的跳跃到达数组末尾。
class Solution_975 {
public:
    int oddEvenJumps(vector<int>& arr) {
        int len = arr.size();
        std::map<int,int> record;
        std::vector<bool> increase(len);
        std::vector<bool> decrease(len);
        increase[len - 1] = true;
        decrease[len - 1] = true;
        record[arr[len - 1]] = len - 1;
        int res = 1;
        for (int i = len - 2;i >= 0;i--){
            //先找大于等于当前的数
            auto it1 = record.lower_bound(arr[i]);
            if(it1 != record.end()){
                int idx = (*it1).second;
                if (decrease[idx]){
                    increase[i] = true;
                    res++;
                }
            }
            auto it2 = record.upper_bound(arr[i]);
            if(it2 != record.begin()){
                it2--;
                int idx = (*it2).second;
                if (increase[idx]){
                    decrease[i] = true;
                }
            }
            record[arr[i]] = i;
        }
        return res;
    }
};