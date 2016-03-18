<div id="content">
<h1>Reset password</h1>
&nbsp;
{{if .flash.error}}
<h3>{{.flash.error}}</h3>
&nbsp;
{{end}}
{{if .Errors}}
{{range $rec := .Errors}}
<h3>{{$rec}}</h3>
{{end}}
&nbsp;
{{end}}


  		<div class="col-md-12 content">
  			  <div class="panel panel-default">
                <div class="panel-heading">
                    Recupera Password
                </div>
                <div class="panel-body">
                        <form method="POST">
                        <table>
                        <tr>
                            <td>Email address:</td>
                            <td><input name="email" type="text" autofocus /></td>
                        </tr>
                        <tr><td>&nbsp;</td></tr>
                        <tr>
                            <td>&nbsp;</td><td><input type="submit" value="Request reset" /></td>
                        </tr>
                        </table>
                        </form>

                </div>
              </div>

  		</div>


