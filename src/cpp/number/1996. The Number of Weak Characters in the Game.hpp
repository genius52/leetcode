#include <vector>
#include <map>
using namespace std;

//如果存在一个其他角色的攻击和防御等级 都严格高于 该角色的攻击和防御等级，则认为该角色为 弱角色
class Solution_1996 {
public:
    int numberOfWeakCharacters(vector<vector<int>>& properties) {
        int len = properties.size();
        //攻击力从小到大排序。攻击力相同时，防御力从大到小排序
        std::sort(properties.begin(),properties.end(),[](const std::vector<int>& v1,const  std::vector<int>& v2)->bool{
            if(v1[0] < v2[0])
                return true;
            else if(v1[0] > v2[0])
                return false;
            else
                return v1[1] > v2[1];//避免[5,2],[5,4]这种顺序
        });
        int res = 0;
        int max_def = -2147483648;//记录最大的防御指数
        //if j > i,properties[j] 不可能比 properties[i]弱小
        for(int i = len - 1;i >= 0;i--){
            //判断当前是不是弱小角色
            if(properties[i][1] < max_def){
                res++;
            }
            max_def = max(max_def,properties[i][j]);
        }
        return res;
    }
};