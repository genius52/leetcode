#include <vector>
#include <map>
#include <unordered_map>
using namespace std;

class TopVotedCandidate {
    std::map<int,int> record;
public:
    TopVotedCandidate(vector<int>& persons, vector<int>& times) {
        std::unordered_map<int,int> person_cnt;
        int len = persons.size();
        int lead_cnt = 0;
        int lead_person = -1;
        for(int i = 0;i < len;i++){
            person_cnt[persons[i]]++;
            if (person_cnt[persons[i]] >= lead_cnt){
                lead_cnt = person_cnt[persons[i]];
                lead_person = persons[i];
            }
            record[times[i]] = lead_person;
        }
    }

    int q(int t) {
        auto it = record.upper_bound(t);
        it--;
        return it->second;
    }
};

