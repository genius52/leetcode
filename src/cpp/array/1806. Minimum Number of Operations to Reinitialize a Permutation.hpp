#include <vector>
//Input: n = 4
//Output: 2
//Explanation: prem = [0,1,2,3] initially.
//After the 1st operation, prem = [0,2,1,3]
//After the 2nd operation, prem = [0,1,2,3]
//So it takes only 2 operations.
//If i % 2 == 0, then arr[i] = perm[i / 2].
//If i % 2 == 1, then arr[i] = perm[n / 2 + (i - 1) / 2].
class Solution_1806 {
public:
    int reinitializePermutation(int n) {
        std::vector<int> data(n);
        for(int i = 0;i < n;i++){
            data[i] = i;
        }
        int steps = 1;
        while(true){
            std::vector<int> old = data;
            for(int i = 0;i < n;i++){
                if((i | 1) == i){//odd
                    data[i] = old[n/2 + (i - 1)/2];
                }else{
                    data[i] = old[i/2];
                }
            }
            bool match = true;
            for(int i = 0;i < n;i++){
                if(data[i] != i){
                    match = false;
                    break;
                }
            }
            if(match){
                return steps;
            }
            steps++;
        }
        return steps;
    }
};