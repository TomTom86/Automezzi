<!DOCTYPE html>
<!--[if IE 8]><html class="no-js lt-ie9" lang="en" ><![endif]-->
<!--[if gt IE 8]><!--><html class="no-js" ><!--<![endif]-->
<html>
<head>
<meta charset="utf-8">
<meta http-equiv="X-UA-Compatible" content="IE=edge">
<meta name="viewport" content="width=device-width, initial-scale=1.0">
<title>Portale Servizi E' Cos&igrave;</title>
<!-- Fogli di stile -->
<link href="/static/bootstrap/3.3.6/css/bootstrap.css" rel="stylesheet" media="screen">
<link href="/static/bootstrap/3.3.6/css/stili-custom.css" rel="stylesheet" media="screen">
<link href="/static/bootstrap/3.3.6/css/default.css" rel="stylesheet" media="screen">
<!-- Modernizr -->
<script src="/static/bootstrap/3.3.6/js/modernizr.custom.js"></script>
<!-- respond.js per IE8 -->
<!--[if lt IE 9]>
<script src="/static/bootstrap/3.3.6/js/respond.min.js"></script>
<![endif]-->
</head>
<body>

{{if .InSession}}     
<nav class="navbar navbar-default navbar-static-top">
    <div class="container-fluid">
		<!-- Brand and toggle get grouped for better mobile display -->
		<div class="navbar-header">
			<button type="button" class="navbar-toggle navbar-toggle-sidebar collapsed">
			MENU
			</button>
			<button type="button" class="navbar-toggle collapsed" data-toggle="collapse" data-target="#bs-example-navbar-collapse-1">
				<span class="sr-only">Toggle navigation</span>
				<span class="icon-bar"></span>
				<span class="icon-bar"></span>
				<span class="icon-bar"></span>
			</button>
			<a class="navbar-brand" href="#">
				<img id="brand" src="/static/img/logo80.png"> </img>
			</a>
	    </div>

		<!-- Collect the nav links, forms, and other content for toggling -->
            <div class="collapse navbar-collapse" id="bs-example-navbar-collapse-1">      

                <ul class="nav navbar-nav navbar-right">
                         
                        <li class="dropdown ">
                                <a href="#" class="dropdown-toggle" data-toggle="dropdown" role="button" aria-expanded="false">
                                    <span class="glyphicon glyphicon-user"></span> Utente
                                    <span class="caret"></span></a>
                                    <ul class="dropdown-menu" role="menu">
                                        <li class=""><a href="http://{{.domainname}}/user/profile"><span class="fa glyphicon glyphicon-cog  fa-lg"></span> Aggiorna Profilo </a></li>
                                        <li class=""><a href="#"><span class="fa glyphicon glyphicon-envelope fa-lg"></span> Messaggi <span class="fa badge  fa-lg pull-right"> 42 </span></a></li>
                                        <li class=""><a href="#">Autorizzazioni</a></li>
                                        <li class="divider"></li>
                                        <li><a href="http://{{.domainname}}/user/logout"><span class="fa glyphicon glyphicon-log-out fa-lg"></span> Logout</a></li>
                                    </ul>
                            
                        </li>
                </ul>
			</div><!-- /.navbar-collapse -->
	</div><!-- /.container-fluid -->
</nav> 
    <div class="container-fluid main-container">
  		<div class="col-md-2 sidebar">
  			<div class="row">
	
                <!-- uncomment code for absolute positioning tweek see top comment in css -->
                <div class="absolute-wrapper"> </div>
                <!-- Menu -->
                <div class="side-menu">
                    <nav class="navbar navbar-default" role="navigation">
                        <!-- Main Menu -->
                        <div class="side-menu-container">
                            <ul class="nav navbar-nav">
                                <li><a href="http://{{.domainname}}/"><span class="glyphicon glyphicon-dashboard"></span> Home</a></li>

                                <!-- Dropdown-->
                                {{if .Automezzi}}
                                <li class="panel panel-default" id="dropdown">
                                    <a data-toggle="collapse" href="#dropdown-lvl1">
                                        <span class="fa glyphicon glyphicon-road fa-lg"></span> Automezzi <span class="caret"></span>
                                    </a>

                                    <!-- Dropdown level 1 -->
                                    <div id="dropdown-lvl1" class="panel-collapse collapse">
                                        <div class="panel-body">
                                            <ul class="nav navbar-nav">
                                                <li><a href="#">Veicoli</a></li>
                                                <li><a href="#">Movimenti</a></li>
                                                <li><a href="#">Rifornimenti</a></li>
                                                <li><a href="#">Spese</a></li>
                                                <li><a href="#">Incidenti</a></li>
                                                <li><a href="#">Multe</a></li>

                                                <!-- Dropdown level 2 -->
                                                <li class="panel panel-default" id="dropdown">
                                                    <a data-toggle="collapse" href="#dropdown-lvl2">
                                                        <span class="glyphicon glyphicon-off"></span> Veicoli <span class="caret"></span>
                                                    </a>
                                                    <div id="dropdown-lvl2" class="panel-collapse collapse">
                                                        <div class="panel-body">
                                                            <ul class="nav navbar-nav">
                                                                <li><a href="#">Aggiungi</a></li>
                                                                <li><a href="#">Visualizza</a></li>
                                                                <li><a href="#">Esporta</a></li>
                                                            </ul>
                                                        </div>
                                                    </div>
                                                </li>
                                            </ul>
                                        </div>
                                    </div>
                                </li>
                                {{end}}
                                {{if .Admin}}
                                <li><a href="http://{{.domainname}}/manage/id!0!id__gte,0"><span class="glyphicon glyphicon-cog"></span> Sicurezza</a></li>
                                <li><a href="http://{{.domainname}}/appadmin/index/id!0!id__gte,0"><span class="fa glyphicon glyphicon-cog fa-lg"></span> Pannello Admin</a></li>
                                {{end}}
                            </ul>
                        </div><!-- /.navbar-collapse -->
                    </nav>

                </div>
            </div> 
        </div>

{{end}}