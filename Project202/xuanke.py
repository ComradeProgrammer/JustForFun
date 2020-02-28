import urllib.request
import re
from http import cookiejar
from urllib import parse 
import sys

headers = {
    "User-Agent": "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/80.0.3987.116 Safari/537.36"
}
#cookie
cook_jar = cookiejar.CookieJar()
cook_hanlder = urllib.request.HTTPCookieProcessor(cook_jar)
opener = urllib.request.build_opener(cook_hanlder)
#教务发POST包模板
basic_form={
                "kcdmpx":"",
                "kcmcpx":"",
                "rlpx":"",
                "rwh":"",
                "zy":"",
                "qz":"",
                "token":"",
                "pageKclb":"",
                "pageXklb":"xslbxk",
                "pageXkmkdm":"ZYL",
                "pageXnxq":"",
                "pageKkxiaoqu":"",
                "pageKkyx":"",
                "pageKcmc":""
            }
def buaa_vpn_login():
    global headers
    global opener
    #请求登录页面
    init_url="https://e1.buaa.edu.cn/users/sign_in"
    init_req=urllib.request.Request(init_url)
    init_response=opener.open(init_req)
    init_page=init_response.read().decode("utf-8")
    #提取csrf-token字段
    init_regex=r'name=\"csrf-token\"\s+content=\"(.*)\"'
    csrf_token=re.search(init_regex,init_page).group(1)
    login_form={
        "utf8":"✓",
        "authenticity_token":csrf_token,
        "user[login]":"*************************",
        "user[password]":"********************",
        "user[dymatice_code]":"unknown",
        "commit":"登录 Login" 
    }
    login_str = parse.urlencode(login_form).encode('utf-8')
    login_request = urllib.request.Request(init_url,headers=headers,data=login_str)
    login_page=opener.open(login_request).read().decode("utf-8")
    print("vpn登录成功")
    # with open("test.html","w",encoding="utf-8") as f:
    #     f.write(login_page)

def vpn_jiaowu_login():
    global headers,opener
    jiaowu_url="https://10-200-21-61-7001.e1.buaa.edu.cn/ieas2.1/welcome"
    jiaowu_request = urllib.request.Request(jiaowu_url,headers=headers)
    jiaowu_response = opener.open(jiaowu_request)
    # jiaowu_page = jiaowu_response.read().decode("utf-8")
    # with open("test2.html","w",encoding="utf-8") as f:
    #     f.write(jiaowu_page)
    # print("VPN-教务登录成功")

def jiaowu_xuanke():
    global headers,opener,courseType,course,semester,basic_form
    token_regex=r'id=\"token\"\s*name=\"token\"\s*value=\"(.*)\"'

    #发包进入教务-专业课程选课-核心专业或一般专业
    zhuanye_url="https://10-200-21-61-7001.e1.buaa.edu.cn/ieas2.1/xslbxk/queryXsxk?pageXkmkdm=ZYL"
    zhuanye_request = urllib.request.Request(zhuanye_url,headers=headers)
    zhuanye_response = opener.open(zhuanye_request)
    zhuanye_page = zhuanye_response.read().decode("utf-8")
    print("进入选课界面")

    xuanke_url="https://10-200-21-61-7001.e1.buaa.edu.cn/ieas2.1/xslbxk/saveXsxk"
    query_url="https://10-200-21-61-7001.e1.buaa.edu.cn/ieas2.1/xslbxk/queryXsxkList"

    #发包点击查询按钮
    token=re.search(token_regex,zhuanye_page).group(1)
    query_form=basic_form.copy()
    query_form["token"]=token
    query_form["pageKclb"]=courseType
    query_form["pageXnxq"]=semester
    query_str = parse.urlencode(query_form).encode('utf-8')
    query_request = urllib.request.Request(query_url,
            headers=headers,data=query_str)
    query_page = opener.open(query_request).read().decode("utf-8")

    #开 始 爆 破
    while True:
        #提取csrf-token
        token=re.search(token_regex,query_page).group(1)
        #Сука блядь
        
        #构造选课请求并发出
        xuanke_form=basic_form.copy()
        xuanke_form["rwh"]=course
        xuanke_form["token"]=token
        xuanke_form["pageKclb"]=courseType
        xuanke_str = parse.urlencode(xuanke_form).encode('utf-8')
        xuanke_request = urllib.request.Request(xuanke_url,headers=headers,data=xuanke_str)
        xuanke_response=opener.open(xuanke_request)
        print("status"+str(xuanke_response.status))

        #重新模拟点击查询按钮为下一次发包做好准备
        token=re.search(token_regex,xuanke_response.read().decode("utf-8")).group(1)
        query_form=basic_form.copy()
        query_form["token"]=token
        query_form["pageKclb"]=courseType
        query_form["pageXnxq"]=semester
        query_request = urllib.request.Request(query_url,
        headers=headers,data=query_str)
        query_page = opener.open(query_request).read().decode("utf-8")
#主程序
# python xuanke.py "2019-20202" "J" "2019-2020-2-B3J063220-001"
semester=sys.argv[1]
courseType=sys.argv[2]  
course=sys.argv[3]
while True:
    try:
        buaa_vpn_login()
        vpn_jiaowu_login()
        jiaowu_xuanke()
    except Exception as e:
        print(e)



