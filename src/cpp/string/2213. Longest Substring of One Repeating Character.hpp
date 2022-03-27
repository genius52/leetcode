#include <vector>
#include <string>
#include <map>
#include <set>
using namespace std;

//输入：s = "babacc", queryCharacters = "bcb", queryIndices = [1,3,3]
//输出：[3,3,4]
//解释：
//- 第 1 次查询更新后 s = "bbbacc" 。由单个字符重复组成的最长子字符串是 "bbb" ，长度为 3 。
//- 第 2 次查询更新后 s = "bbbccc" 。由单个字符重复组成的最长子字符串是 "bbb" 或 "ccc"，长度为 3 。
//- 第 3 次查询更新后 s = "bbbbcc" 。由单个字符重复组成的最长子字符串是 "bbbb" ，长度为 4 。
//因此，返回 [3,3,4]

//"mef"
//"dmdaucfak"
//[2,0,0,1,2,0,2,0,2]
class Solution_2213 {
public:
    vector<int> longestRepeating(string s, string queryCharacters, vector<int>& queryIndices) {
        std::map<int,int> char_position[26];//char_position[0]
        std::vector<std::multiset<int>>  max_len_record(26);
        int str_len = s.size();
        int left = 0;
        int right = 0;
        while (left < str_len){
            while (right < str_len && s[right] == s[left]){
                right++;
            }
            char_position[s[left] - 'a'][left] = right - left;
            max_len_record[s[left] - 'a'].insert(right - left);
            left = right;
        }
        //int pre_max_len = 0;
        int l = queryIndices.size();
        std::vector<int> res(l);
        for(int i = 0;i < l;i++){
            //查看char_position[queryCharacters[i] - 'a']前面和后面的位置，判断最大长度是否增长
            //std::cout<<"before1 : "<<i << std::endl;
            int cur_pos = queryIndices[i];
            char pre_char = s[queryIndices[i]];
            s[cur_pos] = queryCharacters[i];
            std::cout<<s<< std::endl;
            //查找之前字符小于当前位置的
            //auto find_before_prechar = std::upper_bound( char_position[pre_char - 'a'].begin() , char_position[pre_char - 'a'].end() , cur_pos, std::greater<int>());
            auto find_before_prechar = char_position[pre_char - 'a'].lower_bound(cur_pos);
            if(find_before_prechar == char_position[pre_char - 'a'].begin()){
                find_before_prechar = char_position[pre_char - 'a'].end();
            }else{
                find_before_prechar--;
            }
            //std::cout<<"before2 : "<<i << std::endl;
            if(find_before_prechar != char_position[pre_char - 'a'].end()){//删除作为结尾的字符
                if(find_before_prechar->first + find_before_prechar->second == cur_pos - 1){//字符是最后一个
                    max_len_record[pre_char - 'a'].erase(find_before_prechar->second);
                    find_before_prechar->second--;
                    max_len_record[pre_char - 'a'].insert(find_before_prechar->second);
                }else if(find_before_prechar->first + find_before_prechar->second > cur_pos){//字符在中间
                    max_len_record[pre_char - 'a'].erase(find_before_prechar->second);
                    char_position[pre_char - 'a'][cur_pos + 1] = find_before_prechar->first + find_before_prechar->second - 1 - cur_pos;
                    find_before_prechar->second = cur_pos - find_before_prechar->first;
                    max_len_record[pre_char - 'a'].insert(find_before_prechar->second);
                    max_len_record[pre_char - 'a'].insert(char_position[pre_char - 'a'][cur_pos + 1]);
                }
            }
            //std::cout<<"start : "<<i << std::endl;
            auto find_start_prechar = char_position[pre_char - 'a'].find(cur_pos);//删除作为开头的字符
            if(find_start_prechar != char_position[pre_char - 'a'].end()){
                max_len_record[pre_char - 'a'].erase(find_start_prechar->second);
                if(find_start_prechar->second > 1){
                    char_position[pre_char - 'a'][cur_pos + 1] = find_start_prechar->second - 1;
                    max_len_record[pre_char - 'a'].insert(find_start_prechar->second - 1);
                }
                char_position[pre_char - 'a'].erase(find_start_prechar);
            }
            char replace_char = queryCharacters[i];
            int next_pos = queryIndices[i] + 1;
            //auto find_last = std::lower_bound( char_position[replace_char - 'a'].begin() , char_position[replace_char - 'a'].end() , cur_pos, std::greater<int>());
            auto find_last = char_position[replace_char - 'a'].lower_bound(cur_pos);
            if(find_last == char_position[replace_char - 'a'].begin()){
                find_last = char_position[replace_char - 'a'].end();
            }else{
                find_last--;
            }
            auto find_next = char_position[replace_char - 'a'].find(next_pos); //找后面的
            if(find_last != char_position[replace_char - 'a'].end() &&
                    find_last->first + find_last->second == cur_pos){
                if(find_next != char_position[replace_char - 'a'].end()){//合并前后
                    int cur_len = find_next->first + find_next->second - find_last->first;
                    find_last->second = cur_len;
                    max_len_record[replace_char - 'a'].erase(find_last->second);
                    max_len_record[replace_char - 'a'].erase(find_next->second);
                    max_len_record[replace_char - 'a'].insert(cur_len);
                    char_position[replace_char - 'a'].erase(find_next);
                }else{//合并前面的
                    max_len_record[replace_char - 'a'].erase(find_last->second);
                    find_last->second++;
                    max_len_record[replace_char - 'a'].insert(find_last->second);
                }
            }else if(find_next != char_position[replace_char - 'a'].end()){//合并后面的
                max_len_record[replace_char - 'a'].erase(find_next->second);
                max_len_record[replace_char - 'a'].insert(find_next->second + 1);
                char_position[replace_char - 'a'][cur_pos] = find_next->second + 1;
                char_position[replace_char - 'a'].erase(find_next);
            }else{//插入单个
                max_len_record[replace_char - 'a'].insert(1);
                char_position[replace_char - 'a'][cur_pos] = 1;
            }
            //std::cout<<"end : "<<i << std::endl;
            int cur_max_len = 0;
            for (int j = 0; j < 26; ++j) {
                if(!max_len_record[j].empty())
                    cur_max_len = max(cur_max_len,*max_len_record[j].rbegin());
            }
            res[i] = cur_max_len;
        }
        return res;
    }
};