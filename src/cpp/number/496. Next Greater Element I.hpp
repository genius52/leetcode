#include <vector>
#include <stack>
#include <unordered_map>
using namespace std;

class Solution_496 {
public:
    vector<int> nextGreaterElement(vector<int>& nums1, vector<int>& nums2) {
        int len2 = nums2.size();
        std::stack<int> s;
        std::unordered_map<int,int> m;
        for(int i = 0;i < len2;i++){
            if(s.empty() || s.top() > nums2[i]){
                s.push(nums2[i]);
            }else{
                while (!s.empty() && s.top() < nums2[i]){
                    m[s.top()] = nums2[i];
                    s.pop();
                }
                s.push(nums2[i]);
            }
        }
        int len1 = nums1.size();
        std::vector<int> res(len1);
        for(int i = 0;i < len1;i++){
            if(m.find(nums1[i]) != m.end()){
                res[i] = m[nums1[i]];
            }else{
                res[i] = -1;
            }
        }
        return res;
    }
};