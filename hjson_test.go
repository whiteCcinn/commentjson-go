package commentjson_go_test

import (
	"./"
	"fmt"
	"io/ioutil"
	"os"
	"testing"
)

func TestHJSON(t *testing.T) {
	cases := []struct {
		orig string
		want string
	}{
		{ // 1
			orig: `{"foo": "bar", }`,
			want: `{"foo":"bar"}`,
		},
		{ // 2
			orig: `  "foo": "bar",  `,
			want: `{"foo":"bar"}`,
		},
		{ // 3
			orig: `  "foo": "bar"  `,
			want: `{"foo":"bar"}`,
		},
		{ // 4
			orig: `{foo: bar}`,
			want: `{"foo":"bar"}`,
		},
		{ // 5
			orig: `{foo: bar bar}`,
			want: `{"foo":"bar bar"}`,
		},
		{ // 6
			orig: `{"foo" : "bar", "ding": "bat"}`,
			want: `{"foo":"bar","ding":"bat"}`,
		},
		{ // 7
			orig: `{ foo : bar, ding: bat}`,
			want: `{"foo":"bar","ding":"bat"}`,
		},
		{ // 8
			orig: `{foo:  false  }`,
			want: `{"foo":false}`,
		},
		{ // 9
			orig: `{ foo:  [1,2,3,4]  }`,
			want: `{"foo":[1,2,3,4]}`,
		},
		{ // 10
			orig: `{ foo:  [ 1 , 2 , 3 , 4 ]  }`,
			want: `{"foo":[1,2,3,4]}`,
		},
		{ // 11
			orig: "{ foo:  [  \n 1 \n 2  \n 3  \n  4 \n]\n}",
			want: `{"foo":[1,2,3,4]}`,
		},
		{ // 12
			orig: `{ foo:  [ "1", "2", "3", "4",  ]  }`,
			want: `{"foo":["1","2","3","4"]}`,
		},
		{ // 13
			orig: `{ 日本語:  [ "1", "2", "3", "4",  ]  }`,
			want: `{"日本語":["1","2","3","4"]}`,
		},
		{ // 14
			orig: `
# junk
foo: "bar",
`,
			want: `{"foo":"bar"}`,
		},
		{ // 15
			orig: `
foo: '''
bar
''',
`,
			want: `{"foo":"bar"}`,
		},
		{ // 16
			orig: `
// comment
foo: """
bar
"""
// another
`,
			want: `{"foo":"bar"}`,
		},
		{ // 17
			orig: `
/* comment
  whatever */
foo: """
bar
"""
/* another
`,
			want: `{"foo":"bar"}`,
		},
		{ // 18
			orig: `
foo: /Users/nickg
`,
			want: `{"foo":"/Users/nickg"}`,
		},
		{ // 19
			orig: `
[
    "tbllog_event",
    "tbllog_login",
    "tbllog_online",
    "tbllog_pay",
    "tbllog_player",
    "tbllog_quit",
    "tbllog_role",
    "t_log_test_1"
]
`,
			want: `["tbllog_event","tbllog_login","tbllog_online","tbllog_pay","tbllog_player","tbllog_quit","tbllog_role","t_log_test_1"]`,
		},
		{ // 20
			orig: `
 // test
{
    "8"      : "uiaVtWPYBlofk" 
} 
`,
			want: `{"8":"uiaVtWPYBlofk"}`,
		},
		{ // 21
			orig: `
 // app_id 规则
// 内部 1000-2000
// 外部 2299-3000
{
    "8"      : "uiaVtWPYBlofk" 
} 
`,
			want: `{"8":"uiaVtWPYBlofk"}`,
		},
		{ // 22
			orig: `
/*
`,
			want: ``,
		},
		{ // 23
			orig: `
/**
`,
			want: ``,
		},
		{ // 24
			orig: `
/*/
`,
			want: ``,
		},
		{ // 25
			orig: `
// app_id 规则
// 内部：1-10000
// 外部：100001-199999
{
        // 这里是注释
    "8"      : "uiaVtWPYBlofk" // 这里是注释
   /* 测试 */,
	"9": "uiaVtWPYBlofk" /*继续测试*/,
    "10": /*测试*/"uiaVtWPYBlofk"
//
} 
`,
			want: `{"8":"uiaVtWPYBlofk","9":"uiaVtWPYBlofk","10":"uiaVtWPYBlofk"}`,
		},
	}

	for num, tt := range cases {
		got := commentjson_go.ToJSON([]byte(tt.orig))
		if tt.want != string(got) {
			t.Errorf("%d: want %s got %s", num+1, tt.want, got)
		}
	}
}

func TestJsonFIle(t *testing.T) {
	// 打开json文件
	jsonFile, err := os.Open("test.json")

	// 最好要处理以下错误
	if err != nil {
		fmt.Println(err)
	}

	// 要记得关闭
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	res := commentjson_go.ToJSON(byteValue)

	fmt.Println(string(res))
}
