#include <vector>
#include <set>
#include <stack>
using namespace std;

//1,17,18,0,18,10,20,0
class Solution_2454 {
public:
    vector<int> secondGreaterElement(vector<int>& nums) {
        int len = nums.size();
        std::vector<int> res(len,-1);
        std::stack<int> s1;//没有发现更大数字的索引，降序
        std::stack<int> s2;//发现了一个更大数字的索引，降序
        for (int i = 0; i < len; ++i) {
            while(!s2.empty() && nums[i] > nums[s2.top()]){
                res[s2.top()] = nums[i];
                s2.pop();
            }
            std::stack<int> tmp;
            while(!s1.empty() && nums[i] > nums[s1.top()]){
                tmp.push(s1.top());
                s1.pop();
            }
            while(!tmp.empty()){
                s2.push(tmp.top());
                tmp.pop();
            }
            s1.push(i);
        }
        return res;
    }
};