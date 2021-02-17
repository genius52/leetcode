#include <string>
#include <vector>
#include <unordered_set>
using namespace std;

class Solution_752 {
public:
    int openLock(vector<string>& deadends, string target) {
        std::unordered_set<std::string> visited;
        for(auto s : deadends){
            visited.insert(s);
        }
        std::queue<std::string> q;
        q.push("0000");
        if(visited.find("0000") != visited.end())
            return -1;
        visited.insert("0000");
        int steps = 0;
        while(!q.empty()){
            int len = q.size();
            for(int i = 0;i < len;i++){
                auto num = q.front();
                q.pop();
                if(num == target)
                    return steps;
                for(int j = 0;j < 4;j++){
                    int old_bit = num[j] - '0';
                    std::string add_num;
                    std::string minus_num;
                    for(int k = 0;k < 4;k++){
                        if(k == j){
                            add_num += std::to_string((old_bit + 1) % 10);
                            minus_num += std::to_string((old_bit + 10 - 1) % 10);
                        }
                        else{
                            add_num += num[k];
                            minus_num += num[k];
                        }
                    }
                    if(visited.find(add_num) == visited.end())
                        q.push(add_num);
                    if(visited.find(minus_num) == visited.end())
                        q.push(minus_num);
                    visited.insert(add_num);
                    visited.insert(minus_num);
                }
            }
            steps++;
        }
        return -1;
    }
};