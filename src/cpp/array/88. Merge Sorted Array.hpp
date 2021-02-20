#include <vector>
using namespace std;

class Solution_88 {
public:
    void merge(vector<int>& nums1, int m, vector<int>& nums2, int n) {
        int p1 = 0;
        int p2 = 0;
        int cur = 0;
        vector<int> v;
        v.resize(m+n);
        while(p1 < m && p2 < n){
            if(nums1[p1] <= nums2[p2])
                v[cur++] = nums1[p1++];
            else
                v[cur++] = nums2[p2++];
        }
        while(p1 < m){
            v[cur++] = nums1[p1++];
        }
        while(p2 < n){
            v[cur++] = nums2[p2++];
        }
        for(int i = 0;i < (m+n);i++){
            nums1[i] = v[i];
        }
    }
};