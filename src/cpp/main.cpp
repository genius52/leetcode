#include <iostream>
#include <set>
#include <thread>
#include <condition_variable>
#include "array.h"
#include "diagram.h"
#include "mylist.h"
#include "other.h"
#include "mystring.h"
#include "tree.h"
#include "./number/229. Majority Element II.hpp"
#include "./number/220. Contains Duplicate III.hpp"
#include "./number/1441. Build an Array With Stack Operations.hpp"
#include "./number/1442. Count Triplets That Can Form Two Arrays of Equal XOR.hpp"
#include "./number/1447. Simplified Fractions.hpp"
#include "./number/368. Largest Divisible Subset.hpp"
#include "./tree/1443. Minimum Time to Collect All Apples in a Tree.cpp"
#include "./tree/1457. Pseudo-Palindromic Paths in a Binary Tree.hpp"
#include "./graph/332. Reconstruct Itinerary.cpp"
#include "./graph/207. Course Schedule.hpp"
#include "./graph/210. Course Schedule II.hpp"
#include "./graph/1444. Number of Ways of Cutting a Pizza.hpp"
#include "./string/1446. Consecutive Characters.cpp"
#include "./string/1451. Rearrange Words in a Sentence.hpp"
#include "./string/1455. Check If a Word Occurs As a Prefix of Any Word in a Sentence.hpp"
#include "./string/1456. Maximum Number of Vowels in a Substring of Given Length.hpp"
#define x 9999
#define max 9999
int data[10][10];
int dist[10];//记录最短路径为多少
int path[10];//记录最短路径
void fpath(int a[][10]);
int froute(int a[][10]);
int getlastlength(std::string& s);
std::vector<std::string> format_string(std::vector<std::string> v);
std::mutex mtx;
std::condition_variable gcv;
bool g_ready = false;

//Input: target = [2,3,4], n = 4
//Output: ["Push", "Pop", "Push", "Push", "Push"]

class Solution_5406 {
public:
    long find_target(int start, int end, int limit, std::map<int, vector<int>> record,std::vector<bool>& visited) {
        if (start == end) {
            return 0;
        }
        if (visited[start])
            return 100000;
        visited[start] = true;
        long cur_min_steps = 100000;;
        int subway_len = record[start].size();
        for (int i = 0; i < subway_len; i++) {
            int res = 0;
            int new_start = record[start][i];
            res = 1 + find_target(new_start, end, limit, record, visited);
            if (res < cur_min_steps) {
                cur_min_steps = res;
            }
        }
        visited[start] = false;
        return cur_min_steps;
    }

    int minTime(int n, vector<vector<int>>& edges, vector<bool>& hasApple) {
        std::map<int, vector<int>> record;
        int len = edges.size();
        for (int i = 0; i < len; i++) {
            if (record.find(edges[i][0]) == record.end()) {
                record[edges[i][0]] = std::vector<int>{ edges[i][1] };
            }
            else {
                record[edges[i][0]].push_back(edges[i][1]);
            }

            if (record.find(edges[i][1]) == record.end()) {
                record[edges[i][1]] = std::vector<int>{ edges[i][0] };
            }
            else {
                record[edges[i][1]].push_back(edges[i][0]);
            }
        }
        int apple_len = hasApple.size();
        int total_steps = 0;
        int start_index = 0;
        for (int i = 0; i < apple_len; i++) {
            if (!hasApple[i])
                continue;
            std::vector<bool> visited(n, false);
            total_steps += find_target(start_index, i, n, record, visited);
            start_index = i;
        }
        std::vector<bool> visited(n, false);
        total_steps += find_target(start_index, 0, n, record, visited);
        return total_steps;
    }
};


