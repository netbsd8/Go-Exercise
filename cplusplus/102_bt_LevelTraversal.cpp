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
    vector<vector<int> > levelOrder(TreeNode* root) {
        vector<vector<int> > res;
        if (root == NULL) {
            return res;
        }

        queue<TreeNode*> q;
        q.push(root);

        while(!q.empty()) {
            vector<int> out;
            int length = q.size();

            for (int i=0; i<length; i++) {
                TreeNode* t = q.front();
                q.pop();
                out.push_back(t->val);
                if (t->left) q.push(t->left);
                if (t->right) q.push(t->right);
            }

            res.push_back(out);
        }

        return res;
    }

    vector<vector<int> > levelOrderRecursive(TreeNode* root) {
        vector<vector<int> > res;
        helper(res, 0, root);
        return res;
    }

    void helper(vector<vector<int> >& res, int level, TreeNode* root) {
        if (root == NULL) return;

        if (level == res.size()) res.push_back(vector<int>());

        res[level].push_back(root->val);
        if (root->left) helper(res, level+1, root->left);
        if (root->right) helper(res, level+1, root->right);
    }
};