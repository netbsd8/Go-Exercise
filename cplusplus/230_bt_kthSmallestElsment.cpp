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
    int kthSmallest(TreeNode* root, int k) {
        int leftCount = 0;

        helper(root->left, leftCount);
        if (leftCount + 1 == k) {
            return root->val;
        }
        else if (leftCount + 1 > k) {
            return kthSmallest(root->left, k);
        }
        else {
            return kthSmallest(root->right, k-leftCount-1);
        }
    }

    void helper(TreeNode* root, int& index) {
        if (root == NULL) {
            return;
        }

        helper(root->left, index);
        index = index + 1;
        helper(root->right, index);
    }
};