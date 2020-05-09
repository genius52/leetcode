//
// Created by 陶诚程 on 2019/2/21.
//
#include "array.h"
#include <map>
#include <vector>

int remove_duplicate_number(std::vector<int>& nums)
{
    auto length = nums.size();
    if(length == 1)
        return 1;
    int first = 0,second = 0;
    while ( second < length) {
        if(nums[first] != nums[second])
        {
            nums[++first] = nums[second++];
        }
        else
        {
            second++;
        }
    }
    return first + 1;
    //nums.erase(nums[first],nums[second]);
}

int fib(int n)
{
    if(n == 1 || n == 2)
        return 1;
    return fib(n-1) + fib(n-2);
}

int fibex(int n)
{
    if(n == 1 || n == 2)
    {
        return 1;
    }
    std::map<int,int> m;
    m.insert(std::make_pair(1,1));
    m.insert(std::make_pair(2,1));
    int sum = 0;
    for (int i = 1; i <= n; ++i) {
        if(m.find(i) != m.end())
        {
            sum += m[i];
        }
        else
        {
            sum = m[i-1] + m[i-2];
            m[i] = sum;
        }
    }
    return sum;
}

int fibex2(int n)
{
    if(n == 1 || n == 2)
    {
        return 1;
    }
    int last1 = 1;
    int last2 = 1;
    int sum = 0;
    for(int i = 1;i < n;i++)
    {
        sum = last1 + last2;
        last1 = last2;
        last2 = sum;
    }
    return sum;
}
/*
[3, 0, 8, 4],
[2, 4, 5, 7],
[9, 2, 6, 3],
[0, 3, 1, 0]

gridNew = [
[8, 4, 8, 7],
[7, 4, 7, 7],
[9, 4, 8, 7],
[3, 3, 3, 3] ]
*/
#include <math.h>
#include <algorithm>
int maxIncreaseKeepingSkyline(std::vector<std::vector<int>>& grid) {
    int sum = 0;
    if(grid.size() == 0)
        return sum;

    std::vector<int> max_height_rows;
    for(int i = 0;i < grid.size();i++){
        int max_height_current_row = 0;
        for (int j = 0; j < grid[0].size(); ++j) {
            if(grid[i][j] > max_height_current_row)
                max_height_current_row = grid[i][j];
        }
        max_height_rows.push_back(max_height_current_row);
    }

    std::vector<int> max_height_columns;
    for(int i = 0;i < grid[0].size();i++){
        int max_height_current_column = 0;
        for (int j = 0; j < grid.size(); ++j) {
            if(grid[j][i] > max_height_current_column)
                max_height_current_column = grid[j][i];
        }
        max_height_columns.push_back(max_height_current_column);
    }

    for(int i = 0;i < grid.size();i++){
        for (int j = 0; j < grid[0].size(); ++j) {
            auto available_height = max_height_rows[i]<max_height_columns[j]?max_height_rows[i]:max_height_columns[j];
            if(available_height > grid[i][j])
                sum += abs(available_height-grid[i][j]);
        }
    }
    return sum;
}

struct point{
    int x;
    int y;
    int dis = 0;
};

std::vector<std::vector<int>> allCellsDistOrder(int R, int C, int r0, int c0) {
    std::vector<std::vector<int>> res;
    if(r0 > R || c0 > C)
        return res;
    std::vector<point> vp;
    for(int i = 0;i < R;i++){
        for(int j = 0;j < C;j++){
            point p;
            p.x = j;
            p.y = i;
            p.dis = abs(i - r0) + abs(j - c0);
            vp.push_back(p);
        }
    }
    std::sort(vp.begin(),vp.end(),[](const point& p1,const point& p2)->bool{
       return p1.dis > p2.dis;
    });
    for(auto it = vp.begin();it != vp.end();it++){
        std::vector<int> v;
        v.push_back(it->x);
        v.push_back(it->y);
        res.push_back(v);
    }
    return res;
}

#include <set>
int lastStoneWeight(std::vector<int>& stones) {
    std::multiset<int,std::greater<int>> s;
    for(auto it = stones.begin();it != stones.end();it++){
        s.insert(*it);
    }
    while(s.size()>0){
        auto first = s.begin();
        auto biggest = *first;
        s.erase(first++);
        if(first == s.end())
            return biggest;
        else{
            if(*first != biggest)
                s.insert(biggest - *first);
            s.erase(first);
        }
    }
    return 0;
}

bool cmp(std::string& s1,std::string& s2,std::string& pattern){
    int len1 = s1.length();
    int len2 = s2.length();
    int min_len = len1 > len2?len2:len1;
    for(int i = 0;i < min_len;i++){
        auto c1 = s1[i];
        auto c2 = s2[i];
        auto pos1 = pattern.find(c1);
        auto pos2 = pattern.find(c2);
        if(pos1 > pos2)
            return false;
    }
    if(len1 > len2)
        return false;
    return true;
}

bool isAlienSorted(std::vector<std::string>& words, std::string order) {
    int len = words.size();
    if(len <= 1)
        return true;
    for(int i = 0;i < len - 1;i++){
        auto first = words[i];
        auto second = words[i+1];
        if(!cmp(first,second,order))
            return false;
    }
    return true;
}

bool canThreePartsEqualSum(std::vector<int>& A) {

    return true;
}

class Solution {
public:
    std::vector<int> relativeSortArray(std::vector<int>& A, std::vector<int>& B) {

        int key[1001], idx = 0;
        for (int i = 0; i < 1001; i++) key[i] = 1000+i;
        for (int b : B) key[b] = idx++;
        sort(A.begin(), A.end(), [&](int c, int d){
            return key[c] < key[d];
        });
        return A;
    }
};


bool uniqueOccurrences(std::vector<int>& arr) {
    int count[2000]{};
    for(auto it = arr.begin();it != arr.end(); it++){
        if(count[*it + 1000] == 1)
            return false;
        else
            count[*it + 1000] = 1;
    }
    return true;
}

//241
//Input: [[1,1],2,[1,1]]
//Output: [1,1,2,1,1]
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


namespace s1366{
    std::string Solution::rankTeams(std::vector<std::string>& votes) {
        auto l = votes.size();
        if (l == 0){
            return "";
        }
        auto nums = votes[0].size();
        std::map<char,int> record;
        for (int i = 0;i < l;i++){
            for (int j = 0;j < nums;j++){
                if (record.find(votes[i][j]) == record.end()){
                    record[votes[i][j]] = nums - j;
                }
                else{
                    record[votes[i][j]] += nums - j;
                }
            }
        }
        std::vector<std::pair<char,int> > vec;
        for (auto it :record){
            vec.push_back(it);
        }
        std::string res;
        std::sort(vec.begin(),vec.end(),[](const std::pair<char,int>& a, const std::pair<char,int>& b)->bool{
            return a.second > b.second;
        });
        for (auto it = vec.begin();it != vec.end();it++){
            res += it->first;
        }
        return res;
    }
}