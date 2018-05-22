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

class Solution1 {
public:
    vector<int> preorderTraversal(TreeNode* root) {
        vector<int> ans;

        helper(root, ans);

        return ans;
    }

    void helper(TreeNode* root, vector<int> &ans) {
        if (root == NULL) {
            return;
        }

        ans.push_back(root->val);

        if (root->left) {
            helper(root->left, ans);
        }

        if (root->right) {
            helper(root->right, ans);
        }
    }
};

class Solution2 {
public:
    vector<int> preorderTraversal(TreeNode* root) {
        vector<int> ans;
        stack<TreeNode*> s;
        s.push(root);

        TreeNode* cur = root;
        while (!s.empty() || cur) {
            if (cur) {
                ans.push_back(cur->val);
                s.push(cur);
                cur = cur->left;
            } else {
                TreeNode* temp = s.top();
                s.pop();
                if (temp->right) {
                    cur = temp->right;
                }
            }
        }

        return ans;
    }
};