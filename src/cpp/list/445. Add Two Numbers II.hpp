/**
 * Definition for singly-linked list.
 * struct ListNode {
 *     int val;
 *     ListNode *next;
 *     ListNode() : val(0), next(nullptr) {}
 *     ListNode(int x) : val(x), next(nullptr) {}
 *     ListNode(int x, ListNode *next) : val(x), next(next) {}
 * };
 */
#include <deque>
#include "../define.h"

//Input: (7 -> 2 -> 4 -> 3) + (5 -> 6 -> 4)
//Output: 7 -> 8 -> 0 -> 7
class Solution_445 {
public:
    ListNode* addTwoNumbers(ListNode* l1, ListNode* l2) {
        if(l1 == nullptr)
            return l2;
        if(l2 == nullptr)
            return l1;
        std::deque<int> q1;
        std::deque<int> q2;
        while(l1){
            q1.push_back(l1->val);
            l1 = l1->next;
        }
        while(l2){
            q2.push_back(l2->val);
            l2 = l2->next;
        }
        bool upgrade = false;
        ListNode* head = nullptr;
        ListNode* last = nullptr;
        while(!q1.empty() && !q2.empty()){
            int e1 = q1.back();
            int e2 = q2.back();
            q1.pop_back();
            q2.pop_back();
            ListNode* visit = nullptr;
            int sum = e1 + e2 + upgrade;
            if(sum < 10){
                upgrade = false;
                visit = new ListNode(sum);
            }
            else{
                upgrade = true;
                visit = new ListNode(sum % 10);
            }
            if(last != nullptr)
                visit->next = last;
            last = visit;
            head = visit;
        }
        while(!q1.empty()){
            int e1 = q1.back();
            q1.pop_back();
            ListNode* visit = nullptr;
            int sum = e1 + upgrade;
            if(sum < 10){
                upgrade = false;
                visit = new ListNode(sum);
            }
            else{
                upgrade = true;
                visit = new ListNode(0);
            }
            visit->next = last;
            last = visit;
            head = visit;
        }
        while(!q2.empty()){
            int e2 = q2.back();
            q2.pop_back();
            int sum = e2 + upgrade;
            ListNode* visit = nullptr;
            if(sum < 10){
                upgrade = false;
                visit = new ListNode(sum);
            }
            else{
                upgrade = true;
                visit = new ListNode(0);
            }
            visit->next = last;
            last = visit;
            head = visit;
        }
        if(upgrade){
            ListNode* visit = new ListNode(1);
            visit->next = last;
            head = visit;
        }
        return head;
    }
};