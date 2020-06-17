#include <vector>
using namespace std;

class SubrectangleQueries {
    vector<vector<int>> data;
    int rows;
    int columns;
public:
    SubrectangleQueries(vector<vector<int>>& rectangle) {
        rows = rectangle.size();
        columns = rectangle[0].size();
        data = rectangle;
    }

    void updateSubrectangle(int row1, int col1, int row2, int col2, int newValue) {
        for(int i = row1;i <= row2;i++){
            std::fill(data[i].begin()+ col1,data[i].begin() + col2 + 1,newValue);
        }
    }

    int getValue(int row, int col) {
        return data[row][col];
    }
};