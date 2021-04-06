#include <vector>
#include <unordered_map>
using namespace std;

class Solution_454 {
public:
    int fourSumCount(vector<int>& A, vector<int>& B, vector<int>& C, vector<int>& D) {
        int len = A.size();
        unordered_map<int,int> map;
        for(int i = 0;i < len;i++)
            for(int j = 0;j < len;j++){
                auto sum = -(A[i]+B[j]);
                map[sum]++;
            }

        int cnt = 0;
        for(int i = 0;i < len;i++)
            for(int j = 0;j < len;j++){
                auto sum = C[i] + D[j];
                if(map.find(sum) != map.end()){
                    cnt += map[sum];
                }
            }
        return cnt;
    }
};