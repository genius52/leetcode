#include <vector>
using namespace std;

class Solution_1823 {
public:
    int findTheWinner(int n, int k) {
        std::vector<int> record(n);
        for(int i = 0;i < n;i++){
            record[i] = i + 1;
        }
        int cnt = 0;
        int pos = 0;
        while(cnt < (n - 1)){
            int cur_len = record.size();
            int next = (pos + k - 1) % cur_len;
            if(next == cur_len - 1){
                pos = 0;
            }else{
                pos = next;
            }
            std::vector<int>::iterator it = record.begin() + next;
            record.erase(it);
            cnt++;
        }
        return record[0];
    }
};