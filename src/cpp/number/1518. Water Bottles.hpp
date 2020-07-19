

class Solution_1518 {
public:
    int numWaterBottles(int numBottles, int numExchange) {
        int res = numBottles;
        int rest_bottles = numBottles;
        while(rest_bottles >= numExchange){
            int fullbottles = rest_bottles / numExchange;
            res += fullbottles;
            rest_bottles = rest_bottles % numExchange + fullbottles;
        }
        return res;
    }
};