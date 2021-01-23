#include <vector>
#include <algorithm>
using namespace std;

struct point{
    int x;
    int y;
    int dis = 0;
};

std::vector<std::vector<int>> allCellsDistOrder(int R, int C, int r0, int c0) {
    std::vector<std::vector<int>> res;
    if(r0 > R || c0 > C)
        return res;
    std::vector<point> vp;
    for(int i = 0;i < R;i++){
        for(int j = 0;j < C;j++){
            point p;
            p.x = j;
            p.y = i;
            p.dis = abs(i - r0) + abs(j - c0);
            vp.push_back(p);
        }
    }
    std::sort(vp.begin(),vp.end(),[](const point& p1,const point& p2)->bool{
        return p1.dis > p2.dis;
    });
    for(auto it = vp.begin();it != vp.end();it++){
        std::vector<int> v;
        v.push_back(it->x);
        v.push_back(it->y);
        res.push_back(v);
    }
    return res;
}


