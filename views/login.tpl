<html>
	<body>
		<h1>login</h1>
		<form action="/login" method="post" name = "mycount">
			用户名:<input type="text" name="username"><br>
			密码:<input type="password" name="pwd"><br>
           <label style="font-size:20px; color:yellow; left:150px;position:absolute; top:200px; diplay:block;">记住密码</label>
           <input type="checkbox" name="autologin" style="position:absolute; left:260px; top:230px;">
			<input type="submit" value="登录"><br>
			<a href="/regist">注册</a>
		</form>
	</body>
</html>