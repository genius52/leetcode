
//Given a positive integer n, break it into the sum of at least two positive integers and
// maximize the product of those integers. Return the maximum product you can get.

//Input: 10
//Output: 36
//Explanation: 10 = 3 + 3 + 4, 3 × 3 × 4 = 36.
class Solution_343 {
public:
    int calc(int num,int divide){
        int sub_num = num / divide;
        int rest = num % divide;
        int res = 1;
        for(int i = 0;i < divide;i++){
            if(rest > 0){
                res *= (sub_num + 1);
                rest--;
            }else{
                res *= sub_num;
            }
        }
        return res;
    }

    int integerBreak(int n) {
        int min_part = 2;
        int max_part = n - 1;
        int max = 1;
        for(int i = min_part;i <= max_part;i++){
            auto res = calc(n,i);
            if(res > max){
                max = res;
            }
        }
        return max;
    }
};