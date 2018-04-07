#include <iostream>
#include <vector>
#include <algorithm>
#include <string>

using namespace std;

class Solution {
public:
    Solution() {

    }

    string getPermutation(int n, int k) {
      vector<string> ans;
      string out;
      vector<bool> visited(n+1, false);

      permutSeqHelper(ans, out, visited, n);

      sort(ans.begin(), ans.end());
        
      return ans[k-1];
    }

    void permutSeqHelper(vector<string> &ans, string &out, vector<bool> &visited, int n) {
        if (out.size() == n) {
            ans.push_back(out);

            return;
        }

        for (int i=1; i<=n; i++) {
            if (visited[i]) {
                continue;
            }
            visited[i] = true;
            out.push_back('0'+i);
            permutSeqHelper(ans, out, visited, n);
            out.pop_back();
            visited[i] = false;
        }
    }
};

int main() {
    cout << "hello world" << endl;
    Solution* s = new Solution();
    cout << s->getPermutation(3,2) << endl;
    return 0;
}