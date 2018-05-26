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
    TreeNode* lowestCommonAncestor(TreeNode *root, TreeNode *p, TreeNode* q) {
        if (root == NULL) {
            return NULL;
        }

        if (root == p) {
            return p;
        }
        if (root == q) {
            return q;
        }

        TreeNode *l = lowestCommonAncestor(root->left, p, q);
        TreeNode *r = lowestCommonAncestor(root->right, p, q);
        if (l && r) {
            return root;
        }

        return l ? l : r;
    }
};