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

int main() {
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