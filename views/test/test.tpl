<div id="content">
<h2>Test.</h2>
<h1>Registrazione Veicolo</h1>
&nbsp;

{{if .flash.error}}
<h3>{{.flash.error}}</h3>
&nbsp;
{{end}}{{if .flash.notice}}
<h3>{{.flash.notice}}</h3>
&nbsp;
{{end}}
<form method="POST">
<table>
<tr>
    <td>Targa: </td>
    <td><input name="first" type="text" value="{{.User.First}}" autofocus /> {{if .Errors.First}}{{.Errors.First}}{{end}}</td>
</tr>
<tr>
    <td>Data In Flotta:</td>
    <td><input name="last" type="text" value="{{.User.Last}}" /></td>
</tr>
<tr>
    <td>Data Fine Flotta: </td>
    <td><input name="email" type="text" value="{{.User.Email}}" /> {{if .Errors.Email}}{{.Errors.Email}}{{end}}</td>
</tr>
<tr>      
    <td>Note: </td>
    <td><input name="password" type="password" /> {{if .Errors.Password}}{{.Errors.Password}}{{end}}</td>
</tr>
<tr>      
    <td>MatriculationYear: </td>
    <td><input name="password2" type="password" /> {{if .Errors.Confirm}}{{.Errors.Confirm}}{{end}}</td>
</tr>
<tr>      
    <td>NLibretto: </td>
    <td><input name="password2" type="password" /> {{if .Errors.Confirm}}{{.Errors.Confirm}}{{end}}</td>
</tr>
<tr>      
    <td>NTelaio: </td>
    <td><input name="password2" type="password" /> {{if .Errors.Confirm}}{{.Errors.Confirm}}{{end}}</td>
</tr>
<tr>      
    <td>Marca: </td>
    <td><input name="password2" type="password" /> {{if .Errors.Confirm}}{{.Errors.Confirm}}{{end}}</td>
</tr>
<tr>      
    <td>Modello: </td>
    <td><input name="password2" type="password" /> {{if .Errors.Confirm}}{{.Errors.Confirm}}{{end}}</td>
</tr>
<tr>      
    <td>NorEuro: </td>
    <td><input name="password2" type="password" /> {{if .Errors.Confirm}}{{.Errors.Confirm}}{{end}}</td>
</tr>
<tr>      
    <td>Kw: </td>
    <td><input name="password2" type="password" /> {{if .Errors.Confirm}}{{.Errors.Confirm}}{{end}}</td>
</tr>
<tr>      
    <td>Cilindrata : </td>
    <td><input name="password2" type="password" /> {{if .Errors.Confirm}}{{.Errors.Confirm}}{{end}}</td>
</tr>
<tr>      
    <td>ConsumoTeorico: </td>
    <td><input name="password2" type="password" /> {{if .Errors.Confirm}}{{.Errors.Confirm}}{{end}}</td>
</tr>
<tr>      
    <td>KmAnno: </td>
    <td><input name="password2" type="password" /> {{if .Errors.Confirm}}{{.Errors.Confirm}}{{end}}</td>
</tr>
<tr>      
    <td>CostoKm: </td>
    <td><input name="password2" type="password" /> {{if .Errors.Confirm}}{{.Errors.Confirm}}{{end}}</td>
</tr>
<tr>      
    <td>Pneumatici : </td>
    <td><input name="password2" type="password" /> {{if .Errors.Confirm}}{{.Errors.Confirm}}{{end}}</td>
</tr>
<tr><td>&nbsp;</td></tr>
<tr>
    <td>&nbsp;</td><td><input type="submit" value="Register" /></td>
</tr>
</table>
</form>
</div>