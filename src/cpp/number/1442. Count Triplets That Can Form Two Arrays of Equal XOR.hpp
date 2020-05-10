
class Solution_5405 {
public:
    int countTriplets(vector<int>& arr) {
        int cnt = 0;
        int len = arr.size();
        bool** visited = new bool* [len];
        int** dp = new int* [len];
        for (int i = 0; i < len; i++) {
            dp[i] = new int[len];
            visited[i] = new bool[len];
            for (int j = 0; j < len; j++) {
                visited[i][j] = false;
            }
        }

        for (int i = 0; i < len - 1; i++) {
            for (int j = i + 1; j < len; j++) {
                for (int k = j; k < len; k++) {
                    int a = 0;
                    if (visited[i][j - 1]) {
                        a = dp[i][j - 1];
                    }
                    else {
                        for (int m = i; m < j; m++) {
                            a ^= arr[m];
                        }
                        dp[i][j - 1] = a;
                        visited[i][j - 1] = true;
                    }
                    int b = 0;
                    if (visited[j][k]) {
                        b = dp[j][k];
                    }
                    else {
                        for (int n = j; n <= k; n++) {
                            b ^= arr[n];
                        }
                        dp[j][k] = b;
                        visited[j][k] = true;
                    }
                    if (a == b)
                        cnt++;
                }
            }
        }
        return cnt;
    }
};