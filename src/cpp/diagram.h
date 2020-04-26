//
// Created by 陶诚程 on 2019/4/16.
//

#ifndef LEET_DIAGRAM_H
#define LEET_DIAGRAM_H

#endif //LEET_DIAGRAM_H
#include <vector>
#include <map>
#include <queue>

void bfs();
void dfs();

int bfs_short_route(int array[][4],int start,int end);

int my_short_distance(int a[][10],int start,int end);

//612
class Solution_612 {
public:
    int leastInterval(std::vector<char>& tasks, int n) {
        std::priority_queue<int>  q;
        std::map<char,int> record;

        for(auto t : tasks){
            if (record.find(t) == record.end()){
                record[t] = 1;
            }else{
                record[t]++;
            }
        }
        for (auto r : record){
            q.push(r.second);
        }
        int steps = 0;
        while (!q.empty()){
            int i = 0;
            std::priority_queue<int> q_bak;
            while (i < n + 1 && !q.empty()){
                int e = q.top();
                q.pop();
                e--;
                steps++;
                if(e > 0)
                    q_bak.push(e);
                i++;
                if (q.empty()){
                    if(!q_bak.empty())
                        steps += n  + 1 - i;
                    break;
                }
            }
            while (!q_bak.empty()){
                q.push(q_bak.top());
                q_bak.pop();
            }
        }
        return steps;
    }
};

//[4,2,0,3,2,5]
class Solution_84 {
public:
    int largestRectangleArea(std::vector<int> &height) {
        int ret = 0;
        height.push_back(0);
        std::vector<int> index;
        for(int i = 0; i < height.size(); i++)
        {
            while(index.size() > 0 && height[index.back()] >= height[i])
            {
                int h = height[index.back()];
                index.pop_back();

                int sidx = index.size() > 0 ? index.back() : -1;
                if(h * (i-sidx-1) > ret)
                    ret = h * (i-sidx-1);
            }
            index.push_back(i);
        }
        return ret;
    }
};

class Solution_85 {
public:
    int maximalRectangle(std::vector<std::vector<char> >& matrix) {
        auto rows = matrix.size();
        auto columns = matrix[0].size();
        std::vector<std::vector<int> > dp_width(rows,std::vector<int>(columns,0));
        std::vector<std::vector<int> > dp_height(rows,std::vector<int>(columns,0));
        std::vector<std::vector<int> > dp_area(rows,std::vector<int>(columns,0));
        int max = 0;
        for(int i = 0;i < rows;i++){
            dp_width[i][0] = matrix[i][0] == '1'? 1: 0;
            for(int j = 0;j < columns;j++){
                dp_height[0][j] = matrix[0][j] == '1'? 1: 0;
                if (matrix[i][j] == '1'){
                    if (j >= 1){
                        dp_width[i][j] = 1 + dp_width[i][j - 1];
                    }
                    if(i >= 1){
                        dp_height[i][j] = 1 + dp_height[i - 1][j];
                    }
                    if (i == 0 || j == 0){
                        dp_area[i][j] = dp_width[i][j] * dp_height[i][j];
                    }else{

                    }
                    if (dp_area[i][j] > max){
                        max = dp_area[i][j];
                    }
                }
            }
        }
        return max;
    }
};