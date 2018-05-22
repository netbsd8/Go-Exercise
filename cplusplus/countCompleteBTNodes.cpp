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
    int countNodes(TreeNode* root) {
        if (root == NULL) {
            return 0;
        }

        int lh = leftHeight(root);
        int rh = rightHeight(root);
        if (lh == rh) {
            return pow(2, lh) - 1;
        }

        int count = countNodes(root->left) + countNodes(root->right) + 1;
        return count;
    }

    int leftHeight(TreeNode* root) {
        if (root == NULL) {
            return 0;
        }

        int count = 1 + leftHeight(root->left);
        return count;
    }

    int rightHeight(TreeNode* root) {
        if (root == NULL) {
            return 0;
        }

        int count = 1 + rightHeight(root->right);
        return count;
    }

    int countNodes2(TreeNode* root) {
        int res = 0;
        if (root == NULL) {
            return 0;
        }

        res = 1 + countNodes2(root->left) + countNodes2(root->right);
        return res;
    }
};