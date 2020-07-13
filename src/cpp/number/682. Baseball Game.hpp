#include <vector>
#include <deque>
#include <sstream>
using namespace std;

class Solution_682 {
public:
    int calPoints(vector<string>& ops) {
        int len = ops.size();

        int res;//number which will contain the result
        stringstream convert(ops[0]); // stringstream used for the conversion initialized with the contents of Text
        convert >> res;
        std::deque<int> stack;
        stack.push_back(res);
        for(int i = 1;i < len;i++){
            if(ops[i] == "+"){
                auto last_score = stack.back();
                stack.pop_back();
                auto last_score2 = stack.back();
                stack.push_back(last_score);
                auto current_score = last_score + last_score2;
                res += current_score;
                stack.push_back(current_score);
            }else if(ops[i] == "D"){
                auto scores = stack.back();
                res += scores * 2;
                stack.push_back(scores * 2);
            }else if(ops[i] == "C"){
                auto score = stack.back();
                res -= score;
                stack.pop_back();
            }else{
                stringstream convert(ops[i]);
                int num;
                convert >> num;
                res += num;
                stack.push_back(num);
            };
        }
        return res;
    }
};