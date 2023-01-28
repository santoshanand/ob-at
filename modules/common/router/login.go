package router

import (
	"net/http"

	"github.com/santoshanand/at/modules/common/utils"
)

func (p *routes) login(w http.ResponseWriter, r *http.Request) {
	res := utils.H{"working": "working"}
	p.writeJSON(w, 200, res)
}
