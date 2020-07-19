#include <string>
#include <vector>
using namespace std;

//Input:
//widths = [4,10,10,10,10,10,10,10,10,10,10,10,10,10,10,10,10,10,10,10,10,10,10,10,10,10]
//S = "bbbcccdddaaa"
//Output: [2, 4]
class Solution_806 {
public:
    vector<int> numberOfLines(vector<int>& widths, string S) {
        std::vector<int> res;
        res.resize(2);
        int lines = 1;
        int cur_length = 0;
        for(auto c : S){
            int char_len = widths[c - 'a'];
            if(cur_length + char_len > 100){
                lines++;
                cur_length = char_len;
            }else{
                cur_length += char_len;
            }
        }
        res[0] = lines;
        res[1] = cur_length;
        return res;
    }
};