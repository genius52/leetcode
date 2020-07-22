#include <vector>
#include <string>
#include <map>
#include <unordered_map>
using namespace std;

//Input: ["abcd","cdab","cbad","xyzz","zzxy","zzyx"]
//Output: 3
//Explanation:
//One group is ["abcd", "cdab", "cbad"], since they are all pairwise special equivalent, and none of the other strings are all pairwise special equivalent to these.
//
//The other two groups are ["xyzz", "zzxy"] and ["zzyx"].  Note that in particular, "zzxy" is not special equivalent to "zzyx".
class Solution_893 {
public:
    int numSpecialEquivGroups(vector<string>& A) {
        int len = A.size();
        int word_len = A[0].size();
        std::unordered_map<string,int> record;
        for(int i = 0;i < len;i++){
            std::map<char,int> even_map;//偶数
            std::map<char,int> odd_map;//奇数
//            std::vector<int> even_map;
//            std::vector<int> odd_map;
//            even_map.resize(26);
//            odd_map.resize(26);
            for(int j = 0;j < word_len;j++){
                if(j % 2 == 0)
                    even_map[A[i][j]]++;
                else{
                    odd_map[A[i][j]]++;
                }
            }
            std::string word;
            for(auto it = even_map.begin();it != even_map.end();it++){
                std::string w((*it).second,(*it).first);
                word += w;
            }
            for(auto it = odd_map.begin();it != odd_map.end();it++){
                std::string w((*it).second,(*it).first);
                word += w;
            }
            record[word]++;
        }
        return record.size();
    }
};