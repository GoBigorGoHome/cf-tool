/* 用正则表达式获得Codeforces的提交语言列表。
*/

#include <regex>
#include <iostream>


using namespace std;

#include <regex>
#include <iostream>

int main() {
    string line;
    string s = R"_([\s\S]*?value="(.+?)"[\s\S]*?>([\s\S]+?)<[\s\S]*?)_";
    cout << s << '\n';
    regex Pattern(s);
    // getline(cin, line) 返回 cin。
    while (getline(cin, line)) {
        smatch Match;
        if (regex_search(line, Match, Pattern)) {
            cout << '"' << Match[1] << "\": \"" << Match[2] << "\",\n";
        }
    }
}