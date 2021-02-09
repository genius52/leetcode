#include <vector>
using namespace std;

class Solution_688 {
public:
    double dp_knightProbability(int N, int K, int r, int c,std::vector<std::vector<int>> dir,std::vector<std::vector<std::vector<double>>>& memo){
        if (r < 0 || r >= N || c < 0 || c >= N){
            return 0.0;
        }
        if (K == 0){
            return 1;
        }
        if(memo[r][c][K] != 0){
            return memo[r][c][K];
        }
        double res = 0.0;
        for(int i = 0;i < 8;i++){
            res += 0.125 * dp_knightProbability(N,K - 1,r + dir[i][0],c + dir[i][1],dir,memo);
        }
        memo[r][c][K] = res;
        return res;
    }

    double knightProbability(int N, int K, int r, int c) {
        std::vector<std::vector<int>> dir{{-2,1},{-1,2},{1,2},{2,1},{2,-1},{1,-2},{-1,-2},{-2,-1}};
        std::vector<std::vector<std::vector<double>>> memo(N,std::vector<std::vector<double>>(N,std::vector<double>(K + 1)));
        return dp_knightProbability(N,K,r,c,dir,memo);
    }
};