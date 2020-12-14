# commentjson-go


## Api

- commentjson_go.ToJSON(a []byte)

```json
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
```
compression

```json
["tbllog_event","tbllog_login","tbllog_online","tbllog_pay","tbllog_player","tbllog_quit","tbllog_role","t_log_test_1"]
```


```json
[
  {
    ## 这是注释
    "log_name": "abc", ## 这也是注释
    "event_name": "ddd",
    "event_type": "sss",
    "metadata": {
      "#account_id": {
        "field": "dd",
        "default": "{dd}"
      },
      "#time": {
        "field": "null",
        "default": ""
      }
    }
  },## 这里注释
  {
    "log_name": "null",
    "event_name": "null",
    "event_type": "null",
    "metadata": {
      "#account_id": {
        "field": "null",
        "default": "{null}"
      },
      "#time": {
        "field": "null",
        "default": ""
      }
    }
  }
]
```

compression

```json
[{"log_name":"abc","event_name":"ddd","event_type":"sss","metadata":{"#account_id":{"field":"dd","default":"{dd}"},"#time":{"field":"null","default":""}}},{"log_name":"null","event_name":"null","event_type":"null","metadata":{"#account_id":{"field":"null","default":"{null}"},"#time":{"field":"null","default":""}}}]
```
