#include <iostream>
#include <set>
#include <thread>
#include <condition_variable>
#include "define.h"
#include "array/239. Sliding Window Maximum.hpp"
#include "array/396. Rotate Function.hpp"
#include "array/417. Pacific Atlantic Water Flow.hpp"
#include "array/1470. Shuffle the Array.hpp"
#include "array/1471. The k Strongest Values in an Array.hpp"
#include "array/1465. Maximum Area of a Piece of Cake After Horizontal and Vertical Cuts.hpp"
#include "array/1481. Least Number of Unique Integers after K Removals.hpp"
#include "array/1477. Find Two Non-overlapping Sub-arrays Each With Target Sum.hpp"
#include "array/435. Non-overlapping Intervals.hpp"
#include "array/436. Find Right Interval.hpp"
#include "array/1493. Longest Subarray of 1's After Deleting One Element.hpp"
#include "array/1497. Check If Array Pairs Are Divisible by k.hpp"
#include "array/456. 132 Pattern.hpp"
#include "array/498. Diagonal Traverse.hpp"
#include "array/525. Contiguous Array.hpp"
#include "array/846. Hand of Straights.hpp"
#include "array/983. Minimum Cost For Tickets.hpp"
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
#include "array/1626. Best Team With No Conflicts.hpp"
#include "array/1673. Find the Most Competitive Subsequence.hpp"
#include "array/1562. Find Latest Group of Size M.hpp"
#include "array/315. Count of Smaller Numbers After Self.hpp"
#include "array/1365. How Many Numbers Are Smaller Than the Current Number.hpp"
#include "array/324. Wiggle Sort II.hpp"
#include "array/84. Largest Rectangle in Histogram.hpp"
#include "array/1738. Find Kth Largest XOR Coordinate Value.hpp"
#include "array/1508. Range Sum of Sorted Subarray Sums.hpp"
#include "array/322. Coin Change.hpp"
#include "array/1743. Restore the Array From Adjacent Pairs.hpp"
#include "array/1546. Maximum Number of Non-Overlapping Subarrays With Sum Equals Target.hpp"
#include "array/1705. Maximum Number of Eaten Apples.hpp"
#include "array/37. Sudoku Solver.hpp"
#include "array/4. Median of Two Sorted Arrays.hpp"
#include "array/688. Knight Probability in Chessboard.hpp"
#include "array/1769. Minimum Number of Operations to Move All Balls to Each Box.hpp"
#include "array/1770. Maximum Score from Performing Multiplication Operations.hpp"
#include "array/1775. Equal Sum Arrays With Minimum Number of Operations.hpp"
#include "array/350. Intersection of Two Arrays II.hpp"
#include "array/406. Queue Reconstruction by Height.hpp"
#include "array/407. Trapping Rain Water II.hpp"
#include "array/410. Split Array Largest Sum.hpp"
#include "array/1806. Minimum Number of Operations to Reinitialize a Permutation.hpp"
#include "array/1812. Determine Color of a Chessboard Square.hpp"
#include "array/1817. Finding the Users Active Minutes.hpp"
#include "array/1818. Minimum Absolute Sum Difference.hpp"
#include "array/480. Sliding Window Median.hpp"
#include "array/493. Reverse Pairs.hpp"
#include "number/229. Majority Element II.hpp"
#include "number/220. Contains Duplicate III.hpp"
#include "number/1441. Build an Array With Stack Operations.hpp"
#include "number/1442. Count Triplets That Can Form Two Arrays of Equal XOR.hpp"
#include "number/1447. Simplified Fractions.hpp"
#include "number/368. Largest Divisible Subset.hpp"
#include "number/343. Integer Break.hpp"
#include "number/373. Find K Pairs with Smallest Sums.hpp"
#include "number/376. Wiggle Subsequence.hpp"
#include "number/386. Lexicographical Numbers.hpp"
#include "number/390. Elimination Game.hpp"
#include "number/397. Integer Replacement.hpp"
#include "number/1464. Maximum Product of Two Elements in an Array.hpp"
#include "number/1475. Final Prices With a Special Discount in a Shop.hpp"
#include "number/1492. The kth Factor of n.hpp"
#include "number/1498. Number of Subsequences That Satisfy the Given Sum Condition.hpp"
#include "number/464. Can I Win.hpp"
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
#include "number/219. Contains Duplicate II.hpp"
#include "number/31. Next Permutation.hpp"
#include "number/166. Fraction to Recurring Decimal.hpp"
#include "number/241. Different Ways to Add Parentheses.hpp"
#include "number/659. Split Array into Consecutive Subsequences.hpp"
#include "number/781. Rabbits in Forest.hpp"
#include "number/1619. Mean of Array After Removing Some Elements.hpp"
#include "number/853. Car Fleet.hpp"
#include "number/295. Find Median from Data Stream.hpp"
#include "number/1647. Minimum Deletions to Make Character Frequencies Unique.hpp"
#include "number/612. Shortest Distance in a Plane.hpp"
#include "number/1734. Decode XORed Permutation.hpp"
#include "number/1736. Latest Time by Replacing Hidden Digits.hpp"
#include "number/477. Total Hamming Distance.hpp"
#include "number/60. Permutation Sequence.hpp"
#include "number/357. Count Numbers with Unique Digits.hpp"
#include "number/486. Predict the Winner.hpp"
#include "number/1742. Maximum Number of Balls in a Box.hpp"
#include "number/1760. Minimum Limit of Balls in a Bag.hpp"
#include "number/1774. Closest Dessert Cost.hpp"
#include "number/1792. Maximum Average Pass Ratio.hpp"
#include "number/347. Top K Frequent Elements.hpp"
#include "number/363. Max Sum of Rectangle No Larger Than K.hpp"
#include "number/414. Third Maximum Number.hpp"
#include "number/378. Kth Smallest Element in a Sorted Matrix.hpp"
#include "number/1823. Find the Winner of the Circular Game.hpp"
#include "number/518. Coin Change 2.hpp"
#include "number/575. Distribute Candies.hpp"
#include "number/1834. Single-Threaded CPU.hpp"
#include "number/1838. Frequency of the Most Frequent Element.hpp"
#include "number/629. K Inverse Pairs Array.hpp"
#include "number/632. Smallest Range Covering Elements from K Lists.hpp"
#include "number/1850. Minimum Adjacent Swaps to Reach the Kth Smallest Number.hpp"
#include "number/667. Beautiful Arrangement II.hpp"
#include "number/672. Bulb Switcher II.hpp"
#include "number/1854. Maximum Population Year.hpp"
#include "tree/1443. Minimum Time to Collect All Apples in a Tree.cpp"
#include "tree/1457. Pseudo-Palindromic Paths in a Binary Tree.hpp"
#include "tree/450. Delete Node in a BST.hpp"
#include "tree/449. Serialize and Deserialize BST.hpp"
#include "tree/987. Vertical Order Traversal of a Binary Tree.hpp"
#include "tree/988. Smallest String Starting From Leaf.hpp"
#include "tree/297. Serialize and Deserialize Binary Tree.hpp"
#include "graph/207. Course Schedule.hpp"
#include "graph/210. Course Schedule II.hpp"
#include "graph/1444. Number of Ways of Cutting a Pizza.hpp"
#include "graph/310. Minimum Height Trees.hpp"
#include "graph/355. Design Twitter.hpp"
#include "graph/1466. Reorder Routes to Make All Paths Lead to the City Zero.hpp"
#include "graph/399. Evaluate Division.hpp"
#include "graph/433. Minimum Genetic Mutation.hpp"
#include "graph/1472. Design Browser History.hpp"
#include "graph/1476. Subrectangle Queries.hpp"
#include "graph/1496. Path Crossing.hpp"
#include "graph/452. Minimum Number of Arrows to Burst Balloons.hpp"
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
#include "graph/1584. Min Cost to Connect All Points.hpp"
#include "graph/1727. Largest Submatrix With Rearrangements.hpp"
#include "graph/51. N-Queens.hpp"
#include "graph/52. N-Queens II.hpp"
#include "graph/1366. Rank Teams by Votes.hpp"
#include "graph/743. Network Delay Time.hpp"
#include "graph/787. Cheapest Flights Within K Stops.hpp"
#include "graph/684. Redundant Connection.hpp"
#include "graph/332. Reconstruct Itinerary.hpp"
#include "graph/502. IPO.hpp"
#include "graph/514. Freedom Trail.hpp"
#include "graph/1824. Minimum Sideway Jumps.hpp"
#include "graph/497. Random Point in Non-overlapping Rectangles.hpp"
#include "graph/630. Course Schedule III.hpp"
#include "graph/675. Cut Off Trees for Golf Event.hpp"
#include "graph/699. Falling Squares.hpp"
#include "string/1446. Consecutive Characters.cpp"
#include "string/1451. Rearrange Words in a Sentence.hpp"
#include "string/1455. Check If a Word Occurs As a Prefix of Any Word in a Sentence.hpp"
#include "string/1456. Maximum Number of Vowels in a Substring of Given Length.hpp"
#include "string/424. Longest Repeating Character Replacement.hpp"
#include "string/474. Ones and Zeroes.hpp"
#include "string/482. License Key Formatting.hpp"
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
#include "string/942. DI String Match.hpp"
#include "string/1079. Letter Tile Possibilities.hpp"
#include "string/535. Encode and Decode TinyURL.hpp"
#include "string/1419. Minimum Number of Frogs Croaking.hpp"
#include "string/567. Permutation in String.hpp"
#include "string/44. Wildcard Matching.hpp"
#include "string/3. Longest Substring Without Repeating Characters.hpp"
#include "string/316. Remove Duplicate Letters.hpp"
#include "string/97. Interleaving String.hpp"
#include "string/752. Open the Lock.hpp"
#include "string/1768. Merge Strings Alternately.hpp"
#include "string/242. Valid Anagram.hpp"
#include "string/516. Longest Palindromic Subsequence.hpp"
#include "string/1839. Longest Substring Of All Vowels in Order.hpp"
#include "string/726. Number of Atoms.hpp"
#include "list/1483. Kth Ancestor of a Tree Node.hpp"
#include "list/445. Add Two Numbers II.hpp"
#include "list/138. Copy List with Random Pointer.hpp"
#include "list/25. Reverse Nodes in k-Group.hpp"
#include "list/729. My Calendar I.hpp"
#include "list/731. My Calendar II.hpp"
#include "list/146. LRU Cache.hpp"
#include "list/206. Reverse Linked List.hpp"
#include "list/352. Data Stream as Disjoint Intervals.hpp"
#include "list/460. LFU Cache.hpp"
#include "list/432. All O`one Data Structure.hpp"
#include "list/503. Next Greater Element II.hpp"
#include "list/649. Dota2 Senate.hpp"
#include "list/705. Design HashSet.hpp"
#include "list/706. Design HashMap.hpp"
#include "list/707. Design Linked List.hpp"
#include "list/710. Random Pick with Blacklist.hpp"
#include "list/715. Range Module.hpp"
#include "thread/1115. Print FooBar Alternately.hpp"
#include "thread/1116. Print Zero Even Odd.hpp"

