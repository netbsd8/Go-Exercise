#include <iostream>
#include <vector>
#include <unordered_set>
#include <stack>
#include <string>
#include <algorithm>
#include <bits/stdc++.h>

using namespace std;

struct TreeNode {
    TreeNode    *left;
    TreeNode    *right;
    int         val;
    TreeNode(int x): val(x), left(NULL), right(NULL) {}
};

class Solution {
public:
    vector<vector<int > > zigzagLevelOrder(TreeNode* root) {
        vector<vector<int > > res;
        if (!root) {
            return res;
        }

        int level = 1;
        queue<TreeNode*> q;
        q.push(root);

        while(!q.empty()) {
            int size = q.size();
            vector<int> ans;
            for (int i=0; i<size; i++) {
                TreeNode* t = q.front();
                q.pop();
                if (t->left) {
                    q.push(t->left);
                }
                if (t->right) {
                    q.push(t->right);
                }

                if (level%2 == 0) {
                    ans.insert(ans.begin(), t->val);
                }
                else {
                    ans.push_back(t->val);
                }
            }
            level++;
            res.push_back(ans);
        }

        return res;
    }
};