package service

import (
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize/v2"
	"github.com/easysoft/zendata/src/model"
	constant "github.com/easysoft/zendata/src/utils/const"
	i118Utils "github.com/easysoft/zendata/src/utils/i118"
	logUtils "github.com/easysoft/zendata/src/utils/log"
	"github.com/easysoft/zendata/src/utils/vari"
	"github.com/mattn/go-runewidth"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"os"
	"strings"
)

const (
	size = 4
)

func ListRes() {
	res := map[string][size]string{}
	path := vari.ExeDir + "data" + string(os.PathSeparator) + "system"
	GetFilesAndDirs(path, &res)

	msg := ""
	nameWidth := 0
	titleWidth := 0
	for key, arr := range res {
		path := arr[0]
		if key == "yaml" {
			arr[2], arr[3] = readYamlInfo(path)
		} else if key == "excel" {
			arr[2], arr[3] = readExcelInfo(path)
		}

		res[key] = arr
		lent := runewidth.StringWidth(arr[1])
		if lent > nameWidth {
			nameWidth = lent
		}
		lent2 := runewidth.StringWidth(arr[2])
		if lent2 > titleWidth {
			titleWidth = lent2
		}
	}

	for _, arr := range res {
		name := arr[1]
		name = name  + strings.Repeat(" ", nameWidth - runewidth.StringWidth(name))
		title := arr[2]
		title = title  + strings.Repeat(" ", titleWidth - runewidth.StringWidth(title))

		msg = msg + fmt.Sprintf("%s  %s  %s\n", name, title, arr[3])
	}

	logUtils.Screen(msg)
}

func GetFilesAndDirs(path string, res *map[string][size]string)  {
	dir, err := ioutil.ReadDir(path)
	if err != nil {
		return
	}

	for _, fi := range dir {
		if fi.IsDir() {
			GetFilesAndDirs(path + constant.PthSep + fi.Name(), res)
		} else {
			name := fi.Name()
			arr := [size]string{}
			if strings.HasSuffix(name, ".yaml"){
				arr[0] = path + constant.PthSep + name
				arr[1] = path[strings.LastIndex(path, "system"):] + constant.PthSep + name
				(*res)["yaml"] = arr
			} else if strings.HasSuffix(name, ".xlsx") {
				arr[0] = path + constant.PthSep + name
				arr[1] = path[strings.LastIndex(path, "system"):] + constant.PthSep + name
				(*res)["excel"] = arr
			}
		}
	}
}

func readYamlInfo(path string) (title string, desc string) {
	configDef := model.DefData{}

	yamlContent, err := ioutil.ReadFile(path)
	if err != nil {
		logUtils.Screen(i118Utils.I118Prt.Sprintf("fail_to_read_file", path))
		return
	}
	err = yaml.Unmarshal(yamlContent, &configDef)
	if err != nil {
		logUtils.Screen(i118Utils.I118Prt.Sprintf("fail_to_parse_file", path))
		return
	}

	title = configDef.Title
	desc = configDef.Desc
	return
}

func readExcelInfo(path string) (title string, desc string) {
	excel, err := excelize.OpenFile(path)
	if err != nil {
		logUtils.Screen(i118Utils.I118Prt.Sprintf("fail_to_read_file", path))
		return
	}

	for index, sheet := range excel.GetSheetList() {
		if index > 0 {
			title = title + "|"
		}
		title = title + sheet
	}

	desc = i118Utils.I118Prt.Sprintf("excel_data")
	return
}