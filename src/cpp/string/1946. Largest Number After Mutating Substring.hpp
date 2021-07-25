#include <vector>
#include <string>
using namespace std;

class Solution_1946 {
public:
    string maximumNumber(string num, vector<int>& change) {
        int l = num.size();
        std::string res;
        bool start = false;
        bool end = false;
        for (int i = 0; i < l; i++) {
            int n1 = int(num[i] - '0');
            int n2 = change[n1];
            if (n1 == n2) {
                res += num[i];
            }
            else if (n2 > n1 && !end) {
                start = true;
                res += std::to_string(n2);
            }
            else {
                if (start) {
                    end = true;
                }
                res += num[i];
            }
        }
        return res;
    }
};
