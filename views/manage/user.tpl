

  		<div class="col-md-10 content">
  			  <div class="panel panel-default">
                <div class="panel-heading">
                    <h1>{{.UFirst}} {{.ULast}}</h1>
                </div>
                <div class="panel-body">
                        <div id="content">
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
                        <form method="POST">
                        <table>
                        <tr>
                            <td>First name:</td>
                            <td><input name="first" type="text" value="{{.UFirst}}" /></td>
                        </tr>
                        <tr>
                            <td>Last name:</td>
                            <td><input name="last" type="text" value="{{.ULast}}"/></td>
                        </tr>
                        <tr>
                            <td>Email address:</td>
                            <td><input name="email" type="text" value="{{.UEmail}}"/></td>
                        </tr>
                        <tr>
                            <td>Privilegi Utente:</td>
                            <td><select name="userlvl" size="1">
                            {{.Userlvllist}}
                            </select>
                            </td>
                        </tr>
                        <tr>
                            <td>Utente Bloccato:</td>
                            {{.Checkbloccato}}
                            </td>
                        </tr>
                        <tr><td>&nbsp;</td></tr>
                        <tr>
                        <td>Optional:</td>
                        </tr>
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
                        <td>
                            <fieldset>
                            <legend>Strumenti:</legend>
                            {{.Checkautomezzi}}
                            {{.Checkservizi}}
                            </fieldset>
                        </td>
                        </tr>
                        <tr><td>&nbsp;</td></tr>
                        <tr>
                            <td>&nbsp;</td><td><input type="submit" value="Update" /></td>
                        </tr>
                        </table>
                        <a href="http://{{.domainname}}/manage/id!0!id__gte,0" >Indietro</a>

                        </form>
                        </div>
                </div>
              </div>

  		</div>