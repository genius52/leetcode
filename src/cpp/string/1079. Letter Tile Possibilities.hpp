#include <string>
#include <set>
using namespace std;

class Solution_1079{
public:
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

};
