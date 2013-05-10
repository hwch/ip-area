package main

import (
        "flag"
        "fmt"
        "io/ioutil"
        "net/http"
        "strconv"
        "strings"
)

func UStringToRune(s string) ([]rune, error) {

        s0 := strings.Replace(s, "u", "", -1)
        sb := strings.Split(s0, "\\")
        r := make([]rune, len(sb))
        i := 0
        for _, v := range sb {
                if v == "" {
                        continue
                }
                vv, err := strconv.ParseUint(v, 16, 32)
                if err != nil {
                        return nil, err
                }
                r[i] = rune(vv)
                i++
        }
        return r, nil
}

func main() {
        ip := flag.String("ip", "", "The IP address")
        flag.Parse()
        if *ip == "" {
                flag.PrintDefaults()
                return
        }

        /* e, _ := UStringToRune("\\u534e\\u4e2d\\u79d1\\u6280\\u5927\\u5b66\\u4e1c\\u6821\\u533a\\u6559\\u80b2\\u7f51")
           fmt.Printf("%s\n", string(e))
           return
        */

        req_str := fmt.Sprintf("http://int.dpool.sina.com.cn/iplookup/iplookup.php?format=js&ip=%s", *ip)
        resp, err := http.Get(req_str)
        if err != nil {
                fmt.Printf("Error:%v\n", err)
                return
        }
        defer resp.Body.Close()
        body, err := ioutil.ReadAll(resp.Body)
        tlen := len(body)
        ns := strings.Index(string(body), "{")
        if ns == -1 {
                fmt.Printf("未知IP地址[%s]\n", *ip)
                return
        }

        s1 := string(body[ns+1 : tlen-2])
        qnv := strings.Replace(s1, "\"", "", -1)
        nv := strings.Split(qnv, ",")
        country := ""
        province := ""
        city := ""
        district := ""
        ISP := ""
        vtype := ""
        desc := ""

        for _, v := range nv {
                vv := strings.Split(v, ":")
                switch vv[0] {
                case "ret":
                        if len(vv[1]) == 0 || vv[1] == "0" {
                                fmt.Printf("未知IP地址[%s]\n", *ip)
                                return
                        }
                case "country":
                        if len(vv[1]) != 0 {
                                if v, err := UStringToRune(vv[1]); err != nil {
                                        fmt.Print("返回字符串格式非法\n")
                                        return
                                } else {
                                        country = string(v)
                                }
                        }
                case "province":
                        if len(vv[1]) != 0 {
                                if v, err := UStringToRune(vv[1]); err != nil {
                                        fmt.Print("返回字符串格式非法\n")
                                        return
                                } else {
                                        province = string(v)
                                }
                        }
                case "city":
                        if len(vv[1]) != 0 {
                                if v, err := UStringToRune(vv[1]); err != nil {
                                        fmt.Print("返回字符串格式非法\n")
                                        return
                                } else {
                                        city = string(v)
                                }
                        }
                case "district":
                        if len(vv[1]) != 0 {
                                if v, err := UStringToRune(vv[1]); err != nil {
                                        fmt.Print("返回字符串格式非法\n")
                                        return
                                } else {
                                        district = string(v)
                                }
                        }
                case "ISP":
                        if len(vv[1]) != 0 {
                                if v, err := UStringToRune(vv[1]); err != nil {
                                        fmt.Print("返回字符串格式非法\n")
                                        return
                                } else {
                                        ISP = string(v)
                                }
                        }
                case "type":
                        if len(vv[1]) != 0 {
                                if v, err := UStringToRune(vv[1]); err != nil {
                                        fmt.Print("返回字符串格式非法\n")
                                        return
                                } else {
                                        vtype = string(v)
                                }
                        }
                case "desc":
                        if len(vv[1]) != 0 {
                                if v, err := UStringToRune(vv[1]); err != nil {
                                        fmt.Print("返回字符串格式非法\n")
                                        return
                                } else {
                                        desc = string(v)
                                }
                        }
                }
        }

        fmt.Printf("IP 详细信息:\n\tIP: %s\n\t国家:%s\n\t省份:%s\n\t"+
                "城市:%s\n\t区:%s\n\tISP:%s\n\t类型:%s\n\t其他:%s\n",
                *ip, country, province, city, district, ISP, vtype, desc)
}
