<div id="content">
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
    <td><input name="targa" type="text" value="{{.VeicoliDG.Targa}}" autofocus /> {{if .Errors.Targa}}{{.Errors.Targa}}{{end}}</td>
</tr>
<tr>
    <td>Data In Flotta:</td>
    <td><input name="datainflotta" type="text" value="{{.VeicoliDG.DataInFlotta}}" /></td>
</tr>
<tr>
    <td>Data Fine Flotta: </td>
    <td><input name="datafineflotta" type="text" value="{{.VeicoliDG.DataFineFlotta}}" /> {{if .Errors.DataFineFlotta}}{{.Errors.DataFineFlotta}}{{end}}</td>
</tr>
<tr>      
    <td>Note: </td>
    <td><input name="note" type="text" value="{{.VeicoliDG.Note}}" /> {{if .Errors.Note}}{{.Errors.Note}}{{end}}</td>
</tr>
<tr>      
    <td>Tipo Veicolo: </td>
    <td><input name="tipoveicolo" type="text" value="{{.VeicoliDG.TipiVeicolo}}" /> {{if .Errors.TipiVeicolo}}{{.Errors.TipiVeicolo}}{{end}}</td>
</tr>
<tr>      
    <td>Settori: </td>
    <td><input name="settori" type="text" value="{{.VeicoliDG.Settori}}" /> {{if .Errors.Settori}}{{.Errors.Settori}}{{end}}</td>
</tr>
<tr>      
    <td>Condizioni: </td>
    <td><input name="condizioni" type="text" value="{{.VeicoliDG.Condizioni}}" /> {{if .Errors.Condizioni}}{{.Errors.Condizioni}}{{end}}</td>
</tr>
<tr>      
    <td>Impieghi: </td>
    <td><input name="impieghi" type="text" value="{{.VeicoliDG.Impieghi}}" /> {{if .Errors.Impieghi}}{{.Errors.Impieghi}}{{end}}</td>
</tr>
<tr>      
    <td>Conducenti: </td>
    <td><input name="conducenti" type="text" value="{{.VeicoliDG.Conducenti}}" /> {{if .Errors.Conducenti}}{{.Errors.Conducenti}}{{end}}</td>
</tr>

<tr><td>&nbsp;</td></tr>
<tr>
    <td>&nbsp;</td><td><input type="submit" value="RegisterVeicle" /></td>
</tr>
</table>
</form>
</div>