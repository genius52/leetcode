struct Listnode{
    int val = 0;
    Listnode* next = nullptr;
};

class MyLinkedList {
    Listnode* head = nullptr;
    Listnode* tail = nullptr;
    int size = 0;
public:
    /** Initialize your data structure here. */
    MyLinkedList() {

    }

    /** Get the value of the index-th node in the linked list. If the index is invalid, return -1. */
    int get(int index) {
        if(index >= size)
            return -1;
        auto visit = head;
        while(index > 0){
            visit = visit->next;
            index--;
        }
        return visit->val;
    }

    /** Add a node of value val before the first element of the linked list. After the insertion, the new node will be the first node of the linked list. */
    void addAtHead(int val) {
        auto new_node = new Listnode();
        new_node->val = val;
        new_node->next = head;
        head = new_node;
        size++;
        if(size == 1){
            tail = head;
        }
    }

    /** Append a node of value val to the last element of the linked list. */
    void addAtTail(int val) {
        auto new_node = new Listnode();
        new_node->val = val;
        if(tail == nullptr){
            head = new_node;
            tail = new_node;
        }else{
            tail->next = new_node;
            tail = new_node;
        }
        size++;
    }

    /** Add a node of value val before the index-th node in the linked list. If index equals to the length of linked list, the node will be appended to the end of linked list. If index is greater than the length, the node will not be inserted. */
    void addAtIndex(int index, int val) {
        if(index > size)
            return;
        if(index == size){//add tail
            auto new_node = new Listnode();
            new_node->val = val;
            if(size == 0){//empty list
                head = new_node;
            }else{
                tail->next = new_node;
            }
            tail = new_node;
        }else if(index < size){
            if(index == 0){
                auto new_node = new Listnode();
                new_node->val = val;
                new_node->next = head;
                head = new_node;
            }else{
                auto visit = head;
                Listnode* pre = nullptr;
                while(index > 0){
                    pre = visit;
                    visit = visit->next;
                    index--;
                }
                auto new_node = new Listnode();
                new_node->val = val;
                new_node->next = visit;
                pre->next = new_node;
            }
        }
        size++;
    }

    /** Delete the index-th node in the linked list, if the index is valid. */
    void deleteAtIndex(int index) {
        if(index >= size)
            return;
        int old_index = index;
        if(index == 0){
            auto del = head;
            head = head->next;
            delete del;
            del = nullptr;
        }else{
            auto visit = head;
            Listnode* pre = nullptr;
            while(index > 0){
                pre = visit;
                visit = visit->next;
                index--;
            }
            pre->next = visit->next;
            delete visit;
            visit = nullptr;
            if(old_index == size - 1){
                tail = pre;
            }
        }
        size--;
    }
};