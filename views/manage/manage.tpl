<div id="content" class="manage">
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
		
		<tr>
		<th style="width:100px"><a href="http://{{.domainname}}/manage/{{if eq .order "id"}}-{{end}}id!{{.offset}}!{{.query}}">Id</a></th>
		<th style="width:100px"><a href="http://{{.domainname}}/manage/{{if eq .order "first"}}-{{end}}first!{{.offset}}!{{.query}}">Nome</a></th>
		<th style="width:100px"><a href="http://{{.domainname}}/manage/{{if eq .order "last"}}-{{end}}last!{{.offset}}!{{.query}}">Cognome</a></th>
		<th style="width:100px"><a href="http://{{.domainname}}/manage/{{if eq .order "email"}}-{{end}}email!{{.offset}}!{{.query}}">Email</a></th>
		<th style="width:100px">Modifica</th>
		</tr>
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