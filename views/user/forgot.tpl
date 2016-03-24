    <div class="modal-dialog">
            <div class="checkmodal-container">
            <div id="brand">&nbsp</div>
            </br>
  			  <div class="panel panel-default">
                <div class="panel-heading">
                     Recupera Password
                </div>                
                <div class="panel-body">
                        <form method="POST">
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
                        {{if .flash.notice}}
                        <h3>{{.flash.notice}}</h3>
                        &nbsp;
                        {{end}}    
                        <table>
                        <tr>
                            <td>Email address:</td>
                            <td><input name="email" type="email" autofocus /></td>
                        </tr>
                        <tr><td>&nbsp;</td></tr>
                        <tr>
                            <td>&nbsp;</td><td><input type="submit" value="Richiedi Reset" /></td>
                        </tr>
                        </table>
                        </form>

                </div>
                <div class="check-help">
                    <a href="http://localhost:8080/">Home</a> 
                </div>                
              </div>

  		</div>









