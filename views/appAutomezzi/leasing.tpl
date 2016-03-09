<div id="content">
<h1>Registrazione Contratto Leasing</h1>
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
    <td><input name="ncontratto" type="text" value="{{.ContrLeasing.NContratto}}" autofocus /> {{if .Errors.NContratto}}{{.Errors.NContratto}}{{end}}</td>
</tr>
<tr>      
    <td>Fornitori: </td>
    <td><input name="fornitori" type="text" value="{{.ContrLeasing.Fornitori}}" /> {{if .Errors.Fornitori}}{{.Errors.Fornitori}}{{end}}</td>
</tr>
<tr>
    <td>Data Conttratto:</td>
    <td><input name="datacont" type="text" value="{{.ContrLeasing.DataCont}}" /></td>
</tr>
<tr>      
    <td>Fine Contratto: </td>
    <td><input name="finecont" type="text" value="{{.ContrLeasing.FineCont}}" /> {{if .Errors.FineCont}}{{.Errors.FineCont}}{{end}}</td>
</tr>
<tr>      
    <td>PrimaRata: </td>
    <td><input name="primarata" type="text" value="{{.ContrLeasing.PrimaRata}}" /> {{if .Errors.PrimaRata}}{{.Errors.PrimaRata}}{{end}}</td>
</tr>
<tr>
    <td>Rata Successiva: </td>
    <td><input name="ratasucc" type="text" value="{{.ContrLeasing.RataSucc}}" /> {{if .Errors.RataSucc}}{{.Errors.RataSucc}}{{end}}</td>
</tr>
<tr>      
    <td>Numero Rate: </td>
    <td><input name="nrate" type="text" value="{{.ContrLeasing.NRate}}" /> {{if .Errors.NRate}}{{.Errors.NRate}}{{end}}</td>
</tr>
<tr>      
    <td>Riscatto: </td>
    <td><input name="riscatto" type="text" value="{{.ContrLeasing.Riscatto}}" /> {{if .Errors.Riscatto}}{{.Errors.Riscatto}}{{end}}</td>
</tr>
<tr>      
    <td>DataRiscatto: </td>
    <td><input name="datariscatto" type="text" value="{{.ContrLeasing.DataRiscatto}}" /> {{if .Errors.DataRiscatto}}{{.Errors.DataRiscatto}}{{end}}</td>
</tr>
<tr>      
    <td>ImportoTot: </td>
    <td><input name="importotot" type="text" value="{{.ContrLeasing.ImportoTot}}" /> {{if .Errors.ImportoTot}}{{.Errors.ImportoTot}}{{end}}</td>
</tr>
<tr>      
    <td>FineGaranzia: </td>
    <td><input name="finegaranzia" type="text" value="{{.ContrLeasing.FineGaranzia}}" /> {{if .Errors.FineGaranzia}}{{.Errors.FineGaranzia}}{{end}}</td>
</tr>
<tr>      
    <td>KmInizioGest: </td>
    <td><input name="kminiziogest" type="text" value="{{.ContrLeasing.KmInizioGest}}" /> {{if .Errors.KmInizioGest}}{{.Errors.KmInizioGest}}{{end}}</td>
</tr>
<tr>      
    <td>KmFineGest: </td>
    <td><input name="kmfineGest" type="text" value="{{.ContrLeasing.KmFineGest}}" /> {{if .Errors.KmFineGest}}{{.Errors.KmFineGest}}{{end}}</td>
</tr>
<tr>      
    <td>Note: </td>
    <td><input name="note" type="text" value="{{.ContrLeasing.Note}}" /> {{if .Errors.Note}}{{.Errors.Note}}{{end}}</td>
</tr>
<tr>      
    <td>Allegati: </td>
    <td><input name="allegati" type="text" value="{{.ContrLeasing.Allegati}}" /> {{if .Errors.Allegati}}{{.Errors.Allegati}}{{end}}</td>
</tr>

<tr><td>&nbsp;</td></tr>
<tr>
    <td>&nbsp;</td><td><input type="submit" value="RegisterAcquisto" /></td>
</tr>
</table>
</form>
</div>