#include <vector>
#include <map>
using namespace std;


// func LargestSubmatrix(matrix [][]int) int {
//	var rows int = len(matrix)
//	var columns int = len(matrix[0])
//	var record [][]int = make([][]int,rows)
//	for i := 0;i < rows;i++{
//		record[i] = make([]int,columns)
//	}
//	for j := 0;j < columns;j++{
//		for i := 0;i < rows;i++{
//			if i == 0{
//				record[i][j] = matrix[i][j]
//			}else{
//				if matrix[i][j] == 1{
//					record[i][j] = 1 + record[i - 1][j]
//				}
//			}
//		}
//	}
//	var max_area int = 0
//	for i := 0;i < rows;i++{
//		var calc map[int]int = make(map[int]int)
//		for j := 0;j < columns;j++{
//			if record[i][j] == 0{
//				continue
//			}
//			//for h,_ := range calc{
//			//	if h < record[i][j]{
//			//		calc[h]++
//			//	}
//			//}
//			if _,ok := calc[record[i][j]];ok{
//				calc[record[i][j]]++
//			}else{
//				calc[record[i][j]] = 1
//			}
//		}
//		for h,w := range calc{
//			area := h * w
//			if area > max_area{
//				max_area = area
//			}
//		}
//	}
//	return max_area
//}
class Solution_1727 {
public:
    int largestSubmatrix(vector<vector<int>>& matrix) {
        int rows  = matrix.size();
        int columns = matrix[0].size();
        std::vector<std::vector<int>> record(rows,std::vector<int>(columns));
        for (int j = 0;j < columns;j++){
            for (int i = 0;i < rows;i++){
                if (i == 0){
                    record[i][j] = matrix[i][j];
                }else{
                    if (matrix[i][j] == 1){
                        record[i][j] = 1 + record[i - 1][j];
                    }
                }
            }
        }

        int max_area = 0;
        for (int i = 0;i < rows;i++){
            std::map<int,int> calc;
            for (int j = 0;j < columns;j++){
                if (record[i][j] == 0){
                    continue;
                }
                calc[record[i][j]]++;
            }

            int prev_len = 0;
            for (auto it = calc.rbegin();it != calc.rend();it++){
                int height = (*it).first;
                int width = (*it).second;
                int area = height * (width + prev_len);
                if (area > max_area){
                    max_area = area;
                }
                prev_len += width;
            }
        }
        return max_area;
    }
};