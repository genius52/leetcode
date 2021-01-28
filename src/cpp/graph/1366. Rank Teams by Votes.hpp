#include <string>
#include <vector>
#include <map>
#include <unordered_set>
using namespace std;

//Input: votes = ["BCA","CAB","CBA","ABC","ACB","BAC"]
//Output: "ABC"
//Explanation:
//Team A was ranked first by 2 voters, second by 2 voters and third by 2 voters.
//Team B was ranked first by 2 voters, second by 2 voters and third by 2 voters.
//Team C was ranked first by 2 voters, second by 2 voters and third by 2 voters.
//There is a tie and we rank teams ascending by their IDs.
class Solution_1366 {
public:
    string rankTeams(vector<string> &votes) {
        int len = votes.size();
        if(len == 0){
            return "";
        }
        if(len == 1){
            return votes[0];
        }
        int count = votes[0].size();
        if(count == 1){
            return votes[0];
        }
        std::map<int,std::vector<int>> record;
        for(int pos = 0;pos < count;pos++){
            for(int i = 0;i < len;i++){
                auto player = votes[i][pos];
                if(record.find(player) == record.end()){
                    std::vector<int> v(count);
                    record[player] = v;
                }
                record[player][pos]++;
            }
        }
        std::vector<int> scores;
        for(auto c : votes[0]){
            scores.push_back(c);
        }
        sort(scores.begin(),scores.end(),[&](int a,int b)->bool{
            for(int i = 0; i < count; i++){
                if(record[a][i] != record[b][i]){
                    return record[a][i] > record[b][i];
                }
            }
            return a < b;
        });
        std::string res;
        for(auto it = scores.begin();it != scores.end();it++){
            std::string s;
            s += *it;
            res += s;
        }
        return res;
    }
};