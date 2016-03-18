<div id="content">
<h1>Registrazione Utente</h1>
&nbsp;

{{if .flash.error}}
<h3>{{.flash.error}}</h3>
&nbsp;
{{end}}{{if .flash.notice}}
<h3>{{.flash.notice}}</h3>
&nbsp;
{{end}}


  		<div class="col-md-12 content">
  			  <div class="panel panel-default">
                <div class="panel-heading">
                    Registrazione Utente
                </div>
                <div class="panel-body">
                        <form method="POST">
                        <table>
                        <tr>
                            <td>First name: </td>
                            <td><input name="first" type="text" value="{{.User.First}}" autofocus /> {{if .Errors.First}}{{.Errors.First}}{{end}}</td>
                        </tr>
                        <tr>
                            <td>Last name:</td>
                            <td><input name="last" type="text" value="{{.User.Last}}" /></td>
                        </tr>
                        <tr>
                            <td>Email address: </td>
                            <td><input name="email" type="text" value="{{.User.Email}}" /> {{if .Errors.Email}}{{.Errors.Email}}{{end}}</td>
                        </tr>
                        <tr>      
                            <td>Password (must be at least 6 characters): </td>
                            <td><input name="password" type="password" /> {{if .Errors.Password}}{{.Errors.Password}}{{end}}</td>
                        </tr>
                        <tr>      
                            <td>Confirm password: </td>
                            <td><input name="password2" type="password" /> {{if .Errors.Confirm}}{{.Errors.Confirm}}{{end}}</td>
                        </tr>
                        <tr><td>&nbsp;</td></tr>
                        <tr>
                            <td>&nbsp;</td><td><input type="submit" value="Register" /></td>
                        </tr>
                        </table>
                        </form>
                </div>
              </div>

  		</div>


