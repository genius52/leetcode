#include <string>
using namespace std;

class BrowserHistory {
    std::vector<std::string> data;
    int cur = 0;
public:
    BrowserHistory(string homepage) {
        data.push_back(homepage);
        cur = 0;
    }

    void visit(string url) {
        auto begin = data.begin() + cur + 1;
        data.erase(begin, data.end());
        data.push_back(url);
        cur++;
    }

    string back(int steps) {
        int i = 0;
        std::string res = data[cur];
        while (cur > 0 && i < steps) {
            cur--;
            i++;
            res = data[cur];
        }
        return res;
    }

    string forward(int steps) {
        int len = data.size();
        int i = 0;
        std::string res = data[cur];
        while (cur < (len - 1) && i < steps) {
            cur++;
            i++;
            res = data[cur];
        }
        return res;
    }
};