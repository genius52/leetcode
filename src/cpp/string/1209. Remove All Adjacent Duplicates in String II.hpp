#include <string>
#include <deque>

using namespace std;

//Input: s = "deeedbbcccbdaa", k = 3
//Output: "aa"
class Solution_1209 {
public:
    string removeDuplicates(string s, int k) {
        std::deque<std::pair<char,int>> q;
        std::pair<char,int> p;
        p.first = '#';
        p.second = 1;
        q.push_back(p);
        int len = s.length();
        for(int i = 0;i < len;i++){
            if (q.back().first == s[i]){
                q.back().second++;
            }else{
                q.push_back({s[i],1});
            }
            if(q.back().second == k)
                q.pop_back();
        }
        q.pop_front();
        std::string res;
        while(!q.empty()){
            auto front = q.front();
            q.pop_front();
            res.append(front.second,front.first);
        }
        return res;
    }
};