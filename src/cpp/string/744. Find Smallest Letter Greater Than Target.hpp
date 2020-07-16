

class Solution_744 {
public:
    char nextGreatestLetter(vector<char>& letters, char target) {
        int len = letters.size();
        int low = 0;
        int high = len - 1;
        if(target >= letters[high])
            return letters[0];
        target++;
        while(low < high){
            int mid = low + (high - low)/2;
            if(letters[mid] == target){
                return letters[mid];
            }
            else if(letters[mid] < target){
                low = mid + 1;
            }else if(letters[mid] > target){
                high = mid;
            }
        }
        return letters[high];
    }
};