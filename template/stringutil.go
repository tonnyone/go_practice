package template

import (
	"bytes"
	"fmt"
	"regexp"
	ttemplate "text/template"
	"text/template/parse"
)

// SuperReplace 字符串部分匹配算法, obj的值不能有换行符
func SuperReplace(ticket string, tmpl string, obj map[string]any) (string, error) {
	t := ttemplate.New("super_replace").Delims("[[[", "]]]")
	p, err := t.Parse(tmpl)
	if err != nil {
		return "", err
	}
	var b bytes.Buffer
	if err := p.Execute(&b, obj); err != nil {
		return "", err
	}
	replaceStr := b.String()
	fmt.Printf("replStr: %s", replaceStr)
	regex := genCoverMarchRegex(p.Tree)
	fmt.Printf("regex: %s", regex)
	result, err := regexReplace(regex, ticket, replaceStr)
	if err != nil {
		return "", err
	}
	return result, nil
}

func genMarchRegex(tree *parse.Tree) string {
	var regex string
	regexHasFlag := "[\\s\\S]+"
	regexAllFlag := "[\\s\\S]*"
	regex = fmt.Sprintf("%s\n", regexAllFlag)
	for _, n := range tree.Root.Nodes {
		if n.Type() != parse.NodeText {
			regex = fmt.Sprintf("%s%s", regex, regexHasFlag)
		} else {
			regex = fmt.Sprintf("%s%s", regex, n.String())
		}
	}
	regex = fmt.Sprintf("%s%s", regex, regexAllFlag)
	return regex
}

func genCoverMarchRegex(tree *parse.Tree) string {
	var start, end string
	for i, n := range tree.Root.Nodes {
		if i == 0 && n.Type() == parse.NodeText {
			start = n.String()
		}
		if i != 0 {
			if n.Type() == parse.NodeText {
				end = n.String()
			} else {
				end = ""
			}
		}
	}
	regex := fmt.Sprintf("%s%s%s", start, "[\\s\\S]+", end)
	return regex
}

func regexReplace(regex string, origin, replaceStr string) (string, error) {
	re := regexp.MustCompile(regex)
	originSubmatch := re.FindStringSubmatch(origin)
	replaceStrSubMatch := re.FindStringSubmatch(replaceStr)
	originSubmatchLen := len(originSubmatch)
	repalceStrSubMatchLen := len(replaceStrSubMatch)
	if originSubmatchLen != repalceStrSubMatchLen {
		return "", fmt.Errorf("regex replace group not equal:regex: \n%s \norigin%d: \n%s \nreplace%d: \n%s",
			regex, originSubmatchLen, origin, repalceStrSubMatchLen, replaceStr)
	}
	result := re.ReplaceAllString(origin, replaceStr)
	return result, nil
}

func genMarchGroupValues(origin string, regex string) []string {
	fmt.Printf("regex: %s", regex)
	re := regexp.MustCompile(regex)
	subMatch := re.FindStringSubmatch(origin)
	fmt.Printf("%v,%d", subMatch, len(subMatch))
	return subMatch[1:]
}

func findStarEndFlag(tree *parse.Tree) (string, string) {
	var start, end string
	for i, n := range tree.Root.Nodes {
		if i == 0 && n.Type() == parse.NodeText {
			start = n.String()
			continue
		}
		if n.Type() == parse.NodeText {
			end = n.String()
		} else {
			end = ""
		}
	}
	return start, end
}

func ReplaceAllStringSubmatchFunc(re *regexp.Regexp, str string, repl func([]string) string) string {
	result := ""
	lastIndex := 0

	for _, v := range re.FindAllSubmatchIndex([]byte(str), -1) {
		groups := []string{}
		for i := 0; i < len(v); i += 2 {
			groups = append(groups, str[v[i]:v[i+1]])
		}

		result += str[lastIndex:v[0]] + repl(groups)
		lastIndex = v[1]
	}

	return result + str[lastIndex:]
}
