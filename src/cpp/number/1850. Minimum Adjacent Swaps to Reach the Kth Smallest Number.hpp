#include <string>
using namespace std;

class Solution_1850 {
public:
    int getMinSwaps(string num, int k) {
        std::string num2(num);
        while(k > 0){
            std::next_permutation(num2.begin(),num2.end());
            k--;
        }
        int len = num.size();
        int res = 0;
        for(int i = 0;i < len;i++){
            if(num[i] == num2[i])
                continue;
            //int target = num[i];
            for(int j = i + 1;j < len;j++){
                if(num[i] == num2[j]){
                    for (auto z = j; z > i; --z) {
                        std::swap(num2[z], num2[z-1]);
                        res++;
                    }
                    //std::swap(num2[i],num2[j]);
                    //res += j - i;
                    break;
                }
            }
        }
        return res;
    }
};