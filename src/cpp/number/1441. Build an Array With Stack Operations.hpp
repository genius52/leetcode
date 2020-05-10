
class Solution_5404 {
public:
    vector<string> buildArray(vector<int>& target, int n) {
        std::vector<std::string> res;
        int len = target.size();
        int num = 1;
        for (int i = 0; i < len;) {
            if (target[i] == num) {
                res.push_back("Push");
                num++;
                i++;
            }
            else {
                while (num < target[i]) {
                    res.push_back("Push");
                    res.push_back("Pop");
                    num++;
                }
            }
        }
        return res;
    }
};