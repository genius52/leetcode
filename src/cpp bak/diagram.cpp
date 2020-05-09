//
// Created by 陶诚程 on 2019/4/16.
//
#include "diagram.h"
#include <iostream>
#include <vector>
#include <queue>
#include <map>
#define MAX_POINT_NUMBER 100
std::vector<int> point_map[MAX_POINT_NUMBER];//调用bfs前构建好点和点之间的相邻关系
int points_visit[MAX_POINT_NUMBER] = {0};

void bfs(int root,int n)
{
    std::queue<int> queue;
    queue.push(root);                   //出发节点入队列
    points_visit[root] = 1;
    while(!queue.empty())
    {
        auto val = queue.front();
        queue.pop();
        auto len = point_map[val].size();//获取该节点所有的相邻节点
        for (int i = 0; i < len; ++i) {//查看相邻节点之前是否已经访问过
            if (points_visit[point_map[val][i]] == 0)//如果没访问过，设置访问过标志位，并入队列
            {
                points_visit[point_map[val][i]] = 1;
                queue.push(point_map[val][i]);
            }
        }
    }
}

void DFS(int root)
{
    points_visit[root]=1;//将当前节点设置未已访问
    int len = point_map[root].size();//获取当前节点的所有相邻节点

    for(int i=0;i<len;++i)
        if(points_visit[point_map[root][i]] == 0)//遍历所有相邻节点，如果该节点未访问过，继续深度优先遍历
            DFS(point_map[root][i]);
}

//int route[4][4] = {
//        {0,4,2,0},
//        {4,0,0,1},
//        {2,0,0,1},
//        {0,1,1,0}
//};

int bfs_short_route(int arr[][4],int start,int end)
{
//    int minlength = 0;
//    for(int i = 0;i < 4;i++)
//        for(int j = 0;j < 4;j++)
//            std::cout<<array[i][j]<<std::endl;
    int points_visit[4] = {0};
    std::queue<int> queue;
    queue.push(start);
    points_visit[start] = 1;
    int length[4] = {0};
    int last = start;
    while(!queue.empty())
    {
        auto cur_node = queue.front();
        queue.pop();
        for (int i = 0; i < 4; ++i) {
            if(i == cur_node)
                continue;
            if(arr[cur_node][i] == 0)
                continue;
            if(points_visit[i] == 0)
            {
                points_visit[i] = 1;
                length[cur_node] += arr[cur_node][i];
                queue.push(i);
                std::cout<<cur_node<<","<<i<<" length = " <<arr[cur_node][i]<<std::endl;
            }
        }
    }
    return -1;
}
#define x 9999

int my_short_distance(int a[][10],int start,int end)
{
    int distance[10];
    for (int k = 0; k < 10; ++k) {
        distance[k] = 0;
    }
    distance[start] = 0;
    //dist[i] = i 到 start 的距离
    for (int i = 0; i < 10; ++i) {
        distance[i] = x;
        if(i == start)
            continue;
        //1. 算出start到相邻结点的距离
        for (int j = 0; j < i; ++j) {
            if(a[j][i]!= x)
                if((distance[j] + a[j][i])< distance[i])
                    distance[i] = distance[j]+a[j][i];  // 更新从start到i的最短距离
        }

    }
    return distance[end];
}


//#include <iostream>
//#include <string>
//#include <list>
//using namespace std;
//int main()
//{
//    double pi[3] = { 0.2, 0.4, 0.4 };
//    double C[3][3] = { 0.5, 0.2, 0.3, 0.3, 0.5, 0.2, 0.2, 0.3, 0.5 };
//    double E[3][2] = { 0.5, 0.5, 0.4, 0.6, 0.7, 0.3 };
//    string output[4] = { "R", "W", "R","W" };
//    int state[3] = { 1, 2, 3 };
//    int row = 3;
//    int column = 3;
//    //开辟数组空间
//    double **delta = new double *[row];
//    int **path = new int *[row];    int i, j,k;
//    for (i = 0; i < 3; i++)
//    {
//        delta[i] = new double[3];        path[i] = new int[3];
//    }    //将输出状态转为数组
//    int outnum[4];
//    for (i = 0; i < row; i++)
//    {
//        if (output[i] == "R")
//            outnum[i] = 0;
//        else if (output[i] == "W")
//            outnum[i] = 1;
//    }    //初始化
//    for (i = 0; i < column; i++)
//    {
//        path[i][0] = 0;
//        delta[i][0] = pi[i] * E[i][0];
//        cout << delta[i][0] << endl;
//    }    //递归
//    for (j = 1; j < row; j++)//序列遍历，矩阵列遍历
//    {
//        for (k = 0; k < column; k++)//状态遍历
//        {
//            double pro = 0;
//            int sta = 0;
//            for (i = 0; i < column; i++)//矩阵行遍历
//            {
//                double temp = delta[i][j - 1] *C[i][k]* E[k][outnum[j]] ;//delta[i][j-1]*转移概率*发射概率
//                cout << delta[i][j - 1] << " " << E[k][outnum[j]] << endl;
//                cout << temp << endl;
//                if (temp > pro)
//                {
//                    pro = temp;
//                    sta = i;//记录路径信息
//                }
//            }
//            delta[k][j] = pro;
//            path[k][j]= sta;
//        }
//    }    //终止，找到概率最大值
//     double max = 0;    int sta = 0;
//     for (i = 0; i < column; i++)
//     {
//         if (delta[i][row - 1] > max)
//        {
//             max = delta[i][row - 1];
//             sta = i;
//        }
//     }    //回溯，找到最大路径及其对应的状态值
//     list<int> listPath;
//     listPath.push_back(sta+1);
//     for (j = row - 1; j > 0; j--)
//     {
//         sta = path[sta][j];
//         listPath.push_back(sta+1);
//     }    //输出
//     cout << "max probability: " << max << endl;
//     list<int> ::iterator itepath;
//     for (itepath = listPath.begin(); itepath != listPath.end(); itepath++)
//         cout << *itepath << " ";
//}