int main() {
    {
        Solution_743 s743;
        std::vector<std::vector<int>> times{{1,2,1},{2,3,2},{1,3,4}};
        int n = 3;
        int k = 1;
        auto res = s743.networkDelayTime(times,n,k);
        std::cout << "743 res = " << res << std::endl;
    }
    {
        Solution_726 s726;
        std::string formula = "K4(ON(SO3)2)2";
        auto res = s726.countOfAtoms(formula);
        std::cout << "726 res = " << res << std::endl;
    }
    {
        //addRange(10, 20): null
        //removeRange(14, 16): null
        //queryRange(10, 14): true (Every number in [10, 14) is being tracked)
        //queryRange(13, 15): false (Numbers like 14, 14.03, 14.17 in [13, 15) are not being tracked)
        //queryRange(16, 17): true (The number 16 in [16, 17) is still being tracked, despite the remove operation)

//        RangeModule s715;
//        s715.addRange(10,20);
//        s715.addRange(1,2);
//        s715.removeRange(14,16);
//        auto res = s715.queryRange(10,14);
//        res = s715.queryRange(13,15);
//        res = s715.queryRange(16,17);
//        std::cout << "715 res = " << res << std::endl;
    }
    {
        int n = 5;
        std::vector<int> blacklist{2,1,0};
        Solution_710 s710(n,blacklist);
        auto res = s710.pick();
        std::cout << "710 res = " << res << std::endl;
    }
    {
//        MyLinkedList* myLinkedList = new MyLinkedList();
//        myLinkedList->addAtHead(1);
//        myLinkedList->addAtTail(3);
//        myLinkedList->addAtIndex(1, 2);    // linked list becomes 1->2->3
//        auto res = myLinkedList->get(1);              // return 2
//        myLinkedList->deleteAtIndex(1);    // now the linked list is 1->3
//        res = myLinkedList->get(1);              // return 3

//        MyLinkedList* myLinkedList = new MyLinkedList();
//        myLinkedList->addAtIndex(0,10);
//        myLinkedList->addAtIndex(0,20);
//        myLinkedList->addAtIndex(1, 30);    // linked list becomes 1->2->3
//        auto res = myLinkedList->get(0);              // return 2         // return 3

//        MyLinkedList* myLinkedList = new MyLinkedList();
//        myLinkedList->addAtHead(0);
//        myLinkedList->addAtIndex(1,4);
//        myLinkedList->addAtTail(8);
//        myLinkedList->addAtHead(5);
//        myLinkedList->addAtIndex(4,3);
//        myLinkedList->addAtTail(0);
//        myLinkedList->addAtTail(5);
//        myLinkedList->addAtIndex(6,3);
//        myLinkedList->deleteAtIndex(7);
//        myLinkedList->deleteAtIndex(5);
//        myLinkedList->addAtTail(4);    // linked list becomes 1->2->3

        //["MyLinkedList","addAtHead","addAtIndex","get"]
        //[[],[2],[0,1],[1]]
        MyLinkedList* myLinkedList = new MyLinkedList();
        myLinkedList->addAtHead(2);
        myLinkedList->addAtIndex(0,1);
        auto res = myLinkedList->get(1);
    }
    {
        MyHashMap s706;
        s706.put(1, 1); // The map is now [[1,1]]
        s706.put(2, 2); // The map is now [[1,1], [2,2]]
        auto res = s706.get(1);    // return 1, The map is now [[1,1], [2,2]]
        res = s706.get(3);
    }
    {
        MyHashSet s705;
        s705.add(1);      // set = [1]
        s705.add(2);      // set = [1, 2]
        auto res = s705.contains(1); // return True
        res = s705.contains(3); // return False, (not found)
        s705.add(2);      // set = [1, 2]
    }
    {
        Solution_699 s699;
        std::vector<std::vector<int>> positions{{9,7},{1,9},{3,1}};
        auto res = s699.fallingSquares(positions);
        std::cout << "699 res = " << res[0] << std::endl;
    }
    {
        Solution_684 s684;
        //std::vector<std::vector<int>> edges{{9,10},{5,8},{2,6},{1,5},{3,8},{4,9},{8,10},{4,10},{6,8},{7,9}};
        std::vector<std::vector<int>> edges{{1,5},{3,4},{3,5},{4,5},{2,4}};
        auto res = s684.findRedundantConnection(edges);
        std::cout << "684 res = " << res[0] << std::endl;
    }
    {
        Solution_1854 s1854;
        std::vector<std::vector<int>> logs{{1993,1999},{2000,2010}};
        auto res = s1854.maximumPopulation(logs);
        std::cout<<"1854 res = "<<res<<std::endl;
    }
    {
        Solution_675 S675;
        std::vector<std::vector<int>> forest{{1,2,3},{0,1,0},{7,6,5}};
        auto res = S675.cutOffTree(forest);
        std::cout<<"675 res = "<<res<<std::endl;
    }
    {
        Solution_672 s672;
        int n = 1000;
        int presses = 1000;
        auto res = s672.flipLights(n,presses);
        std::cout<<"672 res = "<<res<<std::endl;
    }
    {
        Solution_667 s667;
        int n = 10;
        int k = 5;
        auto res = s667.constructArray(n,k);
        std::cout<<"667 res = "<<res.size()<<std::endl;
    }
    {
        Solution_661 s661;
        std::vector<std::vector<int>> input{{1,1,1},{1,0,1},{1,1,1}};
        auto res = s661.imageSmoother(input);
        std::cout<<"661 res = "<<res.size()<<std::endl;
    }
    {
        Solution_1850 s1850;
        std::string num = "5489355142";
        int k = 4;
        auto res = s1850.getMinSwaps(num,k);
        std::cout<<"1850 res = "<<res<<std::endl;
    }
    {
        Solution_649 s649;
        std::string senate = "DRRDRDRDRDDRDRDR";
        auto res = s649.predictPartyVictory(senate);
        std::cout<<"649 res = "<<res<<std::endl;
    }
    {
        Solution_632 s632;
        std::vector<std::vector<int>> nums{{4,10,15,24,26},{0,9,12,20},{5,18,22,30}};
        auto res = s632.smallestRange(nums);
        std::cout<<"632 res = "<<res[0]<<std::endl;
    }
    {
        Solution_630 s630;
        //[[100, 200], [200, 1300], [1000, 1250], [2000, 3200]]
        std::vector<std::vector<int>> courses{{2000, 3200},{100, 200}, {200, 1300}, {1000, 1250}};
        auto res = s630.scheduleCourse(courses);
        std::cout<<"630 res = "<<res<<std::endl;
    }
    {
        Solution_629 s629;
        int n = 10;
        int k = 8;
        auto res = s629.kInversePairs(n,k);
        std::cout<<"629 res = "<<res<<std::endl;
    }
    {
        Solution_1839 s1839;
        std::string word = "aeiaaioaaaaeiiiiouuuooaauuaeiu";
        auto res = s1839.longestBeautifulSubstring(word);
        std::cout<<"1839 res = "<<res<<std::endl;
    }
    {
        Solution_1838 s1838;
        std::vector<int> nums{1,4,4,4,4,4,4,13};
        int k = 5;
        auto res = s1838.maxFrequency(nums,k);
        std::cout<<"1838 res = "<<res<<std::endl;
    }
    {
        Solution_1834 s1834;
        std::vector<std::vector<int>> tasks{{7,10},{7,12},{7,5},{7,4},{7,2}};
        auto res = s1834.getOrder(tasks);
        std::cout<<"1834 res = "<<res.size()<<std::endl;
    }
    {
        Solution_575 s575;
        std::vector<int> candyType{1,1,2,3};
        auto res = s575.distributeCandies(candyType);
        std::cout<<"575 res = "<<res<<std::endl;
    }
    {
        Solution_518 s518;
        int amount = 5;
        std::vector<int> coins{1,2,5};
        auto res = s518.change(amount,coins);
        std::cout<<"518 res = "<<res<<std::endl;
    }
    {
        Solution_516 s516;
        std::string s = "bbbab";
        auto res = s516.longestPalindromeSubseq(s);
        std::cout<<"516 res = "<<res<<std::endl;
    }
    {
        Solution_507 s507;
        int num = 6;
        auto res = s507.checkPerfectNumber(num);
        std::cout<<"507 res = "<<res<<std::endl;
    }
    {
        Solution_503 s503;
        std::vector<int> nums{1,2,3,4,3};
        auto res = s503.nextGreaterElements(nums);
        std::cout<<"503 res = "<<res[0]<<std::endl;
    }
    {
        std::vector<std::vector<int>> rects{{-2,-2,-1,-1},{1,0,3,0}};
        Solution_497 s497(rects);
        auto res = s497.pick();
        std::cout<<"497 res = "<<res[0]<<std::endl;
    }
    {
        Solution_1824 s1824;
        std::vector<int> obstacles{0,2,2,1,0,3,0,3,0,1,3,1,1,0,1,3,1,1,1,0,2,0,0,3,3,0,3,2,2,0,0,3,3,3,0,0,2,0,0,3,3,0,3,3,0,0,3,1,0,1,0,2,3,1,1,0,3,3,0,3,1,3,0,2,2,0,1,3,0,1,0,3,0,1,3,1,2,2,0,0,3,0,1,3,2,3,2,1,0,3,2,2,0,3,3,0,3,0,0,1,0,2,0,0,0,2,1,2,0,2,2,3,3,3,0,0,1,1,3,0,0,0,1,2,2,1,2,1,3,2,2,3,1,3,0,1,1,1,3,0,0,0,2,0,2,0,3,1,2,3,3,2,2,2,0,0,0,1,0,0,0,0,3,0,0,0,3,0,2,1,2,3,1,0,3,3,2,0,1,1,0,1,0,2,2,2,1,3,0,3,0,2,1,1,3,1,0,1,2,2,0,2,2,0,0,3,3,1,3,0,1,1,0,3,0,2,1,2,2,0,0,0,1,2,3,1,2,1,1,2,2,1,1,0,2,3,3,3,0,2,3,2,0,0,0,1,0,2,2,0,0,2,0,2,0,1,1,0,3,1,3,3,0,1,0,3,0,3,1,2,3,1,0,0,2,3,2,0,0,3,1,2,3,2,2,3,1,3,3,2,0,1,3,0,3,2,2,3,2,1,2,2,0,3,2,0,2,1,2,2,3,1,3,2,2,0,0,1,0,3,1,3,3,0,0,2,2,2,2,0,1,0,3,1,3,3,3,0,2,3,2,0,3,3,3,3,3,3,2,2,1,1,0,3,1,3,2,3,0,0,0,2,1,1,3,1,3,2,1,3,0,1,1,3,2,2,1,0,0,3,2,1,3,2,3,3,2,1,2,0,2,2,0,2,2,3,2,0,2,3,3,1,1,2,0,1,1,1,2,3,2,1,2,1,0,2,3,1,1,3,3,2,0,1,3,2,3,3,0,1,2,3,2,1,1,2,1,0,0,1,0,3,1,1,1,0,2,0,2,2,3,0,1,0,2,0,0,3,1,1,2,0,0,2,1,1,0,2,2,2,3,1,2,0,1,2,0,1,2,1,2,3,1,1,1,1,0,3,3,2,1,0,0,1,0,3,0,0,2,2,2,1,1,2,1,2,1,1,2,0,3,1,3,2,1,2,2,3,1,0,1,1,1,0,0,0,1,3,3,2,2,1,2,0,0,0,3,1,3,2,3,1,3,2,3,1,3,2,0,1,2,1,1,2,1,3,0,1,1,1,3,3,1,0,0,3,2,2,3,1,1,0,3,0,0,3,0,3,1,2,0,2,3,2,3,0,3,2,3,0,2,2,3,0,3,3,3,1,0,1,2,2,0,3,3,1,3,2,2,3,2,1,1,0,0,0,0,2,1,0,1,1,1,1,0,3,0,1,0,0,1,0,2,0,0,1,2,0,0,0,3,3,1,0,3,2,1,2,3,2,0,3,3,0,2,3,1,1,0,2,2,3,3,0,1,0,0,3,1,2,3,0,1,2,3,2,2,0,1,2,0,3,0,3,0,1,1,3,2,2,2,3,1,2,0,0,3,0,2,3,3,1,0,3,3,0,0,0,3,2,1,1,3,1,1,1,1,1,1,3,3,3,0,0,1,1,1,1,2,2,0,1,0,2,2,0,2,1,3,1,1,1,2,1,1,0,3,1,0,2,3,0,1,2,0,0,3,1,2,3,0,0,3,1,0,2,2,0,1,1,2,2,1,3,1,2,1,0,1,2,3,2,3,0,3,1,3,0,2,0,3,1,1,0,3,2,0,3,0,2,0,0,3,3,1,1,1,0,0,1,1,1,2,3,1,3,1,2,0,0,3,3,0,3,3,0,0,0,3,3,0,3,3,2,3,3,3,3,1,1,1,3,1,1,3,3,1,0,0,3,1,2,0,2,0,3,0,2,1,0,1,0,2,3,3,3,2,3,3,2,0,0,0,2,2,3,0,0,3,0,2,3,0,1,3,2,1,2,0,1,3,2,2,0,1,1,3,3,0,2,3,0,3,3,1,2,3,2,1,0,2,3,2,2,2,3,0,1,1,3,1,0,2,1,3,2,2,2,3,3,1,1,1,3,2,3,1,0,2,3,0,2,3,0,1,3,3,1,1,1,1,0,1,1,2,2,0,2,1,1,0,1,0,3,1,1,1,3,3,2,1,2,3,2,2,3,1,0,3,2,0,1,0,1,3,3,3,0,3,3,2,3,1,2,2,1,1,0,0,3,0};
        auto res = s1824.minSideJumps(obstacles);
        std::cout<<"1824 res = "<<res<<std::endl;
    }
    {
        Solution_1823 s1823;
        int n = 6;
        int k = 5;
        auto res = s1823.findTheWinner(n,k);
        std::cout<<"1823 res = "<<res<<std::endl;
    }
    {
        Solution_514 s514;
        std::string ring = "godding";
        std::string key = "gd";
        auto res = s514.findRotateSteps(ring,key);
        std::cout<<"502 res = "<<res<<std::endl;
    }
    {
        Solution_502 s502;
        int k = 1;
        int W = 0;
        std::vector<int> Profits{1,2,3};
        std::vector<int> Capital{1,1,2};
        auto res = s502.findMaximizedCapital(k,W,Profits,Capital);
        std::cout<<"502 res = "<<res<<std::endl;
    }
    {
        Solution_493 s493;
        std::vector<int> nums{1,3,2,3,1};
        auto res = s493.reversePairs(nums);
        std::cout<<"493 res = "<<res<<std::endl;
    }
    {
        Solution_480 s480;
        std::vector<int> nums{5,5,8,1,4,7,1,3,8,4};
        int k = 8;
        auto res = s480.medianSlidingWindow(nums,k);
        std::cout<<"480 res = "<<res.size()<<std::endl;
    }
    {
        Solution_378 s378;
//        std::vector<std::vector<int>> matrix{{1,5,9},{10,11,13},{12,13,15}};
//        int k = 8;
        std::vector<std::vector<int>> matrix{{-5}};
        int k = 1;
        auto res = s378.kthSmallest(matrix,k);
        std::cout<<"378 res = "<<res<<std::endl;
    }
    {
        Solution_1818 s1818;
        std::vector<int> nums1{1,10,4,4,2,7};
        std::vector<int> nums2{9,3,5,1,7,4};
        auto res = s1818.minAbsoluteSumDiff(nums1,nums2);
        std::cout<<"5724 res = "<<res<<std::endl;
    }
    {
        Solution_1817 s1817;
        std::vector<std::vector<int>> logs{{0,5},{1,2},{0,2},{0,5},{1,3}};
        int k = 5;
        auto res = s1817.findingUsersActiveMinutes(logs,k);
        std::cout<<"5732 res = "<<res.size()<<std::endl;
    }
    {
        Solution_1812 s1812;
        std::string s = "h3";
        auto res = s1812.squareIsWhite(s);
        std::cout<<"1812 res = "<<res<<std::endl;
    }
    {
        AllOne s432;
        s432.inc("a");
        s432.inc("a");
        s432.inc("b");
        auto res = s432.getMaxKey();
        res = s432.getMinKey();
        s432.dec("b");
        res = s432.getMinKey();
        std::cout<<"432 res = "<<res<<std::endl;
    }
    {
        //["LFUCache", "put", "put", "get", "put", "get", "get", "put", "get", "get", "get"]
        //[[2], [1, 1], [2, 2], [1], [3, 3], [2], [3], [4, 4], [1], [3], [4]]
        //Output
        //[null, null, null, 1, null, -1, 3, null, -1, 3, 4]
        LFUCache lfu(2);
        lfu.put(1,1);
        lfu.put(2,2);
        auto res = lfu.get(1);
        lfu.put(3,3);
        res = lfu.get(2);
        res = lfu.get(3);
        lfu.put(4,4);
        res = lfu.get(1);
        res = lfu.get(3);
        res = lfu.get(4);
        std::cout<<"460 res = "<<res<<std::endl;
    }
    {
        Solution_1806 s1806;
        int n = 4;
        auto res = s1806.reinitializePermutation(n);
        std::cout<<"1806 res = "<<res<<std::endl;
    }
    {
        std::vector<int> nums{1,2,3,4,5};
        int m = 2;
        Solution_410 s410;
        auto res = s410.splitArray(nums,m);
        std::cout<<"410 res = "<<res<<std::endl;
    }
    {
        Solution_414 s414;
        std::vector<int> nums{3,2,2};
        auto res = s414.thirdMax(nums);
        std::cout<<"414 res = "<<res<<std::endl;
    }
    {
        Solution_407 s407;
        std::vector<std::vector<int>> heightMap{{1,4,3,1,3,2},{3,2,1,3,2,4},{2,3,3,2,3,1}};
        auto res = s407.trapRainWater(heightMap);
        std::cout<<"407 res = "<<res<<std::endl;
    }
    {
        Solution_406 s406;
        std::vector<std::vector<int>> people{{7,0},{4,4},{7,1},{5,0},{6,1},{5,2}};
        auto res = s406.reconstructQueue(people);
        std::cout<<"406 res = "<<res.size()<<std::endl;
    }
    {
        //nums1 = [1,7,11], nums2 = [2,4,6]
        Solution_373 s373;
        std::vector<int> nums1{1,7,11};
        std::vector<int> nums2{2,4,6};
        int k = 3;
        auto res = s373.kSmallestPairs(nums1,nums2,k);
        std::cout<<"373 res = "<<res.size()<<std::endl;
    }
    {
        Solution_363 s363;
        //std::vector<std::vector<int>> matrix{{1,0,1},{0,-2,3}};
        std::vector<std::vector<int>> matrix{{2,2,-1}};
        int k = 5;
        auto res = s363.maxSumSubmatrix(matrix,k);
        std::cout << "363 res = "<< res <<std::endl;
    }
    {
        SummaryRanges s352;
        s352.addNum(1);
        s352.addNum(3);
        s352.addNum(7);
        s352.addNum(2);
        s352.addNum(6);
        auto res = s352.getIntervals();
        std::cout << "352 res = "<< res.size() <<std::endl;
    }
    {
        Solution_350 s350;
        //[4,9,5]
        //[9,4,9,8,4]
        std::vector<int> nums1{4,9,5};
        std::vector<int> nums2{9,4,9,8,4};
        auto res = s350.intersect(nums1,nums2);
        std::cout << "350 res = "<< res.size() <<std::endl;
    }
    {
        Solution_347 s347;
        std::vector<int> nums{1,1,1,2,2,3};
        int k = 2;
        auto res = s347.topKFrequent(nums,k);
        std::cout << "347 res = "<< res.size() <<std::endl;
    }
    {
        //        //[["JFK","SFO"],["JFK","ATL"],["SFO","ATL"],["ATL","JFK"],["ATL","SFO"]]
        //        //[["JFK","KUL"],["JFK","NRT"],["NRT","JFK"]]
        Solution_332 s332;
        //std::vector<std::vector<std::string>> input = {std::vector<std::string>{"JFK","KUL"}, std::vector<std::string>{"JFK","NRT"},

        //std::vector<std::vector<std::string>> input{{"JFK","SFO"},{"JFK","ATL"},{"SFO","ATL"},{"ATL","JFK"},{"ATL","SFO"}};
        std::vector<std::vector<std::string>> input{{"MUC","LHR"},{"JFK","MUC"},{"SFO","SJC"},{"LHR","SFO"}};
        auto res = s332.findItinerary(input);
        std::cout << "332 res = "<< res.size() <<std::endl;
    }
    {
        Solution_1792 s1792;
        std::vector<std::vector<int>> classes{{2,4},{3,9},{4,5},{2,10}};
        int extraStudents = 4;
        auto res = s1792.maxAverageRatio(classes,extraStudents);
        std::cout << "1792 res = " << res << std::endl;
    }
    {
        std::string s = "anagram";
        std::string t = "nagaram";
        Solution_242 s242;
        auto res = s242.isAnagram(s,t);
        std::cout << "242 res = " << res << std::endl;
    }
//    {
//        Solution5699 s5699;
//        int n = 5;
//        std::vector<std::vector<int>> edges{ {1,2,3},{1,3,3},{2,3,1},{1,4,2},{5,2,2},{3,5,1},{5,4,10} };
//        auto res = s5699.countRestrictedPaths(n, edges);
//        std::cout << "5699 res = " << res << std::endl;
//    }
    {
        Solution_239 s239;
        std::vector<int> nums{1,3,-1,-3,5,3,6,7};
        int k = 3;
        auto res = s239.maxSlidingWindow(nums,k);
        std::cout << "239 res = " << res[0] << std::endl;
    }
    {
        std::vector<int> nums{1,0,1,1};
        int k = 1;
        Solution_219 s219;
        auto res = s219.containsNearbyDuplicate(nums,k);
        std::cout << "219 res = " << res << std::endl;
    }
    {
        Solution206 s206;
        ListNode l1(1);
        ListNode l2(2);
        ListNode l3(3);
        l1.next = &l2;
        l2.next = &l3;
        auto res = s206.reverseList(&l1);
        std::cout << "206 res = " << res->val << std::endl;
    }
    {
        Solution_1775 s1775;
        std::vector<int> n1{ 1};
        std::vector<int> n2{ 6,6};
        auto res = s1775.minOperations(n1,n2);
        std::cout << "1775 res = " << res << std::endl;
    }
    {
        Solution_1774 s1774;
        std::vector<int> baseCosts{ 3,10 };
        std::vector<int> toppingCosts{ 2,5};
        int target = 9;
        auto res = s1774.closestCost(baseCosts, toppingCosts, target);
        std::cout << "1774 res = " << res << std::endl;
    }
    {
        //[[2],[1,1],[2,2],[1],[3,3],[2],[4,4],[1],[3],[4]]
        LRUCache lRUCache(2);
        lRUCache.put(1, 1);
        lRUCache.put(2, 2);
        auto res = lRUCache.get(1);
        lRUCache.put(3, 3);
        res = lRUCache.get(2);
        lRUCache.put(4, 4);
        res = lRUCache.get(1);
        res = lRUCache.get(3);
        res = lRUCache.get(4);
        //lRUCache.put(3, 3);
        //res = lRUCache.get(2);
        std::cout << "146 res = " << res << std::endl;
    }
    {
        Solution_1770 s1770;
        std::vector<int> nums{ -5, -3, -3, -2, 7, 1 };
        std::vector<int> multipliers{-10, -5, 3, 4, 6};
        //std::vector<int> nums{ 822, 963, 170, 305, 258, 841, 204, 748, 850, 721, 824, 405, 108, 843, 76, 78, 848, 318, 105, 530, 703, 38, 875, 785, 223, 733, 310, 402, 450, 447, 712, 253, 369, 189, 232, 478, 990, 681, 847, 137, 706, 789, 241, 457, 351, 91, 445, 823, 780, 555, 144, 866, 286, 808, 147, 334, 818, 751, 286, 60, 90, 88, 53, 850, 513, 694, 727, 100, 335, 137, 92, 37, 806, 963, 566, 948, 193, 838, 299, 299, 655, 777, 50, 353, 681, 65, 47, 768, 562, 854, 553, 939, 334, 419, 570, 704, 601, 735, 123, 602, 621, 910, 821, 327, 526, 613, 455, 979, 32, 876, 974, 359, 765, 248, 313, 380, 127, 382, 551, 987, 720, 753, 382, 962, 415, 7, 973, 648, 194, 399, 379, 717, 462, 682, 488, 835, 781, 599, 841, 890, 520, 24, 296, 702, 2, 690, 613, 983, 653, 485, 907, 643, 458, 811, 100, 254, 964, 990, 708, 413, 543, 632, 584, 324, 843, 415, 849, 976, 727, 894, 380, 318, 513, 818, 228, 971, 990, 13, 373, 882, 433, 84, 986, 303, 891, 789, 775, 612, 612, 470, 509, 718, 566, 416, 236, 710, 875, 861, 905, 758, 400, 467, 116, 297, 140, 527, 43, 950, 794, 475, 954, 140, 743, 403, 506, 252, 623, 479, 960, 379, 161, 140, 702, 677, 700, 52, 77, 428, 258, 545, 926, 999, 45, 436, 99, 810, 823, 733, 683, 199, 625, 766, 547, 574, 483, 535, 625, 954, 173, 136, 470, 149, 956, 509, 929, 119, 228, 194, 146, 968, 901, 883, 530, 161, 885, 949, 137, 726, 697, 313, 717, 781, 169, 678, 377, 53, 114, 273, 517, 866, 595, 546, 288, 448, 16, 855, 748, 754, 88, 960, 451, 872, 52, 911, 178, 652, 70, 108, 116, 461, 72, 590, 11, 71, 633, 531, 276, 298, 425, 452, 338, 325, 191, 123, 386, 21, 120, 288, 263, 511, 639, 279, 195, 865, 957, 606, 180, 587, 344, 526, 707, 383, 671, 764, 985, 85, 116, 506, 835, 691, 702, 705, 292, 220, 926, 480, 82, 953, 612, 191, 127, 943, 191, 82, 998, 920, 171, 824, 12, 487, 357, 684, 918, 939, 738, 495, 132, 239, 550, 329, 831, 127, 523, 91, 624, 760, 577, 996, 283, 492, 127, 955, 747, 701, 629, 5, 723, 383, 528, 97, 914, 211, 573, 320, 416, 52, 190, 137, 695, 283, 91, 559, 695, 348, 679, 855, 380, 561, 705, 66, 183, 847, 568, 579, 805, 423, 926, 904, 934, 73, 539, 382, 698, 398, 953, 831, 157, 923, 747, 68, 312, 825, 517, 463, 365, 884, 179, 549, 811, 299, 125, 34, 734, 844, 40, 604, 282, 55, 420, 234, 195, 324, 170, 605, 842, 300, 943, 451, 507, 276, 156, 97, 131, 691, 103, 782, 886, 583, 485, 269, 636, 415, 276, 105, 141, 86, 229, 858, 923, 526, 431, 469, 647, 678, 673, 316, 97, 397, 546, 658, 395, 158, 162, 939, 136, 953, 519, 475, 986, 611, 378, 133, 23, 897, 112, 558, 245, 732, 732, 683, 131, 648, 719, 688, 967, 415, 666, 2, 68, 206, 915, 43, 346, 477, 645, 841, 747, 200, 533, 81, 216, 785, 204, 832, 390, 755, 782, 629, 925, 523, 954, 602, 628, 299, 490, 30, 406, 183, 157, 428, 232, 206, 887, 411, 796, 397, 843, 574, 277, 665, 963, 497, 108, 759, 763, 406, 646, 397, 100, 419, 299, 718, 517, 315, 189, 604, 623, 792, 928, 72, 689, 955, 604, 262, 148, 706, 955, 169, 956, 311, 955, 903, 562, 600, 840, 403, 346, 115, 107, 815, 135, 488, 363, 629, 889, 404, 316, 595, 30, 934, 121, 194, 246, 444, 465, 878, 942, 886, 545, 8, 746, 516, 42, 726, 883, 821, 827, 386, 749, 945, 313, 386, 237, 612, 108, 965, 280, 390, 136, 74, 989, 692, 764, 91, 335, 268, 979, 13, 349, 46, 403, 592, 168, 198, 27, 191, 801, 80, 41, 256, 228, 127, 521, 998, 257, 93, 779, 484, 922, 924, 931, 894, 336, 133, 492, 815, 972, 524, 136, 918, 96, 192, 25, 367, 86, 784, 378, 583, 677, 971, 507, 550, 638, 436, 1000, 684, 31, 565, 627, 915, 517, 221, 930, 523, 900, 472, 776, 892, 779, 576, 490, 27, 87, 517, 301, 570, 747, 481, 596, 360, 276, 884, 891, 784, 951, 250, 99, 711, 538, 560, 951, 578, 725, 686, 765, 543, 224, 399, 458, 558, 593, 447, 367, 56, 778, 264, 644, 725, 961, 598, 829, 45, 433, 286, 747, 343, 94, 486, 462, 135, 877, 469, 822, 137, 843, 501, 18, 939, 182, 618, 486, 442, 444, 907, 549, 455, 216, 461, 684, 29, 440, 365, 506, 481, 301, 855, 225, 195, 112, 894, 398, 909, 878, 226, 480, 514, 742, 722, 344, 113, 401, 941, 297, 349, 32, 558, 397, 989, 961, 478, 56, 302, 350, 53, 658, 529, 444, 888, 40, 468, 104, 895, 629, 365, 611, 56, 639, 716, 181, 988, 91, 742, 712, 596, 588, 339, 795, 406, 127, 558, 181, 701, 24, 861, 338, 509, 680, 203, 730, 444, 331, 913, 770, 902, 720, 186, 812, 226, 350, 760, 343, 995, 330, 367, 734, 963, 758, 858, 591, 847, 68, 780, 812, 255, 695, 268, 955, 919, 347, 60, 86, 566, 156, 134, 59, 738, 491, 804, 367, 88, 780, 364, 306, 967, 772, 43, 130, 957, 565, 57, 485, 418, 169, 861, 599, 28, 109, 440, 194, 833, 799, 129, 945, 443, 147, 21, 368, 641, 798, 32, 9, 638, 366, 694, 836, 1000, 811, 515, 898, 741, 717, 889, 458, 209, 556, 611, 825, 288, 755, 450, 942, 799, 404, 801, 892, 29, 375, 545, 859, 207, 573, 194, 885, 227, 983, 148, 219, 903, 875, 319, 668, 492, 732, 747, 410, 217, 473, 718, 564, 897, 876, 8, 305, 204, 576, 747, 290, 604, 315, 25, 156, 28, 48, 897, 735, 90, 760, 962, 26, 135, 545, 535, 775, 917, 330, 564, 738, 298, 19, 99, 278, 540, 781, 632, 659, 343, 193, 478, 608, 600, 657, 1000, 816, 412, 802, 931, 655, 896, 157, 865, 218, 87, 683, 271, 144, 964, 695, 907, 460, 865, 901, 759, 657, 887, 486, 564, 545, 188, 258, 827, 403, 173, 750, 55, 859, 857, 511, 758, 300, 832, 975, 37, 258, 332, 74, 179, 26, 117, 852, 624, 250, 229, 54, 139, 963, 453, 728, 440, 247, 442, 901, 253, 592, 301, 557, 367, 558, 732, 151, 975, 279, 815, 520, 424, 810, 215, 396, 362, 259, 786, 356, 777, 57, 407, 76, 330, 751, 178, 440, 802, 329, 851, 112, 788, 938, 697, 399, 249, 478, 258, 474, 978, 699, 576, 616, 932, 576, 682, 361, 502, 357, 894, 957, 803, 529, 521, 443, 208, 87, 441, 247, 531, 990, 148, 123, 527, 798, 207, 227, 43, 46, 850, 935, 393, 271, 633, 897, 955, 781, 584, 463, 309, 36, 864, 960, 152, 754, 998, 679, 356, 785, 46, 228, 55, 692, 460, 733, 492, 378, 592, 135, 925, 797, 748, 593, 932, 393, 353, 223, 947, 396, 953, 685, 186, 748, 383, 824, 262, 76, 758, 730, 725, 950, 462, 464, 286, 513, 765, 343, 723, 873, 671, 934, 86, 325, 335, 73, 754, 726, 863, 179, 880, 916, 339, 256, 541, 200, 837, 932, 624, 639, 88, 36, 209, 841, 736, 662, 849, 586, 738, 516, 776, 743, 831, 757, 480, 818, 382, 993, 290, 163, 464, 602, 907, 414, 348, 501, 665, 928, 225, 124, 257, 922, 839, 305, 890, 568, 870, 106, 898, 454, 169, 559, 790, 236, 646, 347, 41, 949, 333, 887, 211, 984, 349, 26, 622, 900, 156, 461, 226, 617, 333, 976, 32, 342, 442, 995, 99, 726, 951, 114, 329, 67, 602, 358, 306, 844, 124, 465, 383, 548, 265, 213, 106, 720, 917, 526, 759, 316, 619, 18, 83, 461, 649, 579, 75, 493, 568, 604, 482, 796, 431, 503, 67, 169, 312, 509, 772, 591, 509, 372, 635, 620, 359, 253, 137, 535, 732, 919, 746, 186, 95, 441, 28, 309, 112, 924, 260, 389, 1, 414, 130, 640, 162, 177, 106, 867, 302, 879, 7, 553, 790, 762, 796, 977, 647, 111, 342, 866, 944, 657, 974, 434, 894, 641, 382, 982, 891, 407, 779, 838, 138, 659, 432, 706, 725, 969, 361, 631, 256, 33, 226, 678, 840, 820, 582, 847, 816, 658, 611, 190, 84, 627, 420, 239, 870, 121, 1000, 754, 665, 612, 767, 541, 780, 22, 141, 954, 245, 287, 687, 5, 939, 783, 373, 76, 995, 316, 160, 601, 284, 808, 833, 880, 870, 85, 838, 129, 276, 977, 810, 295, 890, 299, 718, 376, 980, 685, 196, 439, 90, 371, 573, 682, 868, 370, 298, 616, 689, 492, 502, 212, 260, 320, 219, 310, 997, 972, 42, 989, 216, 318, 240, 227, 519, 461, 161, 829, 147, 451, 345, 493, 408, 44, 1, 601, 288, 889, 794, 569, 646, 105, 110, 428, 233, 765, 953, 674, 967, 236, 227, 600, 400, 532, 770, 110, 192, 637, 925, 659, 730, 478, 957, 864, 560, 85, 173, 488, 645, 853, 49, 556, 85, 553, 781, 478, 576, 926, 38, 931, 536, 337, 534, 161, 166, 606, 318, 390, 491, 972, 95, 162, 841, 856, 13, 967, 334, 673, 354, 481, 654, 914, 456, 393, 303, 145, 288, 144, 651, 33, 584, 698, 754, 78, 614, 700, 96, 719, 838, 947, 974, 513, 884, 133, 23, 746, 879, 672, 922, 893, 609, 620, 468, 121, 966, 42, 994, 935, 199, 56, 481, 285, 292, 568, 776, 830, 163, 742, 663, 768, 974, 634, 293, 934, 803, 413, 364, 880, 38, 922, 241, 387, 389, 810, 341, 466, 500, 486, 285, 57, 453, 951, 804, 563, 874, 138, 382, 234, 290, 645, 64, 51, 977, 322, 59, 424, 265, 451, 10, 226, 503, 34, 445, 10, 614, 256, 336, 285, 512, 31, 583, 732, 61, 744, 162, 624, 226, 905, 516, 788, 898, 587, 426, 38, 22, 130, 708, 653, 531, 56, 803, 440, 154, 698, 646, 684, 943, 108, 455, 595, 469, 450, 78, 130, 536, 227, 766, 384, 904, 295, 116, 298, 838, 987, 636, 273, 51, 191, 402, 947, 607, 314, 794, 459, 920, 208, 351, 722, 737, 305, 399, 221, 798, 707, 550, 869, 614, 988, 680, 392, 196, 378, 436, 916, 533, 697, 836, 681, 518, 493, 678, 565, 967, 319, 441, 187, 337, 393, 530, 79, 704, 493, 978, 249, 540, 953, 940, 534, 604, 749, 752, 559, 488, 902, 52, 169, 50, 959, 684, 544, 61, 38, 336, 43, 119, 758, 55, 877, 442, 877, 395, 297, 500, 678, 932, 519, 206, 46, 808, 172, 246, 630, 340, 193, 480, 418, 191, 646, 313, 113, 534, 677, 478, 49, 892, 34, 794, 571, 8, 606, 548, 51, 199, 1000, 154, 615, 235, 772, 897, 377, 106, 704, 128, 635, 692, 17, 964, 915, 555, 799, 326, 939, 846, 249, 329, 606, 259, 996, 616, 137, 165, 420, 882, 9, 364, 848, 209, 447, 746, 697, 699, 690, 861, 870, 244, 851, 942, 779, 538, 516, 448, 224, 98, 480, 869, 430, 90, 933, 308, 417, 515, 807, 585, 625, 632, 620, 50, 173, 417, 689, 67, 862, 833, 322, 722, 719, 701, 190, 400, 201, 9, 358, 126, 910, 652, 335, 867, 182, 602, 214, 670, 931, 581, 520, 970, 121, 139, 472, 107, 416, 631, 914, 118, 547, 297, 263, 405, 192, 674, 817, 46, 774, 828, 370, 644, 985, 755, 960, 226, 514, 110, 766, 296, 670, 762, 179, 662, 862, 390, 416, 616, 405, 731, 822, 924, 606, 102, 609, 586, 281, 226, 812, 780, 318, 363, 797, 970, 338, 940, 732, 188, 439, 512, 409, 104, 671, 548, 324, 816, 273, 269, 120, 642, 394, 318, 962, 425, 811, 647, 370, 960, 361, 702, 475, 516, 690, 74, 597, 943, 174, 200, 802, 813, 919, 323, 567, 195, 873, 891, 316, 657, 203, 718, 250, 204, 239, 924, 216, 253, 921, 393, 863, 37, 233, 870, 264, 770, 674, 581, 579, 984, 938, 20, 245, 831, 364, 292, 972, 294, 236, 761, 66, 486, 694, 863, 489, 601, 305, 981, 894, 798, 718, 518, 62, 451, 766, 457, 348, 244, 678, 479, 551, 141, 808, 894, 89, 994, 550, 398, 24, 762, 262, 344, 463, 323, 951, 158, 789, 310, 599, 968, 883, 163, 614, 271, 637, 403, 35, 389, 437, 449, 187, 671, 596, 428, 449, 546, 450, 755, 966, 141, 788, 866, 53, 388, 241, 195, 367, 279, 279, 878, 292, 731, 31, 996, 284, 31, 46, 547, 80, 510, 763, 962, 887, 298, 195, 186, 58, 25, 185, 568, 34, 199, 861, 188, 602, 820, 63, 106, 967, 365, 887, 66, 388, 61, 168, 400, 356, 255, 101, 918, 93, 443, 29, 448, 796, 601, 545, 905, 6, 355, 762, 247, 870, 180, 17, 20, 656, 50, 964, 196, 583, 647, 516, 862, 964, 560, 590, 124, 119, 429, 330, 138, 341, 844, 61, 935, 781, 158, 205, 603, 775, 201, 629, 493, 715, 168, 881, 413, 986, 902, 56, 503, 304, 505, 984, 156, 262, 100, 72, 879, 679, 401, 415, 715, 102, 240, 116, 287, 737, 225, 160, 902, 572, 52, 240, 731, 812, 938, 884, 167, 997, 365, 218, 675, 193, 531, 690, 298, 270, 85, 787, 949, 767, 8, 214, 564, 937, 548, 220, 781, 45, 280, 988, 261, 533, 455, 460, 159, 9, 191, 45, 529, 901, 416, 152, 408, 907, 624, 432, 932, 324, 9, 791, 872, 729, 929, 743, 690, 412, 270, 356, 545, 201, 130, 863, 908, 103, 65, 622, 468, 907, 802, 942, 540, 589, 625, 457, 32, 229, 818, 299, 99, 943, 20, 293, 147, 764, 291, 438, 588, 107, 521, 953, 235, 200, 130, 607, 114, 126, 732, 573, 472, 586, 562, 8, 533, 354, 413, 865, 112, 895, 678, 195, 763, 489, 163, 878, 83, 486, 19, 171, 970, 596, 948, 705, 166, 751, 780, 202, 289, 553, 995, 116, 838, 517, 585, 930, 117, 896, 546, 886, 1000, 366, 697, 965, 540, 345, 82, 501, 449, 722, 380, 670, 162, 449, 30, 199, 18, 45, 850, 678, 259, 313, 159, 542, 141, 845, 39, 560, 50, 712, 503, 580, 22, 227, 440, 44, 474, 322, 559, 81, 261, 865, 241, 129, 831, 860, 653, 608, 854, 597, 405, 475, 871, 308, 341, 551, 893, 498, 336, 673, 651, 68, 620, 65, 920, 875, 474, 923, 101, 789, 222, 59, 925, 254, 633, 750, 810, 814, 151, 363, 357, 919, 973, 974, 414, 304, 399, 311, 33, 976, 533, 114, 911, 451, 602, 85, 316, 808, 120, 187, 512, 650, 149, 950, 575, 728, 543, 933, 245, 311, 311, 955, 15, 289, 16, 481, 95, 608, 450, 109, 163, 323, 703, 67, 491, 247, 128, 545, 406, 501, 8, 72, 275, 483, 947, 396, 333, 426, 635, 701, 712, 469, 367, 774, 460, 301, 718, 90, 199, 507, 860, 483, 363, 918, 280, 629, 890, 440, 769, 515, 914, 604, 747, 87, 734, 491, 466, 140, 354, 86, 333, 585, 947, 439, 936, 278, 513, 985, 362, 968, 57, 874, 875, 769, 499, 898, 309, 972, 137, 579, 761, 578, 658, 588, 506, 238, 176, 572, 885, 596, 657, 8, 90, 916, 755, 293, 307, 534, 475, 430, 658, 86, 801, 997, 4, 924, 379, 507, 178, 1000, 571, 737, 357, 347, 75, 513, 193, 214, 274, 307, 152, 59, 591, 426, 905, 96, 473, 168, 425, 130, 84, 502, 950, 226, 566, 74, 938, 204, 217, 461, 594, 954, 705, 282, 478, 324, 364, 576, 819, 915, 420, 984, 545, 41, 860, 630, 297, 43, 298, 513, 492, 16, 165, 151, 291, 408, 815, 903, 155, 686, 222, 258, 576, 986, 762, 560, 932, 997, 326, 575, 266, 378, 987, 365, 137, 849, 549, 868, 983, 95, 983, 374, 893, 633, 462, 643, 207, 940, 827, 91, 394, 242, 886, 633, 540, 471, 884, 716, 387, 208, 199, 846, 343, 877, 85, 90, 474, 164, 320, 699, 594, 275, 40, 708, 446, 821, 259, 522, 845, 651, 301, 621, 200, 541, 153, 474, 244, 80, 812, 212, 924, 341, 185, 22, 680, 766, 529, 206, 192, 330, 913, 848, 23, 10, 309, 518, 740, 257, 777, 534, 917, 823, 172, 901, 197, 681, 688, 923, 824, 609, 446, 734, 84, 268, 352, 739, 631, 228, 183, 997, 562, 756, 349, 791, 633, 939, 275, 889, 706, 385, 560, 588, 745, 401, 149, 800, 768, 761, 756, 808, 138, 707, 350, 944, 192, 697, 990, 435, 419, 240, 387, 455, 812, 778, 479, 429, 445, 51, 730, 307, 238, 153, 331, 838, 178, 454, 658, 829, 645, 332, 855, 684, 44, 554, 383, 430, 574, 243, 305, 994, 778, 513, 392, 209, 720, 232, 68, 613, 674, 716, 663, 859, 548, 993, 40, 575, 245, 443, 822, 883, 695, 39, 69, 976, 215, 51, 515, 722, 67, 214, 131, 729, 909, 419, 593, 266, 515, 434, 347, 123, 699, 525, 192, 753, 701, 865, 266, 107, 289, 851, 351, 82, 17, 601, 754, 398, 365, 940, 343, 870, 488, 32, 19, 458, 745, 807, 919, 130, 582, 478, 53, 796, 990, 211, 102, 388, 371, 21, 606, 998, 1, 37, 689, 276, 673, 642, 889, 689, 811, 399, 691, 219, 780, 294, 737, 225, 286, 331, 249, 439, 425, 955, 718, 685, 712, 529, 885, 700, 514, 11, 904, 555, 380, 495, 932, 439, 957, 698, 504, 784, 600, 340, 360, 653, 660, 676, 437, 836, 404, 871, 399, 23, 283, 563, 609, 595, 928, 72, 710, 950, 870, 684, 421, 175, 180, 367, 761, 585, 842, 563, 371, 240, 265, 447, 558, 514, 97, 433, 32, 908, 900, 736, 423, 826, 605, 952, 882, 862, 388, 898, 742, 810, 957, 971, 606, 28, 376, 677, 606, 746, 719, 877, 25, 168, 262, 828, 187, 688, 922, 202, 302, 213, 835, 536, 129, 720, 556, 412, 723, 964, 143, 635, 156, 677, 444, 437, 460, 878, 6, 892, 729, 144, 883, 873, 319, 689, 170, 396, 725, 313, 439, 20, 303, 934, 895, 351, 734, 164, 38, 421, 789, 668, 775, 688, 433, 964, 786, 473, 573, 320, 954, 155, 661, 47, 239, 261, 665, 701, 220, 263, 29, 466, 84, 949, 482, 244, 45, 510, 943, 624, 410, 932, 783, 480, 511, 354, 885, 425, 357, 243, 160, 496, 419, 35, 631, 304, 703, 479, 349, 49, 619, 720, 607, 509, 896, 561, 741, 801, 662, 264, 410, 511, 461, 182, 323, 501, 571, 810, 171, 916, 57, 87, 459, 840, 199, 591, 189, 229, 787, 674, 468, 774, 274, 811, 536, 553, 475, 66, 940, 243, 58, 136, 789, 511, 18, 589, 416, 598, 522, 360, 241, 377, 322, 721, 971, 868, 912, 715, 460, 138, 413, 7, 56, 132, 877, 14, 548, 376, 839, 275, 15, 985, 729, 235, 540, 312, 36, 886, 796, 474, 470, 755, 965, 417, 505, 481, 425, 837, 829, 557, 615, 68, 213, 732, 52, 436, 249, 881, 334, 769, 514, 743, 304, 775, 858, 775, 802, 168, 551, 745, 214, 803, 418, 790, 588, 801, 190, 667, 653, 639, 577, 989, 220, 566, 869, 788, 348, 960, 257, 513, 138, 394, 426, 487, 691, 720, 767, 851, 353, 474, 230, 822, 348, 612, 938, 883, 50, 97, 809, 772, 243, 135, 861, 568, 257, 876, 400, 285, 1, 621, 472, 555, 821, 807, 924, 299, 385, 534, 107, 119, 9, 785, 55, 549, 595, 639, 557, 210, 198, 553, 346, 723, 269, 551, 659, 217, 760, 329, 160, 164, 53, 351, 476, 533, 935, 344, 655, 159, 254, 722, 341, 953, 74, 731, 823, 759, 224, 323, 319, 935, 798, 703, 793, 256, 49, 142, 914, 779, 72, 345, 815, 746, 756, 799, 992, 18, 155, 384, 182, 54, 399, 977, 726, 282, 664, 315, 889, 672, 887, 62, 863, 79, 560, 626, 679, 391, 272, 261, 854, 707, 376, 24, 369, 161, 275, 575, 93, 195, 539, 820, 187, 782, 86, 378, 423, 807, 862, 215, 609, 23, 935, 236, 924, 60, 493, 481, 162, 907, 411, 739, 463, 910, 95, 748, 686, 675, 635, 721, 225, 319, 856, 465, 503, 726, 924, 166, 931, 312, 57, 389, 572, 332, 402, 378, 576, 990, 698, 37, 656, 765, 370, 32, 840, 909, 20, 816, 99, 692, 543, 358, 741, 512, 362, 237, 912, 545, 907, 2, 504, 541, 821, 852, 623, 21, 544, 299, 522, 992, 100, 586, 253, 308, 779, 174, 530, 113, 433, 418, 244, 141, 660, 935, 851, 446, 660, 503, 184, 890, 238, 453, 629, 942, 221, 311, 776, 575, 192, 863, 131, 361, 716, 436, 208, 525, 695, 680, 77, 23, 331, 858, 380, 507, 428, 418, 576, 622, 494, 678, 960, 85, 820, 106, 315, 228, 653, 280, 315, 900, 807, 752, 68, 464, 335, 102, 914, 148, 299, 776, 936, 707, 233, 633, 510, 918, 226, 205, 673, 501, 667, 152, 86, 635, 801, 632, 623, 817, 170, 781, 553, 539, 30, 153, 77, 83, 252, 170, 867, 67, 721, 716, 634, 271, 516, 214, 360, 165, 712, 856, 384, 902, 276, 697, 635, 342, 406, 12, 264, 973, 936, 267, 102, 300, 392, 519, 468, 331, 171, 282, 788, 980, 62, 381, 55, 298, 693, 31, 502, 373, 618, 977, 473, 715, 346, 885, 399, 156, 877, 802, 558, 567, 297, 322, 846, 918, 187, 186, 46, 542, 818, 966, 42, 836, 604, 550, 367, 610, 493, 714, 986, 306, 208, 147, 489, 151, 253, 302, 230, 782, 164, 134, 668, 826, 851, 299, 70, 307, 11, 167, 873, 333, 193, 981, 135, 779, 801, 795, 478, 9, 88, 17, 755, 58, 609, 117, 526, 558, 254, 958, 22, 536, 785, 596, 767, 640, 108, 389, 558, 184, 314, 870, 54, 368, 512, 905, 89, 503, 282, 734, 518, 362, 24, 288, 444, 445, 845, 133, 477, 920, 147, 552, 530, 214, 92, 116, 259, 295, 348, 913, 812, 132, 6, 289, 629, 831, 272, 995, 979, 271, 197, 964, 68, 665, 633, 943, 17, 785, 529, 710, 788, 725, 164, 617, 589, 50, 924, 923, 739, 581, 885, 641, 360, 386, 245, 287, 653, 586, 405, 260, 898, 3, 533, 902, 661, 48, 890, 844, 837, 170, 828, 530, 586, 812, 587, 641, 51, 215, 220, 733, 386, 334, 656, 779, 287, 410, 680, 413, 990, 606, 233, 905, 21, 391, 743, 394, 834, 843, 938, 815, 771, 18, 702, 668, 94, 653, 550, 928, 788, 795, 234, 85, 341, 651, 623, 106, 821, 325, 890, 547, 710, 617, 303, 513, 413, 569, 736, 373, 794, 237, 418, 825, 498, 803, 608, 735, 237, 119, 977, 154, 429, 744, 625, 53, 413, 91, 926, 200, 296, 8, 975, 453, 28, 271, 486, 305, 146, 149, 663, 368, 783, 139, 538, 268, 343, 533, 152, 987, 968, 225, 278, 671, 892, 694, 553, 978, 544, 49, 413, 3, 668, 745, 250, 115, 707, 154, 961, 536, 279, 510, 363, 5, 408, 918, 18, 146, 331, 625, 685, 803, 487, 576, 21, 455, 605, 698, 823, 668, 183, 223, 830, 752, 930, -250, 592, 517, 726, 205, -478, 199, 388, -658, 293, -400, -790, -53, -671, 929, -657, -411, -365, 69, 45, -713, -878, -368, -230, -920, -257, -989, 778, 504, -999, 863, -746, 531, 77, -308, 724, 389, 260, -444, 425, 938, -91, 517, -791, -955, 332, -942, 36, 947, -417, 26, -929, -899, 373, -263, -865, -587, -979, -34, -535, 427, -9, 383, 813, -456, -506, 665, -838, 951, -771, -110, -329, -358, -902, 913, 277, -867, -969, -506, 639, -232 };
        //std::vector<int> multipliers{ -525,571,-905,-670,-892,-87,-697,-566,-606,-375,-749,-294,-312,-765,-714,-833,496,-990,-731,-228,-178,-564,-566,-495,-328,-319,-180,-895,-461,-453,-16,-938,-627,-8,-95,144,-60,-794,-663,-723,-492,-427,-719,-497,-506,-823,-974,-529,-427,-640,-274,-675,-741,899,-349,-283,-584,-521,-634,-346,-690,-868,-603,-409,-479,-895,-783,-532,-748,-611,126,-523,-995,-109,-225,-845,-470,-884,-125,-201,-24,-450,-516,-654,-285,-309,-41,-895,-306,-337,-827,-288,-99,-551,-518,-895,965,-568,-254,-166,-363,-77,-173,520,-649,-428,-684,-358,750,-401,-58,-747,-638,-3,-222,-986,-992,-826,-690,-290,-86,300,-984,532,-564,-231,-677,-160,-912,-740,-447,-172,-164,-269,-150,-28,-193,-419,-803,-838,632,-461,-483,-82,-804,-205,-629,-163,-449,-330,-987,-760,-783,-739,-525,-965,-661,-19,-842,-324,-81,770,-111,-613,-244,-689,-193,-367,-939,-107,-936,-840,-112,-700,-330,-396,-138,-156,-362,310,-437,-848,-5,-624,631,-388,-192,-66,-704,-916,-796,542,-37,-858,-68,-961,-533,-157,-306,-768,-688,-888,-987,-437,-465,244,-542,-976,-173,-23,-945,-378,-456,-564,-764,544,747,-389,-167,-388,-934,-178,-466,-361,-169,-610,-95,-836,-611,-387,-472,-396,-629,-33,-799,-691,-853,-328,-234,-264,-978,-189,-308,-510,-665,-719,-246,-220,-418,-732,982,-521,-708,-790,-683,-793,-169,-335,-584,-429,-421,-355,-295,-150,-888,-394,-431,-149,-243,-394,-56,-774,-170,-906,-811,-712,-456,-541,-757,-373,-40,-278,-132,-79,-774,263,-612,-811,-366,-813,-576,-8,676,-43,-983,-376,-153,-48,-906,-182,-335,-285,-419,-909,-433,-223,-487,60,-766,-356,-701,-623,-672,-872,-320,-782,-5,-747,-415,-385,-835,-393,-693,-22,-91,-638,-786,-133,-14,-218,-713,-560,-725,-200,-890,766,-979,-369,-481,-924,-500,-295,-940,-658,-528,-684,490,-690,-881,-781,-410,-141,-365,-598,-840,-440,-460,-787,-450,-326,-92,-596,-141,-65,-930,-691,-547,-765,-76,-328,-751,-653,-783,-552,-470,-478,-232,-829,-477,462,-831,-713,-851,-228,-254,135,-528,-784,-292,-472,-26,-890,-252,-684,-580,-791,-273,-623,-53,-289,-52,-165,-261,-395,-939,-477,-455,-138,-473,-289,-139,-63,-685,-11,-294,-152,-182,-907,-218,-233,-631,-809,-292,-703,-78,-527,-92,-778,-223,-636,-22,-122,-419,-440,-518,-310,-34,-93,-166,-584,-312,-627,-711,326,-513,-818,-350,-897,-676,-503,-664,-447,-653,-105,502,-678,-734,-614,-334,-170,-152,-409,-707,-410,-295,-78,961,-800,-152,-342,-342,-30,487,-692,-426,947,-111,-454,-184,-168,-105,-460,-994,-565,-944,492,-602,-353,-112,-224,-368,-849,-468,-866,-908,-577,-211,-905,-177,-829,-693,-912,-924,-280,-172,-467,-794,-470,-953,-919,-904,-174,-868,-865,-463,-976,-939,-225,-592,-235,-172,-308,-115,-605,-930,-698,-460,-344,-810,-467,-80,-610,-521,-877,-9,-202,-951,-496,-521,-569,-447,-815,-987,-661,-727,739,-744,-672,-635,-431,-233,-57,-704,-277,-8,-794,-127,-744,-251,-771,-617,-412,-925,-311,-611,-169,-756,-219,-627,-175,-149,-765,-32,-553,-576,-484,-698,-599,-84,-677,-117,636,-253,-950,-208,-893,-622,-8,-477,-4,-981,-581,-406,-59,-89,167,-222,758,-100,536,-688,-952,-57,-797,-649,-983,-442,-828,-544,-842,-473,-133,-548,-514,-889,-430,-119,-835,-863,-231,-754,-533,-134,-832,-785,-537,-205,-870,-729,-641,-71,-915,-789,-340,-501,-641,-483,-525,-146,-100,-645,-543,-780,-466,-231,-964,-315,-311,881,-864,-501,-661,-156,-213,-872,-823,-999,-87,-687,-892,-925,-196,-438,-606,-178,-841,-660,-981,-579,-640,-203,-430,-532,-670,-713,-329,-631,-297,-522,-679,-45,-18,-36,-930,-359,-18,-335,-791,-242,-106,-674,-152,847,-167,-366,-973,-95,-658,-18,-989,-422,-832,-331,-853,-969,-673,-13,-802,-86,-147,-82,-253,522,-403,-463,-964,-57,-344,-434,-976,-385,597,-522,-954,-711,-620,-554,-542,-392,-511,-440,-290,-749,-243,186,-312,-704,-275,-317,-525,-57,-228,-395,177,-916,-390,-57,-811,-555,445,-499,-90,-787,-718,-596,-250,-882,-875,-19,-627,-6,-830,-724,-716,-841,-133,-122,-839,-815,-360,486,-181,-599,-264,-579,-506,-375,-2,-465,-186,-88,-334,-362,-612,-494,-79,-9,-824,-191,-910,-382,-453,-670,-750,-664,-127,-749,-734,-132,-698,699,-563,-985,-399,-961,-602,-733,-495,-305,-375,-790,-367,-136,-314,-979,-303,-32,-109,-920,-677 };
        auto res = s1770.maximumScore(nums, multipliers);
        std::cout << "3 res = " << res << std::endl;
    }
    {
        Solution_1769 s1769;
        std::string boxes = "110";
        auto res = s1769.minOperations(boxes);
        std::cout << "2 res = " << res[0] << std::endl;
    }
    {
        Solution_1768 s1768;
        std::string word1 = "abc";
        std::string word2 = "pqr";
        auto res = s1768.mergeAlternately(word1, word2);
        std::cout << "1 res = " << res << std::endl;
    }
    {
        Solution_1760 s1760;
        std::vector<int> nums{9};
        int maxOperations = 2;
        auto res = s1760.minimumSize(nums,maxOperations);
        std::cout << "1760 res = " << res<< std::endl;
    }
    {
        Solution_752 s752;
        std::vector<std::string> deadends{"0000"};
        std::string target = "0202";
        auto res = s752.openLock(deadends,target);
        std::cout << "752 res = " << res<< std::endl;
    }
    {
        Solution_37 s37;
        std::vector<std::vector<char>> board{{'5','3','.','.','7','.','.','.','.'},
                                             {'6','.','.','1','9','5','.','.','.'},
                                             {'.','9','8','.','.','.','.','6','.'},
                                             {'8','.','.','.','6','.','.','.','3'},
                                             {'4','.','.','8','.','3','.','.','1'},
                                             {'7','.','.','.','2','.','.','.','6'},
                                             {'.','6','.','.','.','.','2','8','.'},
                                             {'.','.','.','4','1','9','.','.','5'},
                                             {'.','.','.','.','8','.','.','7','9'}};
        s37.solveSudoku(board);
        for(int i = 0;i < 9;i++){
            for(int j = 0;j < 9;j++){
                std::cout<<board[i][j];
            }
            std::cout<<std::endl;
        }
    }
    {
        //MyCalendar();
//MyCalendar.book(10, 20); // returns true
//MyCalendar.book(50, 60); // returns true
//MyCalendar.book(10, 40); // returns true
//MyCalendar.book(5, 15); // returns false
//MyCalendar.book(5, 10); // returns true
//MyCalendar.book(25, 55); // returns true
        MyCalendarTwo s731;
        auto res = s731.book(10, 20); // returns true
        res = s731.book(50, 60); // returns true
        res = s731.book(10, 40); // returns true
        res = s731.book(5, 15); // returns false
        res = s731.book(5, 10); // returns true
        res = s731.book(25, 55); // returns true
        std::cout << "731 res = " << res<< std::endl;
    }
    {
        MyCalendar s729;
        auto res = s729.book(10, 20); // returns true
        res = s729.book(15, 25); // returns false
        res = s729.book(20, 30); // returns true
        std::cout << "729 res = " << res<< std::endl;
    }
    {
        Solution_688 s688;
        int N = 3;
        int K = 2;
        int r = 0;
        int c = 0;
        auto res = s688.knightProbability(N,K,r,c);
        std::cout << "688 res = " << res<< std::endl;
    }
    {
        Solution_97 s97;
        std::string s1 = "aabcc";
        std::string s2 = "dbbca";
        std::string s3 = "aadbbcbcac";
        auto res = s97.isInterleave(s1,s2,s3);
        std::cout << "97 res = " << res << std::endl;
    }
    {
        Solution_316 s316;
        std::string s = "bbcaac";
        auto res = s316.removeDuplicateLetters(s);
        std::cout << "316 res = " << res << std::endl;
    }
    {
        std::string s = "pwwkew";
        Solution_3 s3;
        auto res = s3.lengthOfLongestSubstring(s);
        std::cout << "3 res = " << res << std::endl;
    }
    {
        //[1,2,3]
        //[1,2,2]
        Solution_4 s4;
        std::vector<int> nums1{1,2,3};
        std::vector<int> nums2{1,2,2};
        auto res = s4.findMedianSortedArrays(nums1,nums2);
        std::cout << "4 res = " << res << std::endl;
    }
    {
        //"acdcb"
        //"a*cb"
        Solution_44 s44;
        std::string s = "acdcb";
        std::string p = "a*cb";
        auto res = s44.isMatch(s,p);
        std::cout << "44 res = " << res << std::endl;
    }
    {
        Solution_1705 s1705;
//        std::vector<int> apples{1,2,3,5,2};
//        std::vector<int> days{3,2,1,4,2};
//        std::vector<int> apples{3,0,0,0,0,2};
//        std::vector<int> days{3,0,0,0,0,2};
        std::vector<int> apples{5,2,3};
        std::vector<int> days{6,9,10};
        auto res = s1705.eatenApples(apples,days);
        std::cout << "1705 res = " << res << std::endl;
    }
    {
        Solution_787 s787;
        int n = 5;
        std::vector<std::vector<int>> flights{{0,1,1},{1,2,2},{2,3,3},{3,4,4}};
        int src = 0;
        int dst = 1;
        int K = 4;
        auto res = s787.findCheapestPrice(n,flights,src,dst,K);
        std::cout << "787 res = " << res << std::endl;
    }
    {
        Solution_1546 s1546;
        std::vector<int> nums{-2,6,6,3,5,4,1,2,8};
        int target = 10;
        auto res = s1546.maxNonOverlapping(nums,target);
        std::cout << "1546 res = " << res << std::endl;
    }
    {
        Solution_567 s567;
        std::string s1 = "aba";
        std::string s2 = "aadbaaooo";
        auto res = s567.checkInclusion(s1,s2);
        std::cout << "567 res = " << res << std::endl;
    }
    {
        Solution_1743 s1743;
        std::vector<std::vector<int>> adjacentPairs{{2,1},{3,4},{3,2}};
        auto res = s1743.restoreArray(adjacentPairs);
        std::cout << "1743 res = " << res[0] << std::endl;
    }
    {
        Solution_1742 s1742;
        int lowLimit = 5;
        int highLimit = 15;
        auto res = s1742.countBalls(lowLimit,highLimit);
        std::cout << "1742 res = " << res << std::endl;
    }
    {
        Solution_1049 s1049;
        vector<int> stones{2,7,4,1,8,1,6,4,33,68,32};
        auto res = s1049.lastStoneWeightII(stones);
        std::cout << "1049 res = " << res << std::endl;
    }
    {
        Solution_322 s322;
        std::vector<int> coins{1,2,5};
        int amount = 11;
        auto res = s322.coinChange(coins,amount);
        std::cout << "322 res = " << res << std::endl;
    }
    {
        Solution_1419 s1419;
        std::string croakOfFrogs = "croackcrrooakak";
        auto res = s1419.minNumberOfFrogs(croakOfFrogs);
        std::cout << "1419 res = " << res << std::endl;
    }
    {
        Solution_486 s486;
        //std::vector<int> nums{1, 5, 233, 7};
        std::vector<int> nums{1, 5, 2};
        auto res = s486.PredictTheWinner(nums);
        std::cout << "486 res = " << res << std::endl;
    }
    {
        Solution_357 s357;
        int n = 3;
        auto res = s357.countNumbersWithUniqueDigits(n);
        std::cout << "357 res = " << res << std::endl;
    }
    {
        Solution_535 s535;
        std::string url = "https://leetcode.com/problems/design-tinyurl";
        auto encoded_url = s535.encode(url);
        auto decoded_url = s535.decode(encoded_url);
        std::cout << "535 decoded_url = " << decoded_url << std::endl;
    }
    {
        ListNode l1(1);
        ListNode l2(2);
        ListNode l3(3);
        ListNode l4(4);
        ListNode l5(5);
        ListNode l6(6);
        l1.next = &l2;
        l2.next = &l3;
        l3.next = &l4;
        l4.next = &l5;
        l5.next = &l6;
        Solution_25 s25;
        int k = 2;
        auto res = s25.reverseKGroup(&l1,k);
        std::cout << "1366 res = " << res->val << std::endl;
    }
    {
        Solution_1366 s1366;
        std::vector<std::string> votes{"ABC","ACB","ABC","ACB","ACB"};
        auto res = s1366.rankTeams(votes);
        std::cout << "1366 res = " << res << std::endl;
    }
    {
        Solution_60 s60;
        int n = 8;
        int k = 15025;
        auto res = s60.getPermutation(n,k);
        std::cout << "60 res = " << res << std::endl;
    }
    {
        Solution_477 s477;
        std::vector<int> nums{4, 14, 2};
        auto res = s477.totalHammingDistance(nums);
        std::cout << "477 res = " << res << std::endl;
    }
    {
        Solution_1508 s1508;
        std::vector<int> nums{1,2,3,4};
        int n = 4;
        int left = 1;
        int right = 5;
        auto res = s1508.rangeSum(nums,n,left,right);
        std::cout << "1508 res = " << res << std::endl;
    }
    {
        Solution_52 s52;
        int n = 4;
        auto res = s52.totalNQueens(n);
        std::cout << "52 res = " << res << std::endl;
    }
    {
        Solution_51 s51;
        int n = 4;
        auto res = s51.solveNQueens(n);
        std::cout << "51 res = " << res.size() << std::endl;
    }
    {
        Solution_1738 s1738;
        std::vector<std::vector<int>> matrix{{5,2},{1,6}};
        int k = 3;
        auto res = s1738.kthLargestValue(matrix,k);
        std::cout << "1738 res = " << res << std::endl;
    }
    {
        std::string time = "0?:3?";
        Solution_1736 s1736;
        auto res = s1736.maximumTime(time);
        std::cout << "1736 res = " << res << std::endl;
    }
    {
        Solution_1734 s1734;
        std::vector<int> encoded{6,5,4,6};
        auto res = s1734.decode(encoded);
        std::cout << "1734 res = " << res.size() << std::endl;
    }
    {
        Solution_315 s315;
        std::vector<int> nums{5,2,3,8,5,4,3,5,8,9,2,1,0,8,6,1};
        auto res = s315.countSmaller(nums);
        std::cout << "315 res = " << res.size() << std::endl;
    }
    {
        Solution_1562 s1562;
        std::vector<int> arr{3,2,5,6,10,8,9,4,1,7};
        int m = 3;
        auto res = s1562.findLatestStep(arr,m);
        std::cout << "1562 res = " << res << std::endl;
    }
    {
        Solution_1673 s1673;
        std::vector<int> nums{4,3,2,1,6,4,8,9,3,2};
        int k = 5;
        auto res = s1673.mostCompetitive(nums,k);
        std::cout << "1673 res = " << res.size() << std::endl;
    }
    {
        Solution_1727 s1727;
        std::vector<std::vector<int>> matrix{{0,0,1},{1,1,1},{1,0,1}};
        auto res = s1727.largestSubmatrix(matrix);
        std::cout << "1727 res = " << res << std::endl;
    }
    {
        //[1,3,7,3,2,4,10,7,5]
        //[4,5,2,1,1,2,4,1,4]
        //[9,2,8,8,2]
        //[4,1,3,3,5]
        //[1,1,1,1,1,1,1,1,1,1] 
        //[811,364,124,873,790,656,581,446,885,134] 
        Solution_1626 s1626;
        std::vector<int> scores{1,1,1,1,1,1,1,1,1,1};
        std::vector<int> ages{811,364,124,873,790,656,581,446,885,134};
        auto res = s1626.bestTeamScore(scores,ages);
        std::cout << "1626 res = " << res << std::endl;
    }
    {
        Solution_1584 s1584;
        std::vector<std::vector<int>> points{{0,0},{2,2},{3,10},{5,2},{7,0}};
        auto res = s1584.minCostConnectPoints(points);
        std::cout << "1584 res = " << res << std::endl;
    }
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
        //std::this_thread::sleep_for(std::chrono::seconds(20));
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
        //std::this_thread::sleep_for(std::chrono::seconds(20));
    }
    {
        Solution_1054 s1054;
        std::vector<int> barcodes{1,1,1,1,2,2,3,3};
        auto res = s1054.rearrangeBarcodes(barcodes);
        std::cout << "1054 res = " << res.size() << std::endl;
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
        Solution_297 s297;
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
        //using namespace s1365;
        Solution_1365 s1365;
        auto nums = std::vector<int>{37,64,63,2,41,78,51,36,2,20,25,41,72,100,17,43,54,27,34,86,12,48,70,44,87,68,62,98,68,30,8,92,5,10};
        auto res = s1365.smallerNumbersThanCurrent(nums);
        std::cout<<res[0]<<std::endl;
    }
    {
        Solution_138::Node* n1 = new Solution_138::Node(7);
        Solution_138::Node* n2 = new Solution_138::Node(13);
        Solution_138::Node* n3 = new Solution_138::Node(11);
        Solution_138::Node* n4 = new Solution_138::Node(10);
        Solution_138::Node* n5 = new Solution_138::Node(1);
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
        Solution_942 s942;
        auto res = s942.diStringMatch("IDID");
        std::cout<<"942 res = "<<res.size()<<std::endl;
    }
    {
        Solution_1079 s1079;
        std::cout<<s1079.numTilePossibilities("AAB")<<std::endl;
        std::set<std::string> map;
        std::string s = "123";
        map.insert(s);
        auto res = map.find(s);
        if(map.end() == res)
            map.insert(s);
        map.insert(s);
        std::cout<<map.size()<<std::endl;
    }
}