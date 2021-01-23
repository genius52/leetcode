#include <vector>
using namespace std;

class NestedInteger {
public:
    // Return true if this NestedInteger holds a single integer, rather than a nested list.
    bool isInteger() const;

    // Return the single integer that this NestedInteger holds, if it holds a single integer
    // The result is undefined if this NestedInteger holds a nested list
    int getInteger() const;

    // Return the nested list that this NestedInteger holds, if it holds a nested list
    // The result is undefined if this NestedInteger holds a single integer
    const std::vector<NestedInteger> &getList() const;
};

class NestedIterator {
public:
    NestedIterator(std::vector<NestedInteger> &nestedList) {
        dfs(nestedList,data);
        current = data.begin();
    }

    int next() {
        if(current == data.end())
            return 0;
        auto res =  *current;
        current++;
        return res;
    }

    bool hasNext() {
        if(current == data.end())
            return false;
        return true;
    }

private:
    void dfs(std::vector<NestedInteger> &nestedList,std::vector<int>& data){
        for (int i = 0;i < nestedList.size();i++){
            if (nestedList[i].isInteger()){
                data.push_back(nestedList[i].getInteger());
            }
            else{
                std::vector<NestedInteger> childlist = nestedList[i].getList();
                dfs(childlist,data);
            }
        }
    }
    std::vector<int> data;
    std::vector<int>::iterator current;
};

/**
 * Your NestedIterator object will be instantiated and called as such:
 * NestedIterator i(nestedList);
 * while (i.hasNext()) cout << i.next();
 */