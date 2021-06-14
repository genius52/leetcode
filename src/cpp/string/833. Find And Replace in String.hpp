#include <vector>
#include <string>
using namespace std;

class Solution_833 {
public:
    string findReplaceString(string s, vector<int>& indexes, vector<string>& sources, vector<string>& targets) {
        int s_len = s.length();
        int l = indexes.size();
        std::vector<std::pair<int,int>> position_index(l);//first: position in string s, second: position in indexes
        int k = 0;
        for(int i = 0;i < l;i++){
            position_index[k++] = {indexes[i],i};
        }
        std::sort(position_index.rbegin(),position_index.rend());
        std::string res = s;
        int pos = s_len;
        for(int i = 0;i < l;i++){
            //int idx = position_index[i];
            int sub_len = sources[position_index[i].second].length();
            std::string origin_sub = s.substr(position_index[i].first,sub_len);
            if(origin_sub == sources[position_index[i].second]){
                res = res.substr(0,position_index[i].first) + targets[position_index[i].second] + res.substr(position_index[i].first + sub_len);
            }
        }
        return res;
    }
};