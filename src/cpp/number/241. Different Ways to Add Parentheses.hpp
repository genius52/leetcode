#include <vector>
#include <string>
#include <deque>
using namespace std;

class Solution_241 {
public:
    std::vector<int> dp_diffWaysToCompute(string input){
        std::vector<int> record;
        int l = input.length();
        if(input.empty()){
            return record;
        }

        for(int i = 0;i <= l;i++){
            std::vector<int> pre_record;
            std::vector<int> post_record;
            if(input[i] == '+' || input[i] == '-' || input[i] == '*'){
                pre_record = dp_diffWaysToCompute(input.substr(0,i));
                post_record = dp_diffWaysToCompute(input.substr(i+1));
            }
            for(auto pre_num : pre_record){
                for(auto post_num : post_record){
                    if(input[i] == '+'){
                        record.push_back(pre_num + post_num);
                    }
                    else if(input[i] == '-'){
                        record.push_back(pre_num - post_num);
                    }
                    else if(input[i] == '*'){
                        record.push_back(pre_num * post_num);
                    }
                }
            }
        }

        if(record.empty()){
            int val = atoi(input.c_str());
            record.push_back(val);
            return record;
        }
        return record;
    }

    vector<int> diffWaysToCompute(string input) {
        int l = input.size();
        return dp_diffWaysToCompute(input);
    }
};