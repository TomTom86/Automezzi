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
	</div>
</div>