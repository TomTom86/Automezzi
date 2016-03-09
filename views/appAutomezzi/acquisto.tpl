<div id="content">
<h1>Registrazione Contratto Acquisto</h1>
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
    <td>NContratto: </td>
    <td><input name="ncontratto" type="text" value="{{.ContrAcquisti.NContratto}}" autofocus /> {{if .Errors.NContratto}}{{.Errors.NContratto}}{{end}}</td>
</tr>
<tr>
    <td>Data di Acquisto:</td>
    <td><input name="dataacq" type="text" value="{{.ContrAcquisti.DataAcq}}" /></td>
</tr>
<tr>      
    <td>Fornitori: </td>
    <td><input name="fornitori" type="text" value="{{.ContrAcquisti.Fornitori}}" /> {{if .Errors.Fornitori}}{{.Errors.Fornitori}}{{end}}</td>
</tr>
<tr>
    <td>Importo: </td>
    <td><input name="importo" type="text" value="{{.ContrAcquisti.Importo}}" /> {{if .Errors.Importo}}{{.Errors.Importo}}{{end}}</td>
</tr>
<tr>      
    <td>Ammortamento Annuo: </td>
    <td><input name="ammortamentoannuo" type="text" value="{{.ContrAcquisti.AmmortamentoAnnuo}}" /> {{if .Errors.AmmortamentoAnnuo}}{{.Errors.AmmortamentoAnnuo}}{{end}}</td>
</tr>
<tr>      
    <td>FineGaranzia: </td>
    <td><input name="finegaranzia" type="text" value="{{.ContrAcquisti.FineGaranzia}}" /> {{if .Errors.FineGaranzia}}{{.Errors.FineGaranzia}}{{end}}</td>
</tr>
<tr>      
    <td>KmInizioGest: </td>
    <td><input name="kminiziogest" type="text" value="{{.ContrAcquisti.KmInizioGest}}" /> {{if .Errors.KmInizioGest}}{{.Errors.KmInizioGest}}{{end}}</td>
</tr>
<tr>      
    <td>Note: </td>
    <td><input name="note" type="text" value="{{.ContrAcquisti.Note}}" /> {{if .Errors.Note}}{{.Errors.Note}}{{end}}</td>
</tr>
<tr>      
    <td>Allegati: </td>
    <td><input name="allegati" type="text" value="{{.ContrAcquisti.Allegati}}" /> {{if .Errors.Allegati}}{{.Errors.Allegati}}{{end}}</td>
</tr>
<tr><td>&nbsp;</td></tr>
<tr>
    <td>&nbsp;</td><td><input type="submit" value="RegisterAcquisto" /></td>
</tr>
</table>
</form>
</div>