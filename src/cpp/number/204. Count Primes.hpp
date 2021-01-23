
class Solution_204 {
public:
    int countPrimes(int n) {
        if(n == 1)
            return 0;
        if(n == 2)
            return 1;
        int cnt = 0;
        bool find = false;
        for(int i = 2;i <= n;i++){
            int val = static_cast<int>(sqrt(i))+1;
            for(int j = 2; j <= val;j++){
                if(val == 2){
                    find = false;
                    break;
                }
                if(i%j == 0){
                    find = true;
                    break;
                }
            }
            if(!find)
            {
                find = false;
                cnt++;
            }
        }
        return cnt;
    }
};