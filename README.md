### 邮件验证码认证 不需要缓存/redis


```go
	// 连接
	s, err := NewS("smtp.qq.com", 465, "your mail", "your key")
    
	if err != nil {
		println(err.Error())
	} else {
		println("connected")
	}

	// 发送
	encap, err := s.Send("gonorth@qq.com", "[MicroFlow] signup")
	println("验证码加密后: ", encap)

	if err != nil {
		println(err.Error())
	} else {
		println("send")
	}

	// 验证
	same, err := s.Verify(encap, "11111")
	println("验证码相同: ", same)
	if err != nil {
		println(err.Error())
	} else {
		println("verified")
	}
```