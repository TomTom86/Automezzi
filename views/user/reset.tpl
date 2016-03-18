<div id="content">
<h1>Reset password</h1>
&nbsp;
{{if .flash.error}}
<h3>{{.flash.error}}</h3>
&nbsp;
{{end}}{{if .flash.notice}}
<h3>{{.flash.notice}}</h3>
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
                    Home
                </div>
                <div class="panel-body">
                    <form method="POST">
                    <table>
                    <tr>      
                        <td>New password (must be at least 6 characters):</td>
                        <td><input name="password" type="password" /></td>
                    </tr>
                    <tr>      
                        <td>Confirm new password:</td>
                        <td><input name="password2" type="password" /></td>
                    </tr>
                    <tr><td>&nbsp;</td></tr>
                    <tr>
                        <td>&nbsp;</td><td><input type="submit" value="Reset password" /></td>
                    </tr>
                    </table>
                    </form>
                </div>
              </div>

  		</div>