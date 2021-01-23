#include <vector>
using namespace std;

class Solution_85 {
public:
    int maximalRectangle(std::vector<std::vector<char> >& matrix) {
        auto rows = matrix.size();
        auto columns = matrix[0].size();
        std::vector<std::vector<int> > dp_width(rows,std::vector<int>(columns,0));
        std::vector<std::vector<int> > dp_height(rows,std::vector<int>(columns,0));
        std::vector<std::vector<int> > dp_area(rows,std::vector<int>(columns,0));
        int max = 0;
        for(int i = 0;i < rows;i++){
            dp_width[i][0] = matrix[i][0] == '1'? 1: 0;
            for(int j = 0;j < columns;j++){
                dp_height[0][j] = matrix[0][j] == '1'? 1: 0;
                if (matrix[i][j] == '1'){
                    if (j >= 1){
                        dp_width[i][j] = 1 + dp_width[i][j - 1];
                    }
                    if(i >= 1){
                        dp_height[i][j] = 1 + dp_height[i - 1][j];
                    }
                    if (i == 0 || j == 0){
                        dp_area[i][j] = dp_width[i][j] * dp_height[i][j];
                    }else{

                    }
                    if (dp_area[i][j] > max){
                        max = dp_area[i][j];
                    }
                }
            }
        }
        return max;
    }
};