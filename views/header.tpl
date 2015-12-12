<!DOCTYPE html>

<html>
<head>
<meta http-equiv="content-type" content="text/html; charset=utf-8" />
<title>Portale Servizi E' Cos&igrave;</title>
<meta name="keywords" content="" />
<meta name="description" content="" />
<link href="/static/css/default.css" rel="stylesheet" type="text/css" />
</head>

<body>
<!-- start header -->
<div id="header-bg">
	<div id="header">
		<div align="right">{{if .InSession}}
		Welcome, {{.First}} [<a href="http://{{.domainname}}/user/logout">Logout</a>|<a href="http://{{.domainname}}/user/profile">Profile</a>]
		{{else}}
		[<a href="http://{{.domainname}}/user/login/home">Login</a>]
		{{end}}
		</div>
		<div id="logo">
		</div>
			
		<div id="menu">
			<ul>
				<li class="active"><a href="http://{{.domainname}}/home">HOME</a></li>
				<li><a href="#">AUTOMEZZI</a></li>	
				{{if .Group == 2}}<li class="active"><a href="http://{{.domainname}}/user/manage/">SICUREZZA</a></li>{{end}}
			</ul>
		</div>
	</div>
</div>
<!-- end header -->
<!-- start page -->
<div id="page">