#include <string>
#include <unordered_set>
using namespace std;

class Solution_1496 {
    struct position {
        int first;
        int second;
    };
public:
    bool isPathCrossing(string path){
        auto my_hash = [](const position& p) { return p.first + 10001 * p.second; };
        auto my_equal = [](const position& p1, const position& p2) {
            return p1.first == p2.first && p1.second == p2.second;
        };
        std::unordered_set<position, decltype(my_hash), decltype(my_equal)> trace(10001, my_hash, my_equal);
        trace.insert({0,0});
        int x = 0;
        int y = 0;
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
            if(trace.count({x,y}) > 0)
                return true;
            trace.insert({x,y});
        }
        return false;
    }
//    bool isPathCrossing(string path){
//        std::set<std::pair<int,int>> trace;
//        trace.insert({0,0});
//        int x,y = 0;
//        int len = path.size();
//        for(int i = 0;i < len;i++){
//            if(path[i] == 'N')
//                y += 1;
//            else if(path[i] == 'S')
//                y -= 1;
//            else if(path[i] == 'E')
//                x += 1;
//            else if(path[i] == 'W')
//                x -= 1;
//            if(trace.find({x,y}) != trace.end())
//                return true;
//            trace.insert({x,y});
//        }
//        return false;
//    }

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