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
    vector<int> inorderTraversal(TreeNode* root) {
        vector<int> ans;

        helper(root, ans);

        return ans;
    }

    void helper(TreeNode* root, vector<int> &ans) {
        if (root == NULL) {
            return;
        }

        if (root->left) {
            helper(root->left, ans);
        }

        ans.push_back(root->val);

        if (root->right) {
            helper(root->right, ans);
        }
    }
};

class Solution2 {
public:
    vector<int> inorderTraversal(TreeNode* root) {
        vector<int> ans;
        stack<TreeNode*> s;

        TreeNode* cur = root;
        while (!s.empty() || cur) {
            if (cur) {
                s.push(cur);
                cur = cur->left;
            } else {
                TreeNode* temp = s.top();
                s.pop();
                ans.push_back(temp->val);
                if (temp->right) {
                    cur = temp->right;
                }
            }
        }

        return ans;
    }
};