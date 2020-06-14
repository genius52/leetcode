#include <vector>

//A = [4, 3, 2, 6]
//
//F(0) = (0 * 4) + (1 * 3) + (2 * 2) + (3 * 6) = 0 + 3 + 4 + 18 = 25
//F(1) = (0 * 6) + (1 * 4) + (2 * 3) + (3 * 2) = 0 + 4 + 6 + 6 = 16
//F(2) = (0 * 2) + (1 * 6) + (2 * 4) + (3 * 3) = 0 + 6 + 8 + 9 = 23
//F(3) = (0 * 3) + (1 * 2) + (2 * 6) + (3 * 4) = 0 + 2 + 12 + 12 = 26
//
//So the maximum value of F(0), F(1), F(2), F(3) is F(3) = 26.
//TO DO
class Solution_396 {
public:
    int maxRotateFunction(vector<int>& A) {
        int len = A.size();
        if(len <= 1)
            return 0;
        int max = -2147483648;
        for(int i = 0;i < len;i++){
            int sum = 0;
            for(int j = 0;j < len;j++){
                sum += j * A[(j + len - i) % len];
            }
            if(sum > max){
                max = sum;
            }
        }
        return max;
    }
};