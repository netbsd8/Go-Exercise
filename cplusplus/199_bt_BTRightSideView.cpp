#include <iostream>
#include <vector>
#include <unordered_set>
#include <stack>
#include <queue>
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
    vector<int> rightSideView(TreeNode* root) {
        vector<int> res;
        if (root == NULL) {
            return res;
        }

        queue<TreeNode*> q;
        q.push(root);

        while (!q.empty()) {
            int length = q.size();
            for (int i=0; i<length; i++) {
                TreeNode* t = q.front();
                q.pop();
                if (t->left) q.push(t->left);
                if (t->right) q.push(t->right);
                if (i+1 == length) res.push_back(t->val);
            }
        }

        return res;
    }
};