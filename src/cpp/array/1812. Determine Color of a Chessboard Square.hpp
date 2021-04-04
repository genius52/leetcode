#include <iostream>
using namespace std;

class Solution_1812 {
public:
    bool squareIsWhite(string coordinates) {
        int x_index = coordinates[0] - 'a';
        int y_index = coordinates[1] - '1';
        if((((x_index | 1) == x_index) && ((y_index | 1) == y_index)) ||
        (((x_index | 1) != x_index) && ((y_index | 1) != y_index) )){
            return false;
        }
        return true;
    }
};