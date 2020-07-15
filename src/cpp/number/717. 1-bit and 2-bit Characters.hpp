#include <vector>
using namespace std;

//Input:
//bits = [1, 1, 1, 0]
//Output: False
//Explanation:
//The only way to decode it is two-bit character and two-bit character. So the last character is NOT one-bit character.
class Solution_717 {
public:
    bool dfs_isOneBitCharacter(vector<int>& bits,int len,int pos,bool start){
        if(pos == len - 1){
            if(start)
                return true;
            else
                return false;
        }
        if(start){
            if(bits[pos] == 0){
                return dfs_isOneBitCharacter(bits,len,pos + 1,true);
            }
            else{
                return dfs_isOneBitCharacter(bits,len,pos + 1,false);
            }
        }
        else{
            return dfs_isOneBitCharacter(bits,len,pos + 1,true);
        }
    }

    bool isOneBitCharacter(vector<int>& bits) {
        int len = bits.size();
        if(len == 0)
            return false;
        return dfs_isOneBitCharacter(bits,len,0,true);
    }
};