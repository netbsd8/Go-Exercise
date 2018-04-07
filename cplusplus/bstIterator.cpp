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

class BSTIterator {
public:
    BSTIterator(TreeNode *root) {
        while (root) {
            s.push(root);
            root = root->left;
        }
    }

    bool hasNext() {
        return !s.empty();
    }

    int next() {
            TreeNode *t = s.top();
            TreeNode *res = s.top();
            s.pop();
            if (t->right) {
                t = t->right;
                while(t) {
                    s.push(t);
                    t = t->left;
                }
            }

            return res->val;
    }

private:
    stack<TreeNode*> s;
};

/**
 * Your BSTIterator will be called like this:
 * BSTIterator i = BSTIterator(root);
 * while (i.hasNext()) cout << i.next();
 */