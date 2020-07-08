#include <vector>
#include <math.h>
using namespace std;

//Input: 4
//Output: [2, 2]
//Explanation: The target area is 4, and all the possible ways to construct it are [1,4], [2,2], [4,1].
//But according to requirement 2, [1,4] is illegal; according to requirement 3,  [4,1] is not optimal compared to [2,2]. So the length L is 2, and the width W is 2.
class Solution_492 {
public:
    vector<int> constructRectangle(int area) {
        int big = area;
        int small = 1;
        int limit = static_cast<int>(sqrt(area));
        for(int i = 1;i <= limit;i++){
            int rest = area % i;
            if(rest == 0){
                small = i;
                big = area / i;
            }
        }
        return std::vector<int>{big,small};
    }
};