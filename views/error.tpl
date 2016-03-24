    <div class="modal-dialog">
                    <div class="loginmodal-container">
                    <div id="brand">&nbsp</div>
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
                        &nbsp
                        &nbsp
                    <div class="login-help">
                        <a href="http://localhost:8080/">Home</a>
                    </div>
                    </div>
                </div>
            </div>
    </div>