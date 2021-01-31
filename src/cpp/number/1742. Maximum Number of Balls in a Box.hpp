
class Solution_1742 {
public:
    int countBalls(int lowLimit, int highLimit) {
        int count[100] = {0};
        int max_count = 0;
        for(int i = lowLimit;i <= highLimit;i++){
            int sum = 0;
            int n = i;
            while(n != 0){
                int rest = n % 10;
                sum += rest;
                n = n / 10;
            }
            count[sum]++;
            if(count[sum] > max_count){
                max_count = count[sum];
            }
        }
        return max_count;
    }
};