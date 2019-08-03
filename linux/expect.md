脚本一例

```tcl
#!/usr/bin/expect -f
set timeout -1


set IP [lindex $argv 0]

spawn /usr/bin/ssh -l aaa 10.1.1.1  -q

expect  {

    "yes/no?" {

            send "yes\r"
            expect "*assword*" {
			send "pass\r"
			expect "*bbb87 ~]$ " {
				send {kubectl exec -it `kubectl get pods |grep -Eo "aaa-ccc[^ ]*"` bash}
				send "\r"
				expect "*$ "
				send "cd\r"
				expect "*$ "
				send "ps -ef|grep ccc\r"
				expect "*$ "
				send "cd aaa-ccc;pwd;ls\r"
			}
		}
    }
    "*assword*" {
            send "pass\r"
			expect "*bbb87 ~]$ " {
				send {kubectl exec -it `kubectl get pods |grep -Eo "aaa-ccc[^ ]*"` bash}
				send "\r"
				expect "*$ "
				send "cd\r"
				expect "*$ "
				send "ps -ef|grep ccc\r"
				expect "*$ "
				send "cd aaa-ccc;pwd;ls\r"
			}
        }
	"#"{
		interact
	}
	"$"{
		interact
	}
	"logout"{
		return
	}

}

# spawn  su aaa

interact
```

