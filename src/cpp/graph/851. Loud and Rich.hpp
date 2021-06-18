#include <vector>
#include <unordered_map>
using namespace std;

//输入：richer = [[1,0],[2,1],[3,1],[3,7],[4,3],[5,3],[6,3]], quiet = [3,2,5,4,6,1,7,0]
//输出：[5,5,2,5,4,5,6,7]
//解释：
//answer[0] = 5，
//person 5 比 person 3 有更多的钱，person 3 比 person 1 有更多的钱，person 1 比 person 0 有更多的钱。
//唯一较为安静（有较低的安静值 quiet[x]）的人是 person 7，
//但是目前还不清楚他是否比 person 0 更有钱。
//
//answer[7] = 7，
//在所有拥有的钱肯定不少于 person 7 的人中(这可能包括 person 3，4，5，6 以及 7)，
//最安静(有较低安静值 quiet[x])的人是 person 7。
//
//其他的答案也可以用类似的推理来解释。
class Solution_851 {
public:
    std::pair<int,int> dfs_loudAndRich(std::vector<std::vector<int>>& richer_than_me,vector<int>& quiet,int cur,std::unordered_map<int,std::pair<int,int>>& memo){
        int min_quiet = quiet[cur];
        int min_quiet_index = cur;
        int len = richer_than_me[cur].size();
        if(len == 0)
            return {min_quiet_index,min_quiet};
        if(memo.find(cur) != memo.end())
            return memo[cur];
        for(int i = 0;i < len;i++){
            int next_richer = richer_than_me[cur][i];
            auto ret = dfs_loudAndRich(richer_than_me,quiet,next_richer,memo);
            if(ret.second < min_quiet) {
                min_quiet = ret.second;
                min_quiet_index = ret.first;
            }
        }
        memo[cur].first = min_quiet_index;
        memo[cur].second = min_quiet;
        return memo[cur];
    }

    vector<int> loudAndRich(vector<vector<int>>& richer, vector<int>& quiet) {
        int len = quiet.size();
        std::vector<std::vector<int>> richer_than_me(len);
        for(auto r : richer){
            richer_than_me[r[1]].push_back(r[0]);
        }
        std::unordered_map<int,std::pair<int,int>> memo;
        std::vector<int> res(len);
        for(int i = 0;i < len;i++){
            auto p = dfs_loudAndRich(richer_than_me,quiet,i,memo);
            res[i] = p.first;
        }
        return res;
    }
};