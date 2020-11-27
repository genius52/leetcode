#include <vector>
#include <string>
#include <set>

//Input: ["xga","xfb","yfa"]
//Output: 1

//"doeeqiy","yabhbqe","twckqte"
// 3
//Input: ["xc","yb","za"]
//Output: 0
//Explanation:
//A is already in lexicographic order, so we don't need to delete anything.
//Note that the rows of A are not necessarily in lexicographic order:
//ie. it is NOT necessarily true that (A[0][0] <= A[0][1] <= ...)

//["xga","xfb","yfa"]
//1
class Solution_955 {
public:
    int minDeletionSize(vector<string>& A) {
        int rows = A.size();
        if (rows == 1){
            return 0;
        }
        int columns = A[0].length();
        std::set<int> lines;
        for (int i = 1;i < rows;i++){
            lines.insert(i);
        }
        int res = 0;
        for(int col = 0;col < columns;col++) {
            int l = lines.size();
            if (l <= 0)
                break;
            std::set<int> equal_lines;
            auto cur = lines.begin();
            bool should_delete = false;
            for (; cur != lines.end(); cur++) {
                auto prev = *cur - 1;
                if (A[*cur][col] < A[prev][col]){
                    should_delete = true;
                    break;
                }else if (A[*cur][col] == A[prev][col]){
                    equal_lines.insert(*cur);
                    //equal_lines.insert(*prev);
                }
            }
            if (should_delete){
                res++;
            }else{
                if(!equal_lines.empty()){
                    lines = equal_lines;
                }else{
                    break;
                }
            }
        }
        return res;
    }
};