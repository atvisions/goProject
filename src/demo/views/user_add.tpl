<!doctype html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport"
          content="width=device-width, user-scalable=no, initial-scale=1.0, maximum-scale=1.0, minimum-scale=1.0">
    <meta http-equiv="X-UA-Compatible" content="ie=edge">
    <title>{{.title}}</title>
</head>
<body>
<h1>添加用户</h1>

<form action="addDo" method="post">
    ID：<input type="text" name="id">
     用户名：    <input type="text" name="username" value="">
    密 码：<input type="text" name="password" value="">
    <input type="submit" value="提交">


</form>
</body>
</html>