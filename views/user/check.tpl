    <div class="modal-dialog">
            <div class="checkmodal-container">
            <div id="brand">&nbsp</div>
            </br>
  			  <div class="panel panel-default">
                <div class="panel-heading">
                    Verifica Account
                </div>                
                <div class="panel-body">
                </br>
                        {{if .Verified}}
                        <h2>Il tuo account è stato verificato.</h2>
                        {{else}}
                        <h2>Il tuo account <b>NON</b> è verificato.</h2>
                        {{end}}
                        </br>
                </div>
                    <div class="check-help">
                        <a href="http://localhost:8080/">Home</a> 
                    </div>                
              </div>

  		</div>