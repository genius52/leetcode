#include <vector>
#include <unordered_map>
#include <deque>
#include <set>
#include <algorithm>

//Twitter twitter = new Twitter();
//
//// User 1 posts a new tweet (id = 5).
//twitter.postTweet(1, 5);
//
//// User 1's news feed should return a list with 1 tweet id -> [5].
//twitter.getNewsFeed(1);
//
//// User 1 follows user 2.
//twitter.follow(1, 2);
//
//// User 2 posts a new tweet (id = 6).
//twitter.postTweet(2, 6);
//
//// User 1's news feed should return a list with 2 tweet ids -> [6, 5].
//// Tweet id 6 should precede tweet id 5 because it is posted after tweet id 5.
//twitter.getNewsFeed(1);
//
//// User 1 unfollows user 2.
//twitter.unfollow(1, 2);
//
//// User 1's news feed should return a list with 1 tweet id -> [5],
//// since user 1 is no longer following user 2.
//twitter.getNewsFeed(1);

class Twitter {
    class twitter_data{
    public:
        int id;
        int user_id;
    };
    std::unordered_map<int,std::set<int>> relation;
    std::deque<twitter_data*> twitter_queue;
public:
    /** Initialize your data structure here. */
    Twitter() {
    }

    /** Compose a new tweet. */
    void postTweet(int userId, int tweetId) {
        twitter_data* t = new twitter_data();
        t->user_id = userId;
        t->id = tweetId;
        twitter_queue.push_back(t);
    }

    /** Retrieve the 10 most recent tweet ids in the user's news feed.
     * Each item in the news feed must be posted by users who the user followed or by the user herself.
     * Tweets must be ordered from most recent to least recent. */
    vector<int> getNewsFeed(int userId) {
        std::set<int> watch_list;
        if(relation.find(userId) != relation.end()) {
            watch_list = relation[userId];
        }
        int i = 0;
        std::vector<int> res;
        auto it = twitter_queue.rbegin();
        while(i < 10 && it != twitter_queue.rend()){
            if((watch_list.size() != 0 && watch_list.find((*it)->user_id) != watch_list.end()) || userId == (*it)->user_id){
                res.push_back((*it)->id);
                i++;
            }
            it++;
        }
        return res;
    }

    /** Follower follows a followee. If the operation is invalid, it should be a no-op. */
    void follow(int followerId, int followeeId) {
        relation[followerId].insert(followeeId);
    }

    /** Follower unfollows a followee. If the operation is invalid, it should be a no-op. */
    void unfollow(int followerId, int followeeId) {
        relation[followerId].erase(followeeId);
    }
};

/**
 * Your Twitter object will be instantiated and called as such:
 * Twitter* obj = new Twitter();
 * obj->postTweet(userId,tweetId);
 * vector<int> param_2 = obj->getNewsFeed(userId);
 * obj->follow(followerId,followeeId);
 * obj->unfollow(followerId,followeeId);
 */