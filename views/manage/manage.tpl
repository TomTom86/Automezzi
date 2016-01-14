<div id="content">
	{{if .flash.error}}
		<h3>{{.flash.error}}</h3>
		&nbsp;
	{{end}}
	{{if .flash.notice}}
		<h3>{{.flash.notice}}</h3>
		&nbsp;
	{{end}}
		{{if .Errors}}
		{{range $rec := .Errors}}
		<h3>{{$rec}}</h3>
		{{end}}
	&nbsp;
	{{end}}
	
	
	<div id="content">
	<h1>Utenti</h1>
	<div>
	<br>
	</div>
	<table border="1" style="width:600px">
	{{.Rows}}
	</table>
	{{if .ShowNav}}
	<br>
	<div id="progressbar"></div>pointer in data set
	<div align="right">
	<a href="http://{{.domainname}}/manage/{{.order}}!0!{{.query}}">&lt;&lt;Start</a>&nbsp;&nbsp;&nbsp;&nbsp;
	{{if .showprev}}<a href="http://{{.domainname}}/manage/{{.order}}!{{.prev}}!{{.query}}">&lt;Prev</a>&nbsp;&nbsp;&nbsp;&nbsp;{{end}}
	{{if .next}}<a href="http://{{.domainname}}/manage/{{.order}}!{{.next}}!{{.query}}">Next&gt;</a>&nbsp;&nbsp;&nbsp;&nbsp;{{end}}
	<a href="http://{{.domainname}}/manage/{{.order}}!{{.end}}!{{.query}}">End&gt;&gt;</a>
	</div>
	{{end}}
	</div>
</div>