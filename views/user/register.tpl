
    <div class="modal-dialog">
                    <div class="registrationmodal-container">
                    <div id="brand">&nbsp</div>
                    &nbsp;
                    <h1 class="dark-grey">Registrazione</h3>
                                        &nbsp;
                    
                    {{if .flash.error}}
                    <h3>{{.flash.error}}</h3>
                    &nbsp;
                    {{end}}{{if .flash.notice}}
                    <h3>{{.flash.notice}}</h3>
                    &nbsp;
                    {{end}}
                    <form method="POST"> 
                                
                                <div class="col-lg-12">
                                    <div class="form-group col-lg-12">
                                   
                                        <label>Nome</label>     {{if .Errors.First}}<b class="error"> ****{{.Errors.First}}****</b>{{end}}
                                        <input type="text" name="first" class="form-control" value="{{.User.First}}" placeholder="Nome">
                                    </div>
                                    <div class="form-group col-lg-12">
                                        <label>Cognome</label>       {{if .Errors.Last}}<b class="error"> ****{{.Errors.Last}}****</b>{{end}}
                                        <input type="text" name="last" class="form-control" value="{{.User.Last}}" placeholder="Cognome">
                                    </div>
                                    
                                    <div class="form-group col-lg-6">
                                        <label>Password (min. 6 caratteri)</label>      {{if .Errors.Password}}<b class="error"> ****{{.Errors.Password}}****</b>{{end}}
                                        <input type="password" name="password" class="form-control"  placeholder="Password">
                                    </div>
                                    
                                    <div class="form-group col-lg-6">
                                        <label>Ripeti Password</label>      {{if .Errors.Confirm}}<b class="error"><br>****{{.Errors.Confirm}}****</b>{{end}}
                                        <input type="password" name="password2" class="form-control" placeholder="Password">
                                    </div>
                        
                                    <div class="form-group col-lg-12">
                                        <label>Indirizzo Email</label>       {{if .Errors.Email}}<b class="error">  ****{{.Errors.Email}}****</b>{{end}}	
                                        <input type="email" name="email" class="form-control" value="{{.User.Email}}" placeholder="nome@esempio.com">
                                    </div>
                                   
                                </div>
                                &nbsp;
                                <div class="col-lg-12">
                                    <div class="registrationmodal-contratto">
                                    <h1 class="dark-grey">Termini e Condizioni</h3>
                                    <br>
                                    <p>
                                        <b>Premendo su "Registrati" accetti i termini e le condizioni della compagnia</b>
                                    </p>
                                    <p>
                                        
                                    Si impegna:                                    
                                        <li> a non divulgare la suddetta password ad altri soggetti;</li>
                                        <li> a comunicare immediatamente alla E’ Così S.r.l l’eventuale furto, smarrimento o perdita della riservatezza esclusiva della password al fine del suo bloccaggio e/o sostituzione;</li>
                                   <br>
                                    Dichiara di essere a conoscenza:
                                        <li> che username e la password sono nominali;</li>
                                        <li> che tutte le operazioni con essi effettuate sono direttamente attribuibili al proprietario;</li>
                                        <li> che gli accessi e tutte le operazioni effettuate vengono registrati e controllati;</li>
                                        <li> che utilizzi impropri della suddetta password sono puniti a norma di legge;</li>
                                        <li> che l’accesso all’area clienti e l’utilizzo del servizio verranno bloccati in caso di utilizzi impropri della sopra citata password, di una sua divulgazione o di un suo smarrimento, come pure in caso di eventuali violazione di legge commesse mediante l’utilizzo della stessa;</li>
                                        <li> che E’ Così S.r.l. non sarà responsabile della divulgazione dei dati e/o delle informazioni o della perdita di informazioni derivanti da o in qualsiasi modo connessi all’utilizzo non autorizzato dell’area privata;</li>
                                        <li> che E’ Così S.r.l. adotta le misure utili a garantire la sicurezza dei dati dell’utente ed il trattamento dei dati saranno eseguite esclusivamente dai soggetti responsabili del trattamento e/o da Suoi incaricati.</li>
                                    </p>
                                    <p>Il Cliente avuta conoscenza dei diritti ex art. 7 e dell’informativa ex art. 13 D.Lgs. n. 196/2003, consente il trattamento dei dati personali, compresi quelli sensibili ex art. 22 D.Lgs. n. 196/2003, per finalità di gestione del rapporto contrattuale, nonché in base ad obblighi di legge civili, fiscali o di pubblica sicurezza.
                                    </p>
                                    </div>

                                    <div class="col-sm-12">
                                        <button type="submit" value="register" class="btn btn-primary pull-right">Registrati</button>
                                    </div>	                                    
                                    
                                    
                                </div>
                                </form>
                     &nbsp;  
                    <div class="login-help">
                        <a href="http://localhost:8080/">Indietro</a> 
                    </div>
                    </div>
                </div>
            </div>
    </div>
  			  
                
                
                
                
                
                
                
               

