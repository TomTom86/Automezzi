




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
                    <form method="POST">
                        <input type="text" name="email" placeholder="Username">
                        <input type="password" name="password" placeholder="Password">
                        <input type="submit" name="login" value="Login" class="login loginmodal-submit" value="Login">
                    </form>
                        
                    <div class="login-help">
                        <a href="http://localhost:8080/user/register">Registrati</a> - <a href="http://localhost:8080/user/forgot">Password dimenticata</a>
                    </div>
                    </div>
                </div>
            </div>
    </div>




