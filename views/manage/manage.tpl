

  		<div class="col-md-10 content">
  			  <div class="panel panel-default">
                <div class="panel-heading">
                   <h1>Utenti</h1>
                </div>
                <div class="panel-body">
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
                </div>
              </div>

  		</div>
          
          <div class="container">
    <h3>The columns titles are merged with the filters inputs thanks to the placeholders attributes</h3>
    <hr>
    <p>Inspired by this <a href="http://bootsnipp.com/snippets/featured/panel-tables-with-filter">snippet</a></p>
    <div class="row">
        <div class="panel panel-primary filterable">
            <div class="panel-heading">
                <h3 class="panel-title">Users</h3>
                <div class="pull-right">
                    <button class="btn btn-default btn-xs btn-filter"><span class="glyphicon glyphicon-filter"></span> Filter</button>
                </div>
            </div>
            <table class="table">
                <thead>
                    <tr class="filters">
                        <th><input type="text" class="form-control" placeholder="#" disabled></th>
                        <th><input type="text" class="form-control" placeholder="First Name" disabled></th>
                        <th><input type="text" class="form-control" placeholder="Last Name" disabled></th>
                        <th><input type="text" class="form-control" placeholder="Username" disabled></th>
                    </tr>
                </thead>
                <tbody>
                    <tr>
                        <td>1</td>
                        <td>Mark</td>
                        <td>Otto</td>
                        <td>@mdo</td>
                    </tr>
                    <tr>
                        <td>2</td>
                        <td>Jacob</td>
                        <td>Thornton</td>
                        <td>@fat</td>
                    </tr>
                    <tr>
                        <td>3</td>
                        <td>Larry</td>
                        <td>the Bird</td>
                        <td>@twitter</td>
                    </tr>
                </tbody>
            </table>
        </div>
    </div>
</div>