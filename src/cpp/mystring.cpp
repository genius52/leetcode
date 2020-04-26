#include "mystring.h"
#include <set>
#include <vector>
#include <iostream>

bool isPalindrome(string s) {
    if (s.empty())
        return true;
    int len = s.length();
    int low = 0;
    int high = len - 1;
    while (low < high) {
        if (s[low] < 'A' || s[low] > 'z') {
            low++;
            continue;
        }
        if (s[high] < 'A' || s[high] > 'z') {
            high--;
            continue;
        }
        int diff = abs(s[low] - s[high]);
        if (diff != 0 && diff != 32)
            return false;
        low++;
        high--;
    }
    return true;
}


string addStrings(string num1, string num2) {
    string sum;
    int len1 = num1.length();
    if(len1 == 0)
        return num2;
    int len2 = num2.length();
    if(len2 == 0)
        return num1;
    string result;
    int min_len = len1>len2?len2:len1;
    int index = 0;
    int upgrade = 0;
    while(index < min_len){
        int v1 = num1.at(len1 - 1 - index) - '0';
        int v2 = num2.at(len2 - 1 - index) - '0';
        int sum = v1 + v2 + upgrade;
        if(sum >=10)
        {
            result = std::to_string(sum - 10) + result;
            upgrade = 1;
        }
        else{
            result = std::to_string(sum) + result;
            upgrade = 0;
        }
        index++;
    }
    while(index < len1)
    {
        int val = num1.at(len1 - 1 - index) - '0';
        int sum = val + upgrade;
        if(sum >= 10){
            result = std::to_string(sum - 10) + result;
            upgrade = 1;
        }
        else{
            result = std::to_string(sum) + result;
            upgrade = 0;
        }
        index++;
    }
    while(index < len2)
    {
        int val = num2.at(len2 - 1 - index) - '0';
        int sum = val + upgrade;
        if(sum >= 10){
            result = std::to_string(sum - 10) + result;
            upgrade = 1;
        }
        else{
            result = std::to_string(sum) + result;
            upgrade = 0;
        }
        index++;
    }
    if(len1 == len2 && upgrade == 1)
    {
        result = "1" + result;
    }
    return result;
}

bool repeatedSubstringPattern(string s) {
    int len = s.length();
    if(len <= 1)
        return true;
    for(int i = 0 ;i < len/2;i++){
        int sub_len = i + 1;
        string sub = s.substr(0,sub_len);
        bool match = true;
        for(int j = i + 1;j < len ;j = j + sub_len){
            if(s.find(sub,j) != j){
                match = false;
                break;
            }
        }
        if(match)
            return true;
    }
    return false;
}

bool backspaceCompare(string S, string T) {
    std::string s2,t2;
    for(auto it = S.begin(); it != S.end();it++){
        if(*it == '#'){
            if(!s2.empty()){
                s2.erase(s2.end()-1);
            }
        }
        else{
            s2.push_back(*it);
        }
    }
    for(auto it = T.begin(); it != T.end();it++){
        if(*it == '#'){
            if(!t2.empty()){
                t2.erase(s2.end()-1);
            }
        }
        else{
            t2.push_back(*it);
        }
    }
    return s2 == t2;
}

int str_cnt = 0;
std::set<std::string> str_map;
void perm(std::string& tiles,std::string& sub,int begin,int end){
    if(begin == end){
        if(str_map.find(sub) == str_map.end()){
            str_cnt++;
            str_map.insert(sub);
            std::cout<<sub<<" size="<<str_map.size()<<std::endl;
        }
        return;
    }
    for(int i = begin;i < end;i++){
        //sub.clear();
        swap(tiles[begin],tiles[i]);
        sub.push_back(tiles[begin]);
        perm(tiles,sub,begin + 1,end);
        sub.pop_back();
        //perm(tiles,sub,step+1,target_num);
        swap(tiles[begin],tiles[i]);
    }

}

int numTilePossibilities(std::string tiles) {
    int len = tiles.length();
    if(len <= 1)
        return len;
    for(int i = 1;i <= len;i++){
        std::string sub_s;
        perm(tiles,sub_s,0,i);
    }
    return str_cnt;
}
void swap(char& a,char& b){
    char tmp = a;
    a = b;
    b = tmp;
}

bool judgeCircle(string moves) {
    int len = moves.length();
    if(len == 0)
        return true;
    int vertical_cnt = 0;
    int horizon_cnt = 0;
    for(int i = 0;i < len;i++){
        char e = moves[i];
        if(e == 'U')
            vertical_cnt++;
        else if(e == 'D')
            vertical_cnt--;
        else if(e == 'R')
            horizon_cnt++;
        else if(e == 'L')
            horizon_cnt--;
    }
    return vertical_cnt == 0 && horizon_cnt == 0;
}

//Input: "IDID"
//Output: [0,4,1,3,2]
//Input: "DDI"
//Output: [3,2,0,1]
std::vector<int> diStringMatch(string S) {
    vector<int> res;
    for (int l = 0, h = S.size(), i = 0; i <= S.size(); ++i) {
        res.push_back(S[i] == 'I' ? l++ : h--);
    }
    return res;
}

