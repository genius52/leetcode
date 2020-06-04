
//Input:
//n = 9,
//1 2 3 4 5 6 7 8 9
//2 4 6 8
//2 6
//6
//
//Output:
//6
class Solution_390 {
public:
    int lastRemaining(int n) {
        if(n <= 0)
            return 0;
        int step = 1;
        bool left_right = true;
        int head = 1;
        int tail = n;
        while(head != tail){
            if(left_right){
                if(n % 2 == 1){
                    int tmp = tail;
                    tail = head + step;
                    head = tmp - step;
                }
                else{
                    int tmp = tail;
                    tail = head + step;
                    head = tmp;
                }
            }else {
                if (n % 2 == 1) {
                    int tmp = tail;
                    tail = head - step;
                    head = tmp + step;
                } else {
                    int tmp = tail;
                    tail = head - step;
                    head = tmp;
                }
            }
            n = n / 2;
            left_right = !left_right;
            step *= 2;
        }
        return head;
    }
};