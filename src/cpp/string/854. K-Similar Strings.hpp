#include <string>
#include <unordered_set>
#include <queue>

class Solution_854 {
public:
    int kSimilarity(string s1, string s2) {
        int l = s1.length();
        std::unordered_set<std::string> visited;
        int steps = 0;
        std::queue<std::string> q;
        q.push(s1);
        visited.insert(s1);
        while(q.size() > 0){
            int len = q.size();
            for(int i = 0;i < len;i++){
                auto cur = q.front();
                q.pop();
                if(cur == s2)
                    return steps;
                for(int j = 0;j < l;j++){
                    if(cur[j] == s2[j])
                        continue;
                    //从j后面的位置找到cur[k] == s2[j]
                    for(int k = j + 1;k < l;k++){
                        if(cur[k] == s2[k] || cur[k] != s2[j]){
                            continue;
                        }
                        std::swap(cur[j],cur[k]);
                        if(visited.find(cur) == visited.end()){
                            visited.insert(cur);
                            q.push(cur);
                        }
                        std::swap(cur[j],cur[k]);
                    }
                    break;
                }
            }
            steps++;
        }
        return -1;
    }
};