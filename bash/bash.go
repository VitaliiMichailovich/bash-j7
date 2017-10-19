package bash

import (
	"net/http"
	"io/ioutil"
	"strings"
)

func BashQuote() string {
	response, err := http.Get("http://bash.im/forweb/?u")
	if err != nil {
		return err.Error()
	}
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return err.Error()
	}
	bodyP := strings.Replace(string(body), "var borq='';\nborq += '<' + 'div id=\"b_q\">", "", 1)
	bodyP = strings.Replace(bodyP, "<' + 'small><' + 'a href=\"http://bash.im/\" target=\"_blank\" title=\"bash.im откроется в новом окне\">Больше на bash.im!<' + '/a><' + '/small><' + '/div>';", "", 1)
	bodyP = strings.Replace(bodyP, "' + '", "", -1)
	bodyP = strings.Replace(bodyP, "</div>", "", 1)
	bodyP = strings.Replace(bodyP, "document.write(borq);", "", 1)
	bodyZ := strings.Split(bodyP, "<div id=\"b_q_t\" style=\"padding: 1em 0;\">")
	return bodyZ[1]
}