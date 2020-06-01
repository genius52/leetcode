#include <vector>
using namespace std;

class Solution_386 {
    static bool cmp(const int& a, const int& b)
    {
        auto a_str = std::to_string(a);
        auto len_a = a_str.length();
        auto b_str = std::to_string(b);
        auto len_b = b_str.length();
        if (len_a == len_b)
        {
            return a < b;
        }
        else
        {
            int a_pos = 0;
            int b_pos = 0;
            while(a_pos < len_a && b_pos < len_b){
               if(a_str[a_pos] == b_str[b_pos]){
                   a_pos++;
                   b_pos++;
               }else{
                   return a_str[a_pos] < b_str[b_pos];
               }
            }
            return len_a < len_b;
        }
    }
public:
    vector<int> lexicalOrder(int n) {
        std::vector<int> res;
        res.resize(n);
        for(int i = 1;i <= n;i++){
            res[i - 1] = i;
        }
        std::sort(res.begin(),res.end(),cmp);
        return res;
    }
};