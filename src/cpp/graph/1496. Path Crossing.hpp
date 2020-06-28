#include <string>
#include <unordered_set>
using namespace std;

class Solution_1496 {
public:
    bool isPathCrossing(string path) {
        std::unordered_set<std::string> s;
        s.insert("0,0");
        int horizon = 0;
        int vertical = 0;
        int len = path.length();
        for(int i = 0;i < len;i++){
            if(path[i] == 'N')
                horizon++;
            else if(path[i] == 'S')
                horizon--;
            else if(path[i] == 'E')
                vertical++;
            else if(path[i] == 'W')
                vertical--;
            std::string v =  std::to_string(horizon) + "," + std::to_string(vertical);
            if(s.find(v) != s.end())
                return true;
            s.insert(v);
        }
        return false;
    }
};