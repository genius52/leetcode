#include <string>
#include <queue>
using namespace std;

class Solution_1405 {
public:
    struct cmp {
        bool operator()(const std::pair<int, char> a, const std::pair<int, char> b) {
            return a.first < b.first;
        }
    };
    string longestDiverseString(int a, int b, int c) {
        std::priority_queue<std::pair<int, char>, std::vector<std::pair<int, char> >, cmp > q;
        if(a > 0)
            q.push(std::pair<int, char>(a, 'a'));
        if(b > 0)
            q.push(std::pair<int, char>(b, 'b'));
        if(c > 0)
            q.push(std::pair<int, char>(c, 'c'));
        std::string res;
        while (!q.empty() ) {
            auto top1 = q.top();
            q.pop();
            if (q.empty()) {
                if (top1.first >= 2) {
                    res += top1.second;
                    res += top1.second;
                }
                else if (top1.first == 1) {
                    res += top1.second;
                }
                return res;
            }
            auto top2 = q.top();
            q.pop();
            auto len = res.length();
            if ((len == 0) || (len > 0) && res[len - 1] != top1.second) {
                if (top1.first >= 2) {
                    res += top1.second;
                    res += top1.second;
                    top1.first -= 2;
                }
                else if (top1.first == 1) {
                    res += top1.second;
                    top1.first--;
                }
                if (top2.first >= 1) {
                    res += top2.second;
                    top2.first--;
                }
            }
            else {
                if (top2.first >= 2) {
                    res += top2.second;
                    res += top2.second;
                    top2.first -= 2;
                }
                else if (top2.first >= 1) {
                    res += top2.second;
                    top2.first--;
                }
                if (top1.first == 1) {
                    res += top1.second;
                    top1.first--;
                }
            }

            if(top1.first > 0)
                q.push(top1);
            if(top2.first > 0)
                q.push(top2);
        }
        return res;
    }
};