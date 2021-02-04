#include <vector>
#include <algorithm>
using namespace std;

class Solution_4 {
public:
    double findMedianSortedArrays(vector<int>& nums1, vector<int>& nums2){
        std::vector<int> v;
        v.insert(v.end(),nums1.begin(),nums1.end());
        v.insert(v.end(),nums2.begin(),nums2.end());
        int len = v.size();
        if((len | 1) == len){
            int pos = len/2;
            std::nth_element(v.begin(),v.begin() + pos,v.end());
            return v[pos];
        }else{
            int pos1 = (len - 1)/2;
            int pos2 = len/2;
            std::nth_element(v.begin(),v.begin() + pos1,v.end());
            int first = v[pos1];
            std::nth_element(v.begin(),v.begin() + pos2,v.end());
            int second = v[pos2];
            return (double(first) + double(second))/2.0;
        }
    }
//    double findMedianSortedArrays(vector<int>& nums1, vector<int>& nums2) {
//        int len1 = nums1.size();
//        int len2 = nums2.size();
//        bool odd = (((len1 + len2) | 1) == len1 + len2);
//        if(nums1.empty()){
//            if(odd)
//                return nums2[len2/2];
//            else
//                return (double)(nums2[len2/2] + nums2[len2/2 - 1])/2;
//        }
//        if(nums2.empty()){
//            if(odd)
//                return nums1[len1/2];
//            else
//                return (double)(nums1[len1/2] + nums1[len1/2 - 1])/2;
//        }
//
//        int pos = 0;
//        int l1 = 0;
//        int l2 = 0;
//        int mid_num1 = 0;
//        int mid_num2 = 0;
//        while(pos <= (len1 + len2)/2){
//            if(l1 < len1 && l2 < len2){
//                if(nums1[l1] <= nums2[l2]){
//                    if(!odd)
//                        mid_num2 = mid_num1;
//                    mid_num1 = nums1[l1];
//                    l1++;
//                }
//                else{
//                    if(!odd)
//                        mid_num2 = mid_num1;
//                    mid_num1 = nums2[l2];
//                    l2++;
//                }
//                pos++;
//                continue;
//            }
//
//            if(l1 < len1){
//                if(!odd)
//                    mid_num2 = mid_num1;
//                mid_num1 = nums1[l1];
//                l1++;
//                pos++;
//                continue;
//            }
//            if(l2 < len2){
//                if(!odd)
//                    mid_num2 = mid_num1;
//                mid_num1 = nums2[l2];
//                l2++;
//                pos++;
//                continue;
//            }
//        }
//        std::cout<<mid_num1<<std::endl;
//        std::cout<<mid_num2<<std::endl;
//        if(!odd)
//            return (double)(mid_num1 + mid_num2)/2;
//        return mid_num1;
//    }
};