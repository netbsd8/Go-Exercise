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
    bool isValidBST(TreeNode* root) {
        vector<int> res;
        inorderTravel(root, res);

        for (int i=0; i<res.size()-1; i++) {
            if (res[i] >= res[i+1]) {
                return false;
            }
        }

        return true;
    }

    void inorderTravel(TreeNode* root, vector<int>& res) {
        if (root == NULL) {
            return;
        }

        inorderTravel(root->left, res);
        res.push_back(root->val);
        inorderTravel(root->right, res);
    }
};