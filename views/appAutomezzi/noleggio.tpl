<div id="content">
<h1>Registrazione Contratto Noleggio</h1>
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
    <td><input name="ncontratto" type="text" value="{{.ContrNoleggi.NContratto}}" autofocus /> {{if .Errors.NContratto}}{{.Errors.NContratto}}{{end}}</td>
</tr>
<tr>      
    <td>Fornitori: </td>
    <td><input name="fornitori" type="text" value="{{.ContrNoleggi.Fornitori}}" /> {{if .Errors.Fornitori}}{{.Errors.Fornitori}}{{end}}</td>
</tr>
<tr>
    <td>Data Contratto:</td>
    <td><input name="datacont" type="text" value="{{.ContrNoleggi.DataCont}}" /></td>
</tr>
<tr>      
    <td>Inizio Contratto: </td>
    <td><input name="datainizio" type="text" value="{{.ContrNoleggi.DataInizio}}" /> {{if .Errors.DataInizio}}{{.Errors.DataInizio}}{{end}}</td>
</tr>
<tr>      
    <td>Fine Contratto: </td>
    <td><input name="datafine" type="text" value="{{.ContrNoleggi.DataFine}}" /> {{if .Errors.DataFine}}{{.Errors.DataFine}}{{end}}</td>
</tr>
<tr>      
    <td>Riparamentrizzazione: </td>
    <td><input name="riparamentrizzazione" type="text" value="{{.ContrNoleggi.Riparamentrizzazione}}" /> {{if .Errors.Riparamentrizzazione}}{{.Errors.Riparamentrizzazione}}{{end}}</td>
</tr>
<tr>
    <td>CanoneBase: </td>
    <td><input name="canonebase" type="text" value="{{.ContrNoleggi.CanoneBase}}" /> {{if .Errors.CanoneBase}}{{.Errors.CanoneBase}}{{end}}</td>
</tr>
<tr>      
    <td>Numero Rate: </td>
    <td><input name="nrate" type="text" value="{{.ContrNoleggi.NRate}}" /> {{if .Errors.NRate}}{{.Errors.NRate}}{{end}}</td>
</tr>
<tr>      
    <td>CanoneServizi: </td>
    <td><input name="canoneservizi" type="text" value="{{.ContrNoleggi.CanoneServizi}}" /> {{if .Errors.CanoneServizi}}{{.Errors.CanoneServizi}}{{end}}</td>
</tr>
<tr>      
    <td>CanoneAltro: </td>
    <td><input name="canonealtro" type="text" value="{{.ContrNoleggi.CanoneAltro}}" /> {{if .Errors.CanoneAltro}}{{.Errors.CanoneAltro}}{{end}}</td>
</tr>
<tr>      
    <td>CanoneTot: </td>
    <td><input name="canonetot" type="text" value="{{.ContrNoleggi.CanoneTot}}" /> {{if .Errors.CanoneTot}}{{.Errors.CanoneTot}}{{end}}</td>
</tr>
<tr>      
    <td>KmContrattuali: </td>
    <td><input name="kmcontrattuali" type="text" value="{{.ContrNoleggi.KmContrattuali}}" /> {{if .Errors.KmContrattuali}}{{.Errors.KmContrattuali}}{{end}}</td>
</tr>
<tr>      
    <td>AddebitoKmExtra: </td>
    <td><input name="addebitokmextra" type="text" value="{{.ContrNoleggi.AddebitoKmExtra}}" /> {{if .Errors.AddebitoKmExtra}}{{.Errors.AddebitoKmExtra}}{{end}}</td>
</tr>
<tr>      
    <td>ImportoKm: </td>
    <td><input name="importokm" type="text" value="{{.ContrNoleggi.ImportoKm}}" /> {{if .Errors.ImportoKm}}{{.Errors.ImportoKm}}{{end}}</td>
</tr>
<tr>      
    <td>ImportoTot: </td>
    <td><input name="importotot" type="text" value="{{.ContrNoleggi.ImportoTot}}" /> {{if .Errors.ImportoTot}}{{.Errors.ImportoTot}}{{end}}</td>
</tr>
<tr>      
    <td>KmInizioGest: </td>
    <td><input name="kminiziogest" type="text" value="{{.ContrNoleggi.KmInizioGest}}" /> {{if .Errors.KmInizioGest}}{{.Errors.KmInizioGest}}{{end}}</td>
</tr>
<tr>      
    <td>KmFineGest: </td>
    <td><input name="kmfinegest" type="text" value="{{.ContrNoleggi.KmFineGest}}" /> {{if .Errors.KmFineGest}}{{.Errors.KmFineGest}}{{end}}</td>
</tr>
<tr>      
    <td>Note: </td>
    <td><input name="note" type="text" value="{{.ContrNoleggi.Note}}" /> {{if .Errors.Note}}{{.Errors.Note}}{{end}}</td>
</tr>
<tr>      
    <td>Allegati: </td>
    <td><input name="allegati" type="text" value="{{.ContrNoleggi.Allegati}}" /> {{if .Errors.Allegati}}{{.Errors.Allegati}}{{end}}</td>
</tr>

<tr><td>&nbsp;</td></tr>
<tr>
    <td>&nbsp;</td><td><input type="submit" value="RegisterAcquisto" /></td>
</tr>
</table>
</form>
</div>