#include <math.h>
#include <algorithm>

/*
[3, 0, 8, 4],
[2, 4, 5, 7],
[9, 2, 6, 3],
[0, 3, 1, 0]

gridNew = [
[8, 4, 8, 7],
[7, 4, 7, 7],
[9, 4, 8, 7],
[3, 3, 3, 3] ]
*/
int maxIncreaseKeepingSkyline(std::vector<std::vector<int>>& grid) {
    int sum = 0;
    if(grid.size() == 0)
        return sum;

    std::vector<int> max_height_rows;
    for(int i = 0;i < grid.size();i++){
        int max_height_current_row = 0;
        for (int j = 0; j < grid[0].size(); ++j) {
            if(grid[i][j] > max_height_current_row)
                max_height_current_row = grid[i][j];
        }
        max_height_rows.push_back(max_height_current_row);
    }

    std::vector<int> max_height_columns;
    for(int i = 0;i < grid[0].size();i++){
        int max_height_current_column = 0;
        for (int j = 0; j < grid.size(); ++j) {
            if(grid[j][i] > max_height_current_column)
                max_height_current_column = grid[j][i];
        }
        max_height_columns.push_back(max_height_current_column);
    }

    for(int i = 0;i < grid.size();i++){
        for (int j = 0; j < grid[0].size(); ++j) {
            auto available_height = max_height_rows[i]<max_height_columns[j]?max_height_rows[i]:max_height_columns[j];
            if(available_height > grid[i][j])
                sum += abs(available_height-grid[i][j]);
        }
    }
    return sum;
}
