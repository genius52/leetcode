
//Input: 6
//Output: 1
//Explanation:
//6 in binary is 0b110.
class Solution_868 {
public:
    int binaryGap(int N) {
        int last_one_pos = -1;
        int cur_one_pos = 0;
        int max_gap = 0;
        while(N > 0){
            if((N | 1) == N){
                if(last_one_pos == -1){
                    last_one_pos = cur_one_pos;
                }
                else{
                    if(cur_one_pos - last_one_pos > max_gap)
                        max_gap = cur_one_pos - last_one_pos;
                    last_one_pos = cur_one_pos;
                }
            }
            N >>= 1;
            cur_one_pos++;
        }
        return max_gap;
    }
};