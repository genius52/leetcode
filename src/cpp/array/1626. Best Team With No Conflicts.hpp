#include <vector>
#include <queue>
using namespace std;

//Input: scores = [4,5,6,5], ages = [2,1,2,1]
//Output: 16
//Explanation: It is best to choose the last 3 players.
//Notice that you are allowed to choose multiple people of the same age.
class Solution_1626 {
public:
    int bestTeamScore(vector<int>& scores, vector<int>& ages) {
        int size = scores.size();
        std::vector<std::pair<int,int>> player(size);
        //std::vector<std::vector<int>> dp(size);
        for(int i = 0;i < size;i++){
            player[i] = {ages[i],scores[i]};
        }
        auto mycmp = [](const std::pair<int,int>& p1,const std::pair<int,int>& p2){
            if (p1.first < p2.first)
                return true;
            else if(p1.first > p2.first)
                return false;
            else{
                return p1.second < p2.second;
            }
        };
        std::sort(player.begin(),player.end(),mycmp);

        std::vector<int> dp(size);//dp[i] = choose player i,the max scores
        for (int i = 0;i < size;i++){
            auto age = player[i].first;
            auto score = player[i].second;
            int max = 0;
            for (int j = i - 1;j >= 0;j--){
                int cur_scores = 0;
                if (player[j].first < age && player[j].second <= score){
                    cur_scores += dp[j];
                }else if (player[j].first == age){
                    cur_scores += dp[j];
                }
                if(max < cur_scores){
                    max = cur_scores;
                }
            }
            dp[i] = max + score;
        }
        int res = 0;
        for (int i = 0;i < size;i++){
            if(res < dp[i]){
                res = dp[i];
            }
        }
        return res;
    }
};