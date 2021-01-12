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
#include "mylist.h"
#include "./array/396. Rotate Function.hpp"
#include "./array/417. Pacific Atlantic Water Flow.hpp"
#include "./array/1470. Shuffle the Array.hpp"
#include "./array/1471. The k Strongest Values in an Array.hpp"
#include "./array/1465. Maximum Area of a Piece of Cake After Horizontal and Vertical Cuts.hpp"
#include "./array/1481. Least Number of Unique Integers after K Removals.hpp"
#include "./array/1477. Find Two Non-overlapping Sub-arrays Each With Target Sum.hpp"
#include "./array/435. Non-overlapping Intervals.hpp"
#include "./array/436. Find Right Interval.hpp"
#include "./array/1493. Longest Subarray of 1's After Deleting One Element.hpp"
#include "./array/1497. Check If Array Pairs Are Divisible by k.hpp"
#include "./array/456. 132 Pattern.hpp"
#include "./array/498. Diagonal Traverse.hpp"
#include "./array/525. Contiguous Array.hpp"
#include "./array/846. Hand of Straights.hpp"
#include "./array/983. Minimum Cost For Tickets.hpp"
#include "array/532. K-diff Pairs in an Array.hpp"
#include "array/598. Range Addition II.hpp"
#include "array/661. Image Smoother.hpp"
#include "array/932. Beautiful Array.hpp"
#include "array/735. Asteroid Collision.hpp"
#include "array/1536. Minimum Swaps to Arrange a Binary Grid.hpp"
#include "array/870. Advantage Shuffle.hpp"
#include "array/954. Array of Doubled Pairs.hpp"
#include "array/1024. Video Stitching.hpp"
#include "array/1049. Last Stone Weight II.hpp"
#include "array/1054. Distant Barcodes.hpp"
#include "array/1488. Avoid Flood in The City.hpp"
#include "array/1509. Minimum Difference Between Largest and Smallest Value in Three Moves.hpp"
#include "./number/229. Majority Element II.hpp"
#include "./number/220. Contains Duplicate III.hpp"
#include "./number/1441. Build an Array With Stack Operations.hpp"
#include "./number/1442. Count Triplets That Can Form Two Arrays of Equal XOR.hpp"
#include "./number/1447. Simplified Fractions.hpp"
#include "./number/368. Largest Divisible Subset.hpp"
#include "./number/343. Integer Break.hpp"
#include "./number/373. Find K Pairs with Smallest Sums.hpp"
#include "./number/376. Wiggle Subsequence.hpp"
#include "./number/386. Lexicographical Numbers.hpp"
#include "./number/390. Elimination Game.hpp"
#include "./number/397. Integer Replacement.hpp"
#include "./number/1464. Maximum Product of Two Elements in an Array.hpp"
#include "./number/1475. Final Prices With a Special Discount in a Shop.hpp"
#include "./number/1492. The kth Factor of n.hpp"
#include "./number/1498. Number of Subsequences That Satisfy the Given Sum Condition.hpp"
#include "./number/464. Can I Win.hpp"
#include "number/492. Construct the Rectangle.hpp"
#include "number/507. Perfect Number.hpp"
#include "number/703. Kth Largest Element in a Stream.hpp"
#include "number/682. Baseball Game.hpp"
#include "number/717. 1-bit and 2-bit Characters.hpp"
#include "number/762. Prime Number of Set Bits in Binary Representation.hpp"
#include "number/788. Rotated Digits.hpp"
#include "number/812. Largest Triangle Area.hpp"
#include "number/1518. Water Bottles.hpp"
#include "number/868. Binary Gap.hpp"
#include "number/31. Next Permutation.hpp"
#include "number/166. Fraction to Recurring Decimal.hpp"
#include "number/241. Different Ways to Add Parentheses.hpp"
#include "number/659. Split Array into Consecutive Subsequences.hpp"
#include "number/781. Rabbits in Forest.hpp"
#include "number/1619. Mean of Array After Removing Some Elements.hpp"
#include "number/853. Car Fleet.hpp"
#include "number/1647. Minimum Deletions to Make Character Frequencies Unique.hpp"
#include "./tree/1443. Minimum Time to Collect All Apples in a Tree.cpp"
#include "./tree/1457. Pseudo-Palindromic Paths in a Binary Tree.hpp"
#include "./tree/450. Delete Node in a BST.hpp"
#include "./tree/449. Serialize and Deserialize BST.hpp"
#include "tree/987. Vertical Order Traversal of a Binary Tree.hpp"
#include "tree/988. Smallest String Starting From Leaf.hpp"
#include "./graph/332. Reconstruct Itinerary.cpp"
#include "./graph/207. Course Schedule.hpp"
#include "./graph/210. Course Schedule II.hpp"
#include "./graph/1444. Number of Ways of Cutting a Pizza.hpp"
#include "./graph/310. Minimum Height Trees.hpp"
#include "./graph/355. Design Twitter.hpp"
#include "./graph/1466. Reorder Routes to Make All Paths Lead to the City Zero.hpp"
#include "./graph/399. Evaluate Division.hpp"
#include "./graph/433. Minimum Genetic Mutation.hpp"
#include "./graph/1472. Design Browser History.hpp"
#include "./graph/1476. Subrectangle Queries.hpp"
#include "./graph/1496. Path Crossing.hpp"
#include "./graph/452. Minimum Number of Arrows to Burst Balloons.hpp"
#include "graph/5536. Maximal Network Rank.hpp"
#include "graph/874. Walking Robot Simulation.hpp"
#include "graph/883. Projection Area of 3D Shapes.hpp"
#include "graph/892. Surface Area of 3D Shapes.hpp"
#include "graph/547. Friend Circles.hpp"
#include "graph/677. Map Sum Pairs.hpp"
#include "graph/721. Accounts Merge.hpp"
#include "graph/911. Online Election.hpp"
#include "graph/981. Time Based Key-Value Store.hpp"
#include "graph/1146. Snapshot Array.hpp"
#include "./string/1446. Consecutive Characters.cpp"
#include "./string/1451. Rearrange Words in a Sentence.hpp"
#include "./string/1455. Check If a Word Occurs As a Prefix of Any Word in a Sentence.hpp"
#include "./string/1456. Maximum Number of Vowels in a Substring of Given Length.hpp"
#include "./string/424. Longest Repeating Character Replacement.hpp"
#include "./string/474. Ones and Zeroes.hpp"
#include "./string/482. License Key Formatting.hpp"
#include "string/1507. Reformat Date.hpp"
#include "string/744. Find Smallest Letter Greater Than Target.hpp"
#include "string/748. Shortest Completing Word.hpp"
#include "string/806. Number of Lines To Write String.hpp"
#include "string/824. Goat Latin.hpp"
#include "string/893. Groups of Special-Equivalent Strings.hpp"
#include "string/937. Reorder Data in Log Files.hpp"
#include "string/692. Top K Frequent Words.hpp"
#include "string/5535. Maximum Nesting Depth of the Parentheses.hpp"
#include "string/1616. Split Two Strings to Make Palindrome.hpp"
#include "string/955. Delete Columns to Make Sorted II.hpp"
#include "string/1209. Remove All Adjacent Duplicates in String II.hpp"
#include "./list/1483. Kth Ancestor of a Tree Node.hpp"
#include "./list/445. Add Two Numbers II.hpp"
#include "thread/1115. Print FooBar Alternately.hpp"
#include "thread/1116. Print Zero Even Odd.hpp"
//#include "./number/1461. Check If a String Contains All Binary Codes of Size K.hpp"
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
        Solution_1509 s1509;
        std::vector<int> nums{0,0,0,1,1,1};
        auto res = s1509.minDifference(nums);
        std::cout << "1509 res = " << res << std::endl;
    }
    {
        Solution_1488 s1488;
        vector<int> rain{1,0,2,0,3,0,2,0,0,0,1,2,3};
        auto res = s1488.avoidFlood(rain);
        std::cout << "1488 res = " << res.size() << std::endl;
    }
    {
        Solution_1209 s1209;
        std::string s = "pbbcggttciiippooaais";
        int k = 2;
        auto res = s1209.removeDuplicates(s,k);
        std::cout << "1209 res = " << res << std::endl;
    }
    {
//Input: ["SnapshotArray","set","snap","set","get"]
//[[3],[0,5],[],[0,6],[0,0]]
//Output: [null,null,0,null,5]
        SnapshotArray s1146(1);
        s1146.set(0,15);
        auto res = s1146.snap();
        res = s1146.snap();
        res = s1146.snap();
        res = s1146.get(0,2);
        res = s1146.snap();
        res = s1146.snap();
        s1146.get(0,0);
        std::cout << "1146 res = " << res << std::endl;
    }
    {
        ZeroEvenOdd* s1116 = new ZeroEvenOdd(1);
        std::thread t1 = std::thread([s1116](){
            s1116->zero([](int n) {
                            std::cout << n;
                        });
        });
        t1.detach();
        std::thread t2 = std::thread([s1116](){
            s1116->even([](int n) {
                std::cout << n;
            });
        });
        t2.detach();
        std::thread t3 = std::thread([s1116](){
            s1116->odd([](int n) {
                std::cout << n;
            });
        });
        t3.detach();
        std::this_thread::sleep_for(std::chrono::seconds(20));
    }
    {
        FooBar* s1115 = new FooBar(3);
        std::thread bar = std::thread([s1115](){
            s1115->bar([](){
                static int i = 0;
                std::cout << "bar" << i << std::endl;
                i++;
            });
        });
        bar.detach();
        std::this_thread::sleep_for(std::chrono::seconds(1));
        std::thread foo = std::thread([s1115](){
            s1115->foo([](){
                static int i = 0;
                std::cout << "foo";
                i++;
            });
        });
        foo.detach();
        std::this_thread::sleep_for(std::chrono::seconds(20));
    }
    {
        Solution_1054 s1054;
        std::vector<int> barcodes{1,1,1,1,2,2,3,3};
        auto res = s1054.rearrangeBarcodes(barcodes);
        std::cout << "1054 res = " << res.size() << std::endl;
    }
    {
        Solution_1049 s1049;
        std::vector<int> stones{31,26,33,21,40};
        auto res = s1049.lastStoneWeightII(stones);
        std::cout << "1049 res = " << res << std::endl;
    }
    {
        Solution_988 s988;
        TreeNode t1(0);
        TreeNode t2(1);
        TreeNode t3(2);
        TreeNode t4(3);
        TreeNode t5(4);
        TreeNode t6(3);
        TreeNode t7(4);
        t1.left = &t2;
        t1.right = &t3;
        t2.left = &t4;
        t2.right = &t5;
        t3.left = &t6;
        t3.right = &t7;
        auto res = s988.smallestFromLeaf(&t1);
        std::cout << "988 res = " << res << std::endl;
    }
    {
        //3,9,20,null,null,15,7
        Solution_987 s987;
        TreeNode t1(3);
        TreeNode t2(9);
        TreeNode t3(20);
        TreeNode t4(15);
        TreeNode t5(17);

        t1.left = &t2;
        t1.right = &t3;
        t3.left = &t4;
        t3.right = &t5;
        auto res = s987.verticalTraversal(&t1);
        std::cout << "987 res = " << res.size() << std::endl;
    }
    {
        Solution_983 s983;
        std::vector<int> days{1,2,4,5,6,7,8,9,10,11,12,13,14,15,16,17,18,20,21,24,25,27,28,29,30,31,34,37,38,39,41,43,44,45,47,48,49,54,57,60,62,63,66,69,70,72,74,76,78,80,81,82,83,84,85,88,89,91,93,94,97,99};
        std::vector<int> costs{9,38,134};
        auto res = s983.mincostTickets(days,costs);
        std::cout << "981 res = " << res << std::endl;
    }
    {
        //"love","high",10],["love","low",20]
        TimeMap s981;

        s981.set("love","high",10);
        s981.set("love","low",20);
        auto res = s981.get("love",10);
        res = s981.get("love",30);
        std::cout << "981 res = " << res << std::endl;
    }
    {
        std::vector<std::string> A{"abx","agz","bgc","bfc"};
        Solution_955 s955;
        auto res = s955.minDeletionSize(A);
        std::cout << "955 res = " << res << std::endl;
    }
    {
        Solution_954 s954;
        std::vector<int> A{4,-2,2,-4,1,2,0,0};
        auto res = s954.canReorderDoubled(A);
        std::cout << "954 res = " << res << std::endl;
    }
    {
        std::vector<int> persons{0,1,1,0,0,1,0};
        std::vector<int> times{0,5,10,15,20,25,30};
        TopVotedCandidate s911(persons,times);
        auto res = s911.q(25);
        std::cout << "911 res = " << res << std::endl;
    }
    {
        Solution_1647 s1647;
        std::string s = "ceabaacb";
        auto res = s1647.minDeletions(s);
        std::cout << "1647 res = " << res << std::endl;
    }
    {
        Solution_870 s870;
        std::vector<int> A{12,24,8,32};
        std::vector<int> B{13,25,32,11};
        auto res = s870.advantageCount(A,B);
        std::cout << "870 res = " << res.size() << std::endl;
    }
    {
        Solution_853 s853;
        int target = 10;
        std::vector<int> position{6,8};
        std::vector<int> speed{3,2};
        auto res = s853.carFleet(target,position,speed);
        std::cout << "853 res = " << res << std::endl;
    }
    {
        Solution_846 s846;
        std::vector<int> hand{1,2,3,6,2,3,4,7,8};
        int W = 3;
        auto res = s846.isNStraightHand(hand,W);
        std::cout << "846 res = " << res << std::endl;
    }
    {
        Solution_1619 s1619;
        std::vector<int> arr{ 6,0,7,0,7,5,7,8,3,4,0,7,8,1,6,8,1,1,2,4,8,1,9,5,4,3,8,5,10,8,6,6,1,0,6,10,8,2,3,4 };
        auto res = s1619.trimMean(arr);
        std::cout << "781 res = " << res << std::endl;
    }
    {
        Solution_781 s781;
        std::vector<int> answers{};
        auto res = s781.numRabbits(answers);
        std::cout << "781 res = " << res << std::endl;
    }
    {
        Solution_1616 s1616;
        std::string a = "aejbaalflrmkswrydwdkdwdyrwskmrlfqizjezd";
        std::string b = "uvebspqckawkhbrtlqwblfwzfptanhiglaabjea";
        auto res = s1616.checkPalindromeFormation(a,b);
        std::cout << "1616 res = " << res << std::endl;
    }
    {
        int n = 8;
        std::vector<std::vector<int>> roads{ {0,1 }, { 1,2 }, { 2,3 }, { 2,4 }, { 5,6 }, { 5,7 } };
        Solution_5536 s5536;
        auto res = s5536.maximalNetworkRank(n, roads);
        std::cout << "s5536 res = " << res << std::endl;
    }
    {
        Solution_5535 s5535;
        std::string s = "1";
        auto res = s5535.maxDepth(s);
        std::cout << "5535 res = " << res << std::endl;
    }
    {
        Solution_1536 s1536;
        std::vector<std::vector<int>> grid{{1,0,0},{1,1,0},{1,1,1}};
        auto res = s1536.minSwaps(grid);
        std::cout<<"1536 res = "<<res<<std::endl;
    }
    {
        Solution_735 s735;
        std::vector<int> asteroids{1,-2,-2,-2};
        auto res = s735.asteroidCollision(asteroids);
        std::cout<<"735 res = "<<res.size()<<std::endl;
    }
    {
        Solution_721 s721;
        auto accounts = std::vector<std::vector<std::string>>{{"Lily","Lily8@m.co","Lily2@m.co","Lily2@m.co"},
                                                              {"Fern","Fern6@m.co","Fern1@m.co","Fern1@m.co"},
                                                              {"John","John8@m.co","John0@m.co","John2@m.co"},
                                                              {"Kevin","Kevin4@m.co","Kevin8@m.co","Kevin0@m.co"},
                                                              {"Fern","Fern2@m.co","Fern7@m.co","Fern7@m.co"},
                                                              {"John","John4@m.co","John6@m.co","John7@m.co"},
                                                              {"Hanzo","Hanzo8@m.co","Hanzo8@m.co","Hanzo6@m.co"},
                                                              {"Bob","Bob8@m.co","Bob6@m.co","Bob1@m.co"}};
        auto res = s721.accountsMerge(accounts);
        std::cout<<"721 res = "<<res.size()<<std::endl;
    }
    {
        Solution_692 s692;
        std::vector<std::string> words{"i", "love", "leetcode", "i", "love", "coding"};
        int k = 2;
        auto res = s692.topKFrequent(words,k);
        std::cout<<"692 res = "<<res.size()<<std::endl;
    }
    {
        MapSum s677;

        s677.insert("a", 3);
        auto res = s677.sum("ap");
        s677.insert("b", 2);
        res = s677.sum("a");
        std::cout<<"677 res = "<<res<<std::endl;
    }
    {
        Solution_659 s659;
        std::vector<int> nums{1,3,3,4,4,7,8,8,9,10};
        auto res = s659.isPossible(nums);
        std::cout << "659 res = " << res << std::endl;
    }
    {
        Solution_547 s547;
        std::vector<std::vector<int>> input{ {1,1,0 }, { 1,1,0 }, { 0,0,1 } };
        auto res = s547.findCircleNum(input);
        std::cout << "547 res = " << res << std::endl;
    }
    {
        Solution_525 s525;
        std::vector<int> nums{1,1,1,1,1,1,0,0,0,0};
        auto res = s525.findMaxLength(nums);
        std::cout << "525 res = " << res << std::endl;
    }
    {
        Solution_498 s498;
        std::vector<std::vector<int>> input{ {1,2,3,4},{5,6,7,8},{9,10,11,12} };
        auto res = s498.findDiagonalOrder(input);
        std::cout << "498 res = " << res[0] << std::endl;
    }
    {
        Solution_474 s474;
        std::vector<std::string> strs{"10","0001","111001","1","0"};
        int m = 30;
        int n = 30;
        auto res = s474.findMaxForm(strs,m,n);
        std::cout<<"474 res = "<<res<<std::endl;
    }
    {
        Solution_241 s241;
        std::string input = "2*3-4*5";
        auto res = s241.diffWaysToCompute(input);
        std::cout<<"241 res = "<<res[0]<<std::endl;
    }
    {
        Solution_166 s166;
        int numerator = 1;
        int denominator = 333;
        auto res = s166.fractionToDecimal(numerator,denominator);
        std::cout<<"932 res = "<<res[0]<<std::endl;
    }
    {
        Solution_932 s932;
        int N = 4;
        auto res = s932.beautifulArray(N);
        std::cout<<"932 res = "<<res[0]<<std::endl;
    }
    {
        Solution_31 s31;
        std::vector<int> input{3,4,5};
        s31.nextPermutation(input);
        std::cout<<"31 res = "<<input[0]<<std::endl;
    }
    {
        Solution_937 s937;
        std::vector<std::string> input{"dig1 8 1 5 1","let1 art can","dig2 3 6","let2 own kit dig","let3 art zero"};
        auto res = s937.reorderLogFiles(input);
        std::cout<<"937 res = "<<res[0]<<std::endl;
    }
    {
        Solution_893 s893;
        std::vector<string> input{"abc","acb","bac","bca","cab","cba"};
        auto res = s893.numSpecialEquivGroups(input);
        std::cout<<"893 res = "<<res<<std::endl;
    }
    {
        Solution_892 s892;
        std::vector<std::vector<int>> input{{2,2,2},{2,1,2},{2,2,2}};
        auto res = s892.surfaceArea(input);
        std::cout<<"892 res = "<<res<<std::endl;
    }
    {
        Solution_883 s883;
        std::vector<std::vector<int>> grid{{2,2,2},{2,1,2},{2,2,2}};
        auto res = s883.projectionArea(grid);
        std::cout<<"883 res = "<<res<<std::endl;
    }
    {
        Solution_874 s874;
        std::vector<int> commands{4,-1,3};
        std::vector<vector<int>> obstacles{};
        auto res = s874.robotSim(commands,obstacles);
        std::cout<<"874 res = "<<res<<std::endl;
    }
    {
        Solution_868 s868;
        int N = 8;
        auto res = s868.binaryGap(N);
        std::cout<<"868 res = "<<res<<std::endl;
    }
    {
        Solution_824 s824;
        std::string s = "The quick brown fox jumped over the lazy dog";
        auto res = s824.toGoatLatin(s);
        std::cout<<"824 res = "<<res<<std::endl;
    }
    {
        Solution_1518 s1518;
        int numBottles = 2;
        int numExchange = 3;
        auto res = s1518.numWaterBottles(numBottles,numExchange);
        std::cout<<"1518 res = "<<res<<std::endl;
    }
    {
        Solution_812 s812;
        std::vector<std::vector<int>> input{{1,0},{0,0},{0,1}};
        auto res = s812.largestTriangleArea(input);
        std::cout<<"812 res = "<<res<<std::endl;
    }
    {
        Solution_806 s806;
        std::vector<int> width{10,10,10,10,10,10,10,10,10,10,10,10,10,10,10,10,10,10,10,10,10,10,10,10,10,10};
        std::string S = "abcdefghijklmnopqrstuvwxyz";
        auto res = s806.numberOfLines(width,S);
        std::cout<<"806 res = "<<res[0] << res[1]<<std::endl;
    }
    {
        Solution_788 s788;
        auto res = s788.rotatedDigits(20);
        std::cout<<"762 res = "<<res<<std::endl;
    }
    {
        Solution_762 s762;
        int L = 10;
        int R = 15;
        auto res = s762.countPrimeSetBits2(L,R);
        std::cout<<"762 res = "<<res<<std::endl;
    }
    {
        Solution_748 s748;
        std::string licensePlate = "1s3 456";
        std::vector<string> words{"looks", "pest", "stew", "show"};
        auto res = s748.shortestCompletingWord(licensePlate,words);
        std::cout<<"748 res = "<<res<<std::endl;
    }
    {
        Solution_744 s744;
        std::vector<char> input{'c', 'f', 'j'};
        char target = 'c';
        auto res = s744.nextGreatestLetter(input,target);
        std::cout<<"744 res = "<<res<<std::endl;
    }
    {
        Solution_717 s717;
        std::vector<int> input{1,1,1,0};
        auto res = s717.isOneBitCharacter(input);
        std::cout<<"717 res = "<<res<<std::endl;
    }
    {
        Solution_1507 s1507;
        std::string date = "20th Oct 2052";
        auto res = s1507.reformatDate(date);
        std::cout<<"1507 res = "<<res<<std::endl;
    }
    {
        Solution_682 s682;
        std::vector<std::string> input{"5","2","C","D","+"};
        auto res = s682.calPoints(input);
        std::cout<<"682 res = "<<res<<std::endl;
    }
    {
        std::vector<int> input{};
        int k = 1;
        KthLargest s703(k,input);
        auto res = s703.add(3);
        std::cout<<"703 res = "<<res<<std::endl;
    }
    {
        Solution_661 s661;
        std::vector<std::vector<int>> input{{1,1,1},{1,0,1},{1,1,1}};
        auto res = s661.imageSmoother(input);
        std::cout<<"661 res = "<<res.size()<<std::endl;
    }
    {
        Solution_598 s598;
        int m = 3;
        int n = 3;
        std::vector<std::vector<int>> input{{2,2},{3,3}};
        auto res = s598.maxCount(m,n,input);
        std::cout<<"598 res = "<<res<<std::endl;
    }
    {
        Solution_532 s532;
        std::vector<int> input{1, 3, 1, 5, 4};
        int k = 0;
        auto res = s532.findPairs(input,k);
        std::cout<<"532 res = "<<res<<std::endl;
    }
    {
        Solution_507 s507;
        int num = 6;
        auto res = s507.checkPerfectNumber(num);
        std::cout<<"507 res = "<<res<<std::endl;
    }
    {
        Solution_492 s492;
        int area = 12;
        auto res = s492.constructRectangle(area);
        std::cout<<"492 res = "<<res[0]<<std::endl;
    }
    {
        Solution_482 s482;
        std::string s = "--a-a-a-a--";
        int k = 2;
        auto res = s482.licenseKeyFormatting(s,k);
        std::cout<<"482 res = "<<res<<std::endl;
    }
    {
        Solution_464 s464;
        int maxChoosableInteger = 10;
        int desiredTotal = 11;
        auto res = s464.canIWin(maxChoosableInteger,desiredTotal);
        std::cout<<"464 res = "<<res<<std::endl;
    }
    {
        Solution_456 s456;
        std::vector<int> input{-2,1,2,-2,1,2};
        auto res = s456.find132pattern(input);
        std::cout<<"456 res = "<<res<<std::endl;
    }
    {
        Solution_452 s452;
        std::vector<std::vector<int>> input{{9,12},{1,10},{4,11},{8,12},{3,9},{6,9},{6,7}};
        auto res = s452.findMinArrowShots(input);
        std::cout<<"452 res = "<<res<<std::endl;
    }
    {
        Solution_1498 s1498;
        std::vector<int> input{5,2,4,1,7,6,8};
        int target = 16;
        auto res = s1498.numSubseq(input,target);
        std::cout<<"1498 res = "<<res<<std::endl;
    }
    {
        Solution_1497 s1497;
        std::vector<int> input{-1,1,-2,2,-3,3,-4,4};
        int k = 3;
        auto res = s1497.canArrange(input,k);
        std::cout << "1497 res = "<<res<<std::endl;
    }
    {
        Solution_1496 s1496;
        std::string path = "W";
        auto res = s1496.isPathCrossing(path);
        std::cout << "1496 res = "<<res<<std::endl;
    }
    {
        Solution_1493 s1493;
        std::vector<int> input{0,1,1,1,0,1,1,0,1};
        auto res = s1493.longestSubarray(input);
        std::cout<<"1493 res = "<<res<<std::endl;
    }
    {
        Solution_1492 s1492;
        int n = 24;
        int k = 6;
        auto res = s1492.kthFactor(n,k);
        std::cout<<res<<std::endl;
    }
    {
        Codec_449 s449;
        TreeNode t5(5);
        TreeNode t3(3);
        TreeNode t6(6);
        TreeNode t2(2);
        TreeNode t4(4);
        TreeNode t7(7);
        t5.left = &t3;
        t5.right = &t6;
        t3.left = &t2;
        t3.right = &t4;
        t6.right = &t7;
        auto s = s449.serialize(&t5);
        auto t = s449.deserialize(s);
        std::cout<<"449 s = " << s<<std::endl;
    }
    {
        //root = [5,3,6,2,4,null,7]
        //key = 3
        Solution_450 s450;
        TreeNode t5(5);
        TreeNode t3(3);
        TreeNode t6(6);
        TreeNode t2(2);
        TreeNode t4(4);
        TreeNode t7(7);
        t5.left = &t3;
        t5.right = &t6;
        t3.left = &t2;
        t3.right = &t4;
        t6.right = &t7;
        int key = 3;
        auto res = s450.deleteNode(&t5,key);
        std::cout<< "450 res = "<<res->val << std::endl;
    }
    {
        Solution_445 s445;
        //[1]
        //[9,9]
        ListNode l1(1);
        ListNode l2(9);
        ListNode l3(9);
        l2.next = &l3;
        auto res = s445.addTwoNumbers(&l1,&l2);
        std::cout<<res->val<<std::endl;
    }
    {
        Solution_436 s436;
        std::vector<std::vector<int>> input{{1,2}};
        auto res = s436.findRightInterval(input);
        std::cout<<"436 res = "<<res.size()<<std::endl;
    }
    {
        Solution_435 s435;
        std::vector<std::vector<int>> input{{1,100},{11,22},{1,11},{2,12}};
        auto res = s435.eraseOverlapIntervals(input);
        std::cout<<"435 res = "<<res << std::endl;
    }
    {
        Solution_1477 s1477;
        std::vector<int> input{33,15,14,2,4,39,23,1,3,2,62,1,2,1,2,43,7,9,2,3,1,3,45,17,3,2,1,7,6,21,34,65,1,2};
        int target = 68;
        auto res = s1477.minSumOfLengths(input,target);
        std::cout<<"1477 res = "<<res << std::endl;
    }
    {
        std::vector<std::vector<int>> input{{1,2,1},{4,3,4},{3,2,1},{1,1,1}};
        SubrectangleQueries query(input);
        query.updateSubrectangle(0,0,3,2,5);
        auto res = query.getValue(1,2);
        std::cout<<"1476 res = "<<res << std::endl;
    }
    {
        int n = 7;
        std::vector<int> input{-1, 0, 0, 1, 1, 2, 2};
        TreeAncestor tree(n,input);
        auto res = tree.getKthAncestor(6,3);
        std::cout<<"1475 res = "<<res << std::endl;
    }
    {
        Solution_1475 s1475;
        std::vector<int> input{10,1,1,6};
        auto res = s1475.finalPrices(input);
        std::cout<<"1475 res = "<<res.size() << std::endl;
    }
    {
        Solution_1481 s1481;
        std::vector<int> input{ 4, 3, 1, 1, 3, 3, 2 };
        int k = 3;
        auto res = s1481.findLeastNumOfUniqueInts(input,k);
        std::cout << "1481 res = " << res << std::endl;
    }
    {
        BrowserHistory browserHistory("leetcode.com");
        browserHistory.visit("google.com");       // You are in "leetcode.com". Visit "google.com"
        //browserHistory.visit("facebook.com");     // You are in "google.com". Visit "facebook.com"
        //browserHistory.visit("youtube.com");      // You are in "facebook.com". Visit "youtube.com"
        auto res = browserHistory.back(1);                   // You are in "youtube.com", move back to "facebook.com" return "facebook.com"
        res = browserHistory.back(1);                   // You are in "facebook.com", move back to "google.com" return "google.com"
        res = browserHistory.forward(1);                // You are in "google.com", move forward to "facebook.com" return "facebook.com"
        browserHistory.visit("linkedin.com");     // You are in "facebook.com". Visit "linkedin.com"
        res = browserHistory.forward(2);                // You are in "linkedin.com", you cannot move forward any steps.
        res = browserHistory.back(2);                   // You are in "linkedin.com", move back two steps to "facebook.com" then to "google.com". return "google.com"
        res = browserHistory.back(7);                   // You are in "google.com", you can move back only one step to "leetcode.com". return "leetcode.com"
        std::cout << "5430 res = " << res << std::endl;
    }
    {
        Solution_1471 s1471;
        std::vector<int> input{ 1,2,3,4,5 };
        int k = 2;
        auto res = s1471.getStrongest(input, k);
        std::cout << "1471 res = " << res.size() << std::endl;
    }
    {
        Solution_1470 s1470;
        std::vector<int> input{ 1,2,3,4,4,3,2,1 };
        int n = 4;
        auto res = s1470.shuffle(input, n);
        std::cout << "1470 res = " << res.size() << std::endl;
    }
    {
        //start: "AAAAACCC"
        //end:   "AACCCCCC"
        //bank: ["AAAACCCC", "AAACCCCC", "AACCCCCC"]
        //
        //return: 3
        Solution_433 s433;
        std::string start = "AAAAACCC";
        std::string end = "AACCCCCC";
        std::vector<std::string> bank = {"AAAACCCC", "AAACCCCC", "AACCCCCC"};
        auto res = s433.minMutation2(start,end,bank);
        std::cout<<"433 res = "<<res<<std::endl;
    }
    {
        Solution_424 s424;
        std::string s = "ABAB";
        int k = 2;
        auto res = s424.characterReplacement(s,k);
        std::cout<<"424 res = "<<res<<std::endl;
    }
    {
        Solution_417 s417;
        vector<vector<int>> matrix;
//        matrix.push_back({1,2,2,3,5});
//        matrix.push_back({3,2,3,4,4});
//        matrix.push_back({2,4,5,3,1});
//        matrix.push_back({6,7,1,4,5});
//        matrix.push_back({5,1,1,2,4});
        matrix.push_back({3,3,3});
        matrix.push_back({3,1,3});
        matrix.push_back({0,2,4});
        //auto res = s417.pacificAtlantic(matrix);
        auto res = s417.pacificAtlantic2(matrix);
        std::cout << "417 res = " << res.size() << std::endl;
    }
    {
        Solution_399 s399;
        std::vector<std::vector<std::string>> equations;
        equations.push_back(std::vector<std::string>{"a", "b"});
        equations.push_back(std::vector<std::string>{"b", "c"});
        std::vector<double> values{2.0, 3.0};
        std::vector<std::vector<std::string>> queries;
        queries.push_back(std::vector<std::string>{"a", "c"});
        queries.push_back(std::vector<std::string>{"b", "a"});
        queries.push_back(std::vector<std::string>{"a", "e"});
        queries.push_back(std::vector<std::string>{"a", "a"});
        queries.push_back(std::vector<std::string>{"x", "x"});
        auto res = s399.calcEquation(equations,values,queries);
        std::cout << "399 res = " << res.size() << std::endl;
    }
    {
        Solution_397 s397;
        int n = 7;
        auto res = s397.integerReplacement(n);
        std::cout << "397 res = " << res << std::endl;
    }
    {
        Solution_396 s396;
        std::vector<int> input{4, 3, 2, 6};
        auto res = s396.maxRotateFunction(input);
        std::cout << "390 res = " << res << std::endl;
    }
    {
        Solution_390 s390;
        auto res = s390.lastRemaining(19);
        std::cout << "390 res = " << res << std::endl;
    }
    {
        Solution_386 s386;
        int k = 13;
        auto res = s386.lexicalOrder(k);
        std::cout << "386 res = " << res.size() << std::endl;
    }
    {
        Solution_1466 s1466;
        int n = 6;
        std::vector<std::vector<int>> input;
        input.push_back({ 0,1 });
        input.push_back({ 1,3 });
        input.push_back({ 2,3 });
        input.push_back({ 4,0 });
        input.push_back({ 4,5 });
        auto res = s1466.minReorder(n, input);
        std::cout << "5425 res = " << res << std::endl;
    }
    {
        Solution_1465 s1465;
        //int h = 975;
        //int w = 726;
        //std::vector<int> horizontalCuts{ 327,438,434,346,147,64,613,410,130,621,429,559,446,947,440,520,273,109,732,374,172,968,325,647,10,91,483,467,838,886,593,650,937,77,951,422,384,403,972,204,608,178,309,476,171,820,144,150,498,871,431,554,578,703,668,69,625,781,96,689,893,757,55,292,391,164,252,102,251,616,491,709,532,143,207,749,782,34,196,361,900,631,913,328,121,727,8,116,910,592,917,477,484,43,614,304,2,462,773,203,334,856,518,786,840,962,14,572,197,825,299,565,945,712,37,586,734,512,238,821,744,193,720,205,561,426,534,500,95,654,274,700,813,404,137,735,862,801,369,881,533,81,920,928,710,288,594,236,737,472,806,771,370,73,135,659,128,392,475,78,966,270,21,148,817,569,182,504,208,524,503,176,368,523,478,619,673,663,942,65,313,104,811,573,798,187,596,542,278,936,282,816,103,449,923,469,584,582,657,395,644,306,528,690,850,161,40,541,967,72,953,679,632,373,320,126,186,12,716,906,624,289,769,622,379,517,777,589,377,530,451,796,189,221,371,645,343,145,852,687,752,595,845,643,32,682,101,250,347,340,626,267,754,260,680,590,745,723,707,114,249,231,580,363,295,316,653,549,6,683,243,444,753,706,531,418,576,413,93,149,693,225,5,536,120,7,261,808,686,922,730,157,194,831,360,409,873,758,272,918,450,447,362,31,308,45,872,67,511,174,609,466,867,286,18,812,348,86,188,111,697,105,110,215,294,453,725,763,408,222,648,544,878,807,874,651,948,780,129,861,957,676,588,406,253,916,28,837,946,163,480,756,154,353,283,400,639,885,787,543,218,170,925,891,623,784,670,809,411,293,655,550,896,822,729,736,366,743,92,269,356,487,398,265,711,638,551,506,414,48,125,378,432,846,522,956,931,733,567,810,669,577,56,751,19,364,938,627,25,731,912,167,685,63,482,615,779,939,46,302,190,42,210,496,604,173,112,897,127,99,22,834,575,59,492,721,3,889,50,728,905,767,342,563,90,390,863,258,448,959,692,213,571,436,416,776,118,815,857,355,206,855,141,11,457,677,664,934,285,704,887,888,562,235,824,876,122,842,435,312,44,761,510,257,333,932,277,879,722,617,24,785,521,742,66,519,791,68,705,788,860,75,688,332,949,880,660,407,9,439,323,929,540,445,301,570,336,385,597,83,160,902,739,142,708,220,726,229,919,566,964,138,321,944,924,85,58,335,424,658,970,242,585,926,969,298,601,437,790,661,82,26,248,376,546,960,17,598,869,442,971,326,502,76,415,933,88,331,633,795,762,823,958,349,884,895,192,783,974,133,493,539,191,471,61,775,547,507,827,155,646,89,324,279,805,694,715,124,819,310,963,560,459,770,841,481 };
        //std::vector<int> verticalCuts{ 524,256,169,409,632,77,639,564,621,226,682,8,645,206,535,456,252,669,219,143,368,178,6,57,52,56,251,369,155,520,131,129,464,628,389,33,388,209,481,565,76,696,695,559,534,445,658,550,724,499,358,692,223,327,516,329,412,384,450,53,80,479,137,295,574,221,678,152,558,260,441,36,463,706,282,545,713,351,25,643,14,262,640,336,618,110,44,352,216,469,719,642,716,419,563,546,158,255,17,239,62,703,116,59,214,275,372,414,416,699,608,4,406,622,522,136,237,73,395,523,455,193,381,601,613,30,308,614,293,27,594,421,300,313,557,205,556,349,624,527,42,443,497,151,332,491,355,38,596,182,401,641,667,615,330,75,194,680,374,227,224,533,166,506,616,612,383,278,586,353,657,453,132,560,714,61,130,700,424,84,283,267,176,580,405,232,549,186,609,629,570,476,203,364,45,397,500,197,587,7,531,408,218,122,150,547,160,638,514,473,246,13,431,588,272,343,170,707,135,50,652,215,91,229,541,103,83,536,655,286,161,348,393,228,92,67,442,10,148,670,48,220,378,604,482,12,356,451,525,606,654,722,167,529,328,650,519,334,478,114,354,597,323,111,376,507,426,387,386,663,95,126,96,503,242,3,671,357,568,470,139,459,572,492,299,399,279,725,619,257,29,141,718,581,105,319,49,254,712,502,691,15,723,623,578,711,686,710,265,159,380,579,335,636,225,9,717,363,704,697,119,99,672,661,285,561,172,121,173,60,107,140,662,340,1,302,411,43,626,338,168,234,666,373,390,407,475,28,248,688,26,720,134 };
        int h = 5;
        int w = 4;
        std::vector<int> horizontalCuts{3};
        std::vector<int> verticalCuts{3};
        auto res = s1465.maxArea(h,w,horizontalCuts,verticalCuts);
        std::cout << "1465 res = " << res << std::endl;
    }
    {
        Solution_1464 s1464;
        std::vector<int> input{3,7};
        auto res = s1464.maxProduct(input);
        std::cout << "1464 res = " << res << std::endl;
    }
    {
//        Solution_1461 s1461;
//        std::string s = "00110110";
//        int k = 2;
//        auto res = s1461.hasAllCodes(s, k);
//        std::cout << "1461 res = " << res << std::endl;
    }
    {
        Solution_376 s376;
        std::vector<int> input{1,7,4,9,2,5};
        auto res = s376.wiggleMaxLength(input);
        std::cout<<"376 res = "<<res<<std::endl;
    }
    {
        Solution_373 s373;
        std::vector<int> nums1{1,2};
        std::vector<int> nums2{3};
        int k = 3;
        auto res = s373.kSmallestPairs(nums1,nums2,k);
        std::cout<<"373 res = "<<res.size()<<std::endl;
    }
    {
        //["Twitter","postTweet","follow","follow","getNewsFeed","postTweet","getNewsFeed","getNewsFeed","unfollow","getNewsFeed","getNewsFeed","unfollow","getNewsFeed","getNewsFeed"]
        //[[],[1,5],[1,2],[2,1],[2],[2,6],[1],[2],[2,1],[1],[2],[1,2],[1],[2]]
        Twitter twitter;
        twitter.postTweet(1, 5);

        twitter.follow(1, 2);
        twitter.follow(2, 1);

        auto res = twitter.getNewsFeed(1);

        twitter.postTweet(2, 6);
        res = twitter.getNewsFeed(1);
        res = twitter.getNewsFeed(2);
        twitter.unfollow(2, 1);
        res = twitter.getNewsFeed(1);
        res = twitter.getNewsFeed(2);
        twitter.unfollow(1, 2);
        res = twitter.getNewsFeed(1);
        res = twitter.getNewsFeed(2);
//        Twitter twitter;
//        twitter.postTweet(1, 1);
//        twitter.postTweet(1, 2);
//        twitter.postTweet(1, 3);
//        twitter.postTweet(1, 4);
//        twitter.postTweet(1, 5);
//        twitter.postTweet(1, 6);
//        twitter.postTweet(1, 7);
//        twitter.postTweet(1, 8);
//        twitter.postTweet(1, 9);
//        twitter.postTweet(1, 10);
//        twitter.postTweet(1, 11);
//        auto res = twitter.getNewsFeed(1);
        std::cout<<"355 res = "<<res.size()<<std::endl;
    }
    {
        Solution_310 s310;
        int n = 6;
        std::vector<std::vector<int>> input;
        input.push_back(std::vector<int>{0, 3});
        input.push_back(std::vector<int>{1, 3});
        input.push_back(std::vector<int>{2, 3});
        input.push_back(std::vector<int>{4, 3});
        input.push_back(std::vector<int>{5, 4});
        auto res = s310.findMinHeightTrees(n,input);
        std::cout<<"310 res = "<<res.size()<<std::endl;
    }
    {
        Solution_343 s343;
        auto res = s343.integerBreak(5);
        std::cout<<"343 res = "<<res<<std::endl;
    }
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
        std::vector<int> v{58,506,448,780,716,871,427,693,798,672,446,600,998,840,174,282,436,636,968,84,924,888,228,948,848,
                           553,377,245,396,774,650,336,986,55,1005,120,917,551,434,949,519,718,201,987,786,350,588,646,202,463,
                           261,144,405,290,183,690,517,823,306,517,823,456,767,566,201,900,845,860,17,295,310,929,663,261,914,
                           745,379,569,834,213,919,639,488,867,651,619,224,659,627,793,362,65,299,572,791,800,55,911,952,290,666,
                           934,316,123,327,22,337,88,265,973,708,209,533,782,283,747,1008,397,637,952,453,285,216,397,341,885,544,
                           767,223,258,477,371,174,486,328,755,955,593,490,760,786,834,80,2,82,467,385,375,246,51,197,635,702,979,
                           365,73,292,567,787,745,506,690,840,938,226,325,423,207,360,888,528,374,870,560,342,308,98,645,743,732,59,
                           191,132,694,562,183,645,323,966,971,13,392,389,313,188,471,363,650,993,534,503,322,181,940,793,727,462,
                           137,327,213,402,820,678,54,656,285,909,834,207,780,963,474,537,809,304,17,289,44,269,506,247,22,225,601,
                           696,937,273,103,374,501,131,49,178,932,790,908,154,329,467,996,567,486,977,679,374,969,703 };
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
//        Solution_241 s241;
//        std::string input = "2-1-1";
//        auto res = s241.diffWaysToCompute(input);
//        std::cout<<res.size()<<std::endl;
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
        std::vector<std::string> scores = {"FVSHJIEMNGYPTQOURLWCZKAX","AITFQORCEHPVJMXGKSLNZWUY","OTERVXFZUMHNIYSCQAWGPKJL",
                                           "VMSERIJYLZNWCPQTOKFUHAXG","VNHOZWKQCEFYPSGLAMXJIUTR","ANPHQIJMXCWOSKTYGULFVERZ",
                                           "RFYUXJEWCKQOMGATHZVILNSP","SCPYUMQJTVEXKRNLIOWGHAFZ","VIKTSJCEYQGLOMPZWAHFXURN",
                                           "SVJICLXKHQZTFWNPYRGMEUAO","JRCTHYKIGSXPOZLUQAVNEWFM","NGMSWJITREHFZVQCUKXYAPOL",
                                           "WUXJOQKGNSYLHEZAFIPMRCVT","PKYQIOLXFCRGHZNAMJVUTWES","FERSGNMJVZXWAYLIKCPUQHTO",
                                           "HPLRIUQMTSGYJVAXWNOCZEKF","JUVWPTEGCOFYSKXNRMHQALIZ","MWPIAZCNSLEYRTHFKQXUOVGJ",
                                           "EZXLUNFVCMORSIWKTYHJAQPG","HRQNLTKJFIEGMCSXAZPYOVUW","LOHXVYGWRIJMCPSQENUAKTZF",
                                           "XKUTWPRGHOAQFLVYMJSNEIZC","WTCRQMVKPHOSLGAXZUEFYNJI"};
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