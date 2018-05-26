#include <iostream>
#include <vector>
#include <unordered_set>
#include <stack>
#include <string>
#include <algorithm>
#include <bits/stdc++.h>

using namespace std;

struct TreeLinkNode {
    TreeLinkNode    *left;
    TreeLinkNode    *right;
    TreeLinkNode    *next;
    int         val;
    TreeLinkNode(int x): val(x), left(NULL), right(NULL), next(NULL) {}
};

class Solution {
public:
    void connect(TreeLinkNode* root) {
        if (root == NULL) {
            return;
        }

        queue<TreeLinkNode*> q;
        q.push(root);

        while(!q.empty()) {
            int size = q.size();
            TreeLinkNode* dummy = new TreeLinkNode(-1);
            TreeLinkNode* cur = dummy;

            for (int i=0; i<size; i++) {
                cur->next = q.front();
                TreeLinkNode* t = q.front();
                q.pop();
                if (t->left) {
                    q.push(t->left);
                }
                if (t->right) {
                    q.push(t->right);
                }
                cur = cur->next;
            }
        }

        return;
    }

    // recursive
    void method2(TreeLinkNode* root) {
        if (!root) {
            return;
        }

        // find the next for branch's right
        TreeLinkNode* p = root->next;
        while(p) {
            if (p->left) {
                p = p->left;
                break;
            }

            if (p->right) {
                p = p->right;
                break;
            }

            p = p->next;
        }

        if (root->right) {
            root->right->next = p;
        }
        if (root->left) {
            if (root->right) {
                root->left->next = root->right;
            }
            else {
                root->left->next = p;
            }
        }

        method2(root->left);
        method2(root->right);
    }

};