package parser

import (
	"fmt"
	"regexp"
	"time"
)

const (
	cityListRegexp      string = `<a href="(http://www.zhenai.com/zhenghun/[0-9a-z]+)"[^>]*>([^<]+)</a>`
	cityRegexp          string = `<th><a href="(http://album.zhenai.com/u/[0-9]+)"[^>]*>([^<]*)</a></th>`
	userRegexpAge       string = `<td><span class="label">年龄：</span>([^<]+)</td>`
	userRegexpHeight    string = `<td><span class="label">身高：</span>([^<]+)</td>`
	userRegexpEducation string = `<td><span class="label">学历：</span>([^<]+)</td>`
)

var (
	reCityList      *regexp.Regexp = regexp.MustCompile(cityListRegexp)
	reCity          *regexp.Regexp = regexp.MustCompile(cityRegexp)
	reUserAge       *regexp.Regexp = regexp.MustCompile(userRegexpAge)
	reUserHeight    *regexp.Regexp = regexp.MustCompile(userRegexpHeight)
	reUserEducation *regexp.Regexp = regexp.MustCompile(userRegexpEducation)
)

func EmptyParse(data []byte) []ParseRequest {
	return []ParseRequest{}
}

func CityListParseFunc(data []byte) []ParseRequest {

	matches := reCityList.FindAllSubmatch(data, -1)

	var limit int = 2

	var ret []ParseRequest

	for _, m := range matches {

		limit--
		if limit < 0 {
			break
		}

		url := string(m[1])
		city := string(m[2])

		ret = append(ret, ParseRequest{
			URL:       url,
			ParseFunc: CityParseFunc,
		})

		fmt.Printf("City：%s URL：%s\n", city, url)
	}

	return ret
}

func CityParseFunc(data []byte) []ParseRequest {

	matches := reCity.FindAllSubmatch(data, -1)

	var ret []ParseRequest

	var limit int = 5

	for _, m := range matches {

		limit--
		if limit < 0 {
			break
		}

		url := string(m[1])
		name := string(m[2])

		ret = append(ret, ParseRequest{
			URL:       url,
			ParseFunc: UserParseFunc,
		})

		fmt.Printf("Name：%s URL：%s\n", name, url)
	}

	return ret
}

func UserParseFunc(data []byte) []ParseRequest {

	matchAge := reUserAge.FindStringSubmatch(string(data))
	matchHeight := reUserHeight.FindStringSubmatch(string(data))
	matchEducation := reUserEducation.FindStringSubmatch(string(data))

	fmt.Println("Age：", matchAge[1], "Height：", matchHeight[1], "Education：", matchEducation[1])

	time.Sleep(500 * time.Millisecond)

	return []ParseRequest{}
}
