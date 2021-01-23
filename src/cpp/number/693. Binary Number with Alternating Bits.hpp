
class Solution_693 {
public:
    bool hasAlternatingBits(int n) {
        if(n == 1 || n == 0)
            return true;
        long tag = 0b10;
        if(n & 1 == n)
            tag = tag << 1;
        while(tag <= n){
            if(tag | n != n)
                return false;
            tag = tag<<2;
        }
        return true;
    }
};