#include <string>
#include <vector>
//using namespace std;

//Input: pizza = ["A..","AAA","..."], k = 3
//Output: 3
class Solution_1444 {
public:
    bool apple_exist(vector<string>& pizza,int start_row,int start_col,int end_row,int end_col){
        if(start_row == end_row && start_col == end_col)
            return false;
        if(start_row == end_row){
            for(int j = start_col;j < end_col;j++){
                if(pizza[start_row][j] == 'A')
                    return true;
            }
        }else if(start_col == end_col){
            for(int i = start_row;i < end_row;i++){
                if(pizza[i][start_col] == 'A'){
                    return true;
                }
            }
        }else{
            for(int i = start_row;i < end_row;i++){
                for(int j = start_col;j < end_col;j++){
                    if(pizza[i][j] == 'A')
                        return true;
                }
            }
        }
        return false;
    }

    void dfs_ways(vector<string>& pizza,int rows,int columns,int cur_cut,int k,int start_row,int start_col,int end_row,int end_col,int& total){
        if(cur_cut > k)
            return;
        if(cur_cut == k){
            if(apple_exist(pizza,start_row,start_col,end_row,end_col))
                total++;
            return;
        }
        //cut by row
        bool last_row_checked = false;
        for(int mid_row = start_row + 1;mid_row < rows;mid_row++){
            if(!last_row_checked){
                if(apple_exist(pizza,start_row,start_col,mid_row,end_col)){
                    last_row_checked = true;
                }else{
                    continue;
                }
            }
            dfs_ways(pizza,rows,columns,cur_cut + 1,k,mid_row,start_col,end_row,end_col,total);
        };
        //cut by column
        bool last_column_checked = false;
        for(int mid_col = start_col + 1;mid_col < columns;mid_col++){
            if(!last_column_checked){
                if(apple_exist(pizza,start_row,start_col,end_row,mid_col)){
                    last_column_checked = true;
                }else{
                    continue;
                }
            }
            dfs_ways(pizza,rows,columns,cur_cut + 1,k,start_row,mid_col,end_row,end_col,total);
        }
    }

    int ways(vector<string>& pizza, int k) {
        int rows = pizza.size();
        if (rows == 0)
            return 0;
        int columns = pizza[0].length();
        int total = 0;
        dfs_ways(pizza,rows,columns,1,k,0,0,rows,columns,total);
        return total;
    }
};