#include <vector>
#include <math.h>
using namespace std;

//S=(x1y2+x2y3+x3y1-x1y3-x2y1-x3y2) /2;
//S=(x1y2-x1y3+x2y3-x2y1+x3y1-x3y2)
class Solution_812 {
public:
    double largestTriangleArea(vector<vector<int>>& points) {
        int len = points.size();
        double max_area = 0;
        for(int first = 0;first < len;first++){
            for(int second = first + 1;second < len;second++){
                for(int third = second + 1;third < len;third++){
                    double area = abs(points[first][0] * points[second][1] - points[first][0] * points[third][1] +
                                points[second][0] * points[third][1] - points[second][0] * points[first][1] +
                                points[third][0] * points[first][1] - points[third][0] * points[second][1]) / 2.0;
                    if(area > max_area){
                        max_area = area;
                    }
                }
            }
        }
        return max_area;
    }
};