int main() {
    {
        Solution_368 s368;
        std::vector<int> nums{4,8,10,240};
        auto res = s368.largestDivisibleSubset(nums);
        std::cout<<"368 res = "<<res.size()<<std::endl;
    }
    {
        Solution_1444 s1444;
        std::vector<std::string> pizza = {"A..","AAA","..."};
        int k = 3;
        auto res = s1444.ways(pizza,k);
        std::cout<<"1444 res = "<<res<<std::endl;
    }
    {
        TreeNode t1(8);
        TreeNode t2(8);
        TreeNode t3(7);
        TreeNode t4(7);
        TreeNode t5(4);
        TreeNode t6(7);
        t1.left = &t2;
        t2.left = &t3;
        t2.right = &t4;
        t4.right = &t5;
        t5.right = &t6;

        TreeNode t7(2);
        TreeNode t8(8);
        TreeNode t9(1);
        t4.left = &t7;
        t7.right = &t8;
        t8.right = &t9;
        Solution_1457 s1457;
        auto res = s1457.pseudoPalindromicPaths(&t1);
        std::cout<<"1457 res = "<<res<<std::endl;

    }
    {
        Solution_1456 s1456;
        std::string s = "ibpbhixfiouhdljnjfflpapptrxgcomvnb";
        int k = 33;
        auto res = s1456.maxVowels(s,k);
        std::cout<<"1456 res = "<<res<<std::endl;
    }
    {
        Solution_1455 s1455;
        std::string sentence = "i";
        std::string search = "i";
        auto res = s1455.isPrefixOfWord(sentence,search);
        std::cout<<"1455 res = "<<res<<std::endl;
    }
    {
        Solution_210 s210;
        std::vector<std::vector<int>> input;
        input.push_back({1,0});
        auto res = s210.findOrder(2,input);
        std::cout<<"210 res = "<<res.size()<<std::endl;
    }
    {
        Solution_207 s207;
        int num = 2;
        std::vector<std::vector<int>> input;
//        input.push_back({0,1});
//        input.push_back({0,2});
//        input.push_back({1,2});
//        //auto res = s207.canFinish(num,input);
        input.push_back({1,0});
        auto res = s207.canFinish_bfs(num,input);
        std::cout<<"207 res = "<<res<<std::endl;
    }
    {
        Solution_1451 s1451;
        std::string input = "To be or not to be";
        auto res = s1451.arrangeWords(input);
        std::cout<<"res = "<< res << std::endl;
    }
    {
        Solution_1447 s1447;
        auto res = s1447.simplifiedFractions(4);
        std::cout << "1447 res = " << res.size() << std::endl;
    }
    {
        Solution_1446 s1446;
        string s = "aa";
        auto res = s1446.maxPower(s);
        std::cout<<"1446 res = "<<res<<std::endl;
    }
    {
        //[["JFK","SFO"],["JFK","ATL"],["SFO","ATL"],["ATL","JFK"],["ATL","SFO"]]
        //[["JFK","KUL"],["JFK","NRT"],["NRT","JFK"]]
        Solution_332 s332;
        std::vector<std::vector<std::string>> input = {std::vector<std::string>{"JFK","KUL"}, std::vector<std::string>{"JFK","NRT"},
        std::vector<std::string>{"NRT","JFK"}};
        auto res = s332.findItinerary(input);
        std::cout << "332 res = "<< res.size() <<std::endl;
    }
    {
        Solution_1443 s1443;
        int n = 7;
        std::vector<std::vector<int>> edges;
        edges.push_back({0,1});
        edges.push_back({0,2});
        edges.push_back({1,4});
        edges.push_back({1,5});
        edges.push_back({2,3});
        edges.push_back({2,6});
        std::vector<bool> hasApple{false,false,true,false,true,true,false};
        auto res = s1443.minTime(n,edges,hasApple);
        std::cout<<"1443 res = "<<res<<std::endl;
    }
    {
        Solution_5406 s5406;
        int n = 4;
        std::vector<std::vector<int>> edges{ std::vector<int>{0, 1} ,std::vector<int>{1, 2},std::vector < int>{0, 3}};
        std::vector<bool> apples{true,true,true,true};
        auto res = s5406.minTime(n, edges, apples);
        std::cout << res << std::endl;
    }
    {
        std::vector<int> v{58,506,448,780,716,871,427,693,798,672,446,600,998,840,174,282,436,636,968,84,924,888,228,948,848,553,377,245,396,774,650,336,986,55,1005,120,917,551,434,949,519,718,201,987,786,350,588,646,202,463,261,144,405,290,183,690,517,823,306,517,823,456,767,566,201,900,845,860,17,295,310,929,663,261,914,745,379,569,834,213,919,639,488,867,651,619,224,659,627,793,362,65,299,572,791,800,55,911,952,290,666,934,316,123,327,22,337,88,265,973,708,209,533,782,283,747,1008,397,637,952,453,285,216,397,341,885,544,767,223,258,477,371,174,486,328,755,955,593,490,760,786,834,80,2,82,467,385,375,246,51,197,635,702,979,365,73,292,567,787,745,506,690,840,938,226,325,423,207,360,888,528,374,870,560,342,308,98,645,743,732,59,191,132,694,562,183,645,323,966,971,13,392,389,313,188,471,363,650,993,534,503,322,181,940,793,727,462,137,327,213,402,820,678,54,656,285,909,834,207,780,963,474,537,809,304,17,289,44,269,506,247,22,225,601,696,937,273,103,374,501,131,49,178,932,790,908,154,329,467,996,567,486,977,679,374,969,703 };
        Solution_5405 s5405;
        auto res = s5405.countTriplets(v);
        std::cout << res << std::endl;
    }
    {
        Solution_5404 s5404;
        std::vector<int> v{3,6,9 };
        int n = 12;
        auto res = s5404.buildArray(v, n);
        std::cout << res.size() << std::endl;
    }
    {
        Solution_220 s220;
        std::vector<int> v{1,1,1,3,3,2,2,2};
        int k = 3;
        int t = 0;
        auto res = s220.containsNearbyAlmostDuplicate(v,k,t);
        std::cout<<res<<std::endl;
    }
    {
        Solution_241 s241;
        std::string input = "2-1-1";
        auto res = s241.diffWaysToCompute(input);
        std::cout<<res.size()<<std::endl;
    }
//    {
//        std::thread([](){
//            std::unique_lock <std::mutex> lck(mtx);
//            gcv.wait(lck);
//            std::cout<< "999"<<std::endl;
//            //g_ready = true;
//        }).detach();
//
//        std::this_thread::sleep_for(std::chrono::seconds(3));
//        gcv.notify_all();
//
//        std::thread([](){
//            std::unique_lock <std::mutex> lck(mtx);
//            gcv.wait(lck,[]() {return g_ready; });
//            std::cout<< "888"<<std::endl;
//        }).detach();
//    }
    {
        Solution_324 s324;
        std::vector<int> v{1, 3, 2, 2, 3, 1};
        s324.wiggleSort(v);
        std::cout<<"s324"<<std::endl;
    }
    {
        Solution_84 s84;
        std::vector<int> v{4,2,0,3,2,5};
        auto res = s84.largestRectangleArea(v);
        std::cout<<res<<std::endl;
    }

    {
        Solution_612 s612;
        std::vector<char> v{'A','A','A','B','C','D','E','F','G'};
        int n = 2;
        auto res = s612.leastInterval(v,n);
        std::cout<<res<<std::endl;
    }

//    {
//        std::string s;
//        char c;
//        std::cin>>s;
//        std::cin>>c;
//        int res = 0;
//        transform(s.begin(), s.end(), s.begin(), ::tolower);
//        c = std::tolower(c);
//        for (int i = 0;i < s.length();i++){
//            if (s[i] == c){
//                res++;
//            }
//        }
//        std::cout<<res;
//    }
    {
        std::string s = "XSUWHQ";
        auto l = getlastlength(s);
        std::cout<<l<<std::endl;
    }
    {
        TreeNode t1(3);
        TreeNode t2(7);
        TreeNode t3(9);
        TreeNode t4(2);
        TreeNode t5(1);
        t1.left = &t2;
        t1.right = &t3;
        t3.left = &t4;
        t3.right = &t5;
        Codec s297;
        auto res = s297.serialize(nullptr);
        //"0_1,3|1_2,7|1_3,9|3_4,2|3_5,1|"
        auto root = s297.deserialize(res);
        std::cout<<res <<std::endl;
    }
    {
        MedianFinder m;
        m.addNum(1);
        m.addNum(9);
        m.addNum(2);
        m.addNum(5);
        m.addNum(4);
        m.addNum(3);
        auto res = m.findMedian();
        std::cout<<res<<std::endl;
    }
    {
        using namespace s1366;
        s1366::Solution s;
        std::vector<std::string> scores = {"FVSHJIEMNGYPTQOURLWCZKAX","AITFQORCEHPVJMXGKSLNZWUY","OTERVXFZUMHNIYSCQAWGPKJL","VMSERIJYLZNWCPQTOKFUHAXG","VNHOZWKQCEFYPSGLAMXJIUTR","ANPHQIJMXCWOSKTYGULFVERZ","RFYUXJEWCKQOMGATHZVILNSP","SCPYUMQJTVEXKRNLIOWGHAFZ","VIKTSJCEYQGLOMPZWAHFXURN","SVJICLXKHQZTFWNPYRGMEUAO","JRCTHYKIGSXPOZLUQAVNEWFM","NGMSWJITREHFZVQCUKXYAPOL","WUXJOQKGNSYLHEZAFIPMRCVT","PKYQIOLXFCRGHZNAMJVUTWES","FERSGNMJVZXWAYLIKCPUQHTO","HPLRIUQMTSGYJVAXWNOCZEKF","JUVWPTEGCOFYSKXNRMHQALIZ","MWPIAZCNSLEYRTHFKQXUOVGJ","EZXLUNFVCMORSIWKTYHJAQPG","HRQNLTKJFIEGMCSXAZPYOVUW","LOHXVYGWRIJMCPSQENUAKTZF","XKUTWPRGHOAQFLVYMJSNEIZC","WTCRQMVKPHOSLGAXZUEFYNJI"};
        auto res = s.rankTeams(scores);
        std::cout<<res<<std::endl;
    }
    {
        using namespace s1365;
        s1365::Solution s;
        auto nums = std::vector<int>{37,64,63,2,41,78,51,36,2,20,25,41,72,100,17,43,54,27,34,86,12,48,70,44,87,68,62,98,68,30,8,92,5,10};
        auto res = s.smallerNumbersThanCurrent(nums);
        std::cout<<res[0]<<std::endl;
    }
//    {
//        std::vector<int> nums{1,2,3};
//        using namespace s384;
//        s384::Solution s(nums);
//        for (int i = 0;i<200;i++){
//            auto res = s.shuffle();
//            //s.reset();
//            res = s.shuffle();
//            for (auto i : res){
//                std::cout<<i;
//            }
//            std::cout<<std::endl;
//            //s.reset();
//        }
//        s.reset();
//    }
    {
        Node* n1 = new Node(7);
        Node* n2 = new Node(13);
        Node* n3 = new Node(11);
        Node* n4 = new Node(10);
        Node* n5 = new Node(1);
        n1->next = n2;
        n2->next = n3;
        n3->next = n4;
        n4->next = nullptr;

        n1->random = nullptr;
        n2->random = n1;
        n3->random = nullptr;
        n4->random = n3;
        n5->random = n1;

        auto s138 = new Solution_138();
        auto res = s138->copyRandomList(n1);
        std::cout << res->val << std::endl;
    }

    {
        vector<int> v = {2,7,4,1,8,1};
        std::cout<<lastStoneWeight(v)<<std::endl;
    }
    {
        auto res = diStringMatch("IDID");
    }
    {
        std::cout<<numTilePossibilities("AAB")<<std::endl;
        std::set<std::string> map;
        std::string s = "123";
        map.insert(s);
        auto res = map.find(s);
        if(map.end() == res)
            map.insert(s);
        map.insert(s);
        std::cout<<map.size()<<std::endl;
    }
//    {
//        vector<int> v = {3,2,1,6,0,5};
//        auto res = constructMaximumBinaryTree(v);
//        std::cout<<res->val<<std::endl;
//    }
//
//    {
//        std::vector<int> qv{2, 1, 2};
//        quick_sort(qv, 0, 2);
//        std::cout << hasAlternatingBits(715827882) << std::endl;
//        ListNode l5(5);
//        ListNode l4(4);
//        l4.next = &l5;
//        ListNode l3(3);
//        l3.next = &l4;
//        ListNode l2(2);
//        l2.next = &l3;
//        ListNode l1(1);
//        l1.next = &l2;
//        auto new_list = oddEvenList(&l1);
//
//        std::vector<int> nums2{1, 2, 3};
//        auto ret_subsets = subsets(nums2);
//        std::vector<int> nums{1, 2, 3};
//        auto ret_permute = permute(nums);
//        std::cout << backspaceCompare("ad#b", "aa#b");
//        std::cout << repeatedSubstringPattern("abac") << std::endl;
//        auto cnt = countPrimes(10);
//        auto ret = reverseBits(0b11111111111111111111111111111101);
//        std::cout << isPalindrome("OP") << std::endl;
//        std::cout << addStrings("9", "99") << std::endl;
//        Solution s;
//        std::cout << s.addBinary("100", "110010");
//
//        std::vector<std::string> v;
//        std::cout << longestCommonPrefix(v) << std::endl;
//        std::cout << isPalindrome(10) << std::endl;
//        std::cout << reverse_number(1534236469) << std::endl;
//    }

    int i,m;
    int a[10][10]=
            {
                    {x,4,2,3,x,x,x,x,x,x},
                    {x,x,x,x,10,9,x,x,x,x},
                    {x,x,x,x,6,7,10,x,x,x},
                    {x,x,x,x,x,3,8,x,x,x},
                    {x,x,x,x,x,x,x,4,8,x},
                    {x,x,x,x,x,x,x,9,6,x},
                    {x,x,x,x,x,x,x,5,4,x},
                    {x,x,x,x,x,x,x,x,x,8},
                    {x,x,x,x,x,x,x,x,x,4},
                    {x,x,x,x,x,x,x,x,x,x}
            };


    int start = 0;
    int end = 9;
    std::cout<<"from "<<start <<" run to" <<end<<" the shortest distance is "<< my_short_distance(a,start,end)<<std::endl;

    fpath(a);
    for (int j = 0; j < 10; ++j) {
        printf("%d shortest distance is:  %d\n",j,dist[j]);
    }


    m=froute(a);
    for(i=m-1;i>=0;i--)
        printf("pass from: %d\n",path[i]);
    return 0;
}

