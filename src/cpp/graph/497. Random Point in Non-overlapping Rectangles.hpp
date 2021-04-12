#include <vector>
using namespace std;

class Solution_497 {
public:
    int total = 0;
    std::map<int,int> area_index;//range -> index(in rects)
    std::vector<std::vector<int>> rects_;
    int calc_area(std::vector<int> rect){
        return (rect[2] - rect[0] + 1) * (rect[3] - rect[1] + 1);
    }

    Solution_497(vector<vector<int>>& rects) {
        rects_ = rects;
        int len = rects.size();
        int last_range = 0;
        for(int i = 0;i < len;i++){
            int area = calc_area(rects[i]);
            total += area;
            last_range = area;
            area_index[total] = i;
        }
        //total -= last_range;
    }

    vector<int> pick() {
        auto r = rand() % total;
        auto it = area_index.upper_bound(r);
        std::vector<int> res(2);
        if(it == area_index.end())
            return res;
        int left = rects_[it->second][0];
        int bottom = rects_[it->second][1];
        int right = rects_[it->second][2];
        int top = rects_[it->second][3];
        int width = right - left;
        int height = top - bottom;
        int r_height = 0;
        if(height != 0){
            r_height = rand() % (height + 1);
        }
        int r_width = 0;
        if(width != 0 ){
            r_width = rand() % (width + 1);
        }
        res[0] = left + r_width;
        res[1] = bottom + r_height;
        return res;
    }
};