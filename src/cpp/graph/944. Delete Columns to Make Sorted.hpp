#include <vector>
#include <string>
using namespace std;

class Solution_944 {
public:
    int minDeletionSize(vector<string>& A) {
        int rows = A.size();
        int columns = A[0].size();
        int des_cnt = 0;
        for(int i = 0;i < columns;i++){
            for(int j = 0;j < rows - 1;j++){
                if(A[j][i] > A[j + 1][i]){
                    des_cnt++;
                    break;
                }
            }
        }
        return des_cnt;
    }
};