#include <vector>
using namespace std;

//Input: salary = [4000,3000,1000,2000]
//Output: 2500.00000
//Explanation: Minimum salary and maximum salary are 1000 and 4000 respectively.
class Solution_1491 {
public:
    double average(vector<int>& salary) {
        int len = salary.size();
        int min = salary[0];
        int max = salary[0];
        int total = salary[0];
        for(int i = 1;i < len;i++){
            total += salary[i];
            if(salary[i] < min)
                min = salary[i];
            if(salary[i] > max)
                max = salary[i];
        }
        double res = double(total - min - max) / (len - 2);
        return res;
    }
};