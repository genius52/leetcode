
class Solution_190 {
public:
    uint32_t reverseBits(uint32_t n) {
        uint32_t result = 0;
        uint32_t mask = 1;
        for(int i = 0;i < 32;i++){
            result = result << 1;
            result |= mask & n;
            n = n >> 1;
        }
        return result;
    }
};