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
    TreeNode* buildTree(vector<int> &preorder, vector<int> &inorder) {
        int length = preorder.size();
        if (length == 0) {
            return NULL;
        }

        TreeNode* root = new TreeNode(preorder[0]);

        int location = findIndex(inorder, root->val);

        vector<int> ilList(inorder.begin(), inorder.begin()+location-1);
        vector<int> irList(inorder.begin()+location+1, inorder.end());
        vector<int> plList(preorder.begin()+1, preorder.begin()+1 + location);
        vector<int> prList(preorder.begin()+1+location+1, preorder.end());

        root->left = buildTree(plList, ilList);
        root->right = buildTree(prList, irList);

        return root;
    }

    int findIndex(vector<int> &order, int val) {
        for (int i=0; i<order.size(); i++) {
            if (order[i] == val) {
                return i;
            }
        }

        return -1;
    }
};