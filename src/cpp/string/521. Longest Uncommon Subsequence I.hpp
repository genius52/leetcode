#include <string>
#include <math>
using namespace std;

class Solution_521 {
public:
    int findLUSlength(string a, string b) {
        if(a == b)
            return -1;
        return max(a.length(),b.length());
    }
};