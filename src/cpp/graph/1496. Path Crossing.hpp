#include <string>
#include <unordered_set>
using namespace std;

class Solution_1496 {
public:
    bool isPathCrossing(string path){
        std::set<std::pair<int,int>> trace;
        trace.insert({0,0});
        int x,y = 0;
        int len = path.size();
        for(int i = 0;i < len;i++){
            if(path[i] == 'N')
                y += 1;
            else if(path[i] == 'S')
                y -= 1;
            else if(path[i] == 'E')
                x += 1;
            else if(path[i] == 'W')
                x -= 1;
            if(trace.find({x,y}) != trace.end())
                return true;
            trace.insert({x,y});
        }
        return false;
    }
//    bool isPathCrossing(string path) {
//        std::unordered_set<std::string> s;
//        s.insert("0,0");
//        int horizon = 0;
//        int vertical = 0;
//        int len = path.length();
//        for(int i = 0;i < len;i++){
//            if(path[i] == 'N')
//                horizon++;
//            else if(path[i] == 'S')
//                horizon--;
//            else if(path[i] == 'E')
//                vertical++;
//            else if(path[i] == 'W')
//                vertical--;
//            std::string v =  std::to_string(horizon) + "," + std::to_string(vertical);
//            if(s.find(v) != s.end())
//                return true;
//            s.insert(v);
//        }
//        return false;
//    }
};