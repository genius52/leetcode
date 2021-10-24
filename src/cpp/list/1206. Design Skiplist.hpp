#include <random>


const int MAX_LEVEL = 16;

class Skiplist {
    class ListNode{
    public:
        ListNode** next_;
        int val_ = 0;
        int height_ = MAX_LEVEL;
        ListNode(int val){
            val_ = val;
            next_ = new ListNode*[MAX_LEVEL];
            for(int i = 0;i < MAX_LEVEL;i++)
                next_[i] = nullptr;
        }
        ~ListNode(){
            delete[] next_;
        }
    };
    ListNode* head = nullptr;
public:
    Skiplist() {
        head = new ListNode(0);
    }

    bool search(int target) {
        auto visit = head;
        auto level = MAX_LEVEL - 1;
        while(visit){
            if(visit->val_ == target)
                return true;
            if(visit->val_ > target)
                return false;
            auto next = visit->next_[level];
            if(next != nullptr && target > next->val_)
                visit = next;
            else{
                if(next == nullptr || level == 0)
                    return false;
                level--;
            }
        }
        return false;
    }

    void add(int num) {
        ListNode* add = new ListNode(num);
        auto level = MAX_LEVEL - 1;
        auto visit = head;
        ListNode* pre[MAX_LEVEL];
        while(visit){
            if(visit->val_ >= num){
                break;
            }
            auto next = visit->next_[level];
            if(next != nullptr && next->val_ < num)
                visit = next;
            else{
                if(visit != nullptr)
                    pre[level] = visit;
//                if(next == nullptr || level == 0)
//                    break;
                level--;
                if(level < 0)
                    break;
            }
        }
        auto add_level = rand() % MAX_LEVEL;
        for(int i = 0;i < add_level ;i++){
            add[i].next_ = pre[i]->next_;
            pre[i]->next_[i] = add;
        }
    }

    bool erase(int num) {
        auto visit = head;
        ListNode* pre[MAX_LEVEL];
        auto level = MAX_LEVEL - 1;
        bool find = false;
        while(visit){
            if(visit->val_ == num){
                find = true;
                break;
            }
            if(visit->val_ > num)
                return false;
            auto next = visit->next_[level];
            if(next != nullptr && num > next->val_)
                visit = next;
            else{
                if(visit != nullptr)
                    pre[level] = visit;
                if(next == nullptr || level == 0)
                    return false;
                level--;
            }
        }
        if(find){
            for(int i = 0;i < MAX_LEVEL;i++){
                pre[i]->next_ = visit[i].next_;
            }
            delete visit;
            visit = nullptr;
        }
        return false;
    }
};