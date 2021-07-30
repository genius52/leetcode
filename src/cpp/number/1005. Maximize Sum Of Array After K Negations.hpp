#include <vector>
#include <queue>
#include <set>
using namespace std;

class Solution_1005 {
public:
    int largestSumAfterKNegations(vector<int>& A, int K){
        int len = A.size();
        std::multiset<int> s;//保存最小的K个数
        int sum = 0;
        int min_positive = 2147483647;
        for(int i = 0;i < len;i++){
            sum += A[i];
            if(A[i] < 0){
                s.insert(A[i]);
            }else{
                if(A[i] < min_positive){
                    min_positive = A[i];
                }
            }
        }
        if(s.empty()){
            if ((K | 1) == K){
                sum -= min_positive * 2;
            }
        }else{
            auto it = s.begin();
            int i = 0;
            while(i < K && it != s.end()){
                sum += -(*it) * 2;
                i++;
                it++;
            }
            if(i != K){
                int change = K - i;
                if((change | 1) == change){
                    sum -= min(min_positive,abs(*s.rbegin())) * 2;
                }
            }
        }
        return sum;
    }
//    int largestSumAfterKNegations(vector<int>& A, int K) {
//        quick_sort(A,0,A.size()-1);
//        bool zero_exist = false;
//        int i = A.size() - 1;
//        for(;i > 0 && K > 0;i--){
//            if(A[i] < 0){
//                A[i] = -A[i];
//                K--;
//            }
//            if(A[i] == 0)
//                zero_exist = true;
//        }
//        if(i > 0 || zero_exist){
//            return std::accumulate(A.begin(),A.end(),0);
//        }
//        if(K > 0 && (K % 2) ==1){
//
//            int sum = std::accumulate(A.begin(),A.end(),0);
//            sum += -A[0] * 2;
//            return sum;
//        }
//        return std::accumulate(A.begin(),A.end(),0);
//    }
//    void quick_sort(vector<int>& A ,int begin,int end){
//        if(begin >= end)
//            return;
//        int i = begin;
//        int j = end;
//        int temp = A[begin];
//        while(i<j){
//            while(abs(A[j]) >= abs(temp) && i<j)
//                j--;
//            if(i < j)
//                A[i] = A[j];
//            while(abs(A[i]) <= abs(temp) && i<j)
//                i++;
//            if(i < j)
//                A[j] = A[i];
//        }
//        A[i] = temp;
//        quick_sort(A,begin,i -1);
//        quick_sort(A,i + 1,end);
//    }
};