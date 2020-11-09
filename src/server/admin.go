package server

import (
	"encoding/json"
	"github.com/easysoft/zendata/src/model"
	defServer "github.com/easysoft/zendata/src/server/def"
	commonUtils "github.com/easysoft/zendata/src/utils/common"
	"io"
	"io/ioutil"
	"net/http"
)

func AdminHandler(writer http.ResponseWriter, req *http.Request) {
	setupCORS(&writer, req)

	bytes, err := ioutil.ReadAll(req.Body)
	if len(bytes) == 0 {
		return
	}

	reqData := model.ReqData{}
	err = ParserJsonReq(bytes, &reqData)
	if err != nil {
		outputErr(err, writer)
		return
	}

	ret := model.ResData{ Code: 1, Msg: "success"}
	if reqData.Action == "listDef" {
		ret.Data = defServer.List()
	} else if reqData.Action == "getDef" {
		def := defServer.Get(reqData.Id)
		def.Folder = commonUtils.GetFolder(def.Path)
		ret.Data = def
	} else if reqData.Action == "saveDef" {
		def := convertDef(reqData.Data)

		if def.Id == 0 {
			defServer.Create(&def)
		} else {
			defServer.Update(&def)
		}

		ret.Data = def
	}

	jsonStr, _ := json.Marshal(ret)

	io.WriteString(writer, string(jsonStr))
}
