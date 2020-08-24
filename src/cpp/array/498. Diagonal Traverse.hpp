#include <vector>
using namespace std;

//Input:
//[
//    [1, 2, 3],
//    [4, 5, 6],
//    [7, 8, 9]
//]
//Output: [1, 2, 4, 7, 5, 3, 6, 8, 9]
class Solution_498 {
public:
    vector<int> findDiagonalOrder(vector<vector<int>>& matrix) {
        int rows = matrix.size();
        std::vector<int> res;
        if (rows == 0)
            return res;
        int columns = matrix[0].size();
        res.resize(rows * columns);
        int index = 0;
        //res[index++] = matrix[0][0];
        int x_dir = -1;
        int y_dir = 1;
        int visit_row = 0;
        int visit_column = 0;
        int total = rows * columns;
        bool left_to_right = true;
        while (index < total) {
            while (visit_row >= 0 && visit_row < rows && visit_column >= 0 && visit_column < columns) {
                res[index++] = matrix[visit_row][visit_column];
                if (left_to_right) 
                {
                    if (visit_row == 0) {
                        if (visit_column < columns - 1) {
                            visit_column++;
                        }
                        else {
                            visit_row++;
                        }
                        break;
                    }
                    else if (visit_column == columns - 1) {
                        visit_row++;
                        break;
                    }
                }
                else {
                    if (visit_column == 0) {
                        if (visit_row < rows - 1) {
                            visit_row++;
                        }
                        else {
                            visit_column++;
                        }
                        break;
                    }
                    else if (visit_row == rows - 1) {
                        visit_column++;
                        break;
                    }
                }
                visit_row += x_dir;
                visit_column += y_dir;
            }
            x_dir = -x_dir;
            y_dir = -y_dir;
            left_to_right = !left_to_right;
        }
        return res;
    }
};