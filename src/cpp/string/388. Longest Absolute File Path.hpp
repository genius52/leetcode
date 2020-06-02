#include <string>
#include <deque>
#include <map>
using namespace std;

//TO DO
class Solution_388 {
public:
    int lengthLongestPath(string input) {
        std::map<int,int> record;//key:level (tab count). value: max length of this level
        int len = input.length();
        int max = 0;
        int tab_cnt = 0;
        int start = 0;
        int end = 0;
        while(end < len){
            if(input[end] == '\\'){
                if(input[end + 1] == 'n'){

                }
            }else{
                end++;
            }
        }
        return res;
    }
};