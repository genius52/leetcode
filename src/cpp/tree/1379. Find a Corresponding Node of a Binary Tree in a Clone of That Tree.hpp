#include "../define.h"
#include <deque>
using namespace std;

class Solution_1379 {
public:
    TreeNode* getTargetCopy(TreeNode *original, TreeNode *cloned, TreeNode *target) {
        if (original == nullptr || target == nullptr) {
            return nullptr;
        }
        std::deque < TreeNode * > q;
        q.push_back(original);
        std::deque < TreeNode * > q2;
        q2.push_back(cloned);
        while (!q.empty()) {
            auto top = q.front();
            q.pop_front();
            auto clonedtop = q2.front();
            q2.pop_front();
            if (top == target) {
                return clonedtop;
            }
            if (top->left != nullptr) {
                q.push_back(top->left);
            }
            if (top->right != nullptr) {
                q.push_back(top->right);
            }
            if (clonedtop->left != nullptr) {
                q2.push_back(clonedtop->left);
            }
            if (clonedtop->right != nullptr) {
                q2.push_back(clonedtop->right);
            }
        }
        return nullptr;
    }
}