void fpath(int a[][10])
{
    int i,j,k;
    dist[0]=0;
    for(i=1;i<10;i++)
    {
        //k=max;
        dist[i] = max;
        for(j=0;j<i;j++)//从0号索引开始遍历，先遍历1号，一直到9
        {
            //  [j][i] 意思是 j 到 i 的距离
            if(a[j][i]!=x) //j到i有2种可能，1。直接到达 2，d（i,k) + d(k,j)
                if((dist[j]+a[j][i])< dist[i])
                    dist[i]=dist[j]+a[j][i];   // k = （0 到 j的最短距离）+ （j到i的最短距离）
        }
        //dist[i]=k;// 0号 到 i号 的最短路径
    }
}

int froute(int a[][10])
{
    int j,b,k=1,i=9;
    path[0]=10;
    while(i>0)
    {
        for(j=i-1;j>=0;j--)
        {
            if(a[j][i]!=x)
            {
                b=dist[i]-a[j][i];
                if(b==dist[j])
                {
                    path[k++]=j+1;
                    i=j;
                    break;
                }
            }
        }
    }
    return k;
}

int getlastlength(std::string& s){
    auto l = s.length();
    int res = 0;
    bool findword = false;
    for (int i = l - 1;i >= 0;i--){
        if (s[i] != ' '){
            if (!findword){
                findword = true;
            }
            res++;
        }else{
            if (findword)
                break;
        }
    }
    return res;
}

std::vector<std::string> format_string(std::vector<std::string> v){
    std::vector<std::string> res;
    for(auto i = 0;i < v.size();i++){
        auto s1 = v[i];
        while (s1.length() > 8){
            res.push_back(s1.substr(0,8));
            s1 = s1.substr(8,s1.length() - 8);
        }
        s1.append(8 - s1.length(),'0');
        res.push_back(s1);
    }
    return res;
}
//int getlastlength(std::string s){
//    auto l = s.length();
//    int res = 0;
//    bool findword = false;
//    for (int i = l - 1;i >= 0;i--){
//        if (s[i] == ' '){
//            if(findword)
//                break;
//        }else{
//            if (!findword)
//                findword = true;
//            res++;
//        }
//    }
//    return res;
//}