#include <vector>
#include <unordered_map>
#include <unordered_set>
using namespace std;

//{{1,5},{3,4},{3,5},{4,5},{2,4}};
class Solution_684 {
public:
    int find_group(std::vector<int> group,int i){
        if(group[i] != i){
            group[i] = find_group(group,group[i]);
        }
        return group[i];
    }

    vector<int> findRedundantConnection(vector<vector<int>>& edges){
        int len = edges.size();
        std::vector<int> group(len + 1);
        for(int i = 1;i <= len;i++){
            group[i] = i;
        }
        for(int i = 0;i < len;i++){
            //查找该节点属于哪个组
            int group1 = find_group(group,edges[i][0]);
            int group2 = find_group(group,edges[i][1]);
            if(group1 == group2)
                return edges[i];
            else
                group[group1] = group2;
        }
    }
//    vector<int> findRedundantConnection(vector<vector<int>>& edges) {
//        int len = edges.size();
//        std::vector<int> group(len + 1);
//        for(int i = 1;i <= len;i++){
//            group[i] = i;
//        }
//        std::vector<int> res;
//        int group_num = len + 1;
//        for(int i = 0;i < len;i++){
//            if(group[edges[i][0]] == edges[i][0] && group[edges[i][1]] == edges[i][1]){//two points alone
//                group[edges[i][0]] = group_num;
//                group[edges[i][1]] = group_num;
//                group_num++;
//            }else{
//                if(group[edges[i][0]] == edges[i][0]){//one point alone
//                    group[edges[i][0]] = group[edges[i][1]];
//                }else if(group[edges[i][1]] == edges[i][1]){//one point alone
//                    group[edges[i][1]] = group[edges[i][0]];
//                }else if(group[edges[i][0]] == group[edges[i][1]]){//two points in one group
//                    res = edges[i];
//                    break;
//                }else{//each point in different group
//                    int n = group[edges[i][0]];
//                    for(int j = 1;j <= len;j++){
//                        if(group[j] == n){
//                            group[j] = group[edges[i][1]];
//                        }
//                    }
//                }
//            }
//        }
//        return res;
//    }
};