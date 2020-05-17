#include <string>
#include <vector>
using namespace std;

class Solution_1447 {
public:
    int gcd(int x, int y)
    {
        if (y == 0)
            return x;
        if (x < y)
            return gcd(y, x);
        else
            return gcd(y, x % y);
    }

    vector<string> simplifiedFractions(int n) {
        std::vector<std::string> res;
        for (int i = 2; i <= n; i++) {
            for (int j = 1; j < i; j++) {
                if (gcd(i, j) == 1)
                    res.push_back(std::to_string(j) + "/" + std::to_string(i));
            }
        }
        return res;
    }
};