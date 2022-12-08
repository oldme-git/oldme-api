// Golang原生正则不支持零宽断言，该包对regexp2二次封装，方便使用。regexp2:https://github.com/dlclark/regexp2

package uregex

import "github.com/dlclark/regexp2"

func MatchAllString(pattern string, s string) (matches []string, err error) {
	var re = regexp2.MustCompile(pattern, 0)
	m, err := re.FindStringMatch(s)
	if err != nil {
		return
	}
	for m != nil {
		matches = append(matches, m.String())
		m, _ = re.FindNextMatch(m)
	}
	return
}
