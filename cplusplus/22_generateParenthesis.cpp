#include <iostream>
#include <vector>
#include <set>
#include <algorithm>

using namespace std;

class Solution {
public:
    vector<string> generateParenthesis(int n) {
        set<string> t;
        if (n == 0) t.insert("");
        else {
            vector<string> pre = generateParenthesis(n-1);
            for (auto a : pre) {
                //string temp = '(' + a + ')';
                for (int i=0; i<a.size(); i++) {
                    if (a[i] == '(') {
                        string temp = a;
                        temp.insert(temp.begin()+i+1, '(');
                        temp.insert(temp.begin()+i+2, ')');
                        t.insert(temp);
                    }
                }
                t.insert("()" + a);
                t.insert(a + "()");
            }
        }

        return vector<string>(t.begin(), t.end());
    }

    vector<string> generateParenthesis2(int n) {
        vector<string> res;
        helper(n, n, "", res);
        return res;
    }

    void helper(int left, int right, string out, vector<string> &res) {
        if (left > right) {
            cout << "left > right " << left << " " << right << " return" << endl;
            return;
        }
        if (left == 0 && right == 0) {
            cout << "insert out: " << out << endl;
            res.push_back(out);
        } 
        else {
            cout << "1. left: " << left << "; right: " << right << "; out: " << out << endl;
            if (left > 0) helper(left-1, right, out+'(', res);
            cout << "2. left: " << left << "; right: " << right << "; out: " << out << endl;
            if (right > 0) helper(left, right-1, out+')', res);
            cout << "3. left: " << left << "; right: " << right << "; out: " << out << endl;
        }
    }
};

int main() {
    Solution* s = new Solution();
    vector<string> ans = s->generateParenthesis2(2);
    for (auto i=ans.begin(); i!=ans.end(); i++) {
        cout << *i << endl;
    }
    cout << ans.size() << endl;

    return 0;
}