#include <vector>
using namespace std;

class MyCircularQueue {
private:
    std::vector<int> data_;
    int head = 0;
    int tail = 0;
    int size = 0;
    int max_size = 0;
public:
    MyCircularQueue(int k) {
        data_.resize(k);
        max_size = k;
    }
    bool enQueue(int val) {
        if (isFull()) return false;
        size++;
        data_[tail] = val;
        tail = (tail + 1) % max_size;
        return true;
    }
    bool deQueue() {
        if (isEmpty()) return false;
        size--;
        head = (head + 1) % max_size;
        return true;
    }
    int Front() {
        return isEmpty() ? -1 : data_[head];
    }
    int Rear() {
        if(isEmpty())
            return -1;
        return data_[(tail - 1 + max_size) % max_size];
    }
    bool isEmpty() {
        return size == 0;
    }
    bool isFull() {
        return size == max_size;
    }
};

//list Solution
//struct Node {
//public:
//    int val;
//    Node* next;
//    Node(int v, Node* n=nullptr) {
//        val = v;
//        next = n;
//    }
//};
//
//class MyCircularQueue {
//public:
//    MyCircularQueue(int k) {
//        maxSize = k;
//    }
//    bool enQueue(int val) {
//        if (isFull()) return false;
//        Node* newNode = new Node(val);
//        if (isEmpty()) head = newNode, tail = newNode;
//        else tail->next = newNode, tail = tail->next;
//        size++;
//        return true;
//    }
//    bool deQueue() {
//        if (isEmpty()) return false;
//        Node* rem = head;
//        head = head->next;
//        delete rem;
//        size--;
//        return true;
//    }
//    int Front() {
//        return isEmpty() ? -1 : head->val;
//    }
//    int Rear() {
//        return isEmpty() ? -1 : tail->val;
//    }
//    bool isEmpty() {
//        return size == 0;
//    }
//    bool isFull() {
//        return size == maxSize;
//    }
//private:
//    int maxSize, size = 0;
//    Node *head = new Node(0), *tail = new Node(0);
//};